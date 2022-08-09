package commands

import (
	"learn-go/server"
	"log"

	"github.com/spf13/cobra"
)

func RunServer() *cobra.Command {
	command := cobra.Command{
		Use: "Server",
		Run: func(cmd *cobra.Command, args []string) {
			server := server.NewServer()

			log.Fatal(server.Listen(":3000"))
		},
	}

	return &command
}
