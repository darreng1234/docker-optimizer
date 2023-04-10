/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package slim

import (
	"fmt"

	"github.com/spf13/cobra"
)

// slimImageCmd represents the slimImage command
var SlimImageCmd = &cobra.Command{
	Use:   "slim-image",
	Short: "slim-image",
	Long:  `A longer description of slim-image`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("slimImage called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// slimImageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// slimImageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
