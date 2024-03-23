package router

import (
	"fmt"
	"net/http"

	api "AniGo/router/api"

	"github.com/gin-gonic/gin"
)

var HOST string = "127.0.0.1"
var PORT int = 8770
var URL string = HOST + ":" + fmt.Sprintf("%d", PORT)

func Init() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.html")

	api.Init(router)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Homepage",
		})
	})

	fmt.Println("Online at ", "http://"+URL)
	router.Run(URL)
}
