package customBuilder

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"regexp"

	"github.com/darreng1234/docker-optimizer/docker/layer"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/pickme-go/log"
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

	if buildConfigs.Technology == "python" {

		foundv3, err := regexp.MatchString("^3.*", buildConfigs.Version)
		if err != nil {
			log.Error("Regex Error", err)
		}

		foundv2, err := regexp.MatchString("^2.*", buildConfigs.Version)
		if err != nil {
			log.Error("Regex Error", err)
		}

		if foundv3 {
			BuildImageWithVersion(client, buildConfigs, "python3-Dockerfile")
		} else if foundv2 {
			BuildImageWithVersion(client, buildConfigs, "python2-Dockerfile")
		} else {
			log.Error("Not Supported", buildConfigs.Technology, buildConfigs.Version, "version not supported yet")
		}

	} else {
		log.Warn("Not Supported", buildConfigs.Technology, " Not Supported Yet")
	}

}

func imageCompare(client client.Client, rawImageTag string, optimizedImageTag string) (float64, float64) {

	var rawImageSize float64
	var optimizedImageSize float64
	existingImages := layer.GetImagesOnNode(client)

	//fmt.Printf("%v,%v", rawImageTag, optimizedImageTag)

	for _, image := range existingImages {

		if image.ImageTag == rawImageTag {
			rawImageSize = float64(image.ImageSize)
		}
		if image.ImageTag == optimizedImageTag {
			optimizedImageSize = float64(image.ImageSize)
		}
	}

	//fmt.Printf("%v,%v", rawImageSize, optimizedImageSize)
	return rawImageSize, optimizedImageSize

}

func BuildImageWithVersion(client client.Client, buildConfigs BuildConfigs, manifest string) {

	os.Link(buildConfigs.TemplateFiles+manifest, buildConfigs.CodeDir+manifest)

	defer os.Remove(buildConfigs.CodeDir + manifest)

	ctx := context.Background()

	tar, err := archive.TarWithOptions(buildConfigs.CodeDir, &archive.TarOptions{})
	if err != nil {
		log.Error("Not Found", err)
	}

	optsOptimized := types.ImageBuildOptions{
		Dockerfile: manifest,
		Tags:       []string{buildConfigs.Repository + ":" + buildConfigs.Tag + "optimized"},
		Remove:     true,
	}

	res, err := client.ImageBuild(ctx, tar, optsOptimized)
	if err != nil {
		log.Error("Not Found", err)
	}

	fmt.Printf("%v", res)

	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	ctxRaw := context.Background()
	tarRaw, err := archive.TarWithOptions(buildConfigs.CodeDir, &archive.TarOptions{})
	if err != nil {
		log.Error("Tar Err", err)
	}

	optsRaw := types.ImageBuildOptions{
		Dockerfile: buildConfigs.DefaultDockerFileName,
		Tags:       []string{buildConfigs.Repository + ":" + buildConfigs.Tag + "raw"},
		Remove:     true,
	}

	resRaw, err := client.ImageBuild(ctxRaw, tarRaw, optsRaw)
	if err != nil {
		log.Error("Client Error", err)
	}

	scannerRaw := bufio.NewScanner(resRaw.Body)
	for scannerRaw.Scan() {
		fmt.Println(scannerRaw.Text())
	}

	rawImageSize, optimizedImageSize := imageCompare(client, buildConfigs.Repository+":"+buildConfigs.Tag+"raw", buildConfigs.Repository+":"+buildConfigs.Tag+"optimized")

	// time.Sleep(5 * time.Second)
	percentageReduction := ((rawImageSize - optimizedImageSize) / rawImageSize) * 100
	fmt.Printf("User Image Size: %vMB\n", int64(rawImageSize)/1000000)
	fmt.Printf("Optimizer Image Size: %vMB\n", int64(optimizedImageSize)/1000000)
	fmt.Printf("Optimzer image reduction percentage: %.2f%%\n", percentageReduction)

}
