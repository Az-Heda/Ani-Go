package api

import (
	"AniGo/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetDescriptionEndpoints(g *gin.RouterGroup) {
	r := g.Group("/description")

	r.GET("/all", serveDescriptionAll)
	r.GET("/id/:id", serveDescriptionWithId)
}

func serveDescriptionAll(c *gin.Context) {
	data, err := db.SelectAllDescriptions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func serveDescriptionWithId(c *gin.Context) {
	var id string = c.Param("id")
	data, err := db.SelectDescriptionFromId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}
