package pages

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var navbar [][]string = [][]string{
	{"Home", "/"},
	{"Anime list", "/list/anime"},
	{"Genre list", "/list/genres"},
	{"Theme list", "/list/themes"},
	// {"Dashboard", "/dashboard"},
}

func checkPanic(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Init(r *gin.Engine) {
	r.GET(navbar[0][1], serveIndex)
	r.GET(navbar[1][1], serveAnimeList)
	r.GET(navbar[2][1], serveGenreList)
	r.GET(navbar[3][1], serveThemeList)

	r.GET("/anime/:id", serveAnime)
	r.GET("/character/:id", serveCharacter)
	r.GET("/genre/:id", serveGenre)
	r.GET("/theme/:id", serveTheme)
}
