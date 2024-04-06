package api

import (
	"AniGo/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetCharacterEndpoints(g *gin.RouterGroup) {
	r := g.Group("/character")

	r.GET("/all", serveCharacterAll)
	r.GET("/id/:id", serveCharacterWithId)
}

func serveCharacterAll(c *gin.Context) {
	data, err := db.SelectAllCharacters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func serveCharacterWithId(c *gin.Context) {
	var id string = c.Param("id")
	data, err := db.SelectCharacterFromId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}
