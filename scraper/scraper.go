package scraper

import (
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

func FormatUrl(url string) string {
	url = strings.Split(url, "?")[0]
	if url[len(url)-1:] == "/" {
		url = url[0 : len(url)-1]
	}
	return url
}

func Init(url string) {
	url = FormatUrl(url)

	var urls map[string]string = map[string]string{
		"Anime":      url,
		"Pics":       url + "/pics",
		"Episodes":   url + "/episode",
		"Characters": url + "/characters",
		"Pictures":   url + "/pics",
	}

	conn, err := sqlx.Connect("sqlite", "./db/db.sqlite3")
	if err != nil {
		panic(err)
	}

	var tx *sqlx.Tx = conn.MustBegin()

	parseAnimePage(urls["Anime"], conn, tx)
	if err := tx.Commit(); err != nil {
		log.Fatalln(err)
	}

	tx = conn.MustBegin()
	parseAnimePics(urls["Anime"], urls["Pics"], conn, tx)
	parseAnimeEpisodeList(urls["Anime"], urls["Episodes"], conn, tx)
	parseAnimeCharacterList(urls["Anime"], urls["Characters"], conn, tx)

	if err := tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
