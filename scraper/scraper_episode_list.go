package scraper

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gocolly/colly"
	"github.com/jmoiron/sqlx"
)

func parseAnimeEpisodeList(baseUrl string, url string, conn *sqlx.DB, tx *sqlx.Tx) {
	var animeID = GetAnimeID(conn, baseUrl)
	var episodes []ScraperEpisode
	var descriptions []ScraperDescription

	collector := colly.NewCollector()

	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("parseAnimeEpisodeList", e.Error(), url)
		if e.Error() == "Not Found" {
			return
		}
		panic(e)
	})

	collector.OnHTML("table.episode_list", func(table *colly.HTMLElement) {
		table.ForEach("tbody > tr", func(idx int, tr *colly.HTMLElement) {
			var episode ScraperEpisode
			episode.Anime_ID = animeID
			tr.ForEach("td.episode-number", func(idx int, td *colly.HTMLElement) {
				episodeNumber, err := strconv.Atoi(td.Text)
				if err != nil {
					episode.Number = idx
				} else {
					episode.Number = episodeNumber
				}
				episode.Id = GetEpisodeID(conn, episode.Anime_ID, episode.Number)
			})
			tr.ForEach("td.episode-title", func(idx int, td *colly.HTMLElement) {
				td.ForEach("a", func(idx int, a *colly.HTMLElement) {
					episode.Title = a.Text
				})
			})
			tr.ForEach("td.episode-aired", func(idx int, td *colly.HTMLElement) {
				t, err := time.Parse("Jan 2, 2006", td.Text)
				if err != nil {
					episode.Aired = 0
					return
				}
				episode.Aired = t.Unix() * 1000
			})

			var episodeURL = url + "/" + fmt.Sprint(episode.Number)
			var d = parseAnimeEpisode(episodeURL, episode.Id, conn)
			if len(d.description) > 0 {
				var descr ScraperDescription = ScraperDescription{GetEpisodeDescriptionID(conn, episode.Id, d.description), d.description, "", episode.Id, ""}
				descriptions = append(descriptions, descr)
			}
			if d.duration > 0 {
				episode.Duration = d.duration
			}

			episodes = append(episodes, episode)
		})
	})

	collector.Visit(url)

	insertEpisode(conn, tx, episodes)
	insertEpisodeDescription(conn, tx, descriptions)
}
