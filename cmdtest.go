package main

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/davidjohansson/ecmdsolrsearch"
	"github.com/davidjohansson/ecmdinspect"
	"bufio"
	"fmt"
	"strings"
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
				ecmdsolrsearch.Search(c.Args().First())
			},
		},
		{
			Name:      "article",
			Aliases:   []string{"a"},
			Usage:     "inspects an article",
			Flags:	 []cli.Flag {
              cli.StringFlag{
                Name: "fields, f",
                Usage: "comma separated list of fields to display",
              },
    		cli.StringFlag{
			  Name: "relation, r",
			  Usage: "relation to display",
			},
            },

			Action: func(c *cli.Context) {

				ids := c.Args()
			
				if len(ids) == 0 {
					reader := bufio.NewReader(os.Stdin)
		    		text, _ := reader.ReadString('\n')
					ids = strings.Split(text, " ")

					}
					ecmdinspect.Inspect(c.String("fields"), ids)
					fmt.Println(ids)
			},
		},
	}

	app.Run(os.Args)
}
