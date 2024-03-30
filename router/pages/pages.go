package pages

import "github.com/gin-gonic/gin"

var navbar [][]string = [][]string{
	{"Home", "/"},
	{"Anime", "/anime"},
	{"Characters", "/character"},
	{"Dashboard", "/dashboard"},
}

func Init(r *gin.Engine) {
	r.GET("/", serveIndex)
	r.GET("/anime", serveAnime)
}
