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
					fmt.Printf("%q(%q,%q)\n", quote.Qoute, quote.Character, quote.Anime)
					fmt.Println("/-------------------------/")
				}
				return nil
			},
		},
		{
			Name:  "animes",
			Usage: "Lists down available anime",

			Action: func(c *cli.Context) error {
				animes := api.GetAllAnime()
				var number = 1
				for _, anime := range animes {
					fmt.Printf("%v. %q\n", number, anime)
					number += 1
				}

				return nil
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
