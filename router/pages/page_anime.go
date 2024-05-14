package pages

import (
	"AniGo/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}

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

	genres, err := db.SelectGenreFromAnimeId(animeID)
	checkPanic(err)

	episodes, err := db.SelectEpisodeFromAnimeId(animeID)
	checkPanic(err)

	characters, err := db.SelectCharactersByIdWithDefaultImage(animeID)
	checkPanic(err)

	c.HTML(http.StatusOK, "anime.html", gin.H{
		"title":          "Anime",
		"menu":           navbar,
		"activeMenuItem": "",
		"anime":          anime,
		"images":         outImages,
		"genres":         genres,
		"episodes":       episodes,
		"characters":     characters,
	})
}
