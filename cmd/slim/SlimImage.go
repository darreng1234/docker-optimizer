package slim

import (
	"fmt"

	"github.com/darreng1234/docker-optimizer/docker/customBuilder"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configDir string

// slimImageCmd represents the slimImage command
var SlimImageCmd = &cobra.Command{
	Use:   "slim-image",
	Short: "slim-image",
	Long:  `A longer description of slim-image`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigName("build")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(configDir)

		err := viper.ReadInConfig()

		if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("fatal error config file: %w", err))
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
			panic(err)
		}
		defer cli.Close()

		customBuilder.BuildImage(*cli, buildConfigs)

	},
}

func init() {

	SlimImageCmd.Flags().StringVarP(&configDir, "configDir", "c", "", "The directory of the build configs")

	if err := SlimImageCmd.MarkFlagRequired("configDir"); err != nil {
		fmt.Println(err)
	}

}
