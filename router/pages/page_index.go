package pages

import (
	"net/http"

	db "AniGo/db"

	"github.com/gin-gonic/gin"
)

func imageHandler(n int64, status int64) []string {
	randomImages, err := db.SelectRandomNImages(n, status)
	if err != nil {
		panic(err)
	}
	return randomImages
}

func serveIndex(c *gin.Context) {
	var n int64 = 10
	allAnime, err := db.SelectAllAnime()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":          "Homepage",
		"anime":          allAnime,
		"menu":           navbar,
		"activeMenuItem": "Home",
		"randomImages_0": imageHandler(n, 0),
		"randomImages_1": imageHandler(n, 1),
		"randomImages_2": imageHandler(n, 2),
		"randomImages_3": imageHandler(n, 3),
	})
}
