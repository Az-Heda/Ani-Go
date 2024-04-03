package pages

import (
	"net/http"

	db "AniGo/db"

	"github.com/gin-gonic/gin"
)

func serveAnimeList(c *gin.Context) {

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

	c.HTML(http.StatusOK, "anime_list.html", gin.H{
		"title":          "Anime list",
		"menu":           navbar,
		"activeMenuItem": "Anime list",
		"animeList":      allAnime,
		"statusList":     allStatus,
	})
}
