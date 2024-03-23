package dbintegration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/schollz/progressbar/v3"
)

type anime_Description struct {
	Anime_ID       string
	Description_ID string
}

func insertJoinAnimeDescriptions(db *sqlx.DB) {
	content, err := os.ReadFile("./db-integration/data/Anime_Descriptions.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []anime_Description
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	bar := progressbar.Default(int64(len(data)), "Joining Anime and Descriptions")
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec("INSERT INTO Anime_Descriptions (Anime_ID, Description_ID) VALUES (?, ?)",
			item.Anime_ID,
			item.Description_ID,
		)
		bar.Add(1)
	}
	if err = tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
