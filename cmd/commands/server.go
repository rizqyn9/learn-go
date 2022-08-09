package commands

import "github.com/spf13/cobra"

func RunServer() *cobra.Command {
	command := cobra.Command{
		Use: "Server",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	return &command
}
