package pages

import (
	"encoding/json"
	"net/http"
	"time"

	db "AniGo/db"

	"github.com/gin-gonic/gin"
)

func remapAnimeAiring(allAnime []db.DB_Anime) map[string][]db.DB_Anime {
	var data map[string][]db.DB_Anime = map[string][]db.DB_Anime{}

	for _, a := range allAnime {
		var key string
		if a.CurrentStatus != 1 {
			continue
		}
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

func getDataFor_CurrentStautsPieChartData(allStatuses []db.DB_Status, allAnime []db.DB_Anime) string {
	var data map[string]int = map[string]int{}
	var statusMap map[int]string = map[int]string{}

	for _, s := range allStatuses {
		data[s.Name] = 0
		statusMap[s.Id] = s.Name
	}

	for _, a := range allAnime {
		data[statusMap[a.CurrentStatus]] += 1
	}
	b, _ := json.Marshal(data)
	return string(b)
}

func getDataFor_YearMonthHeatmap(allAnime []db.DB_Anime) string {
	var data map[int]map[string]int = map[int]map[string]int{}
	var months []string = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	for _, a := range allAnime {
		if a.Aired > 0 {
			t := time.Unix(0, a.Aired*1e6)
			if data[t.Year()] == nil {
				data[t.Year()] = map[string]int{}
				for _, m := range months {
					data[t.Year()][m] = 0
				}
			}

			data[t.Year()][t.Month().String()[:3]] += 1
		}
	}

	b, _ := json.Marshal(data)
	return string(b)
}

func serveIndex(c *gin.Context) {
	// airingAnime, err := db.SelectAiringAnime()
	allAnime, err := db.SelectAllAnime()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	allStatuses, err := db.SelectAllStatuses()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var animeDays map[string][]db.DB_Anime = remapAnimeAiring(allAnime)
	var chart_status string = getDataFor_CurrentStautsPieChartData(allStatuses, allAnime)
	var chart_heatmap_ymcounter = getDataFor_YearMonthHeatmap(allAnime)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":                           "Homepage",
		"menu":                            navbar,
		"activeMenuItem":                  "Home",
		"airing":                          animeDays,
		"image_size":                      150,
		"days":                            []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday", "Unknown"},
		"charts_pie_CurrentStatusCounter": chart_status,
		"chart_heatmap_YearMonthCounter":  chart_heatmap_ymcounter,
	})
}
