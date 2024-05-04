package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DB_Anime struct {
	Id               string
	Title            string
	AlternativeTitle sql.NullString
	Aired            int64
	Duration         int64
	Url              string
	CurrentStatus    int
	Season_ID        sql.NullString
	Type_ID          sql.NullString
	Image            sql.NullString
	Description      sql.NullString
}

func SelectAllAnime() ([]DB_Anime, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Anime
	rows, err := conn.Query(`
		SELECT a.Id, a.Title, a.AlternativeTitle, a.Aired, a.Duration, a.Url, a.CurrentStatus, a.Season_ID, a.Type_ID,
			  (
				SELECT i.Url
				FROM Images i
				LEFT JOIN Anime_Images ai ON ai.Image_ID = i.Id
				WHERE ai.Anime_ID = a.Id AND ai.IsDefault = 1
				LIMIT 1
			  ) as Image,
			  (
				SELECT GROUP_CONCAT(d.Description, '\n')
				FROM Descriptions d
				WHERE d.Anime_ID = a.Id
			) as Description
		FROM Anime a
		GROUP BY a.Id
		ORDER BY
		CASE WHEN AlternativeTitle IS NOT NULL
			THEN AlternativeTitle
			ELSE Title
		END
	`)
	if err != nil {
		return data, err
	}
	var tx *sqlx.Tx = conn.MustBegin()
	for rows.Next() {
		var d DB_Anime
		if err = rows.Scan(&d.Id, &d.Title, &d.AlternativeTitle, &d.Aired, &d.Duration, &d.Url, &d.CurrentStatus, &d.Season_ID, &d.Type_ID, &d.Image, &d.Description); err != nil {
			return data, err
		}

		if len(d.Image.String) == 0 {
			images, _ := selectAnimeAlternativeImages(tx, d.Id)
			if len(images) > 0 {
				d.Image = sql.NullString{
					String: images[0],
					Valid:  len(images[0]) > 0,
				}
			}
		}
		data = append(data, d)
	}
	tx.Commit()
	return data, nil
}

func SelectAnimeFromId(id string) (DB_Anime, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Anime{}, err
	}
	var data DB_Anime

	rows, err := conn.Query(`
		SELECT a.Id, a.Title, a.AlternativeTitle, a.Aired, a.Duration, a.Url, a.CurrentStatus, a.Season_ID, a.Type_ID,
			  (
				SELECT GROUP_CONCAT(i.Url, '://:')
				FROM Images i
				LEFT JOIN Anime_Images ai ON ai.Image_ID = i.Id
				WHERE ai.Anime_ID = a.Id
				LIMIT 1
			  ) as Image,
			  (
				SELECT GROUP_CONCAT(d.Description, '\n')
				FROM Descriptions d
				WHERE d.Anime_ID = a.Id
			  ) as Description
		FROM Anime a
		WHERE a.Id = ?
		GROUP BY a.Id
		ORDER BY
		CASE WHEN AlternativeTitle IS NOT NULL
			THEN AlternativeTitle
			ELSE Title
		END;
	`, id)
	if err != nil {
		return DB_Anime{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Title, &data.AlternativeTitle, &data.Aired, &data.Duration, &data.Url, &data.CurrentStatus, &data.Season_ID, &data.Type_ID, &data.Image, &data.Description); err != nil {
			return DB_Anime{}, err
		}
	}
	return data, nil
}

func SelectAnimeFromPartName(name string) ([]DB_Anime, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Anime
	var likeFilter string = "%" + name + "%"
	rows, err := conn.Query(`
		SELECT Id, Title, AlternativeTitle, Aired, Duration, Url, CurrentStatus, Season_ID, Type_ID
		FROM Anime
		WHERE Title LIKE ? OR
			  AlternativeTitle LIKE ?
	`, likeFilter, likeFilter)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var d DB_Anime
		if err = rows.Scan(&d.Id, &d.Title, &d.AlternativeTitle, &d.Aired, &d.Duration, &d.Url, &d.CurrentStatus, &d.Season_ID, &d.Type_ID); err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, nil
}

func selectAnimeAlternativeImages(tx *sqlx.Tx, id string) ([]string, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []string = []string{}
	err = conn.Select(&data, `
		SELECT i.Url
		FROM Images i
		LEFT JOIN Anime_Images ai ON
			ai.Image_ID = i.Id AND
			ai.Anime_ID = ?
		WHERE ai.IsDefault = 0
		LIMIT 1
	`, id)
	if err != nil {
		return []string{}, err
	}
	if len(data) == 1 {
		if len(data[0]) > 0 {
			tx.MustExec(`
				UPDATE Anime_Images
					SET IsDefault = 1
					WHERE Anime_ID = ? AND
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
