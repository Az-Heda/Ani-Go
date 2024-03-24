package pages

import (
	"net/http"

	db "AniGo/db"

	"github.com/gin-gonic/gin"
)

func serveIndex(c *gin.Context) {
	allAnime, err := db.SelectAllAnime()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Homepage",
		"anime": allAnime,
	})
}
