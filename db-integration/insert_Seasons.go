package dbintegration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/schollz/progressbar/v3"
)

type season struct {
	Id     string
	Season string
}

func insertSeasons(db *sqlx.DB) {
	content, err := os.ReadFile("./db-integration/data/Seasons.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []season
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	bar := progressbar.Default(int64(len(data)), "Insert Seasons")
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec("INSERT INTO Seasons (Id, Season) VALUES (?, ?)",
			item.Id,
			item.Season,
		)
		bar.Add(1)
	}
	if err = tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
