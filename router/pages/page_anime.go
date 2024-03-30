package pages

import (
	"net/http"

	db "AniGo/db"

	"github.com/gin-gonic/gin"
)

func serveAnime(c *gin.Context) {

	allAnime, err := db.SelectAllAnime()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	allStatus, err := db.SelectAllStatuses()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "anime.html", gin.H{
		"title":          "Anime list",
		"menu":           navbar,
		"activeMenuItem": "Anime",
		"animeList":      allAnime,
		"statusList":     allStatus,
	})
}
