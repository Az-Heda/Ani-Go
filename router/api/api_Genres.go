package api

import (
	"AniGo/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetGenreEndpoints(g *gin.RouterGroup) {
	r := g.Group("/genre")

	r.GET("/all", serveGenreAll)
	r.GET("/id/:id", serveGenreWithId)
	r.GET("/name/:name", serveGenreWithName)
}

func serveGenreAll(c *gin.Context) {
	data, err := db.SelectAllGenres()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func serveGenreWithId(c *gin.Context) {
	var id string = c.Param("id")
	data, err := db.SelectGenreFromId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func serveGenreWithName(c *gin.Context) {
	var name string = c.Param("name")
	data, err := db.SelectGenreFromName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}
