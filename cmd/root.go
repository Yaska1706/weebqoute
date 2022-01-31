package cmd

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Setup() {
	app := cli.NewApp()
	app.Name = "WeebQuote"
	app.Usage = "It provides quotes of your favourite anime randomly"
	app.Version = "1.0.0"
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
