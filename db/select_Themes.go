package db

type DB_Theme struct {
	Id   string
	Name string
}

func SelectAllThemes() ([]DB_Theme, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Theme
	rows, err := conn.Query("SELECT Id, Name FROM Themes ORDER BY Name ASC")
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var d DB_Theme
		if err = rows.Scan(&d.Id, &d.Name); err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, nil
}

func SelectThemeFromId(id string) (DB_Theme, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Theme{}, err
	}
	var data DB_Theme
	rows, err := conn.Query("SELECT Id, Name FROM Themes WHERE Id = ?", id)
	if err != nil {
		return DB_Theme{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Name); err != nil {
			return DB_Theme{}, err
		}
	}
	return data, nil
}

func SelectThemeFromAnimeId(id string) ([]DB_Theme, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Theme
	rows, err := conn.Query("SELECT g.Id, g.Name FROM Themes g LEFT JOIN Anime_Themes ag ON ag.Theme_ID = g.Id WHERE ag.Anime_ID = ? ORDER BY g.Name;", id)
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var d DB_Theme
		if err = rows.Scan(&d.Id, &d.Name); err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, nil
}

func SelectThemeFromName(name string) (DB_Theme, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Theme{}, err
	}
	var data DB_Theme
	rows, err := conn.Query("SELECT Id, Name FROM Themes WHERE Name = ?", name)
	if err != nil {
		return DB_Theme{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Name); err != nil {
			return DB_Theme{}, err
		}
	}
	return data, nil
}
