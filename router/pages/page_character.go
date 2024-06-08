package pages

import (
	"AniGo/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func serveCharacter(c *gin.Context) {
	var characterID string = c.Param("id")

	var outImages []db.DB_Image

	character, err := db.SelectCharacterFromID(characterID)
	checkPanic(err)

	if len(character.Image) > 0 {
		for _, img := range character.Image {
			queryResult, err := db.SelectImageIdFromUrl(img)
			if err == nil {
				outImages = append(outImages, queryResult)
			}
		}
	}

	var descr []string = []string{}
	if character.Description.Valid {
		descr = strings.Split(character.Description.String, "\\n")
	}


	c.HTML(http.StatusOK, "character.html", gin.H{
		"title":          character.Name,
		"menu":           navbar,
		"activeMenuItem": "",
		"character":      character,
		"images":         outImages,
		"description": descr,
	})
}
