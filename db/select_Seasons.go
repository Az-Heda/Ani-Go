package db

type DB_Season struct {
	Id     string
	Season string
}

func SelectAllSeasons() ([]DB_Season, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Season
	rows, err := conn.Query("SELECT Id, Season FROM Seasons ORDER BY Season ASC")
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var d DB_Season
		if err = rows.Scan(&d.Id, &d.Season); err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, nil
}

func SelectSeasonFromId(id string) (DB_Season, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Season{}, err
	}
	var data DB_Season
	rows, err := conn.Query("SELECT Id, Season FROM Seasons WHERE Id = ?", id)
	if err != nil {
		return DB_Season{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Season); err != nil {
			return DB_Season{}, err
		}
	}
	return data, nil
}

func SelectSeasonFromName(name string) (DB_Season, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Season{}, err
	}
	var data DB_Season
	rows, err := conn.Query("SELECT Id, Season FROM Seasons WHERE Season = ?", name)
	if err != nil {
		return DB_Season{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Season); err != nil {
			return DB_Season{}, err
		}
	}
	return data, nil
}
