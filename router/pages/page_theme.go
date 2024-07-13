package pages

import (
	"AniGo/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func serveTheme(c *gin.Context) {
	var themeID string = c.Param("id")

	theme, err := db.SelectThemeFromId(themeID)
	checkPanic(err)

	animes, err := db.SelectAllAnimeFromThemeID(theme)
	checkPanic(err)

	statuses, err := db.SelectAllStatuses()
	checkPanic(err)

	c.HTML(http.StatusOK, "theme.html", gin.H{
		"title":          theme.Name,
		"menu":           navbar,
		"activeMenuItem": "",
		"theme":          theme,
		"animeList":      animes,
		"statuses":       statuses,
	})
}
