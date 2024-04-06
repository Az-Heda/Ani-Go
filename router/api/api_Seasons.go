package api

import (
	"AniGo/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetSeasonEndpoints(g *gin.RouterGroup) {
	r := g.Group("/season")

	r.GET("/all", serveSeasonAll)
	r.GET("/id/:id", serveSeasonWithId)
	r.GET("/name/:name", serveSeasonWithName)
}

func serveSeasonAll(c *gin.Context) {
	data, err := db.SelectAllSeasons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func serveSeasonWithId(c *gin.Context) {
	var id string = c.Param("id")
	data, err := db.SelectSeasonFromId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func serveSeasonWithName(c *gin.Context) {
	var name string = c.Param("name")
	data, err := db.SelectSeasonFromName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}
