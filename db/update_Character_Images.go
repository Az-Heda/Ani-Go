package db

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

func UpdateCharacterImage_RemoveIsDefault(tx *sqlx.Tx, character_id string) {
	var query string = "UPDATE Character_Images SET IsDefault = 0 WHERE IsDefault = 1 AND Character_ID = ?"
	tx.MustExec(query, character_id)
}

func UpdateCharacterImage_IsDefault(tx *sqlx.Tx, character_id string, image_id string, is_default bool) {
	var queries []string = []string{
		`UPDATE Character_Images
			SET IsDefault = 1
			WHERE Character_ID = ? AND
				  Image_ID = ? `,

		`UPDATE Character_Images
			SET IsDefault = 1
			WHERE Character_ID = ? AND
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
	if len(character_id) > 0 && len(image_id) > 0 {
		tx.MustExec(
			queries[queryIndex],
			character_id,
			image_id,
		)
	}
}
