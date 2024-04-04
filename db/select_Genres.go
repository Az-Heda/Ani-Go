package db

type DB_Genre struct {
	Id   string
	Name string
}

func SelectAllGenres() ([]DB_Genre, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Genre
	rows, err := conn.Query("SELECT Id, Name FROM Genres ORDER BY Name ASC")
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var d DB_Genre
		if err = rows.Scan(&d.Id, &d.Name); err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, nil
}

func SelectGenreFromId(id string) (DB_Genre, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Genre{}, err
	}
	var data DB_Genre
	rows, err := conn.Query("SELECT Id, Name FROM Genres WHERE Id = ?", id)
	if err != nil {
		return DB_Genre{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Name); err != nil {
			return DB_Genre{}, err
		}
	}
	return data, nil
}

func SelectGenreFromAnimeId(id string) ([]DB_Genre, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Genre
	rows, err := conn.Query("SELECT g.Id, g.Name FROM Genres g LEFT JOIN Anime_Genres ag ON ag.Genre_ID = g.Id WHERE ag.Anime_ID = ? ORDER BY g.Name;", id)
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var d DB_Genre
		if err = rows.Scan(&d.Id, &d.Name); err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, nil
}

func SelectGenreFromName(name string) (DB_Genre, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Genre{}, err
	}
	var data DB_Genre
	rows, err := conn.Query("SELECT Id, Name FROM Genres WHERE Name = ?", name)
	if err != nil {
		return DB_Genre{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Name); err != nil {
			return DB_Genre{}, err
		}
	}
	return data, nil
}
