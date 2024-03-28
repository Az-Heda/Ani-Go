package router

import (
	"fmt"

	api "AniGo/router/api"
	pages "AniGo/router/pages"

	"github.com/gin-gonic/gin"
)

var HOST string = "127.0.0.1"
var PORT int = 8770
var URL string = HOST + ":" + fmt.Sprintf("%d", PORT)

func Init() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.html")

	router.Static("/static", "./static")

	api.Init(router)
	pages.Init(router)

	fmt.Println("Online at ", "http://"+URL)
	router.Run(URL)
}
