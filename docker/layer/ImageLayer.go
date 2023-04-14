package layer

import (
	"context"

	"github.com/docker/docker/client"
	"github.com/pickme-go/log"
)

func GetImageData(client client.Client, imageId string) ExistingImageDetails {

	ctx := context.Background()

	resp, _, err := client.ImageInspectWithRaw(ctx, imageId)

	var imageData ExistingImageDetails

	//Check for untagged Images or errors thrown by docker sdk
	if err != nil {
		log.Error("Untagged", err)
	} else {
		imageData.ImageId = resp.ID
		imageData.ImageTag = resp.RepoTags[0]
		imageData.ImageLayers = resp.RootFS.Layers
		imageData.ImageSize = resp.Size
	}

	client.Close()

	return imageData

}
