package api

import (
	"AniGo/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetStudioEndpoints(g *gin.RouterGroup) {
	r := g.Group("/studio")

	r.GET("/all", serveStudioAll)
	r.GET("/id/:id", serveStudioWithId)
	r.GET("/name/:name", serveStudioWithName)
}

func serveStudioAll(c *gin.Context) {
	data, err := db.SelectAllStudios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func serveStudioWithId(c *gin.Context) {
	var id string = c.Param("id")
	data, err := db.SelectStudioFromId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func serveStudioWithName(c *gin.Context) {
	var name string = c.Param("name")
	data, err := db.SelectStudioFromName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}
