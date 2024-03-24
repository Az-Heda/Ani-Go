package db

type DB_Type struct {
	Id   string
	Name string
}

func SelectAllTypes() ([]DB_Type, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Type
	rows, err := conn.Query("SELECT Id, Name FROM Types ORDER BY Name ASC")
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var d DB_Type
		if err = rows.Scan(&d.Id, &d.Name); err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, nil
}

func SelectTypeFromId(id string) (DB_Type, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Type{}, err
	}
	var data DB_Type
	rows, err := conn.Query("SELECT Id, Name FROM Types WHERE Id = ?", id)
	if err != nil {
		return DB_Type{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Name); err != nil {
			return DB_Type{}, err
		}
	}
	return data, nil
}

func SelectTypeFromName(name string) (DB_Type, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Type{}, err
	}
	var data DB_Type
	rows, err := conn.Query("SELECT Id, Name FROM Types WHERE Name = ?", name)
	if err != nil {
		return DB_Type{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Name); err != nil {
			return DB_Type{}, err
		}
	}
	return data, nil
}
