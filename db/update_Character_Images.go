package db

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

func UpdateCharacterImage_IsDefault(tx *sqlx.Tx, character_id string, image_id string, is_default bool) {
	var queries []string = []string{
		`UPDATE Character SET Image_Id = ? WHERE Id = ?`,
		`UPDATE Character SET Image_Id = (SELECT i.Id FROM Images i WHERE i.Url = ?) WHERE Id = ?`,
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
