package dockercontrol

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	u "primitivo.fr/applinh/go-docker-compose/utils"
)

type Vol struct {
	Labels map[string]string
}

func CreateNewContainer(composeUuid string) (string, error) {

	dir, _ := os.Getwd()
	fmt.Println(dir)
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println("Unable to create docker client")
		panic(err)
	}

	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: "8000",
	}
	containerPort, err := nat.NewPort("tcp", "80")
	if err != nil {
		panic("Unable to get the port")
	}
	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	cont, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Entrypoint: []string{"/usr/scripts/start.sh"},
			Image:      "docker/compose",
		},
		&container.HostConfig{
			PortBindings: portBinding,
			Binds:        []string{dir + "/scripts:/usr/scripts", "/var/run/docker.sock:/var/run/docker.sock", dir + "/composefiles/" + composeUuid + ":/var/tmp/docker/compose"},
		}, nil, "")
	if err != nil {
		panic(err)
	}

	cli.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
	fmt.Printf("Container %s has started \n", cont.ID)
	return cont.ID, nil
}

func ListContainers() ([]string, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	if len(containers) > 0 {
		r_containers := []string{}
		for _, container := range containers {
			r_containers = append(r_containers, container.ID)
		}
		return r_containers, nil
	} else {
		return nil, &u.ErrorString{S: "NO_STACK"}
	}
}

func StopContainer(containerID string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	err = cli.ContainerStop(context.Background(), containerID, nil)
	if err != nil {
		panic(err)
	}
	return err
}
