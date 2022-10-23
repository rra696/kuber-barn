package pod

import (
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/oci"
)

func NewPod(registryImage, name string) (*Pod, error) {
	ctx, client, err := initContainerdConnection()
	if err != nil {
		return nil, err
	}

	image, err := client.Pull(ctx, registryImage, containerd.WithPullUnpack)
	if err != nil {
		return nil, err
	}

	id := generateNewID(name)

	container, err := client.NewContainer(
		ctx,
		id,
		containerd.WithImage(image),
		containerd.WithNewSnapshot(id+"-snapshot", image),
		containerd.WithNewSpec(oci.WithImageConfig(image)),
	)
	if err != nil {
		return nil, err
	}

	return &Pod{
		ID:        id,
		ctx:       &ctx,
		container: &container,
		client:    client,
	}, nil
}

func ListRunningPods() ([]string, error) {
	ctx, client, err := initContainerdConnection()
	if err != nil {
		return nil, err
	}

	var runningPods []string

	containers, err := client.Containers(ctx)
	if err != nil {
		return runningPods, err
	}

	for _, container := range containers {
		_, err = container.Task(ctx, cio.Load)
		if err == nil {
			runningPods = append(runningPods, container.ID())
		}
	}

	return runningPods, nil
}

func KillPod(name string) (string, error) {
	ctx, client, err := initContainerdConnection()
	if err != nil {
		return "", err
	}

	container, err := client.LoadContainer(ctx, name)
	if err != nil {
		return "", err
	}

	task, err := container.Task(ctx, cio.Load)
	if err != nil {
		return "", err
	}

	exitStatusC, err := task.Wait(ctx)
	if err != nil {
		return "", err
	}

	runningPod := RunningPod{
		Pod: &Pod{
			ID: name,
			client: client,
			ctx: &ctx,
			container: &container,
		},
		task: &task,
		exitStatusC: exitStatusC,
	}

	_, err = runningPod.Kill()
	if err != nil {
		return "", err
	}

	err = runningPod.Pod.Delete()
	if err != nil {
		return "", err
	}

	return name, nil
}
