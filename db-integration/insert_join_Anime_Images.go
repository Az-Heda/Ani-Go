package dbintegration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/schollz/progressbar/v3"
)

type anime_Image struct {
	Anime_ID string
	Image_ID string
}

func insertJoinAnimeImages(db *sqlx.DB) {
	content, err := os.ReadFile("./db-integration/data/Anime_Images.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []anime_Image
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	bar := progressbar.Default(int64(len(data)), "Joining Anime and Images")
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec("INSERT INTO Anime_Images (Anime_ID, Image_ID) VALUES (?, ?)",
			item.Anime_ID,
			item.Image_ID,
		)
		bar.Add(1)
	}
	if err = tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
