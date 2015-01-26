package main

import "github.com/codegangsta/cli"

func main() {
	app := cli.NewApp()
	app.Name = "Anime Tracker"
	app.Usage = "Keep track of the current viewed anime episode"
	app.Commands = []cli.Command{
		{
			Name:      "add",
			ShortName: "a",
			Usage:     "Add a new anime to the database",
			Action: func(c *cli.Context) {
			},
		},
	}
}
