package cliargs

import (
	"fmt"
	"os"
	"strings"
	"time"

	dbintegration "AniGo/db-integration"
	scraper "AniGo/scraper"

	"github.com/jmoiron/sqlx"
)

func showHelp(ec int, keys []string) {
	fmt.Println("CLI Args list:")
	for _, key := range keys {
		fmt.Printf("%s\n", key)
	}
	os.Exit(ec)
}

func initializeDB(ec int) {
	dbintegration.Init()
	os.Exit(ec)
}

func scrape(url string) {
	scraper.Init(url)
}

func importFile(ec int, _ []string) {
	content, err := os.ReadFile("import.txt")
	if err != nil {
		panic(err)
	}

	conn, err := sqlx.Connect("sqlite", "./db/db-empty.sqlite3")

	if err != nil {
		panic(err)
	}

	conn.MustExec(`
		DELETE FROM Anime
		WHERE Id = (
			SELECT Id
			FROM Anime
			ORDER BY ROWID DESC
			LIMIT 1
		)`)

	// dbInitializer, err := os.ReadFile("./db/db-setup-v2.sql")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// var queryInitializer string = string(dbInitializer)
	// conn.MustExec(queryInitializer)

	var rows []string = strings.Split(string(content), "\n")

	for i, row := range rows {
		var url string = scraper.FormatUrl(strings.TrimSpace(row))
		fmt.Println("[" + fmt.Sprint(i+1) + "] - " + url)
		if len(scraper.ExistAnimeID(conn, url)) == 0 {
			scraper.Init(url)
			if i+1 < len(rows) {
				time.Sleep(90 * time.Second)
			}
		}
	}
	os.Exit(ec)
}
