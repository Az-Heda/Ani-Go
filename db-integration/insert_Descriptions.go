package dbintegration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/schollz/progressbar/v3"
)

type description struct {
	Id          string
	Description string
}

func insertDescriptions(db *sqlx.DB) {
	content, err := os.ReadFile("./db-integration/data/Description.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []description
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	bar := progressbar.Default(int64(len(data)), "Insert Descriptions")
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec("INSERT INTO Descriptions (Id, Description) VALUES (?, ?)",
			item.Id,
			item.Description,
		)
		bar.Add(1)
	}
	if err = tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
