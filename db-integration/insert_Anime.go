package dbintegration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/schollz/progressbar/v3"
)

type anime struct {
	Id               string
	Title            string
	AlternativeTitle string
	Aired            int64
	Duration         int64
	Url              string
	Season_ID        string
	Type_ID          string
	CurrentStatus    int
}

func insertAnime(db *sqlx.DB) {
	content, err := os.ReadFile("./db-integration/data/Anime.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []anime
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	bar := progressbar.Default(int64(len(data)), "Insert Anime")
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec("INSERT INTO Anime (Id, Title, AlternativeTitle, Aired, Duration, Url, Season_ID, Type_ID, CurrentStatus) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
			item.Id,
			nullString(item.Title),
			nullString(item.AlternativeTitle),
			item.Aired,
			item.Duration,
			nullString(item.Url),
			nullString(item.Season_ID),
			nullString(item.Type_ID),
			item.CurrentStatus,
		)
		bar.Add(1)
	}
	if err = tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
