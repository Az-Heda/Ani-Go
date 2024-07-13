package pages

import (
	"net/http"

	db "AniGo/db"

	"github.com/gin-gonic/gin"
)

func serveThemeList(c *gin.Context) {

	allThemes, err := db.SelectAllThemesWithCount()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "theme_list.html", gin.H{
		"title":          "Themes list",
		"menu":           navbar,
		"activeMenuItem": "Theme list",
		"themeList":      allThemes,
	})
}
