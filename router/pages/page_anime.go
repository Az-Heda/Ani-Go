package pages

import (
	"AniGo/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type broadcastDay struct {
	Id   int
	Name string
}

func serveAnime(c *gin.Context) {
	var animeID string = c.Param("id")

	anime, err := db.SelectAnimeFromId(animeID)
	checkPanic(err)

	images, err := db.SelectImagesFromAnimeId(animeID)
	checkPanic(err)

	allStatus, err := db.SelectAllStatuses()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var title string
	if title = anime.Title; anime.AlternativeTitle.Valid {
		title = anime.AlternativeTitle.String
	}

	var descr []string = []string{}
	if anime.Description.Valid {
		descr = strings.Split(anime.Description.String, "\\n")
	}

	genres, err := db.SelectGenreFromAnimeId(animeID)
	checkPanic(err)

	themes, err := db.SelectThemeFromAnimeId(animeID)
	checkPanic(err)

	episodes, err := db.SelectEpisodeFromAnimeId(animeID)
	checkPanic(err)

	characters, err := db.SelectCharactersByIdWithDefaultImage(animeID)
	checkPanic(err)

	c.HTML(http.StatusOK, "anime.html", gin.H{
		"title":          title,
		"menu":           navbar,
		"activeMenuItem": "",
		"anime":          anime,
		"images":         images,
		"genres":         genres,
		"themes":         themes,
		"episodes":       episodes,
		"characters":     characters,
		"description":    descr,
		"statusList":     allStatus,
		"broadcastDays": []broadcastDay{
			{Id: 1, Name: "Monday"},
			{Id: 2, Name: "Tuesday"},
			{Id: 3, Name: "Wednesday"},
			{Id: 4, Name: "Thursday"},
			{Id: 5, Name: "Friday"},
			{Id: 6, Name: "Saturday"},
			{Id: 0, Name: "Sunday"},
			{Id: -1, Name: "Unknown"},
		},
	})
}
