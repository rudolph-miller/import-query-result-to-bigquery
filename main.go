package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"

	var lang string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang",
			Value:       "english",
			Usage:       "language for the greeting",
			Destination: &lang,
		},
	}

	app.Action = func(c *cli.Context) {
		println("Hello friend!")
		println(lang)
	}

	app.Run(os.Args)
}
