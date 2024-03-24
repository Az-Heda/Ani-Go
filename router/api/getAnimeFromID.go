package api

import (
	"log"
	"net/http"

	db "AniGo/db"

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

	conn, err := sqlx.Connect("sqlite", db.DatabaseName)
	if err != nil {
		log.Fatalln(err)
	}

	var data []anime

	rows, err := conn.Query(animeQuery, animeID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	for rows.Next() {
		var d anime
		if err = rows.Scan(&d.Id, &d.Title, &d.AlternativeTitle, &d.Aired, &d.Duration, &d.CurrentStatus, &d.Season, &d.Type); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		data = append(data, d)
	}

	c.JSON(http.StatusOK, data)
}
