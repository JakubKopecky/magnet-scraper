package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "magnet-scraper",
		Usage: "magnet-scraper <url>",
		Action: func(ctx *cli.Context) error {
			fmt.Println(extract_magnet(ctx.Args().Get(0)))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func extract_magnet(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	link, ok := doc.Find("[class*='kaGiantButton']").First().Attr("href")
	if ok {
		return link
	}

	return "not found"
}
