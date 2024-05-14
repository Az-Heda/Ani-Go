package pages

import (
	"net/http"

	db "AniGo/db"

	"github.com/gin-gonic/gin"
)

func remapAnimeAiring(allAnime []db.DB_Anime) map[string][]db.DB_Anime {
	var data map[string][]db.DB_Anime = map[string][]db.DB_Anime{}

	for _, a := range allAnime {
		var key string
		switch a.Broadcast {
		case -1:
			key = "Unknown"
		case 0:
			key = "Sunday"
		case 1:
			key = "Monday"
		case 2:
			key = "Tuesday"
		case 3:
			key = "Wednesday"
		case 4:
			key = "Thursday"
		case 5:
			key = "Friday"
		case 6:
			key = "Saturday"
		}
		if _, ok := data[key]; !ok {
			data[key] = []db.DB_Anime{}
		}

		data[key] = append(data[key], a)
	}

	return data
}

func serveIndex(c *gin.Context) {
	airingAnime, err := db.SelectAiringAnime()

	var animeDays map[string][]db.DB_Anime = remapAnimeAiring(airingAnime)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":          "Homepage",
		"menu":           navbar,
		"activeMenuItem": "Home",
		"airing":         animeDays,
		"image_size":     150,
		"days":           []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday", "Unknown"},
	})
}
