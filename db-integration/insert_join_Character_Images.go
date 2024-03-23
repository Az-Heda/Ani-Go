package dbintegration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/schollz/progressbar/v3"
)

type character_Image struct {
	Character_ID string
	Image_ID     string
}

func insertJoinCharacterImages(db *sqlx.DB) {
	content, err := os.ReadFile("./db-integration/data/Character_Images.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []character_Image
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	bar := progressbar.Default(int64(len(data)), "Joining Character and Images")
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec("INSERT INTO Character_Images (Character_ID, Image_ID) VALUES (?, ?)",
			item.Character_ID,
			item.Image_ID,
		)
		bar.Add(1)
	}
	if err = tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
