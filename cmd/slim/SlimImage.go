package slim

import (
	"fmt"

	"github.com/darreng1234/docker-optimizer/docker/customBuilder"
	"github.com/docker/docker/client"
	"github.com/pickme-go/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configDir string

// slimImageCmd represents the slimImage command
var SlimImageCmd = &cobra.Command{
	Use:   "slim-image",
	Short: "Slim your Docker image by passing build configs",
	Long:  `Slim your Docker image by passing build configs`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigName("build")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(configDir)

		err := viper.ReadInConfig()

		// Handle errors reading the config file

		if err != nil {
			log.Error("Config Err", err)
			panic(err)
		}

		buildConfigs := customBuilder.BuildConfigs{
			TemplateFiles:         fmt.Sprintf("%v", viper.Get("templateFiles")),
			Technology:            fmt.Sprintf("%v", viper.Get("buildOpts.technology")),
			Version:               fmt.Sprintf("%v", viper.Get("buildOpts.version")),
			CodeDir:               fmt.Sprintf("%v", viper.Get("buildOpts.codeDir")),
			Repository:            fmt.Sprintf("%v", viper.Get("buildOpts.repository")),
			Tag:                   fmt.Sprintf("%v", viper.Get("buildOpts.tag")),
			DefaultDockerFileName: fmt.Sprintf("%v", viper.Get("buildOpts.defaultDockerFileName")),
		}

		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			log.Error("Client Error", err)
			panic(err)
		}
		defer cli.Close()

		customBuilder.BuildImage(*cli, buildConfigs)

	},
}

func init() {

	SlimImageCmd.Flags().StringVarP(&configDir, "configDir", "c", "", "The directory of the build configs")

	if err := SlimImageCmd.MarkFlagRequired("configDir"); err != nil {
		log.Error("Not Found", err)
	}

}
