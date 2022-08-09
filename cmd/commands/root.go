package commands

import "github.com/spf13/cobra"

func RootCommand() *cobra.Command {
	command := &cobra.Command{
		Use: "learngo",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	command.AddCommand(RunServer())

	return command
}
