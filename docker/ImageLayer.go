package docker

import (
	"context"

	"github.com/docker/docker/client"
)

func GetBuiltImageLayer(client client.Client, imageId string) []string {

	ctx := context.Background()

	resp, _, err := client.ImageInspectWithRaw(ctx, imageId)

	if err != nil {
		panic(err)
	}

	client.Close()
	builtImageLayers := resp.RootFS.Layers

	return builtImageLayers

}
