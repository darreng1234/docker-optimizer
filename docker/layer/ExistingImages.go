package layer

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pickme-go/log"
)

var images = make([]ExistingImageDetails, 0)

type ExistingImageDetails struct {
	ImageId     string
	ImageTag    string
	ImageLayers []string
	ImageSize   int64
}

func GetImagesOnNode(cli client.Client) []ExistingImageDetails {

	ctx := context.Background()

	images, err := cli.ImageList(ctx, types.ImageListOptions{})

	if err != nil {
		log.Error("Build Error", err)
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
				ImageSize:   imageLayer.ImageSize,
			}

			imageMetadata = append(imageMetadata, imageData)
		} else {
			log.Warn("Untagged Image Found", image.ID)
		}

	}

	return imageMetadata

}
