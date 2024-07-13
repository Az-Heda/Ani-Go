package db

type DB_Theme struct {
	Id   string
	Name string
}

type DB_ThemeCount struct {
	Id    string
	Name  string
	Count int
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

func SelectAllThemesWithCount() ([]DB_ThemeCount, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_ThemeCount
	rows, err := conn.Query(`
		SELECT t.Id, t.Name, COUNT(*) AS "Count"
		FROM Themes t 
		LEFT JOIN Anime_Themes at ON t.Id = at.Theme_ID
		LEFT JOIN Anime a ON a.Id = at.Anime_ID
		GROUP BY t.Id
		ORDER BY t.Name
	`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var d DB_ThemeCount
		if err = rows.Scan(&d.Id, &d.Name, &d.Count); err != nil {
			return nil, err
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
