package db

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

func UpdateAnimeImage_RemoveIsDefault(tx *sqlx.Tx, anime_id string) {
	var query string = "UPDATE Anime_Images SET IsDefault = 0 WHERE IsDefault = 1 AND Anime_ID = ?"
	tx.MustExec(query, anime_id)
}

func UpdateAnimeImage_IsDefault(tx *sqlx.Tx, anime_id string, image_id string, is_default bool) {
	var queries []string = []string{
		` UPDATE Anime_Images
			SET IsDefault = 1
			WHERE Anime_ID = ? AND
				  Image_ID = ? `,

		` UPDATE Anime_Images
			SET IsDefault = 1
			WHERE Anime_ID = ? AND
				  Image_ID = (
						SELECT i.Id
						FROM Images i
						WHERE i.Url = ?
					) `,
	}

	var isUrl bool = strings.HasPrefix(image_id, "http")
	var queryIndex int

	if isUrl {
		queryIndex = 1
	} else {
		queryIndex = 0
	}

	if len(anime_id) > 0 && len(image_id) > 0 {
		tx.MustExec(
			queries[queryIndex],
			anime_id,
			image_id,
		)
	}
}
