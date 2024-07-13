package scraper

import (
	"strings"

	db "AniGo/db"

	"github.com/jmoiron/sqlx"
)

func FormatUrl(url string) string {
	url = strings.Split(url, "?")[0]
	if url[len(url)-1:] == "/" {
		url = url[0 : len(url)-1]
	}
	return url
}

func Init(url string) ScrapedAll {
	url = FormatUrl(url)

	var urls map[string]string = map[string]string{
		"Anime":      url,
		"Pics":       url + "/pics",
		"Episodes":   url + "/episode",
		"Characters": url + "/characters",
		"Pictures":   url + "/pics",
	}

	conn, err := db.GetConnection()
	if err != nil {
		panic(err)
	}

	var tx *sqlx.Tx = conn.MustBegin()

	animeData, _ := parseAnimePage(urls["Anime"], conn, tx)
	// if ok {
	// 	b, _ := json.MarshalIndent(animeData, "", "\t")
	// 	os.WriteFile("parsed-anime.json", b, 0777)
	// }

	animeImages, _ := parseAnimePics(animeData.Anime.Id, urls["Pics"], conn, tx)
	// if ok {
	// 	b, _ := json.MarshalIndent(animeImages, "", "\t")
	// 	os.WriteFile("parsed-anime-pics.json", b, 0777)
	// }

	animeCharacters, _ := parseAnimeCharacterList(animeData.Anime.Id, urls["Characters"], conn, tx)
	// if ok {
	// 	b, _ := json.MarshalIndent(animeCharacters, "", "\t")
	// 	os.WriteFile("parse-anime-characters.json", b, 0777)
	// }

	animeEpisodes, _ := parseAnimeEpisodeList(animeData.Anime.Id, urls["Episodes"], conn, tx)
	// if ok {
	// 	b, _ := json.MarshalIndent(animeEpisodes, "", "\t")
	// 	os.WriteFile("parse-anime-episodes.json", b, 0777)
	// }

	var allData ScrapedAll = ScrapedAll{
		ScrapedAnime:           animeData,
		ScrapedAnimePics:       animeImages,
		ScrapedAnimeCharacters: animeCharacters,
		ScrapedAnimeEpisodes:   animeEpisodes,
	}

	InsertScrapedData(allData, conn, tx)

	return allData

	// if err := tx.Commit(); err != nil {
	// 	log.Fatalln(err)
	// }

	// tx = conn.MustBegin()
	// parseAnimePics(urls["Anime"], urls["Pics"], conn, tx)
	// parseAnimeEpisodeList(urls["Anime"], urls["Episodes"], conn, tx)
	// parseAnimeCharacterList(urls["Anime"], urls["Characters"], conn, tx)

	// if err := tx.Commit(); err != nil {
	// 	log.Fatalln(err)
	// }
}

func InsertScrapedData(data ScrapedAll, conn *sqlx.DB, tx *sqlx.Tx) error {

	// parseAnimePage
	insertType(conn, tx, data.ScrapedAnime.AnimeType)
	insertSeason(conn, tx, data.ScrapedAnime.AnimeSeason)
	insertAnimeDescription(conn, tx, data.ScrapedAnime.AnimeDescription)
	insertStudio(conn, tx, data.ScrapedAnime.AnimeStudios)
	insertGenre(conn, tx, data.ScrapedAnime.AnimeGenre)
	insertTheme(conn, tx, data.ScrapedAnime.AnimeTheme)
	insertAnime(conn, tx, data.ScrapedAnime.Anime, data.ScrapedAnimePics.AnimeImages[0])

	insertAnimeStudio(conn, tx, data.ScrapedAnime.Join_Anime_Studio)
	insertAnimeGenre(conn, tx, data.ScrapedAnime.Join_Anime_Genre)
	insertAnimeTheme(conn, tx, data.ScrapedAnime.Join_Anime_Theme)

	// parseAnimePics
	insertImage(conn, tx, data.ScrapedAnimePics.AnimeImages)
	insertAnimeImage(conn, tx, data.ScrapedAnimePics.Join_Anime_Images)

	// parseAnimeCharacterList
	insertCharacter(conn, tx, data.ScrapedAnimeCharacters.AnimeCharacters)
	insertCharacterDescription(conn, tx, data.ScrapedAnime.AnimeDescription)
	insertImage(conn, tx, data.ScrapedAnimeCharacters.CharacterImages)
	insertCharacterImage(conn, tx, data.ScrapedAnimeCharacters.Join_Character_Image)

	// insertCharacterDescription(conn, tx, data.ScrapedAnime.AnimeDescription)
	insertAnimeCharacter(conn, tx, data.ScrapedAnimeCharacters.Join_Anime_Characters)
	insertCharacterImage(conn, tx, data.ScrapedAnimeCharacters.Join_Character_Image)

	// parseAnimeEpisodeList
	insertEpisode(conn, tx, data.ScrapedAnimeEpisodes.AnimeEpisodes)
	insertEpisodeDescription(conn, tx, data.ScrapedAnime.AnimeDescription)

	return tx.Commit()
}
