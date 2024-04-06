package db

func UpdateAnimeStatus(id string, newValue string) error {
	conn, err := GetConnection()
	if err != nil {
		return err
	}
	_, err = conn.Exec("UPDATE Anime SET CurrentStatus = ? WHERE Id = ?;", newValue, id)
	return err
}
