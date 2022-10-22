package pod

import (
	"context"
	"fmt"
	"syscall"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
)

type Pod struct {
	ID        string
	client    *containerd.Client
	ctx       *context.Context
	container *containerd.Container
}

type RunningPod struct {
	Pod         *Pod
	task        *containerd.Task
	exitStatusC <-chan containerd.ExitStatus
}

func (pod *Pod) Run() (*RunningPod, error) {
	task, err := (*pod.container).NewTask(*pod.ctx, cio.LogFile(fmt.Sprintf("%s%s", LogsPath, pod.ID)))
	if err != nil {
		return nil, err
	}

	exitStatusC, err := task.Wait(*pod.ctx)
	if err != nil {
		return nil, err
	}

	err = task.Start(*pod.ctx)
	if err != nil {
		return nil, err
	}

	return &RunningPod{
		Pod:         pod,
		task:        &task,
		exitStatusC: exitStatusC,
	}, nil
}

func (pod *Pod) Delete() error {
	err := (*pod.container).Delete(*pod.ctx, containerd.WithSnapshotCleanup)
	if err != nil {
		return err
	}

	err = pod.client.Close()
	if err != nil {
		return err
	}

	return nil
}

func (runningPod *RunningPod) Kill() (uint32, error) {
	err := (*runningPod.task).Kill(*runningPod.Pod.ctx, syscall.SIGTERM)
	if err != nil {
		return 0, err
	}

	status := <-runningPod.exitStatusC
	code, _, err := status.Result()
	if err != nil {
		return 0, err
	}

	_, err = (*runningPod.task).Delete(*runningPod.Pod.ctx)
	if err != nil {
		return 0, err
	}

	return code, nil
}
