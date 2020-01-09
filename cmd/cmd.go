package main

import (
	"log"

	"github.com/spf13/cobra"
)

func execute() {
	var (
		guppyCmd = &cobra.Command{
			Use:     "guppy",
			Version: "0.1.0",
			Short:   "Install the Go programming language",
			Long:    "",
			Example: "",
			Run:     func(ccmd *cobra.Command, args []string) {},
		}

		installCmd = &cobra.Command{
			Use:     "install",
			Short:   "install a specific version of Go",
			Long:    "",
			Example: "",
			Run:     func(ccmd *cobra.Command, args []string) {},
		}

		updateCmd = &cobra.Command{
			Use:     "update",
			Short:   "update to the latest version of Go",
			Long:    "",
			Example: "",
			Run:     func(ccmd *cobra.Command, args []string) {},
		}
	)

	guppyCmd.AddCommand(installCmd, updateCmd)

	if err := guppyCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	execute()
}
