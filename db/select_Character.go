package db

import (
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"
)

type DB_Character_SingleImage struct {
	Id          string
	Name        string
	Image       sql.NullString
	Description sql.NullString
}
type DB_Character_MultipleImages struct {
	Id          string
	Name        string
	Image       []string
	Description sql.NullString
}

func SelectCharacterFromID(id string) (DB_Character_MultipleImages, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Character_MultipleImages{}, err
	}
	var data DB_Character_MultipleImages = DB_Character_MultipleImages{}
	rows, err := conn.Query(`
			SELECT c.Id, c.Name,
				  (
					SELECT GROUP_CONCAT(i.Url, '://:')
					FROM Images i
					LEFT JOIN Character_Images ci ON ci.Image_ID = i.Id
					WHERE ci.Character_ID = c.Id
				  ) as Image,
				  (
					SELECT GROUP_CONCAT(d.Description, '\n')
					FROM Descriptions d
					WHERE d.Character_ID = c.Id
				  ) as Description
			FROM Character c
			LEFT JOIN Anime_Characters ac ON ac.Character_ID = c.Id
			WHERE ac.Character_ID = ?;
		`, id)
	if err != nil {
		return DB_Character_MultipleImages{}, err
	}
	var tx *sqlx.Tx = conn.MustBegin()
	for rows.Next() {
		var d DB_Character_SingleImage
		if err = rows.Scan(&d.Id, &d.Name, &d.Image, &d.Description); err != nil {
			return DB_Character_MultipleImages{}, err
		}
		var instance DB_Character_MultipleImages = DB_Character_MultipleImages{
			Id:          d.Id,
			Name:        d.Name,
			Image:       strings.Split(d.Image.String, "://:"),
			Description: d.Description,
		}

		if len(d.Image.String) == 0 {
			alternativeImages, _ := selectCharacterAlternativeImages(tx, instance.Id)
			instance.Image = alternativeImages
		}
		data = instance
		// data = append(data, instance)
	}
	tx.Commit()
	return data, nil
}

func SelectCharactersByIdWithDefaultImage(id string) ([]DB_Character_MultipleImages, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Character_MultipleImages = []DB_Character_MultipleImages{}
	rows, err := conn.Query(`
			SELECT c.Id, c.Name,
				  (
					SELECT GROUP_CONCAT(i.Url, '://:')
					FROM Images i
					LEFT JOIN Character_Images ci ON ci.Image_ID = i.Id
					WHERE ci.Character_ID = c.Id AND ci.IsDefault = 1
				  ) as Image,
				  (
					SELECT GROUP_CONCAT(d.Description, '\n')
					FROM Descriptions d
					WHERE d.Character_ID = c.Id
				  ) as Description
			FROM Character c
			LEFT JOIN Anime_Characters ac ON ac.Character_ID = c.Id
			WHERE ac.Anime_ID = ?;
		`, id)
	if err != nil {
		return nil, err
	}
	var tx *sqlx.Tx = conn.MustBegin()
	for rows.Next() {
		var d DB_Character_SingleImage
		if err = rows.Scan(&d.Id, &d.Name, &d.Image, &d.Description); err != nil {
			return nil, err
		}
		var instance DB_Character_MultipleImages = DB_Character_MultipleImages{
			Id:          d.Id,
			Name:        d.Name,
			Image:       strings.Split(d.Image.String, "://:"),
			Description: d.Description,
		}
		if len(d.Image.String) == 0 {
			alternativeImages, _ := selectCharacterAlternativeImages(tx, instance.Id)
			instance.Image = alternativeImages
		}
		data = append(data, instance)
	}
	tx.Commit()
	return data, nil
}

func SelectCharactersByIdWithoutDefaultImage(id string) ([]DB_Character_MultipleImages, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Character_MultipleImages
	rows, err := conn.Query(`
			SELECT c.Id, c.Name,
				  (
					SELECT GROUP_CONCAT(i.Url, '://:')
					FROM Images i
					LEFT JOIN Character_Images ci ON ci.Image_ID = i.Id
					WHERE ci.Character_ID = c.Id AND ci.IsDefault = 0
				  ) as Image,
				  (
					SELECT GROUP_CONCAT(d.Description, '\n')
					FROM Descriptions d
					WHERE d.Character_ID = c.Id
				  ) as Description
			FROM Character c
			LEFT JOIN Anime_Characters ac ON ac.Character_ID = c.Id
			WHERE ac.Anime_ID = ?;
		`, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var d DB_Character_SingleImage
		if err = rows.Scan(&d.Id, &d.Name, &d.Image, &d.Description); err != nil {
			return nil, err
		}
		data = append(data, DB_Character_MultipleImages{
			Id:          d.Id,
			Name:        d.Name,
			Image:       strings.Split(d.Image.String, "://:"),
			Description: d.Description,
		})
	}
	return data, nil
}

func selectCharacterAlternativeImages(tx *sqlx.Tx, id string) ([]string, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []string = []string{}
	err = conn.Select(&data, `
		SELECT i.Url
		FROM Images i
		LEFT JOIN Character_Images ci ON
			ci.Image_ID = i.Id AND
			ci.Character_ID = ?
		WHERE ci.IsDefault = 0
		LIMIT 1
	`, id)
	if err != nil {
		return []string{}, err
	}
	if len(data) == 1 {
		if len(data[0]) > 0 {
			tx.MustExec(`
				UPDATE Character_Images
					SET IsDefault = 1
					WHERE Character_ID = ? AND
						  Image_ID = (
								SELECT i.Id
								FROM Images i
								WHERE i.Url = ?
							)
				`,
				id,
				data[0],
			)
		}
		return data, nil
	}
	return []string{}, err
}

// func SelectAllCharacters() ([]DB_Character_SingleImage, error) {
// 	conn, err := GetConnection()
// 	if err != nil {
// 		return nil, err
// 	}
// 	var data []DB_Character_SingleImage
// 	rows, err := conn.Query("SELECT c.Id, c.Name, (SELECT i.Url FROM Images i LEFT JOIN Character_Images ci ON ci.Image_ID = i.Id WHERE ci.Character_ID = c.Id LIMIT 1) as Image FROM Character c ORDER BY c.Name")
// 	if err != nil {
// 		return data, err
// 	}
// 	for rows.Next() {
// 		var d DB_Character_SingleImage
// 		if err = rows.Scan(&d.Id, &d.Name, &d.Image); err != nil {
// 			return data, err
// 		}
// 		data = append(data, d)
// 	}
// 	return data, nil
// }

// func SelectCharacterFromId(id string) (DB_Character_SingleImage, error) {
// 	conn, err := GetConnection()
// 	if err != nil {
// 		return DB_Character_SingleImage{}, err
// 	}
// 	var data DB_Character_SingleImage
// 	rows, err := conn.Query("SELECT Id, Name FROM Character WHERE Id = ?", id)
// 	if err != nil {
// 		return DB_Character_SingleImage{}, err
// 	}
// 	for rows.Next() {
// 		if err = rows.Scan(&data.Id, &data.Name); err != nil {
// 			return DB_Character_SingleImage{}, err
// 		}
// 	}
// 	return data, nil
// }

// func SelectCharacterFromAnimeId(id string) ([]DB_Character_MultipleImages, error) {
// 	conn, err := GetConnection()
// 	if err != nil {
// 		return nil, err
// 	}
// 	var data []DB_Character_MultipleImages
// 	rows, err := conn.Query(`
// 		SELECT c.Id, c.Name,
// 			  (
// 				SELECT GROUP_CONCAT(i.Url, '://:')
// 				FROM Images i
// 				LEFT JOIN Character_Images ci ON ci.Image_ID = i.Id
// 				WHERE ci.Character_ID = c.Id
// 			  ) as Image,
// 			  (
// 				SELECT GROUP_CONCAT(d.Description, '\n')
// 				FROM Descriptions d
// 				WHERE d.Character_ID = c.Id
// 			  ) as Description
// 		FROM Character c
// 		LEFT JOIN Anime_Characters ac ON ac.Character_ID = c.Id
// 		WHERE ac.Anime_ID = ?;
// 	`, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for rows.Next() {
// 		var d DB_Character_SingleImage
// 		if err = rows.Scan(&d.Id, &d.Name, &d.Image, &d.Description); err != nil {
// 			return nil, err
// 		}
// 		data = append(data, DB_Character_MultipleImages{
// 			d.Id,
// 			d.Name,
// 			strings.Split(d.Image.String, "://:"),
// 			d.Description,
// 		})
// 	}
// 	return data, nil
// }
