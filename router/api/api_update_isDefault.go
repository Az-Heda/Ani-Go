package api

import (
	"net/http"

	db "AniGo/db"

	"github.com/gin-gonic/gin"
)

func SetIsDefaultEndpoints(g *gin.RouterGroup) {
	r := g.Group("/is-default")

	r.GET("/anime/:anime_id/:image_id", serveChangeAnimeIsDefault)
	r.GET("/character/:character_id/:image_id", serveChangeCharacterIsDefault)
}

func serveChangeAnimeIsDefault(c *gin.Context) {
	var anime_id string = c.Param("anime_id")
	var image_id string = c.Param("image_id")
	conn, err := db.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	_, err = conn.Query(`UPDATE Anime SET Image_Id = ? WHERE Id = ?`, image_id, anime_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
}

func serveChangeCharacterIsDefault(c *gin.Context) {
	var character_id string = c.Param("character_id")
	var image_id string = c.Param("image_id")
	conn, err := db.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	_, err = conn.Query(`UPDATE Character SET Image_Id = ? WHERE Id = ?`, image_id, character_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
}
