package dbintegration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/schollz/progressbar/v3"
)

type image struct {
	Id  string
	Url string
}

func insertImages(db *sqlx.DB) {
	content, err := os.ReadFile("./db-integration/data/Images.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []image
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	bar := progressbar.Default(int64(len(data)), "Insert Images")
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec("INSERT INTO Images (Id, Url) VALUES (?, ?)",
			item.Id,
			item.Url,
		)
		bar.Add(1)
	}
	if err = tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
