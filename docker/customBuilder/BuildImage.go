package customBuilder

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/darreng1234/docker-optimizer/docker/layer"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

type BuildConfigs struct {
	TemplateFiles         string
	Technology            string
	Version               string
	CodeDir               string
	Repository            string
	Tag                   string
	DefaultDockerFileName string
}

func BuildImage(client client.Client, buildConfigs BuildConfigs) {

	os.Link(buildConfigs.TemplateFiles, buildConfigs.CodeDir+"python3-Dockerfile")

	ctx := context.Background()

	tar, err := archive.TarWithOptions(buildConfigs.CodeDir, &archive.TarOptions{})
	if err != nil {
		fmt.Printf("%v", err)
	}

	optsOptimized := types.ImageBuildOptions{
		Dockerfile: "python3-Dockerfile",
		Tags:       []string{buildConfigs.Repository + ":" + buildConfigs.Tag + "optimized"},
		Remove:     true,
	}
	res, err := client.ImageBuild(ctx, tar, optsOptimized)
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("%v", res)

	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	ctxRaw := context.Background()
	tarRaw, err := archive.TarWithOptions(buildConfigs.CodeDir, &archive.TarOptions{})
	if err != nil {
		fmt.Printf("%v", err)
	}

	optsRaw := types.ImageBuildOptions{
		Dockerfile: buildConfigs.DefaultDockerFileName,
		Tags:       []string{buildConfigs.Repository + ":" + buildConfigs.Tag + "raw"},
		Remove:     true,
	}

	resRaw, err := client.ImageBuild(ctxRaw, tarRaw, optsRaw)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v", resRaw)

	scannerRaw := bufio.NewScanner(resRaw.Body)
	for scannerRaw.Scan() {
		fmt.Println(scannerRaw.Text())
	}

	rawImageSize, optimizedImageSize := imageCompare(client, buildConfigs.Repository+":"+buildConfigs.Tag+"raw", buildConfigs.Repository+":"+buildConfigs.Tag+"optimized")

	percentageReduction := ((rawImageSize - optimizedImageSize) / rawImageSize) * 100
	fmt.Printf("User Image Size: %vMB\n", int64(rawImageSize)/1000000)
	fmt.Printf("Optimizer Image Size: %vMB\n", int64(optimizedImageSize)/1000000)
	fmt.Printf("Optimzer image reduction percentage: %.2f%%\n", percentageReduction)

}

func imageCompare(client client.Client, rawImageTag string, optimizedImageTag string) (float64, float64) {

	var rawImageSize float64
	var optimizedImageSize float64
	existingImages := layer.GetImagesOnNode(client)

	for _, image := range existingImages {

		if image.ImageTag == rawImageTag {
			rawImageSize = float64(image.ImageSize)
		}
		if image.ImageTag == optimizedImageTag {
			optimizedImageSize = float64(image.ImageSize)
		}
	}

	return rawImageSize, optimizedImageSize

}
