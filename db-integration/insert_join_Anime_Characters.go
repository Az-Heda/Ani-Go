package dbintegration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/schollz/progressbar/v3"
)

type anime_Character struct {
	Anime_ID     string
	Character_ID string
}

func insertJoinAnimeCharacters(db *sqlx.DB) {
	content, err := os.ReadFile("./db-integration/data/Anime_Characters.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []anime_Character
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	bar := progressbar.Default(int64(len(data)), "Joining Anime and Characters")
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec("INSERT INTO Anime_Characters (Anime_ID, Character_ID) VALUES (?, ?)",
			item.Anime_ID,
			item.Character_ID,
		)
		bar.Add(1)
	}
	if err = tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
