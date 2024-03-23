package dbintegration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/schollz/progressbar/v3"
)

type character_Description struct {
	Character_ID   string
	Description_ID string
}

func insertJoinCharacterDescriptions(db *sqlx.DB) {
	content, err := os.ReadFile("./db-integration/data/Character_Descriptions.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []character_Description
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	bar := progressbar.Default(int64(len(data)), "Joining Character and Descriptions")
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec("INSERT INTO Character_Descriptions (Character_ID, Description_ID) VALUES (?, ?)",
			item.Character_ID,
			item.Description_ID,
		)
		bar.Add(1)
	}
	if err = tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
