package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/jmoiron/sqlx"
)

func parseAnimePics(baseUrl string, url string, conn *sqlx.DB, tx *sqlx.Tx) {
	var animeID string = GetAnimeID(conn, baseUrl)
	var images []ScraperImage
	var join_anime_images []ScraperAnimeImage

	collector := colly.NewCollector()

	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println(url)
		panic(e)
	})

	collector.OnHTML("#content > table > tbody > tr > td:nth-child(2) > div.rightside.js-scrollfix-bottom-rel > table", func(table *colly.HTMLElement) {
		table.ForEach("img[data-src]", func(idx int, img *colly.HTMLElement) {
			var href string = img.Attr("data-src")
			var image ScraperImage = ScraperImage{GetImageID(conn, href), href}
			images = append(images, image)
			join_anime_images = append(join_anime_images, ScraperAnimeImage{animeID, image.Id})
		})
	})

	collector.Visit(url)

	insertImage(conn, tx, images)
	insertAnimeImage(conn, tx, join_anime_images)
}
