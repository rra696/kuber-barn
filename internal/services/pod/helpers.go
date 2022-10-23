package pod

import (
	"context"
	"fmt"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
	"github.com/google/uuid"
)

func initContainerdConnection() (context.Context, *containerd.Client, error) {
	client, err := containerd.New(SocketPath)
	if err != nil {
		return nil, nil, err
	}

	ctx := namespaces.WithNamespace(context.Background(), Namespace)

	return ctx, client, err
}

func generateNewID(name string) string {
	id := uuid.New()

	return fmt.Sprintf("%s-%s", name, id)
}
