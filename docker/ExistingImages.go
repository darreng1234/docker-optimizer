package docker

import (
	"context"

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

		imageLayer := GetBuiltImageLayer(cli, image.ID)
		var imageData = ExistingImageDetails{
			ImageId:     image.ID,
			ImageTag:    image.RepoTags[0],
			ImageLayers: imageLayer,
		}

		imageMetadata = append(imageMetadata, imageData)

	}

	return imageMetadata

}
