package pages

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	r.GET("/", serveIndex)
}
