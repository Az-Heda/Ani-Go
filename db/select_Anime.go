package db

import "database/sql"

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
	rows, err := conn.Query("SELECT a.Id, a.Title, a.AlternativeTitle, a.Aired, a.Duration, a.Url, a.CurrentStatus, a.Season_ID, a.Type_ID, (SELECT i.Url FROM Images i LEFT JOIN Anime_Images ai ON ai.Image_ID = i.Id WHERE ai.Anime_ID = a.Id LIMIT 1) as Image, (SELECT GROUP_CONCAT(d.Description, '\n') FROM Descriptions d LEFT JOIN Anime_Descriptions ad ON ad.Description_ID = d.Id WHERE ad.Anime_ID = a.Id) as Description FROM Anime a GROUP BY a.Id ORDER BY CASE WHEN AlternativeTitle IS NOT NULL THEN AlternativeTitle ELSE Title END")
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var d DB_Anime
		if err = rows.Scan(&d.Id, &d.Title, &d.AlternativeTitle, &d.Aired, &d.Duration, &d.Url, &d.CurrentStatus, &d.Season_ID, &d.Type_ID, &d.Image, &d.Description); err != nil {
			return data, err
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
	rows, err := conn.Query("SELECT Id, Title, AlternativeTitle, Aired, Duration, Url, CurrentStatus, Season_ID, Type_ID FROM Anime WHERE Id = ?", id)
	if err != nil {
		return DB_Anime{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Title, &data.AlternativeTitle, &data.Aired, &data.Duration, &data.Url, &data.CurrentStatus, &data.Season_ID, &data.Type_ID); err != nil {
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
	rows, err := conn.Query("SELECT Id, Title, AlternativeTitle, Aired, Duration, Url, CurrentStatus, Season_ID, Type_ID FROM Anime WHERE Title LIKE ? OR AlternativeTitle LIKE ?", likeFilter, likeFilter)
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
