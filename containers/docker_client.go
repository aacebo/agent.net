package containers

import (
	"context"
	"strconv"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type DockerClient struct {
	client *client.Client
}

func NewDockerClient() (*DockerClient, error) {
	docker, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)

	if err != nil {
		return nil, err
	}

	return &DockerClient{
		client: docker,
	}, nil
}

func (self *DockerClient) Create(args ContainerCreateArgs) (string, error) {
	res, err := self.client.ContainerCreate(context.Background(), &container.Config{
		Image:        args.Image,
		ExposedPorts: nat.PortSet{"80/tcp": {}},
		Env:          args.Env,
		Tty:          true,
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port("80/tcp"): []nat.PortBinding{{
				HostIP:   args.IPAddress,
				HostPort: strconv.Itoa(args.Port),
			}},
		},
	}, nil, nil, args.Name)

	if err != nil {
		return "", err
	}

	return res.ID, nil
}

func (self *DockerClient) Start(id string) error {
	return self.client.ContainerStart(
		context.Background(),
		id,
		container.StartOptions{},
	)
}

func (self *DockerClient) Stop(id string) error {
	return self.client.ContainerStop(
		context.Background(),
		id,
		container.StopOptions{},
	)
}
