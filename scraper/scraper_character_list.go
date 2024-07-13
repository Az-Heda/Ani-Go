package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/jmoiron/sqlx"
)

func parseAnimeCharacterList(animeID string, url string, conn *sqlx.DB, tx *sqlx.Tx) (ScrapedAnimeCharacters, bool) {
	var characters []ScraperCharacter
	var descriptions []ScraperDescription
	var images []ScraperImage
	// var join_character_description []ScraperCharacterDescription
	var join_character_image []ScraperCharacterImage
	var join_anime_characters []ScraperAnimeCharacter

	collector := colly.NewCollector()

	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println(url)
		panic(e)
	})

	collector.OnHTML(".anime-character-container table", func(table *colly.HTMLElement) {
		table.ForEach("tr", func(idx int, tr *colly.HTMLElement) {
			tr.ForEach("td:nth-child(2) a", func(idx int, a *colly.HTMLElement) {
				a.ForEach("h3.h3_character_name", func(idx int, title *colly.HTMLElement) {
					var href string = a.Attr("href")
					var character ScraperCharacter = ScraperCharacter{GetCharacterID(conn, href), title.Text, href, ScraperImage{}}

					join_anime_characters = append(join_anime_characters, ScraperAnimeCharacter{animeID, character.Id})

					var extra parseAnimeCharacterType = parseAnimeCharacter(character.Url, character.Id, conn)

					for _, part := range strings.Split(extra.description, "\n\n") {
						part = strings.TrimSpace(part)
						var descr ScraperDescription = ScraperDescription{GetCharacterDescriptionID(conn, character.Id, part), part, "", "", character.Id}
						descriptions = append(descriptions, descr)
						// join_character_description = append(join_character_description, ScraperCharacterDescription{character.Id, descr.Id})
					}

					var firstImage ScraperImage = ScraperImage{"", ""}

					for _, url := range extra.pics {
						var img ScraperImage = ScraperImage{GetImageID(conn, url), url}
						if firstImage.Id == "" {
							firstImage = img
						}
						images = append(images, img)
						join_character_image = append(join_character_image, ScraperCharacterImage{character.Id, img.Id})
					}

					character.Image = firstImage
					characters = append(characters, character)
				})
			})
		})
	})

	collector.Visit(url)

	// insertCharacter(conn, tx, characters)
	// insertCharacterDescription(conn, tx, descriptions)
	// insertImage(conn, tx, images)

	// insertCharacterDescription(conn, tx, join_character_description)
	// insertAnimeCharacter(conn, tx, join_anime_characters)
	// insertCharacterImage(conn, tx, join_character_image)
	return ScrapedAnimeCharacters{
		AnimeCharacters:       characters,
		AnimeDescriptions:     descriptions,
		CharacterImages:       images,
		Join_Character_Image:  join_character_image,
		Join_Anime_Characters: join_anime_characters,
	}, true
}
