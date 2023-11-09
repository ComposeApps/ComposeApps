/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializing ComposeApps dependencies",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !isDockerInstalled() {
			fmt.Println("Docker is not installed or docker compose plugin is missing. Refer to https://docs.docker.com/compose/install/linux/#install-using-the-repository.")
			return
		} else {
			fmt.Println("Docker is installed. Initializing configuration....")
		}
	},
}

func isDockerInstalled() bool {
  _, err := exec.Command("docker", "compose", "--version").Output()
  if err != nil {
    return false
  }
  return err == nil
}

func init() {
	rootCmd.AddCommand(initCmd)
}

