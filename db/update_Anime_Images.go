package db

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

func UpdateAnimeImage_IsDefault(tx *sqlx.Tx, anime_id string, image_id string, is_default bool) {
	var queries []string = []string{
		`UPDATE Anime SET Image_Id = ? WHERE Id = ?`,
		`UPDATE Anime SET Image_Id = (SELECT i.Id FROM Images i WHERE i.Url = ?) WHERE Id = ?`,
	}

	var isUrl bool = strings.HasPrefix(image_id, "http")
	var queryIndex int

	if isUrl {
		queryIndex = 1
	} else {
		queryIndex = 0
	}

	if len(anime_id) > 0 && len(image_id) > 0 {
		res := tx.MustExec(
			queries[queryIndex],
			anime_id,
			image_id,
		)
		raff, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}
		fmt.Println("Update anime: ", fmt.Sprint(raff))

	}
}
