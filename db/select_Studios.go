package db

type DB_Studio struct {
	Id   string
	Name string
}

func SelectAllStudios() ([]DB_Studio, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Studio
	rows, err := conn.Query("SELECT Id, Name FROM Studios ORDER BY Name ASC")
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var d DB_Studio
		if err = rows.Scan(&d.Id, &d.Name); err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, nil
}

func SelectStudioFromId(id string) (DB_Studio, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Studio{}, err
	}
	var data DB_Studio
	rows, err := conn.Query("SELECT Id, Name FROM Studios WHERE Id = ?", id)
	if err != nil {
		return DB_Studio{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Name); err != nil {
			return DB_Studio{}, err
		}
	}
	return data, nil
}

func SelectStudioFromName(name string) (DB_Studio, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Studio{}, err
	}
	var data DB_Studio
	rows, err := conn.Query("SELECT Id, Name FROM Studios WHERE Name = ?", name)
	if err != nil {
		return DB_Studio{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Name); err != nil {
			return DB_Studio{}, err
		}
	}
	return data, nil
}

func SelectStudioFromAnime(anime_id string) ([]DB_Studio, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Studio
	rows, err := conn.Query(`
		SELECT s.Id, s.Name
		FROM Studios s
		LEFT JOIN Anime_Studioses as2 ON s.Id = as2.Studio_ID
		WHERE as2.Anime_ID = ?
		ORDER BY s.Name ASC
	`, anime_id)
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var d DB_Studio
		if err = rows.Scan(&d.Id, &d.Name); err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, nil
}
