package main

import (
	"os"
	"github.com/codegangsta/cli"
	area "github.com/davidjohansson/ecmd/area"
	article "github.com/davidjohansson/ecmd/article"
	solr "github.com/davidjohansson/ecmd/solr"
	"bufio"
	"strings"
	"fmt"
)

func main() {
	app := cli.NewApp()
	app.Name = "ecmd"
	app.Usage = "cli for common Escenic tasks"

	app.Commands = []cli.Command{
		{
			Name:      "search",
			Aliases:     []string{"s"},
			Usage:     "search solr for a given content type",
			Action: func(c *cli.Context) {
				solr.Search(c.Args().First())
			},
		},
		{
			Name:      "article",
			Aliases:   []string{"a"},
			Usage:     "inspects an article",
			Flags:     []cli.Flag{
				cli.StringFlag{
					Name: "fields, f",
					Usage: "Comma separated list of fields to display. Leave empty for all fields.",
				},
				cli.StringFlag{
					Name: "relation, r",
					Usage: "relation to display",
				},
			},
			Action: func(c *cli.Context) {

				ids := c.Args()
				fmt.Println(c.IsSet("fields"))
				if len(ids) == 0 {
					fmt.Println("Waiting for args...")
					piped := make([]string, 0)
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						piped = append(piped, scanner.Text())
					}
					ids = piped
				}

				article.Inspect(c.String("fields"), ids)
			},
		},
		{
			Name:      "area",
			Usage:     "inspects contents of an area",
			Action: func(c *cli.Context) {
				ids := c.Args()
				if len(ids) == 0 {
					reader := bufio.NewReader(os.Stdin)
					reader.ReadLine()
					text, _ := reader.ReadString('\n')
					ids = strings.Split(text, " ")
				}
				area.ListArea(c.Args().Get(0), c.Args().Get(1))
			},
		},
	}

	app.Run(os.Args)
}
