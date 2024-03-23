package dbintegration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/schollz/progressbar/v3"
)

type episode struct {
	Id       string
	Number   int
	Title    string
	Aired    int64
	Duration int64
	Anime_ID string
}

func insertEpisode(db *sqlx.DB) {
	content, err := os.ReadFile("./db-integration/data/Episode.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []episode
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	bar := progressbar.Default(int64(len(data)), "Insert Episodes")
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec("INSERT INTO Episode (Id, Number, Title, Aired, Duration, Anime_ID) VALUES (?, ?, ?, ?, ?, ?)",
			item.Id,
			item.Number,
			nullString(item.Title),
			item.Aired,
			item.Duration,
			item.Anime_ID,
		)
		bar.Add(1)
	}
	if err = tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
