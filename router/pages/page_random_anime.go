package pages

import (
	"AniGo/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func serveRandomAnime(c *gin.Context) {
	anime, err := db.SelectRandomAnime()
	checkPanic(err)

	c.Redirect(http.StatusFound, "/anime/"+anime.Id)
}
