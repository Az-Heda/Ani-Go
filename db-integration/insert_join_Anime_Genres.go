package dbintegration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/schollz/progressbar/v3"
)

type anime_Genre struct {
	Anime_ID string
	Genre_ID string
}

func insertJoinAnimeGenres(db *sqlx.DB) {
	content, err := os.ReadFile("./db-integration/data/Anime_Genres.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []anime_Genre
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	bar := progressbar.Default(int64(len(data)), "Joining Anime and Genres")
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec("INSERT INTO Anime_Genres (Anime_ID, Genre_ID) VALUES (?, ?)",
			item.Anime_ID,
			item.Genre_ID,
		)
		bar.Add(1)
	}
	if err = tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
