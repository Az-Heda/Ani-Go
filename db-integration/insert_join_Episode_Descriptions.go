package dbintegration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/schollz/progressbar/v3"
)

type episode_Description struct {
	Episode_ID     string
	Description_ID string
}

func insertJoinEpisodeDescriptions(db *sqlx.DB) {
	content, err := os.ReadFile("./db-integration/data/Episode_Descriptions.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []episode_Description
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}
	bar := progressbar.Default(int64(len(data)), "Joining Episode and Descriptions")
	tx := db.MustBegin()
	for _, item := range data {
		tx.MustExec("INSERT INTO Episode_Descriptions (Episode_ID, Description_ID) VALUES (?, ?)",
			item.Episode_ID,
			item.Description_ID,
		)
		bar.Add(1)
	}
	if err = tx.Commit(); err != nil {
		log.Fatalln(err)
	}
}
