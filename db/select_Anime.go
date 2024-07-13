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
	Broadcast        int
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
		SELECT
			a.Id, a.Title, a.AlternativeTitle,
			a.Aired, a.Duration, a.Url,
			a.CurrentStatus, a.Season_ID,
			a.Type_ID, a.Broadcast, i.Url as Image,
			(
				SELECT GROUP_CONCAT(d.Description, '\n')
				FROM Descriptions d
				WHERE d.Anime_ID = a.Id
			) as Description
		FROM Anime a
		LEFT JOIN Images i ON a.Image_Id = I.Id
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
	for rows.Next() {
		var d DB_Anime
		if err = rows.Scan(&d.Id, &d.Title, &d.AlternativeTitle, &d.Aired, &d.Duration, &d.Url, &d.CurrentStatus, &d.Season_ID, &d.Type_ID, &d.Broadcast, &d.Image, &d.Description); err != nil {
			return data, err
		}

		if len(d.Image.String) == 0 {
			images, _ := SelectAnimeAlternativeImages(d.Id, true)
			if len(images) > 0 {
				d.Image = sql.NullString{
					String: images[0],
					Valid:  len(images[0]) > 0,
				}
			}
		}
		data = append(data, d)
	}
	return data, nil
}

func SelectAllAnimeFromGenreID(genre DB_Genre) ([]DB_Anime, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Anime
	rows, err := conn.Query(`
		SELECT
			a.Id, a.Title, a.AlternativeTitle,
			a.Aired, a.Duration, a.Url,
			a.CurrentStatus, a.Season_ID,
			a.Type_ID, a.Broadcast, i.Url as Image,
			(
				SELECT GROUP_CONCAT(d.Description, '\n')
				FROM Descriptions d
				WHERE d.Anime_ID = a.Id
			) as Description
		FROM Anime a
		LEFT JOIN Images i ON a.Image_Id = I.Id
		LEFT JOIN Anime_Genres ag on a.Id = ag.Anime_ID
		LEFT JOIN Genres g on g.Id = ag.Genre_ID
		WHERE g.Id = ?
		GROUP BY a.Id
		ORDER BY
		CASE WHEN AlternativeTitle IS NOT NULL
			THEN AlternativeTitle
			ELSE Title
		END
	`, genre.Id)
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var d DB_Anime
		if err = rows.Scan(&d.Id, &d.Title, &d.AlternativeTitle, &d.Aired, &d.Duration, &d.Url, &d.CurrentStatus, &d.Season_ID, &d.Type_ID, &d.Broadcast, &d.Image, &d.Description); err != nil {
			return data, err
		}

		if len(d.Image.String) == 0 {
			images, _ := SelectAnimeAlternativeImages(d.Id, true)
			if len(images) > 0 {
				d.Image = sql.NullString{
					String: images[0],
					Valid:  len(images[0]) > 0,
				}
			}
		}
		data = append(data, d)
	}
	return data, nil
}

func SelectAllAnimeFromThemeID(theme DB_Theme) ([]DB_Anime, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Anime
	rows, err := conn.Query(`
		SELECT
			a.Id, a.Title, a.AlternativeTitle,
			a.Aired, a.Duration, a.Url,
			a.CurrentStatus, a.Season_ID,
			a.Type_ID, a.Broadcast, i.Url as Image,
			(
				SELECT GROUP_CONCAT(d.Description, '\n')
				FROM Descriptions d
				WHERE d.Anime_ID = a.Id
			) as Description
		FROM Anime a
		LEFT JOIN Images i ON a.Image_Id = I.Id
		LEFT JOIN Anime_Themes at on a.Id = at.Anime_ID
		LEFT JOIN Themes t on t.Id = at.Theme_ID
		WHERE t.Id = ?
		GROUP BY a.Id
		ORDER BY
		CASE WHEN AlternativeTitle IS NOT NULL
			THEN AlternativeTitle
			ELSE Title
		END
	`, theme.Id)
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var d DB_Anime
		if err = rows.Scan(&d.Id, &d.Title, &d.AlternativeTitle, &d.Aired, &d.Duration, &d.Url, &d.CurrentStatus, &d.Season_ID, &d.Type_ID, &d.Broadcast, &d.Image, &d.Description); err != nil {
			return data, err
		}

		if len(d.Image.String) == 0 {
			images, _ := SelectAnimeAlternativeImages(d.Id, true)
			if len(images) > 0 {
				d.Image = sql.NullString{
					String: images[0],
					Valid:  len(images[0]) > 0,
				}
			}
		}
		data = append(data, d)
	}
	return data, nil
}

func SelectAnimeFromId(id string) (DB_Anime, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Anime{}, err
	}
	var data DB_Anime

	rows, err := conn.Query(`
		SELECT
			a.Id, a.Title, a.AlternativeTitle,
			a.Aired, a.Duration, a.Url, a.CurrentStatus,
			a.Season_ID, a.Type_ID, a.Broadcast, i.Url as Image,
			(
				SELECT GROUP_CONCAT(d.Description, '\n')
				FROM Descriptions d
				WHERE d.Anime_ID = a.Id
			) as Description
		FROM Anime a
		LEFT JOIN Images i ON i.Id = a.Image_Id
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
		if err = rows.Scan(&data.Id, &data.Title, &data.AlternativeTitle, &data.Aired, &data.Duration, &data.Url, &data.CurrentStatus, &data.Season_ID, &data.Type_ID, &data.Broadcast, &data.Image, &data.Description); err != nil {
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
		SELECT Id, Title, AlternativeTitle, Aired, Duration, Url, CurrentStatus, Season_ID, Type_ID, Broadcast
		FROM Anime
		WHERE Title LIKE ? OR
			  AlternativeTitle LIKE ?
	`, likeFilter, likeFilter)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var d DB_Anime
		if err = rows.Scan(&d.Id, &d.Title, &d.AlternativeTitle, &d.Aired, &d.Duration, &d.Url, &d.CurrentStatus, &d.Season_ID, &d.Type_ID, &d.Broadcast); err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, nil
}

func SelectAnimeAlternativeImages(id string, update bool) ([]string, error) {
	conn, err := GetConnection()
	var tx *sqlx.Tx = conn.MustBegin()
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
	if len(data) == 1 && update {
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
		tx.Commit()
		return data, nil
	}
	return data, err
}

func SelectAiringAnime() ([]DB_Anime, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Anime
	rows, err := conn.Query(`
		SELECT a.Id, a.Title, a.AlternativeTitle, a.Aired, a.Duration, a.Url, a.CurrentStatus, a.Season_ID, a.Type_ID, a.Broadcast,
			  (
				SELECT i.Url
				FROM Images i
				LEFT JOIN Anime_Images ai ON ai.Image_ID = i.Id
				WHERE ai.Anime_ID = a.Id AND ai.IsDefault = 1
				LIMIT 1
			  ) as Image
		FROM Anime a
		WHERE a.CurrentStatus = 1
		GROUP BY a.Id
		ORDER BY a.Broadcast, 
		CASE WHEN AlternativeTitle IS NOT NULL
			THEN AlternativeTitle
			ELSE Title
		END
	`)
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var d DB_Anime
		if err = rows.Scan(&d.Id, &d.Title, &d.AlternativeTitle, &d.Aired, &d.Duration, &d.Url, &d.CurrentStatus, &d.Season_ID, &d.Type_ID, &d.Broadcast, &d.Image); err != nil {
			return data, err
		}

		if len(d.Image.String) == 0 {
			images, _ := SelectAnimeAlternativeImages(d.Id, true)
			if len(images) > 0 {
				d.Image = sql.NullString{
					String: images[0],
					Valid:  len(images[0]) > 0,
				}
			}
		}
		data = append(data, d)
	}
	return data, nil
}
