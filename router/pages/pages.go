package pages

import "github.com/gin-gonic/gin"

var navbar [][]string = [][]string{
	{"Home", "/"},
	{"Anime list", "/list/anime"},
	// {"Dashboard", "/dashboard"},
}

func Init(r *gin.Engine) {
	r.GET("/", serveIndex)
	r.GET("/list/anime", serveAnimeList)

	r.GET("/anime/:id", serveAnime)
}
