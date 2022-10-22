package pod

import (
	"context"
	"fmt"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/containerd/oci"
	"github.com/google/uuid"
)

func NewPod(registryImage, name string) (*Pod, error) {
	client, err := containerd.New(SocketPath)
	if err != nil {
		return nil, err
	}

	ctx := namespaces.WithNamespace(context.Background(), Namespace)

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

func generateNewID(name string) string {
	id := uuid.New()

	return fmt.Sprintf("%s-%s", name, id)
}
