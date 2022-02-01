package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Yaska1706/weebqoute/api"
	"github.com/urfave/cli"
)

func Setup() {
	app := cli.NewApp()
	app.Name = "WeebQuote"
	app.Usage = "Get random quotes from your favourite anime"
	app.Version = "v1.0.0"

	// we create our commands
	app.Commands = []cli.Command{
		{
			Name:  "random",
			Usage: "It provides a random quote from a random anime",
			Action: func(c *cli.Context) error {
				quote := api.GetRandom()
				fmt.Printf("%q(%q,%q)", quote.Qoute, quote.Character, quote.Anime)
				return nil
			},
		},
		{
			Name:  "random10",
			Usage: "It provides a random quote from a random anime",
			Action: func(c *cli.Context) error {
				quotes := api.GetTenRandom()
				for _, quote := range quotes {
					fmt.Printf("Anime: %q\n", quote.Anime)
					fmt.Printf("Character: %q\n", quote.Character)
					fmt.Printf("Quote: %q\n", quote.Qoute)
					fmt.Println("/-------------------------/")
				}
				return nil
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
