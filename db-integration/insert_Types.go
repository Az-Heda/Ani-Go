package dbintegration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/schollz/progressbar/v3"
)

type types struct {
	Id   string
	Name string
}

func insertType(db *sqlx.DB) {
	content, err := os.ReadFile("./db-integration/data/Types.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []types
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	bar := progressbar.Default(int64(len(data)), "Insert Types")
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec("INSERT INTO Types (Id, Name) VALUES (?, ?)",
			item.Id,
			item.Name,
		)
		bar.Add(1)
	}
	if err = tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
