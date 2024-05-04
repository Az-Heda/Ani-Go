package scraper

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/jmoiron/sqlx"
)

type parseAnimeCharacterType struct {
	description string
	pics        []string
}

func parseAnimeCharacter(url string, characterID string, conn *sqlx.DB) parseAnimeCharacterType {
	var returnType parseAnimeCharacterType = parseAnimeCharacterType{}

	collector := colly.NewCollector()

	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println(url)
		panic(e)
	})

	collector.OnHTML("#horiznav_nav > ul > li:nth-child(3) > a", func(picUrl *colly.HTMLElement) {
		returnType.pics = parseAnimeCharacterPics(picUrl.Attr("href"), characterID, conn)
	})

	collector.OnHTML("#content > table > tbody > tr > td:nth-child(2)", func(descr *colly.HTMLElement) {
		returnType.description = descr.Text
		descr.ForEach("div,h2,table,script", func(idx int, item *colly.HTMLElement) {
			returnType.description = strings.Replace(returnType.description, item.Text, "", 1)
		})
		returnType.description = strings.ReplaceAll(returnType.description, "No voice actors have been added to this character. Help improve our database by searching for a voice actor, and adding this character to their roles here.", "")
	})

	collector.Visit(url)
	returnType.description = strings.TrimSpace(returnType.description)

	var re = regexp.MustCompile(`(?m)\(Source:.*?\)`)
	for _, match := range re.FindAllString(returnType.description, -1) {
		var index int = strings.Index(returnType.description, match)
		if index >= 0 {
			returnType.description = returnType.description[:index]
		}
	}
	returnType.description = strings.TrimSpace(returnType.description)

	return returnType
}
