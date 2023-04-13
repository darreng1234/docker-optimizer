/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package manifest

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"

	"github.com/pickme-go/log"
	"github.com/spf13/cobra"
)

var dockerFilePath string
var packageManager string

// manifestCheckerCmd represents the manifestChecker command
var ManifestCheckerCmd = &cobra.Command{
	Use:   "manifest-checker",
	Short: "manifest-checker",
	Long:  `manifest-checker`,
	Run: func(cmd *cobra.Command, args []string) {

		availablePackages := [...]string{"pip", "apt"}

		file, err := ioutil.ReadFile(dockerFilePath)
		if err != nil {
			log.Error("Not Found", err)
		}

		fileString := string(file)

		//Logic to check for slimming parameters

		if strings.Contains(fileString, "slim") || strings.Contains(fileString, "distroless") {
			fmt.Print("*------------------------------*|  Finding Slimming Parameters  |*------------------------------*\n\n")
			time.Sleep(5 * time.Second)
			log.Info("Found", "Slim Image Parameters Found")
			fmt.Print("Few other base image that can be used \n\nAlpine Images\n =>https://hub.docker.com/_/alpine\n\nDistroless Images\n =>https://github.com/GoogleContainerTools/distroless\n\n")
			fmt.Print("*--------------------------------------------------------------------------------------------*\n\n")
		} else {
			fmt.Print("*------------------------------*|  Finding Slimming Parameters  |*------------------------------*\n\n")
			time.Sleep(5 * time.Second)
			log.Warn("Not Found", "Slim Image Parameters Not Found")
			fmt.Print("Few other base image that can be used \n\nAlpine Images\n =>https://hub.docker.com/_/alpine\n\nDistroless Images\n =>https://github.com/GoogleContainerTools/distroless\n\n")
		}

		// Logic to check for package cleanup

		for _, pkg := range availablePackages {
			found, _ := regexp.MatchString(pkg, packageManager)

			if found {
				if pkg == "pip" {
					if strings.Contains(fileString, "--no-cache-dir") {
						fmt.Print("*------------------------------*|  Finding Package Cleanup Params  |*------------------------------*\n\n")
						time.Sleep(5 * time.Second)
						log.Info("Found", "PIP cleanup commands found\n\n")
						//fmt.Print("*--------------------------------------------------------------------------------------------------*\n\n")
					} else {
						fmt.Print("*------------------------------*|  Finding Package Cleanup Params  |*------------------------------*\n\n")
						time.Sleep(5 * time.Second)
						log.Warn("Not Found", "PIP Cleanup commnds not found")
						fmt.Print("Check below documentation for PIP package cleanup \n\n=> PIP cache maintain - https://pip.pypa.io/en/stable/cli/pip_cache/\n\n")
						//fmt.Print("*--------------------------------------------------------------------------------------------------*\n\n")
					}
				} else if pkg == "pip" {
					if strings.Contains(fileString, "autoclean") || strings.Contains(fileString, "autoremove") {
						fmt.Print("*------------------------------*|  Finding Package Cleanup Params  |*------------------------------*\n\n")
						time.Sleep(5 * time.Second)
						log.Info("Found", "APT cleanup commands found")
					} else {
						fmt.Print("*------------------------------*|  Finding Package Cleanup Params  |*------------------------------*\n\n")
						time.Sleep(5 * time.Second)
						log.Info("Checking", "Finding Slimming Params")
						log.Warn("Not Found", "APT Cleanup commnds not found")
						fmt.Print("Check below documentation for APT package cleanup \n\n=> APT unused package remove - => APT Remove - https://bit.ly/3UAB2bn\n\n=> APT unused dependancy remove - => APT Cleanup - https://linux.die.net/man/8/apt-get\n")
					}
				}

			}

		}

		// Logic to check for multistage Builds

		if strings.Contains(fileString, "COPY --from") || strings.Contains(fileString, "AS") || strings.Contains(fileString, "as") {
			fmt.Print("*------------------------------*|  Finding Multistage Build Parameters  |*------------------------------*\n\n")
			time.Sleep(5 * time.Second)
			log.Info("Found", "Miltistage build parameters found")
			fmt.Print("\n*--------------------------------------------------------------------------------------------------*\n\n")
		} else {
			fmt.Print("*------------------------------*|  Finding Multistage Build Parameters  |*------------------------------*\n\n")
			time.Sleep(5 * time.Second)
			log.Warn("Not Found", "Miltistage build parameters found")
			fmt.Print("Check below documentation for multistage builds \n\n=> Building Docker Images by Stages - https://docs.docker.com/build/building/multi-stage/\n\n=> Copy Dependancies from one stage to another - https://docs.docker.com/engine/reference/builder/\n\n")
			fmt.Print("\n*-----------------------------------------------------------------------------------------------------*\n\n")
		}

	},
}

func init() {

	ManifestCheckerCmd.Flags().StringVarP(&dockerFilePath, "file", "f", "", "Docker manifest file path")

	ManifestCheckerCmd.Flags().StringVarP(&packageManager, "package", "p", "", "Package manager used (pip|apt) ")

	err := ManifestCheckerCmd.MarkFlagRequired("file")
	if err != nil {
		log.Error("Log File Not Found", dockerFilePath, err)
	}

	errPkg := ManifestCheckerCmd.MarkFlagRequired("package")
	if errPkg != nil {
		log.Error("Please specify a package type (pip|apt)", packageManager, errPkg)
	}

}
