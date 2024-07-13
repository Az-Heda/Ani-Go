package scraper

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/jmoiron/sqlx"
)

func parseAnimePage(url string, conn *sqlx.DB, tx *sqlx.Tx) (ScrapedAnime, bool) {
	var anime ScraperAnime
	var animeSeason ScraperSeason
	var animeType ScraperTypes
	var animeDescription []ScraperDescription
	var animeStudios []ScraperStudio
	var animeGenres []ScraperGenre
	var animeThemes []ScraperTheme

	var join_anime_studios []ScraperAnimeStudio
	var join_anime_genre []ScraperAnimeGenre
	var join_anime_theme []ScraperAnimeTheme

	collector := colly.NewCollector()

	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println(url)
		panic(e)
	})

	collector.OnHTML("h1.title-name", func(title *colly.HTMLElement) {
		anime.Id = GetAnimeID(conn, url)
		anime.CurrentStatus = 4
		anime.Url = url
		anime.Title = title.Text
	})
	collector.OnHTML("p.title-english", func(title *colly.HTMLElement) {
		anime.AlternativeTitle = title.Text
	})

	collector.OnHTML("div.leftside", func(info *colly.HTMLElement) {
		info.ForEach("div.spaceit_pad", func(idx int, div *colly.HTMLElement) {
			if strings.Contains(div.Text, "Premiered:") {
				div.ForEach("a", func(idx int, a *colly.HTMLElement) {
					animeSeason.Season = a.Text
					animeSeason.Id = GetSeasonID(conn, animeSeason.Season)
					anime.Season_ID = animeSeason.Id
				})
			}

			if strings.Contains(div.Text, "Aired:") {
				var s string = strings.Trim(strings.Split(div.Text, "\n")[2], " ")
				var times []string = strings.Split(s, " to ")
				t, err := time.Parse("Jan 2, 2006", times[0])
				if err != nil {
					anime.Aired = 0
					return
				}
				anime.Aired = t.Unix() * 1000
			}

			if strings.Contains(div.Text, "Type:") {
				div.ForEach("a", func(idx int, a *colly.HTMLElement) {
					animeType.Name = a.Text
					animeType.Id = GetTypeID(conn, animeType.Name)
					anime.Type_ID = animeType.Id
				})
			}

			if strings.Contains(div.Text, "Studios:") {
				div.ForEach("a", func(idx int, a *colly.HTMLElement) {
					var studioData ScraperStudio = ScraperStudio{GetStudioID(conn, a.Text), a.Text}
					animeStudios = append(animeStudios, studioData)
					join_anime_studios = append(join_anime_studios, ScraperAnimeStudio{anime.Id, studioData.Id})
				})
			}

			if strings.Contains(div.Text, "Genres:") {
				div.ForEach("a", func(idx int, a *colly.HTMLElement) {
					var genreData ScraperGenre = ScraperGenre{GetGenreID(conn, a.Text), a.Text}
					animeGenres = append(animeGenres, genreData)
					join_anime_genre = append(join_anime_genre, ScraperAnimeGenre{anime.Id, genreData.Id})
				})
			}

			if strings.Contains(div.Text, "Theme:") || strings.Contains(div.Text, "Themes:") {
				div.ForEach("a", func(idx int, a *colly.HTMLElement) {
					var themeData ScraperTheme = ScraperTheme{GetThemeID(conn, a.Text), a.Text}
					animeThemes = append(animeThemes, themeData)
					join_anime_theme = append(join_anime_theme, ScraperAnimeTheme{anime.Id, themeData.Id})
				})
			}

			if strings.Contains(div.Text, "Duration:") {
				var s string = strings.ReplaceAll(strings.Trim(strings.Split(div.Text, "\n")[2], " "), " per ep.", "")
				var layoutHM string = "3 hr. 4 min."
				var layoutM string = "4 min."
				t, err := time.Parse(layoutHM, s)
				if err != nil {
					t, err = time.Parse(layoutM, s)
					if err != nil {
						anime.Duration = 0
						return
					}
				}
				anime.Duration = (1000 * 60 * t.Minute()) + (1000 * 60 * 60 * t.Hour())
			}

			if strings.Contains(div.Text, "Broadcast:") {
				anime.Broadcast = -1
				if strings.Contains(div.Text, "Sunday") {
					anime.Broadcast = 0
				} else if strings.Contains(div.Text, "Monday") {
					anime.Broadcast = 1
				} else if strings.Contains(div.Text, "Tuesday") {
					anime.Broadcast = 2
				} else if strings.Contains(div.Text, "Wednesday") {
					anime.Broadcast = 3
				} else if strings.Contains(div.Text, "Thursday") {
					anime.Broadcast = 4
				} else if strings.Contains(div.Text, "Friday") {
					anime.Broadcast = 5
				} else if strings.Contains(div.Text, "Saturday") {
					anime.Broadcast = 6
				}
			}
		})
	})

	collector.OnHTML("p[itemprop=\"description\"]", func(description *colly.HTMLElement) {
		var descriptionParts []string = strings.Split(description.Text, "\n\n")
		for _, descr := range descriptionParts {
			descr = strings.TrimSpace(descr)
			var descriptionData ScraperDescription = ScraperDescription{GetAnimeDescriptionID(conn, anime.Id, descr), descr, anime.Id, "", ""}
			animeDescription = append(animeDescription, descriptionData)
		}
	})

	collector.Visit(url)

	if len(anime.Id) > 0 {
		// insertType(conn, tx, animeType)
		// insertSeason(conn, tx, animeSeason)
		// insertAnimeDescription(conn, tx, animeDescription)
		// insertStudio(conn, tx, animeStudios)
		// insertGenre(conn, tx, animeGenres)
		// insertTheme(conn, tx, animeThemes)
		// insertAnime(conn, tx, anime)
		// insertAnimeStudio(conn, tx, join_anime_studios)
		// insertAnimeGenre(conn, tx, join_anime_genre)
		// insertAnimeTheme(conn, tx, join_anime_theme)
		return ScrapedAnime{
			Anime:             anime,
			AnimeSeason:       animeSeason,
			AnimeType:         animeType,
			AnimeDescription:  animeDescription,
			AnimeStudios:      animeStudios,
			AnimeGenre:        animeGenres,
			AnimeTheme:        animeThemes,
			Join_Anime_Studio: join_anime_studios,
			Join_Anime_Genre:  join_anime_genre,
			Join_Anime_Theme:  join_anime_theme,
		}, true
	}
	return ScrapedAnime{}, false
}
