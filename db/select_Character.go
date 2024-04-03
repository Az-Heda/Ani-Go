package db

import "database/sql"

type DB_Character struct {
	Id    string
	Name  string
	Image sql.NullString
}

func SelectAllCharacters() ([]DB_Character, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var data []DB_Character
	rows, err := conn.Query("SELECT c.Id, c.Name, (SELECT i.Url FROM Images i LEFT JOIN Character_Images ci ON ci.Image_ID = i.Id WHERE ci.Character_ID = c.Id LIMIT 1) as Image FROM Character c ORDER BY c.Name")
	if err != nil {
		return data, err
	}
	for rows.Next() {
		var d DB_Character
		if err = rows.Scan(&d.Id, &d.Name, &d.Image); err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, nil
}

func SelectCharacterFromId(id string) (DB_Character, error) {
	conn, err := GetConnection()
	if err != nil {
		return DB_Character{}, err
	}
	var data DB_Character
	rows, err := conn.Query("SELECT Id, Name FROM Character WHERE Id = ?", id)
	if err != nil {
		return DB_Character{}, err
	}
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Name); err != nil {
			return DB_Character{}, err
		}
	}
	return data, nil
}
