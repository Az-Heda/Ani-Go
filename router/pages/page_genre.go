package pages

import (
	"AniGo/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func serveGenre(c *gin.Context) {
	var genreID string = c.Param("id")

	genre, err := db.SelectGenreFromId(genreID)
	checkPanic(err)

	animes, err := db.SelectAllAnimeFromGenreID(genre)
	checkPanic(err)

	statuses, err := db.SelectAllStatuses()
	checkPanic(err)

	c.HTML(http.StatusOK, "genre.html", gin.H{
		"title":          genre.Name,
		"menu":           navbar,
		"activeMenuItem": "",
		"genre":          genre,
		"animeList":      animes,
		"statuses":       statuses,
	})
}
