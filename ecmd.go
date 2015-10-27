package main

import (
	"os"
	"github.com/codegangsta/cli"
	area "github.com/davidjohansson/ecmd/area"
	article "github.com/davidjohansson/ecmd/article"
	solr "github.com/davidjohansson/ecmd/solr"
	"bufio"
//	"strings"

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
					Usage: "Comma separated list of fields to display.",
				},
				cli.StringFlag{
					Name: "responsedata, c",
					Usage: "Comma separated list of response fields to display, e g 'contentType'",
				},
				cli.StringFlag{
					Name: "relations, r",
					Usage: "Name of relation to show, e g 'topContentRel'",
				},
				cli.StringFlag{
					Name: "meta, m",
					Usage: "Comma separated list of metadata to show, e g 'createdDate'",
				},
			},
			Action: func(c *cli.Context) {

				ids := c.Args()



				if len(ids) == 0 {
					piped := make([]string, 0)
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						piped = append(piped, scanner.Text())
					}
					ids = piped
				}



				article.Inspect(c.String("fields"), c.String("responsedata"), c.String("relations"), c.String("meta"), ids)
			},
		},
		{
			Name:      "area",
			Usage:     "inspects contents of an area",
			Flags:     []cli.Flag{
				cli.BoolFlag{
					Name: "list, l",
					Usage: "List all areas",
				},
			},

			Action: func(c *cli.Context) {
				listAreas := c.Bool("list")
				if !listAreas {
					reqArea := c.Args().Get(0)
					section := c.Args().Get(1)
					area.ListArea(reqArea, section, listAreas)
				} else{
					section := c.Args().Get(0)
					area.ListArea("", section, listAreas)
				}
			},
		},
	}

	app.Run(os.Args)
}
