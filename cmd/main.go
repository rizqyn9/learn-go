package main

import (
	"learn-go/cmd/commands"

	"github.com/rs/zerolog/log"
)

func main() {
	err := commands.RootCommand().Execute()

	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
