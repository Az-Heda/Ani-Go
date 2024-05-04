package api

import (
	"net/http"

	db "AniGo/db"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
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

	var tx *sqlx.Tx = conn.MustBegin()
	db.UpdateAnimeImage_RemoveIsDefault(tx, anime_id)
	db.UpdateAnimeImage_IsDefault(tx, anime_id, image_id, true)
	err = tx.Commit()
	if err != nil {
		panic(err)
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

	var tx *sqlx.Tx = conn.MustBegin()
	db.UpdateCharacterImage_RemoveIsDefault(tx, character_id)
	db.UpdateCharacterImage_IsDefault(tx, character_id, image_id, true)
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
