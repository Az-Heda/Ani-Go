package dbintegration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/schollz/progressbar/v3"
)

type status struct {
	Id   int
	Name string
}

func insertStatuses(db *sqlx.DB) {
	content, err := os.ReadFile("./db-integration/data/Statuses.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []status
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	bar := progressbar.Default(int64(len(data)), "Insert Statuses")
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec("INSERT INTO Statuses (Id, Name) VALUES (?, ?)",
			item.Id,
			item.Name,
		)
		bar.Add(1)
	}
	if err = tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
