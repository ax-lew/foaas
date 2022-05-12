package cmd

import (
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "foaas",
		Short: "FOAAS service",
	}
	cmd.AddCommand(
		ServeCmd(),
	)
	return cmd
}
