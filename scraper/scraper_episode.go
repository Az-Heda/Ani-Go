package scraper

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/jmoiron/sqlx"
)

type parseAnimeEpisodeType struct {
	duration    int64
	description string
}

func parseAnimeEpisode(url string, episodeID string, conn *sqlx.DB) parseAnimeEpisodeType {
	var returnType parseAnimeEpisodeType = parseAnimeEpisodeType{}
	collector := colly.NewCollector()

	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("parseAnimeEpisode", e.Error(), url)
		panic(e)
	})

	collector.OnHTML("td[valign=\"top\"]:nth-child(2) div.pt8.pb8", func(div *colly.HTMLElement) {
		returnType.description = strings.TrimSpace(div.Text)
		returnType.description = strings.Replace(returnType.description, "Synopsis", "", 1)
		returnType.description = strings.TrimSpace(returnType.description)
		if strings.Contains(returnType.description, "Sorry, this episode doesn't seem to have a synopsis yet") {
			returnType.description = ""
		} else {
			var re = regexp.MustCompile(`(?m)\s*\(Source: .*?\)`)
			returnType.description = re.ReplaceAllString(returnType.description, " ")
			returnType.description = strings.TrimSpace(returnType.description)
		}
	})

	collector.OnHTML("td[valign=\"top\"] div.ar", func(div *colly.HTMLElement) {
		returnType.duration = 0
		var re *regexp.Regexp = regexp.MustCompile(`(?m)[0-9]{2}:[0-9]{2}:[0-9]{2}`)
		for _, match := range re.FindAllString(div.Text, -1) {
			t, err := time.Parse("15:04:05", match)
			if err != nil {
				returnType.duration = 0
				return
			}
			returnType.duration += 1000 * int64(t.Second())
			returnType.duration += 1000 * 60 * int64(t.Minute())
			returnType.duration += 1000 * 60 * 60 * int64(t.Hour())
		}
	})

	collector.Visit(url)
	return returnType
}
