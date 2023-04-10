package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var images = make([]ExistingImageDetails, 0)

type ExistingImageDetails struct {
	ImageId     string
	ImageTag    string
	ImageLayers []string
}

func GetImagesOnNode(cli client.Client) []ExistingImageDetails {

	ctx := context.Background()

	images, err := cli.ImageList(ctx, types.ImageListOptions{})

	if err != nil {
		panic(err)
	}

	var imageMetadata []ExistingImageDetails

	for _, image := range images {

		if image.RepoTags[0] != "<none>:<none>" {
			imageLayer := GetImageData(cli, image.ID)

			var imageData = ExistingImageDetails{
				ImageId:     imageLayer.ImageId,
				ImageTag:    imageLayer.ImageTag,
				ImageLayers: imageLayer.ImageLayers,
			}

			imageMetadata = append(imageMetadata, imageData)
		} else {
			fmt.Printf("Untagged Image Found: %v \n", image.ID)
		}

	}

	return imageMetadata

}
