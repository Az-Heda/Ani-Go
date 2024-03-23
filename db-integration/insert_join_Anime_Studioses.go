package dbintegration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/schollz/progressbar/v3"
)

type anime_Studio struct {
	Anime_ID  string
	Studio_ID string
}

func insertJoinAnimeStudios(db *sqlx.DB) {
	content, err := os.ReadFile("./db-integration/data/Anime_Studioses.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []anime_Studio
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	bar := progressbar.Default(int64(len(data)), "Joining Anime and Studios")
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec("INSERT INTO Anime_Studioses (Anime_ID, Studio_ID) VALUES (?, ?)",
			item.Anime_ID,
			item.Studio_ID,
		)
		bar.Add(1)
	}
	if err = tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
