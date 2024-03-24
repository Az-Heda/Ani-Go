package db

type DB_Description struct {
	Id          string
	Description string
}

func SelectAllDescriptions() ([]DB_Description, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Description
	rows, err := conn.Query("SELECT Id, Description FROM Descriptions ORDER BY Description ASC")
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var d DB_Description
		if err = rows.Scan(&d.Id, &d.Description); err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, nil
}

func SelectDescriptionFromId(id string) (DB_Description, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Description{}, err
	}
	var data DB_Description
	rows, err := conn.Query("SELECT Id, Description FROM Descriptions WHERE Id = ?", id)
	if err != nil {
		return DB_Description{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Description); err != nil {
			return DB_Description{}, err
		}
	}
	return data, nil
}
