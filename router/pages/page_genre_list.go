package pages

import (
	"net/http"

	db "AniGo/db"

	"github.com/gin-gonic/gin"
)

func serveGenreList(c *gin.Context) {

	allGenres, err := db.SelectAllGenresWithCount()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "genre_list.html", gin.H{
		"title":          "Genres list",
		"menu":           navbar,
		"activeMenuItem": "Genre list",
		"genreList":      allGenres,
	})
}
