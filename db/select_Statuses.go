package db

type DB_Status struct {
	Id        int
	Name      string
	IsVisible int
}

func SelectAllStatuses() ([]DB_Status, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Status
	rows, err := conn.Query("SELECT Id, Name, IsVisible FROM Statuses ORDER BY Id ASC")
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var d DB_Status
		if err = rows.Scan(&d.Id, &d.Name, &d.IsVisible); err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, nil
}

func SelectStatusFromId(id string) (DB_Status, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Status{}, err
	}
	var data DB_Status
	rows, err := conn.Query("SELECT Id, Name, IsVisible FROM Statuses WHERE Id = ?", id)
	if err != nil {
		return DB_Status{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Name, &data.IsVisible); err != nil {
			return DB_Status{}, err
		}
	}
	return data, nil
}

func SelectStatusFromName(name string) (DB_Status, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Status{}, err
	}
	var data DB_Status
	rows, err := conn.Query("SELECT Id, Name, IsVisible FROM Statuses WHERE Name = ?", name)
	if err != nil {
		return DB_Status{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Name, &data.IsVisible); err != nil {
			return DB_Status{}, err
		}
	}
	return data, nil
}
