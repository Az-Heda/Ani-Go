package pages

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var navbar [][]string = [][]string{
	{"Home", "/"},
	{"Anime list", "/list/anime"},
	// {"Dashboard", "/dashboard"},
}

func checkPanic(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Init(r *gin.Engine) {
	r.GET("/", serveIndex)
	r.GET("/list/anime", serveAnimeList)

	r.GET("/anime/:id", serveAnime)
	r.GET("/character/:id", serveCharacter)
}
