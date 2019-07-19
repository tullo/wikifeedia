// Command wikifeedia does something...
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/awoods187/wikifeedia/db"
	"github.com/awoods187/wikifeedia/wikipedia"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "wikifeedia"
	app.Usage = "runs one of the main actions"
	app.Action = func(c *cli.Context) error {
		println("run a subcommand")
		return nil
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "pgurl",
			Value: "pgurl://root@localhost:26257?sslmode=disable",
		},
	}
	app.Commands = []cli.Command{
		{
			Name: "setup",
			Action: func(c *cli.Context) error {
				pgurl := c.GlobalString("pgurl")
				fmt.Println("Setting up database at", pgurl)
				_, err := db.New(pgurl)
				return err
			},
		},
		{
			Name:        "fetch-top-articles",
			Description: "debug command to exercise the wikipedia client functionality.",
			Action: func(c *cli.Context) error {
				ctx := context.Background()
				wiki := wikipedia.New()
				top, err := wiki.FetchTopArticles(ctx)
				if err != nil {
					return err
				}
				n := c.Int("num-articles")
				for i := 0; i < len(top.Articles) && i < n; i++ {
					article, err := wiki.GetArticle(ctx, top.Articles[i].Article)
					if err != nil {
						return err
					}
					if i > 0 {
						fmt.Println()
					}
					fmt.Printf("%d. %s (%d)\n\n%s\n", i+1, article.Summary.Titles.Normalized, top.Articles[i].Views, article.Summary.Extract)
				}
				return nil
			},
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "num-articles,n",
					Value: 10,
					Usage: "number of articles to fetch",
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run command: %v\n", err)
		os.Exit(1)
	}
}
