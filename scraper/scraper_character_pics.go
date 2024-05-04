package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/jmoiron/sqlx"
)

func parseAnimeCharacterPics(url string, characterID string, conn *sqlx.DB) []string {
	var images []string

	collector := colly.NewCollector()

	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println(url)
		panic(e)
	})

	collector.OnHTML("#content > table > tbody > tr > td:nth-child(2) > table > tbody", func(table *colly.HTMLElement) {
		table.ForEach("img[data-src]", func(idx int, img *colly.HTMLElement) {
			images = append(images, img.Attr("data-src"))
		})
	})

	collector.Visit(url)

	return images
}
