package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)
import "github.com/docker/docker/api/types/swarm"

func main() {
	fmt.Println("starting")

	var cli *client.Client
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.24", nil, defaultHeaders)
	if err != nil {
		fmt.Println("tried to create docker client")
		panic(err)
	}

	serviceSpec := swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name: "foo",
		},
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: swarm.ContainerSpec{
				Image: "alpine:latest",
			},
			RestartPolicy: &swarm.RestartPolicy{
				Condition: swarm.RestartPolicyCondition("any"),
			},
		},
		Networks: []swarm.NetworkAttachmentConfig{
			swarm.NetworkAttachmentConfig{Target: "swarmnet"},
		},
	}

	var resp types.ServiceCreateResponse
	resp, err = cli.ServiceCreate(context.Background(), serviceSpec, types.ServiceCreateOptions{})
	if err != nil {
		fmt.Println("tried to create service", resp)
		panic(err)
	}
	fmt.Println("successfully created service", resp)
	fmt.Println("job's done!")
}
