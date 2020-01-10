package main

import (
	"fmt"
	"log"

	"github.com/j4ng5y/guppy/install"
	"github.com/spf13/cobra"
)

func isValidArg(arg string, validArgs []string) bool {
	for _, i := range validArgs {
		if i == arg {
			return true
		}
	}
	return false
}

func execute() {
	var (
		guppyCmd = &cobra.Command{
			Use:     "guppy",
			Version: "0.1.0",
			Short:   "Guppy - A tiny installer for the Go programming language",
			Args:    cobra.NoArgs,
			Run: func(ccmd *cobra.Command, args []string) {
				ccmd.Usage()
			},
		}

		installCmd = &cobra.Command{
			Use:     "install",
			Short:   "Install a specific version of Go",
			Long:    guppyInstallLong,
			Example: "  guppy install go1\n  guppy install go1.13\n  guppy install go1.13.6",
			Args: func(ccmd *cobra.Command, args []string) error {
				switch {
				case !isValidArg(args[0], validGoVers):
					return fmt.Errorf("%s is not a valid Go version", args[0])
				case len(args) < 1:
					return fmt.Errorf("only one or zero arguments are allowed")
				}
				return nil
			},
			Run: func(ccmd *cobra.Command, args []string) {
				I := install.New(args[0])
				if err := I.Run(); err != nil {
					log.Fatal(err)
				}
			},
		}

		updateCmd = &cobra.Command{
			Use:   "update",
			Short: "Update to the latest version of Go",
			Run:   func(ccmd *cobra.Command, args []string) {},
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

const guppyInstallLong = `Guppy - Install
  The install command will install Go on your system.
  
  -  If no argument is provided, Guppy will install the latest version of Go unless another version is detected.
  -  If only a Go major version is provided, Guppy will install the lasest minor.patch version of that major release.
  -  If only a Go major.minor version is provided, Guppy will install the lastest patch version of that major.minor release.
  -  If a full Go major.minor.match version is provided, Guppy will install that version.
`

var validGoVers = []string{
	"go1",
	"go1.0.0",
	"go1.0.1",
	"go1.0.2",
	"go1.0.3",
	"go1.1",
	"go1.1.0",
	"go1.1.1",
	"go1.1.2",
	"go1.2",
	"go1.2.0",
	"go1.2.1",
	"go1.2.2",
	"go1.3",
	"go1.3.0",
	"go1.3.1",
	"go1.3.2",
	"go1.3.3",
	"go1.4",
	"go1.4.0",
	"go1.4.1",
	"go1.4.2",
	"go1.4.3",
	"go1.5",
	"go1.5.0",
	"go1.5.1",
	"go1.5.2",
	"go1.5.3",
	"go1.5.4",
	"go1.6",
	"go1.6.0",
	"go1.6.1",
	"go1.6.2",
	"go1.6.3",
	"go1.6.4",
	"go1.7",
	"go1.7.0",
	"go1.7.1",
	"go1.7.2",
	"go1.7.3",
	"go1.7.4",
	"go1.7.5",
	"go1.7.6",
	"go1.8",
	"go1.8.0",
	"go1.8.1",
	"go1.8.2",
	"go1.8.3",
	"go1.8.4",
	"go1.8.5",
	"go1.8.6",
	"go1.8.7",
	"go1.9",
	"go1.9.0",
	"go1.9.1",
	"go1.9.2",
	"go1.9.3",
	"go1.9.4",
	"go1.9.5",
	"go1.9.6",
	"go1.9.7",
	"go1.10",
	"go1.10.0",
	"go1.10.1",
	"go1.10.2",
	"go1.10.3",
	"go1.10.4",
	"go1.10.5",
	"go1.10.6",
	"go1.10.7",
	"go1.10.8",
	"go1.11",
	"go1.11.0",
	"go1.11.1",
	"go1.11.2",
	"go1.11.3",
	"go1.11.4",
	"go1.11.5",
	"go1.11.6",
	"go1.11.7",
	"go1.11.8",
	"go1.11.9",
	"go1.11.10",
	"go1.11.11",
	"go1.11.12",
	"go1.11.13",
	"go1.12",
	"go1.12.0",
	"go1.12.1",
	"go1.12.2",
	"go1.12.3",
	"go1.12.4",
	"go1.12.5",
	"go1.12.6",
	"go1.12.7",
	"go1.12.8",
	"go1.12.9",
	"go1.12.10",
	"go1.12.11",
	"go1.12.12",
	"go1.12.13",
	"go1.12.14",
	"go1.12.15",
	"go1.13",
	"go1.13.0",
	"go1.13.1",
	"go1.13.2",
	"go1.13.3",
	"go1.13.4",
	"go1.13.5",
	"go1.13.6",
}
