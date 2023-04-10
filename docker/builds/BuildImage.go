package docker

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

func BuildImage() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	os.Link("/Users/darren/Desktop/Optimzer/docker-optimizer/templates/python/3.8-python-Dockerfile", "/Users/darren/Desktop/Optimzer/sample-python-app/user-service/3.8-python-Dockerfile")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	tar, err := archive.TarWithOptions("/Users/darren/Desktop/Optimzer/sample-python-app/user-service/", &archive.TarOptions{})
	if err != nil {
		fmt.Printf("%v", err)
	}

	opts := types.ImageBuildOptions{
		Dockerfile: "3.8-python-Dockerfile",
		Tags:       []string{"test-auto-build" + "/node-hello"},
		Remove:     true,
	}
	res, err := cli.ImageBuild(ctx, tar, opts)
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("%v", res)

	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		//lastLine = scanner.Text()
		fmt.Println(scanner.Text())
	}

}
