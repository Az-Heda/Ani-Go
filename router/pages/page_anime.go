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
	var images []string

	anime, err := db.SelectAnimeFromId(animeID)
	checkPanic(err)

	if anime.Image.Valid {
		images = strings.Split(anime.Image.String, "://:")
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
		"images":         images,
		"genres":         genres,
		"episodes":       episodes,
		"characters":     characters,
	})
}
