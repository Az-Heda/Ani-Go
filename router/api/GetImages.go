package api

import (
	"AniGo/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetImageEndpoints(g *gin.RouterGroup) {
	r := g.Group("/image")

	r.GET("/all", serveImageAll)
	r.GET("/id/:id", serveImageWithId)
}

func serveImageAll(c *gin.Context) {
	data, err := db.SelectAllImages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func serveImageWithId(c *gin.Context) {
	var id string = c.Param("id")
	data, err := db.SelectImageFromId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}
