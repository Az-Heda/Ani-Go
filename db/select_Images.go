package db

type DB_Image struct {
	Id  string
	Url string
}

func SelectAllImages() ([]DB_Image, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Image
	rows, err := conn.Query("SELECT Id, Url FROM Images ORDER BY Url ASC")
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var d DB_Image
		if err = rows.Scan(&d.Id, &d.Url); err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, nil
}

func SelectImageFromId(id string) (DB_Image, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Image{}, err
	}
	var data DB_Image
	rows, err := conn.Query("SELECT Id, Url FROM Images WHERE Id = ?", id)
	if err != nil {
		return DB_Image{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Url); err != nil {
			return DB_Image{}, err
		}
	}
	return data, nil
}

func SelectRandomNImages(n int64, status int64) ([]string, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []string
	conn.Select(&data, "SELECT i.Url FROM Anime a LEFT JOIN Anime_Images ai ON ai.Anime_ID = a.Id LEFT JOIN Images i ON ai.Image_ID = i.Id WHERE a.CurrentStatus = ? ORDER BY RANDOM() LIMIT ?", status, n)
	return data, nil
}
