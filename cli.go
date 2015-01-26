package main

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func RunCli() {
	app := cli.NewApp()
	app.Name = "Anime Tracker"
	app.Usage = "Keep track of the current viewed anime episode"
	app.Commands = []cli.Command{
		{
			Name:      "add",
			ShortName: "a",
			Usage:     "Add a new anime to the database",
			Action: func(c *cli.Context) {
				if err := addAnime(c.Args().First(), c.Args()[1]); err != nil {
					fmt.Printf("Can't add new anime because: %v\n", err)
				}
			},
		},
		{
			Name:      "remove",
			ShortName: "r",
			Usage:     "Remove an anime",
			Action: func(c *cli.Context) {
				if err := Remove(c.Args().First()); err != nil {
					fmt.Printf("Can't Remove because: %v", err)
				}
			},
		},
		{
			Name:      "show",
			ShortName: "s",
			Usage:     "Show all anime in the database",
			Action: func(c *cli.Context) {
				for _, anime := range animeList {
					fmt.Printf("Title: %s\nEpisode: %s\n\n", anime.Title, anime.Episode)
				}
			},
		},
		{
			Name:      "increment",
			ShortName: "i",
			Usage:     "Increment the last viewed episode",
			Action: func(c *cli.Context) {
				if err := Increment(c.Args().First()); err != nil {
					fmt.Printf("Cant increment because: %v", err)
				}
			},
		},
	}

	app.RunAndExitOnError()
}
