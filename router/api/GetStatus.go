package api

import (
	"AniGo/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetStatusEndpoints(g *gin.RouterGroup) {
	r := g.Group("/status")

	r.GET("/all", serveStatusAll)
	r.GET("/id/:id", serveStatusWithId)
	r.GET("/name/:name", serveStatusWithName)
}

func serveStatusAll(c *gin.Context) {
	data, err := db.SelectAllStatuses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func serveStatusWithId(c *gin.Context) {
	var id string = c.Param("id")
	data, err := db.SelectStatusFromId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func serveStatusWithName(c *gin.Context) {
	var name string = c.Param("name")
	data, err := db.SelectStatusFromName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}
