package main

import (
	"context"
	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/container"
	"docker.io/go-docker/api/types/network"
	"fmt"
	"io"
)

func init() {
}

func main() {

	ctx := context.Background()

	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containerConfig := container.Config{
		Tty:          true,
		Cmd:          []string{"ping", "-c", "4", "8.8.8.8"},
		Image:        "centos",
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
	}

	cBody, err := cli.ContainerCreate(
		ctx,
		&containerConfig,
		&container.HostConfig{},
		&network.NetworkingConfig{},
		"practice-container",
	)
	defer cli.ContainerRemove(ctx, cBody.ID, types.ContainerRemoveOptions{})
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, cBody.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, cBody.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	logOptions := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Timestamps: true,
		Follow:     true,
		Details:    true,
	}
	readCloser, err := cli.ContainerLogs(ctx, cBody.ID, logOptions)
	defer readCloser.Close()
	if err != nil {
		panic(err)
	}

	buff := make([]byte, 256)
	for {
		_, err := readCloser.Read(buff)
		fmt.Printf("%s", string(buff))
		if err == io.EOF {
			break
		}
	}

	fmt.Printf("%s %v\n", cBody.ID[:10], cBody.Warnings)

}

func startContainer(ctx context.Context) (containerID int, err error) {
}
