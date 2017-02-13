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

	service, _, err := cli.ServiceInspectWithRaw(context.Background(), "foo")
	if err != nil {
		fmt.Println("could not find service foo")
		panic(err)
	}

	var resp2 types.ServiceUpdateResponse
	resp2, err = cli.ServiceUpdate(context.Background(), service.ID, service.Version, serviceSpec, types.ServiceUpdateOptions{})
	if err != nil {
		fmt.Println("tried to update service1", resp2)
		panic(err)
	}

	fmt.Println("successfully updated service", resp2)
	fmt.Println("job's done!")
}
