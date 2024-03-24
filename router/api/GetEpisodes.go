package api

import (
	"AniGo/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetEpisodeEndpoints(g *gin.RouterGroup) {
	r := g.Group("/episode")

	r.GET("/all", serveEpisodeAll)
	r.GET("/id/:id", serveEpisodeWithId)
	r.GET("/anime/:id", ServeAnimeEpisodes)
}

func serveEpisodeAll(c *gin.Context) {
	data, err := db.SelectAllEpisodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func serveEpisodeWithId(c *gin.Context) {
	var id string = c.Param("id")
	data, err := db.SelectEpisodeFromId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func ServeAnimeEpisodes(c *gin.Context) {
	var id string = c.Param("id")
	data, err := db.SelectEpisodeFromAnimeId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}
