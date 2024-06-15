package pages

import (
	"AniGo/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func serveCharacter(c *gin.Context) {
	var characterID string = c.Param("id")

	character, err := db.SelectCharacterFromID(characterID)
	checkPanic(err)

	images, err := db.SelectImagesFromCharacterId(characterID)
	checkPanic(err)

	var descr []string = []string{}
	if character.Description.Valid {
		descr = strings.Split(character.Description.String, "\\n")
	}

	c.HTML(http.StatusOK, "character.html", gin.H{
		"title":          character.Name,
		"menu":           navbar,
		"activeMenuItem": "",
		"character":      character,
		"images":         images,
		"description":    descr,
	})
}
