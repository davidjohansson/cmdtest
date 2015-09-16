package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "dsas"
	app.Usage = "cli for common Escenic tasks"

	app.Commands = []cli.Command{
		{
			Name:      "search",
			Aliases:     []string{"a"},
			Usage:     "search solr for a given content type",
			Action: func(c *cli.Context) {
				println("searching solr for: ", c.Args().First())
			},
		},
		{
			Name:      "list",
			Aliases:     []string{"c"},
			Usage:     "list values for a given field",
			Action: func(c *cli.Context) {
				println("list for field: ", c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}
