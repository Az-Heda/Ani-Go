package pages

import (
	"AniGo/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func serveAnime(c *gin.Context) {
	var animeID string = c.Param("id")
	var images []string

	anime, err := db.SelectAnimeFromId(animeID)
	if err != nil {
		panic(err)
	}

	if anime.Image.Valid {
		images = strings.Split(anime.Image.String, "://:")
	}

	genres, err := db.SelectGenreFromAnimeId(animeID)
	if err != nil {
		panic(err)
	}

	episodes, err := db.SelectEpisodeFromAnimeId(animeID)
	if err != nil {
		panic(err)
	}

	c.HTML(http.StatusOK, "anime.html", gin.H{
		"title":          "Anime",
		"menu":           navbar,
		"activeMenuItem": "",
		"anime":          anime,
		"images":         images,
		"genres":         genres,
		"episodes":       episodes,
	})
}
