package db

import "database/sql"

type DB_Episode struct {
	Id          string
	Number      int
	Title       string
	Aired       int64
	Duration    int64
	Anime_ID    string
	Description sql.NullString
}

func SelectAllEpisodes() ([]DB_Episode, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Episode
	rows, err := conn.Query("SELECT Id, Number, Title, Aired, Duration, Anime_ID FROM Episode ORDER BY Anime_ID ASC, Number ASC")
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var d DB_Episode
		if err = rows.Scan(&d.Id, &d.Number, &d.Title, &d.Aired, &d.Duration, &d.Anime_ID); err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, nil
}

func SelectEpisodeFromId(id string) (DB_Episode, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Episode{}, err
	}
	var data DB_Episode
	rows, err := conn.Query("SELECT Id, Number, Title, Aired, Duration, Anime_ID FROM Episode WHERE Id = ?", id)
	if err != nil {
		return DB_Episode{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Number, &data.Title, &data.Aired, &data.Duration, &data.Anime_ID); err != nil {
			return DB_Episode{}, err
		}
	}
	return data, nil
}

func SelectEpisodeFromAnimeId(id string) ([]DB_Episode, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Episode
	rows, err := conn.Query("SELECT e.Id, e.Number, e.Title, e.Aired, e.Anime_Id, (SELECT GROUP_CONCAT(d.Description, '\n') FROM Episode_Descriptions ed LEFT JOIN Descriptions d ON ed.Description_ID = d.Id WHERE ed.Episode_ID = e.Id) AS Description FROM Episode e WHERE e.Anime_ID = ?", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var d DB_Episode
		if err = rows.Scan(&d.Id, &d.Number, &d.Title, &d.Aired, &d.Anime_ID, &d.Description); err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, nil
}
