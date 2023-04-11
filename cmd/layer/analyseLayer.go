/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package layer

import (
	"fmt"

	docker "github.com/darreng1234/docker-optimizer/docker/layer"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

type imageSimilarityData struct {
	imageId          string
	ImageTag         string
	imageLayers      []string
	similarImageTags []string
}

var imageId string

// analyseLayerCmd represents the analyseLayer command
var AnalyseLayerCmd = &cobra.Command{
	Use:   "analyse-layer",
	Short: "Layer Analysis",
	Long:  `Layer Analysis Long Desc`,
	Run: func(cmd *cobra.Command, args []string) {

		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}
		defer cli.Close()

		buildImageData := docker.GetImageData(*cli, imageId)

		if buildImageData.ImageId != "" {
			existingImageMetadata := docker.GetImagesOnNode(*cli)

			var similarImageTags []string

			for _, existingImageData := range existingImageMetadata {
				if checkImageSimilarity(buildImageData.ImageLayers, existingImageData.ImageLayers) {
					similarImageTags = append(similarImageTags, existingImageData.ImageTag)
					image := imageSimilarityData{
						imageId:          buildImageData.ImageId,
						ImageTag:         buildImageData.ImageTag,
						imageLayers:      buildImageData.ImageLayers,
						similarImageTags: similarImageTags,
					}

					fmt.Printf("The Image: %v has %v common layers with tags: %v\n", image.ImageTag, len(image.similarImageTags), image.similarImageTags)
				}
			}
		} else {
			fmt.Println("Application Exitting Now ...")
		}

	},
}

func checkImageSimilarity(builtImageLayers []string, existingImageLayers []string) bool {
	var similarityDetected bool
	for _, builtImageLayer := range builtImageLayers {
		for _, existingImageLayer := range existingImageLayers {
			if builtImageLayer == existingImageLayer {
				similarityDetected = true
			} else {
				similarityDetected = false
			}
		}
	}
	return similarityDetected
}

func init() {

	AnalyseLayerCmd.Flags().StringVarP(&imageId, "imageId", "i", "", "The Image Id")

	if err := AnalyseLayerCmd.MarkFlagRequired("imageId"); err != nil {
		fmt.Println(err)
	}

}
