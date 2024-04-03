package pages

import (
	"AniGo/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func serveAnime(c *gin.Context) {
	var animeID string = c.Param("id")

	anime, err := db.SelectAnimeFromId(animeID)
	if err != nil {
		panic(err)
	}

	c.HTML(http.StatusOK, "anime.html", gin.H{
		"title":          "Anime",
		"menu":           navbar,
		"activeMenuItem": "Anime list",
		"anime":          anime,
	})
}
