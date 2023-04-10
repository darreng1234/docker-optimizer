package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

func GetImageData(client client.Client, imageId string) ExistingImageDetails {

	ctx := context.Background()

	resp, _, err := client.ImageInspectWithRaw(ctx, imageId)

	var imageData ExistingImageDetails

	//Check for untagged Images or errors thrown by docker sdk
	if err != nil {
		fmt.Printf("%v \n", err)
	} else {
		imageData.ImageId = resp.ID
		imageData.ImageTag = resp.RepoTags[0]
		imageData.ImageLayers = resp.RootFS.Layers
	}

	client.Close()

	return imageData

}
