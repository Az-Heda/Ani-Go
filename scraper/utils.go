package scraper

import (
	"database/sql"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func nullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func generateID() string {
	var uuids []string = []string{
		uuid.New().String(),
		uuid.New().String(),
		uuid.New().String(),
	}
	return strings.Join(uuids, "-")
}

func ExistSeasonID(conn *sqlx.DB, value string) string {
	var data []string = []string{}
	err := conn.Select(&data, "SELECT Id FROM Seasons WHERE Season = ? LIMIT 1", value)
	if err != nil {
		return ""
	}
	if len(data) == 1 {
		return data[0]
	}
	return ""
}
func GetSeasonID(conn *sqlx.DB, value string) string {
	var id string = ExistSeasonID(conn, value)
	if len(id) == 0 {
		return generateID()
	}
	return id
}

func ExistTypeID(conn *sqlx.DB, value string) string {
	var data []string = []string{}
	err := conn.Select(&data, "SELECT Id FROM Types WHERE Name = ?", value)
	if err != nil {
		return ""
	}
	if len(data) == 1 {
		return data[0]
	}
	return ""
}
func GetTypeID(conn *sqlx.DB, value string) string {
	var id string = ExistTypeID(conn, value)
	if len(id) == 0 {
		return generateID()
	}
	return id
}

func ExistStudioID(conn *sqlx.DB, value string) string {
	var data []string = []string{}
	err := conn.Select(&data, "SELECT Id FROM Studios WHERE Name = ?", value)
	if err != nil {
		return ""
	}
	if len(data) == 1 {
		return data[0]
	}
	return ""
}
func GetStudioID(conn *sqlx.DB, value string) string {
	var id string = ExistStudioID(conn, value)
	if len(id) == 0 {
		return generateID()
	}
	return id
}

func ExistGenreID(conn *sqlx.DB, value string) string {
	var data []string = []string{}
	err := conn.Select(&data, "SELECT Id FROM Genres WHERE Name = ?", value)
	if err != nil {
		return ""
	}
	if len(data) == 1 {
		return data[0]
	}
	return ""
}
func GetGenreID(conn *sqlx.DB, value string) string {
	var id string = ExistGenreID(conn, value)
	if len(id) == 0 {
		return generateID()
	}
	return id
}

func ExistThemeID(conn *sqlx.DB, value string) string {
	var data []string = []string{}
	err := conn.Select(&data, "SELECT Id FROM Themes WHERE Name = ?", value)
	if err != nil {
		return generateID()
	}
	if len(data) == 1 {
		return data[0]
	}
	return generateID()
}
func GetThemeID(conn *sqlx.DB, value string) string {
	var id string = ExistThemeID(conn, value)
	if len(id) == 0 {
		return generateID()
	}
	return id
}

func ExistAnimeDescriptionID(conn *sqlx.DB, animeID string, descr string) string {
	var data []string = []string{}
	err := conn.Select(&data, "SELECT Id FROM Descriptions WHERE Description = ? AND Anime_ID = ? LIMIT 1", descr, animeID)
	if err != nil {
		return ""
	}

	if len(data) == 1 {
		return data[0]
	}
	return ""
}
func GetAnimeDescriptionID(conn *sqlx.DB, animeID string, descr string) string {
	return generateID()
	// var id string = ExistAnimeDescriptionID(conn, animeID, descr)
	// if len(id) == 0 {
	// 	return generateID()
	// }
	// return id
}

func ExistEpisodeDescriptionID(conn *sqlx.DB, episodeID string, descr string) string {
	var data []string = []string{}
	err := conn.Select(&data, "SELECT Id FROM Descriptions WHERE Description = ? AND Episode_ID = ? LIMIT 1", descr, episodeID)
	if err != nil {
		return ""
	}

	if len(data) == 1 {
		return data[0]
	}
	return ""
}
func GetEpisodeDescriptionID(conn *sqlx.DB, episodeID string, descr string) string {
	return generateID()
	// var id string = ExistEpisodeDescriptionID(conn, episodeID, descr)
	// if len(id) == 0 {
	// 	return generateID()
	// }
	// return id
}

func ExistCharacterDescriptionID(conn *sqlx.DB, characterID string, descr string) string {
	var data []string = []string{}
	err := conn.Select(&data, "SELECT Id FROM Descriptions WHERE Description = ? AND Character_ID = ? LIMIT 1", descr, characterID)
	if err != nil {
		return ""
	}

	if len(data) == 1 {
		return data[0]
	}
	return ""
}
func GetCharacterDescriptionID(conn *sqlx.DB, characterID string, descr string) string {
	return generateID()
	// var id string = ExistCharacterDescriptionID(conn, characterID, descr)
	// if len(id) == 0 {
	// 	return generateID()
	// }
	// return id
}

func ExistImageID(conn *sqlx.DB, url string) string {
	var data []string = []string{}
	err := conn.Select(&data, "SELECT Id FROM Images WHERE Url = ?", url)
	if err != nil {
		return ""
	}
	if len(data) == 1 {
		return data[0]
	}
	return ""
}
func GetImageID(conn *sqlx.DB, url string) string {
	var id string = ExistImageID(conn, url)
	if len(id) == 0 {
		return generateID()
	}
	return id
}

func ExistAnimeID(conn *sqlx.DB, url string) string {
	var data []string = []string{}
	err := conn.Select(&data, "SELECT Id FROM Anime WHERE Url = ?", url)
	if err != nil {
		return ""
	}
	if len(data) == 1 {
		return data[0]
	}
	return ""
}
func GetAnimeID(conn *sqlx.DB, url string) string {
	var id string = ExistAnimeID(conn, url)
	if len(id) == 0 {
		return generateID()
	}
	return id
}

func ExistCharacterID(conn *sqlx.DB, url string) string {
	var data []string = []string{}
	err := conn.Select(&data, "SELECT Id FROM Character WHERE Url = ?", url)
	if err != nil {
		return ""
	}
	if len(data) == 1 {
		return data[0]
	}
	return ""
}
func GetCharacterID(conn *sqlx.DB, url string) string {
	var id string = ExistCharacterID(conn, url)
	if len(id) == 0 {
		return generateID()
	}
	return id
}

func ExistEpisodeID(conn *sqlx.DB, animeID string, episodeNumber int) string {
	var data []string = []string{}
	err := conn.Select(&data, "SELECT Id FROM Episode WHERE Anime_ID = ? AND Number = ?", animeID, episodeNumber)
	if err != nil {
		return ""
	}
	if len(data) == 1 {
		return data[0]
	}
	return ""
}
func GetEpisodeID(conn *sqlx.DB, animeID string, episodeNumber int) string {
	var id string = ExistEpisodeID(conn, animeID, episodeNumber)
	if len(id) == 0 {
		return generateID()
	}
	return id
}

func insertAnime(conn *sqlx.DB, tx *sqlx.Tx, data ScraperAnime, image ScraperImage) {
	tx.MustExec(`
		INSERT INTO Anime (Id, Title, AlternativeTitle, Aired, Duration, Url, Broadcast, Season_ID, Type_ID, Image_ID, CurrentStatus)
		SELECT ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		WHERE NOT EXISTS (
			SELECT 1
			FROM Anime
			WHERE Id = ?
		)`,
		data.Id,
		data.Title,
		nullString(data.AlternativeTitle),
		data.Aired,
		data.Duration,
		data.Url,
		data.Broadcast,
		nullString(data.Season_ID),
		nullString(data.Type_ID),
		image.Id,
		data.CurrentStatus,
		data.Id,
	)
}

func insertSeason(conn *sqlx.DB, tx *sqlx.Tx, data ScraperSeason) {
	tx.MustExec(`
		INSERT INTO Seasons (Id, Season)
		SELECT ?, ?
		WHERE NOT EXISTS (
			SELECT 1
			FROM Seasons
			WHERE Id = ?
		)`,
		data.Id,
		data.Season,
		data.Id,
	)
}

func insertType(conn *sqlx.DB, tx *sqlx.Tx, data ScraperTypes) {
	tx.MustExec(`
		INSERT INTO Types (Id, Name)
		SELECT ?, ?
		WHERE NOT EXISTS (
			SELECT 1
			FROM Types
			WHERE Id = ?
		)`,
		data.Id,
		data.Name,
		data.Id,
	)
}

func insertAnimeDescription(conn *sqlx.DB, tx *sqlx.Tx, data []ScraperDescription) {
	for _, item := range data {
		tx.MustExec(`
			INSERT INTO Descriptions (Id, Description, Anime_ID, Episode_ID, Character_ID)
			SELECT ?, ?, ?, null, null
			WHERE NOT EXISTS (
				SELECT 1
				FROM Descriptions
				WHERE Id = ?
			)`,
			item.Id,
			item.Description,
			item.Anime_ID,
			item.Id,
			item.Anime_ID,
		)

	}
}

func insertEpisodeDescription(conn *sqlx.DB, tx *sqlx.Tx, data []ScraperDescription) {
	for _, item := range data {
		tx.MustExec(`
			INSERT INTO Descriptions (Id, Description, Anime_ID, Episode_ID, Character_ID)
			SELECT ?, ?, null, ?, null
			WHERE NOT EXISTS (
				SELECT 1
				FROM Descriptions
				WHERE Id = ?
			)`,
			item.Id,
			item.Description,
			item.Episode_ID,
			item.Id,
			item.Episode_ID,
		)

	}
}

func insertCharacterDescription(conn *sqlx.DB, tx *sqlx.Tx, data []ScraperDescription) {
	for _, item := range data {
		tx.MustExec(`
			INSERT INTO Descriptions (Id, Description, Anime_ID, Episode_ID, Character_ID)
			SELECT ?, ?, null, null, ?
			WHERE NOT EXISTS (
				SELECT 1
				FROM Descriptions
				WHERE Id = ?
			)`,
			item.Id,
			item.Description,
			item.Character_ID,
			item.Id,
			item.Character_ID,
		)

	}
}

func insertStudio(conn *sqlx.DB, tx *sqlx.Tx, data []ScraperStudio) {
	for _, item := range data {
		tx.MustExec(`
			INSERT INTO Studios (Id, Name)
			SELECT ?, ?
			WHERE NOT EXISTS (
				SELECT 1
				FROM Studios
				WHERE Id = ?
			)`,
			item.Id,
			item.Name,
			item.Id,
		)
	}
}

func insertGenre(conn *sqlx.DB, tx *sqlx.Tx, data []ScraperGenre) {
	for _, item := range data {
		tx.MustExec(`
			INSERT INTO Genres (Id, Name)
			SELECT ?, ?
			WHERE NOT EXISTS (
				SELECT 1
				FROM Genres
				WHERE Id = ?
			)`,
			item.Id,
			item.Name,
			item.Id,
		)
	}
}

func insertTheme(conn *sqlx.DB, tx *sqlx.Tx, data []ScraperTheme) {
	for _, item := range data {
		tx.MustExec(`
			INSERT INTO Themes (Id, Name)
			SELECT ?, ?
			WHERE NOT EXISTS (
				SELECT 1
				FROM Themes
				WHERE Id = ?
			)`,
			item.Id,
			item.Name,
			item.Id,
		)
	}
}

func insertEpisode(conn *sqlx.DB, tx *sqlx.Tx, data []ScraperEpisode) {
	for _, item := range data {
		tx.MustExec(`
			INSERT INTO Episode (Id, Number, Title, Aired, Duration, Anime_ID)
			SELECT ?, ?, ?, ?, ?, ?
			WHERE NOT EXISTS (
				SELECT 1
				FROM Episode
				WHERE Id = ?
			)`,
			item.Id,
			item.Number,
			nullString(item.Title),
			item.Aired,
			item.Duration,
			item.Anime_ID,
			item.Id,
		)
	}
}

func insertCharacter(conn *sqlx.DB, tx *sqlx.Tx, data []ScraperCharacter) {
	for _, item := range data {
		tx.MustExec(`
			INSERT INTO Character (Id, Name, Url, Image_ID)
			SELECT ?, ?, ?, ?
			WHERE NOT EXISTS (
				SELECT 1
				FROM Character
				WHERE Id = ?
			)`,
			item.Id,
			item.Name,
			item.Url,
			item.Image.Id,
			item.Id,
		)
	}
}

func insertImage(conn *sqlx.DB, tx *sqlx.Tx, data []ScraperImage) {
	for _, item := range data {
		tx.MustExec(`
			INSERT INTO Images (Id, Url)
			SELECT ?, ?
			WHERE NOT EXISTS (
				SELECT 1
				FROM Images 
				WHERE Id = ?
			)`,
			item.Id,
			item.Url,
			item.Id,
		)
	}
}

func insertAnimeStudio(conn *sqlx.DB, tx *sqlx.Tx, data []ScraperAnimeStudio) {
	for _, item := range data {
		tx.MustExec(`
			INSERT INTO Anime_Studioses (Anime_ID, Studio_ID)
			SELECT ?, ?
			WHERE NOT EXISTS (
				SELECT 1
				FROM Anime_Studioses
				WHERE Anime_ID = ? AND
					  Studio_ID = ?
			)`,
			item.Anime_ID,
			item.Studio_ID,
			item.Anime_ID,
			item.Studio_ID,
		)
	}
}

func insertAnimeGenre(conn *sqlx.DB, tx *sqlx.Tx, data []ScraperAnimeGenre) {
	for _, item := range data {
		tx.MustExec(`
			INSERT INTO Anime_Genres (Anime_ID, Genre_ID)
			SELECT ?, ?
			WHERE NOT EXISTS (
				SELECT 1
				FROM Anime_Genres
				WHERE Anime_ID = ? AND
					  Genre_ID = ?
			)`,
			item.Anime_ID,
			item.Genre_ID,
			item.Anime_ID,
			item.Genre_ID,
		)
	}
}

func insertAnimeTheme(conn *sqlx.DB, tx *sqlx.Tx, data []ScraperAnimeTheme) {
	for _, item := range data {
		tx.MustExec(`
			INSERT INTO Anime_Themes (Anime_ID, Theme_ID)
			SELECT ?, ?
			WHERE NOT EXISTS (
				SELECT 1
				FROM Anime_Themes
				WHERE Anime_ID = ? AND
					  Theme_ID = ?
			)`,
			item.Anime_ID,
			item.Theme_ID,
			item.Anime_ID,
			item.Theme_ID,
		)
	}
}

func insertAnimeCharacter(conn *sqlx.DB, tx *sqlx.Tx, data []ScraperAnimeCharacter) {
	for _, item := range data {
		tx.MustExec(`
			INSERT INTO Anime_Characters (Anime_ID, Character_ID)
			SELECT ?, ?
			WHERE NOT EXISTS (
				SELECT 1
				FROM Anime_Characters
				WHERE Anime_ID = ? AND
					  Character_ID = ?
			)`,
			item.Anime_ID,
			item.Character_ID,
			item.Anime_ID,
			item.Character_ID,
		)
	}
}

func insertAnimeImage(conn *sqlx.DB, tx *sqlx.Tx, data []ScraperAnimeImage) {
	for _, item := range data {
		tx.MustExec(`
			INSERT INTO Anime_Images (Anime_ID, Image_ID)
			SELECT ?, ?
			WHERE NOT EXISTS (
				SELECT 1
				FROM Anime_Images
				WHERE Anime_ID = ? AND
					  Image_ID = ?
			)`,
			item.Anime_ID,
			item.Image_ID,
			item.Anime_ID,
			item.Image_ID,
		)
	}
}

func insertCharacterImage(conn *sqlx.DB, tx *sqlx.Tx, data []ScraperCharacterImage) {
	for _, item := range data {
		tx.MustExec(`
			INSERT INTO Character_Images (Character_ID, Image_ID)
			SELECT ?, ?
			WHERE NOT EXISTS (
				SELECT 1
				FROM Character_Images
				WHERE Character_ID = ? AND
					  Image_ID = ?
			)`,
			item.Character_ID,
			item.Image_ID,
			item.Character_ID,
			item.Image_ID,
		)
	}
}
