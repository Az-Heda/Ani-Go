package dbintegration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/schollz/progressbar/v3"
)

type studio struct {
	Id   string
	Name string
}

func insertStudios(db *sqlx.DB) {
	content, err := os.ReadFile("./db-integration/data/Studios.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []studio
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	bar := progressbar.Default(int64(len(data)), "Insert Studios")
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec("INSERT INTO Studios (Id, Name) VALUES (?, ?)",
			item.Id,
			item.Name,
		)
		bar.Add(1)
	}
	if err = tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
