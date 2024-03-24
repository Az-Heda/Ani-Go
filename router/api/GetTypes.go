package api

import (
	"AniGo/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetTypeEndpoints(g *gin.RouterGroup) {
	r := g.Group("/type")

	r.GET("/all", serveTypeAll)
	r.GET("/id/:id", serveTypeWithId)
	r.GET("/name/:name", serveTypeWithName)
}

func serveTypeAll(c *gin.Context) {
	data, err := db.SelectAllTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func serveTypeWithId(c *gin.Context) {
	var id string = c.Param("id")
	data, err := db.SelectTypeFromId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func serveTypeWithName(c *gin.Context) {
	var name string = c.Param("name")
	data, err := db.SelectTypeFromName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}
