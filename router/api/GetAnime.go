package api

import (
	"AniGo/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetAnimeEndpoints(g *gin.RouterGroup) {
	r := g.Group("/anime")

	r.GET("/all", serveAnimeAll)
	r.GET("/id/:id", serveAnimeWithId)
	r.GET("/search/:name", serveAnimeWithName)
}

func serveAnimeAll(c *gin.Context) {
	data, err := db.SelectAllAnime()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func serveAnimeWithId(c *gin.Context) {
	var id string = c.Param("id")
	data, err := db.SelectAnimeFromId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func serveAnimeWithName(c *gin.Context) {
	var name string = c.Param("name")
	data, err := db.SelectAnimeFromPartName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}
