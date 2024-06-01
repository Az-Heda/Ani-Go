package cliargs

import (
	"fmt"
	"os"
	"strings"
	"time"

	db "AniGo/db"
	dbintegration "AniGo/db-integration"
	scraper "AniGo/scraper"

	"github.com/jmoiron/sqlx"
)

func showHelp(ec int, keys []string) {
	fmt.Println("CLI Args list:")
	for _, key := range keys {
		fmt.Printf("%s\n", key)
	}
	os.Exit(ec)
}

func initializeDB(ec int) {
	dbintegration.Init()
	os.Exit(ec)
}

func scrape(url string) {
	scraper.Init(url)
}

func importFile(ec int, _ []string) {
	content, err := os.ReadFile("import.txt")
	if err != nil {
		panic(err)
	}

	conn, err := sqlx.Connect("sqlite", "./db/db.sqlite3")

	if err != nil {
		panic(err)
	}

	// conn.MustExec(`
	// 	DELETE FROM Anime
	// 	WHERE Id = (
	// 		SELECT Id
	// 		FROM Anime
	// 		ORDER BY ROWID DESC
	// 		LIMIT 1
	// 	)`)

	// dbInitializer, err := os.ReadFile("./db/db-setup-v2.sql")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// var queryInitializer string = string(dbInitializer)
	// conn.MustExec(queryInitializer)

	var rows []string = strings.Split(string(content), "\n")

	for i, row := range rows {
		var url string = scraper.FormatUrl(strings.TrimSpace(row))
		fmt.Println("[" + fmt.Sprint(i+1) + "] - " + url)
		if len(scraper.ExistAnimeID(conn, url)) == 0 {
			scraper.Init(url)
			if i+1 < len(rows) {
				time.Sleep(90 * time.Second)
			}
		}
	}
	os.Exit(ec)
}

func createVault(ec int, _ []string) {
	var vaultDir string = "Collection"
	var template string = vaultDir + "/main.template"

	content, err := os.ReadFile(template)
	if err != nil {
		panic(err)
	}
	var structure string = string(content)

	allAnime, _ := db.SelectAllAnime()
	allStatuses, _ := db.SelectAllStatuses()

	var bannedChars []string = []string{"\"", "'", "/", "(", ")", "[", "]", ":", ";", ",", " "}
	for _, anime := range allAnime {
		allStudios, _ := db.SelectStudioFromAnime(anime.Id)
		allGenres, _ := db.SelectGenreFromAnimeId(anime.Id)
		allThemes, _ := db.SelectThemeFromAnimeId(anime.Id)
		allDescription, _ := db.SelectDescriptionFromAnimeID(anime.Id)
		allEpisodes, _ := db.SelectEpisodeFromAnimeId(anime.Id)
		var currentStructure string = structure
		var title string = anime.Title
		if anime.AlternativeTitle.Valid {
			title = anime.AlternativeTitle.String
		}

		currentStructure = strings.Replace(currentStructure, "--original.title--", anime.Title, -1)
		currentStructure = strings.Replace(currentStructure, "--title--", title, -1)
		for _, s := range allStatuses {
			if s.Id == anime.CurrentStatus {
				var status string = strings.Replace("#"+toTitle(s.Name), " ", "", -1)
				currentStructure = strings.Replace(currentStructure, "--status--", status, -1)
				break
			}
		}
		var studioText string = ""
		for _, s := range allStudios {
			fmt.Println("Studio: ", s.Name)
			studioText += "#" + toTitle(s.Name)
		}
		if len(studioText) == 0 {
			studioText = "---"
		}
		currentStructure = strings.Replace(currentStructure, "--studios--", studioText, -1)

		if anime.Type_ID.Valid {
			typeData, _ := db.SelectTypeFromId(anime.Type_ID.String)
			currentStructure = strings.Replace(currentStructure, "--type--", "#"+toTitle(typeData.Name), -1)
		} else {
			currentStructure = strings.Replace(currentStructure, "--type--", "---", -1)
		}

		var genreText string = ""
		var themeText string = ""
		var descrText string = ""
		var episodeText string = ""

		for _, genre := range allGenres {
			genreText += "- #" + toTitle(genre.Name) + "\n"
		}
		for _, theme := range allThemes {
			themeText += "- #" + toTitle(theme.Name) + "\n"
		}
		for _, d := range allDescription {
			descrText += d.Description + "\n"
		}
		for _, e := range allEpisodes {
			var aired time.Time = time.Unix(0, e.Aired*1e6)
			episodeText += "| " + fmt.Sprint(e.Number) + " | " + e.Title + " | " + aired.Format("02/01/2006") + " |\n"
		}

		currentStructure = strings.Replace(currentStructure, "--genres--", genreText, -1)
		currentStructure = strings.Replace(currentStructure, "--themes--", themeText, -1)
		currentStructure = strings.Replace(currentStructure, "--description--", descrText, -1)
		currentStructure = strings.Replace(currentStructure, "--episodes--", episodeText, -1)

		for _, bc := range bannedChars {
			title = strings.Replace(toTitle(title), bc, "", -1)
		}
		fmt.Println(vaultDir + "/Anime/" + title + ".md")
		os.WriteFile(vaultDir+"/Anime/"+title+".md", []byte(currentStructure), 0777)
		break
	}
	os.Exit(1)
}

func toTitle(item string) string {
	var parts []string = strings.Split(item, " ")
	for i := range parts {
		parts[i] = strings.ToUpper(parts[i][0:1]) + strings.ToLower(parts[i][1:])
	}
	return strings.Join(parts, " ")
}
