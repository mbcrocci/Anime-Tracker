package main

import (
	"fmt"
	"log"

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
					log.Printf("Can't add new anime because: %v\n", err)
				}
			},
		},
		{
			Name:      "remove",
			ShortName: "r",
			Usage:     "Remove an anime",
			Action: func(c *cli.Context) {
				if err := Remove(c.Args().First()); err != nil {
					log.Printf("Can't Remove because: %v", err)
				}
			},
		},
		{
			Name:      "show",
			ShortName: "s",
			Usage:     "Show all anime in the database",
			Action: func(c *cli.Context) {
				if c.Args().First() != "" {
					anime, err := Search(c.Args().First())
					if err != nil {
						log.Println(err)
					}
					fmt.Printf("Title: %s\nEpisode: %d\n\n", anime.Title, anime.Episode)

				} else {
					for _, anime := range animeList {
						fmt.Printf("Title: %s\nEpisode: %d\n\n", anime.Title, anime.Episode)
					}
				}

			},
		},
		{
			Name:      "increment",
			ShortName: "i",
			Usage:     "Increment the last viewed episode",
			Action: func(c *cli.Context) {
				if err := Increment(c.Args().First()); err != nil {
					log.Printf("Cant increment because: %v", err)
				}
			},
		},
		{
			Name:      "web",
			ShortName: "w",
			Usage:     "Start a server that serves the anime database",
			Action: func(c *cli.Context) {
				if err := RunWeb(); err != nil {
					log.Println("Can't start server: %v", err)
				}
			},
		},
	}

	app.RunAndExitOnError()
}
