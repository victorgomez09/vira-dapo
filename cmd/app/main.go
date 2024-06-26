package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	localCli "github.com/victorgomez09/vira-dapo/cmd/cli"
)

func main() {
	app := &cli.App{
		Name: "vira-dapo",

		Commands: []*cli.Command{
			localCli.AppCommands(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
