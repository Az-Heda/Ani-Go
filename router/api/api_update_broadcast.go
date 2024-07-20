package api

import (
	"net/http"
	"strconv"

	db "AniGo/db"

	"github.com/gin-gonic/gin"
)

func SetBroadcastEndpoints(g *gin.RouterGroup) {
	r := g.Group("/broadcast")

	r.PATCH("/:anime_id/:day_id", serveChangeAnimeBroadcast)
}

func serveChangeAnimeBroadcast(c *gin.Context) {
	var anime_id string = c.Param("anime_id")
	var day_id string = c.Param("day_id")
	conn, err := db.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	dayValue, err := strconv.ParseInt(day_id, 10, 8)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if dayValue >= -1 && dayValue < 7 { // 0-6 Sun-Sat ; -1: Unknown

		_, err = conn.Query(`UPDATE Anime SET Broadcast = ? WHERE Id = ?`, day_id, anime_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		return
	}
	c.String(http.StatusBadRequest, "%v", "Invalid parameter")
}
