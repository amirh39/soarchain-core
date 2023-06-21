package cmd

import (
	"fmt"

	"soarchain/version"

	"github.com/spf13/cobra"
)

func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the application binary version information",
		RunE: func(cmd *cobra.Command, _ []string) error {
			fmt.Print(version.Version())
			return nil
		},
	}
}
