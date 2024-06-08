package pages

import (
	"AniGo/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func serveAnime(c *gin.Context) {
	var animeID string = c.Param("id")

	var outImages []db.DB_Image

	anime, err := db.SelectAnimeFromId(animeID)
	checkPanic(err)

	if anime.Image.Valid {
		for _, img := range strings.Split(anime.Image.String, "://:") {
			queryResult, err := db.SelectImageIdFromUrl(img)
			if err == nil {
				outImages = append(outImages, queryResult)
			}
		}
	}

	var title string
	if title = anime.Title; anime.AlternativeTitle.Valid {
		title = anime.AlternativeTitle.String
	}
	
	var descr []string = []string{}
	if anime.Description.Valid {
		descr = strings.Split(anime.Description.String, "\\n")
	}

	genres, err := db.SelectGenreFromAnimeId(animeID)
	checkPanic(err)

	themes, err := db.SelectThemeFromAnimeId(animeID)
	checkPanic(err)

	episodes, err := db.SelectEpisodeFromAnimeId(animeID)
	checkPanic(err)

	characters, err := db.SelectCharactersByIdWithDefaultImage(animeID)
	checkPanic(err)

	c.HTML(http.StatusOK, "anime.html", gin.H{
		"title":          title,
		"menu":           navbar,
		"activeMenuItem": "",
		"anime":          anime,
		"images":         outImages,
		"genres":         genres,
		"themes":         themes,
		"episodes":       episodes,
		"characters":     characters,
		"description":    descr,
	})
}
