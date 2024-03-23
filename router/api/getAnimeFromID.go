package api

import (
	dbintegration "AniGo/db-integration"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var animeQuery string = `
SELECT 	a.Id as "Id",
		a.Title as "Title",
		a.AlternativeTitle as "AlternativeTitle",
		a.Aired as "Aired",
		a.Duration as "Duration",
		a.CurrentStatus as "CurrentStatus",
		s.Season as "Season",
		t.Name as "Type"
FROM Anime a
LEFT JOIN Seasons s ON s.Id = a.Season_ID
LEFT JOIN Types t ON t.Id = a.Type_ID 
WHERE a.Id = ?
`

func getAnimeFromID(c *gin.Context) {
	var animeID string = c.Param("id")

	db, err := sqlx.Connect("sqlite", dbintegration.DatabaseName)
	if err != nil {
		log.Fatalln(err)
	}

	var data []anime

	err = db.Select(&data, animeQuery, animeID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}
