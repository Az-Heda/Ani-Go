package dbintegration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/schollz/progressbar/v3"
)

type genre struct {
	Id   string
	Name string
}

func insertGenres(db *sqlx.DB) {
	content, err := os.ReadFile("./db-integration/data/Genres.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []genre
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	bar := progressbar.Default(int64(len(data)), "Insert Genres")
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec("INSERT INTO Genres (Id, Name) VALUES (?, ?)",
			item.Id,
			item.Name,
		)
		bar.Add(1)
	}
	if err = tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
