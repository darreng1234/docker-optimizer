/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/darreng1234/docker-optimizer/cmd/layer"
	"github.com/darreng1234/docker-optimizer/cmd/manifest"
	"github.com/darreng1234/docker-optimizer/cmd/slim"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "docker-optimizer",
	Short: "Choose an optimization option to continue with the process",
	Long:  `Choose an optimization option to continue with the process`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommandPallette() {
	rootCmd.AddCommand(layer.AnalyseLayerCmd)
	rootCmd.AddCommand(slim.SlimImageCmd)
	rootCmd.AddCommand(manifest.ManifestCheckerCmd)
}

func init() {
	addSubCommandPallette()
}
