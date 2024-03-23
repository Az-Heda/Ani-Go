package api

import (
	"github.com/gin-gonic/gin"
)

func Init(eng *gin.Engine) {
	api := eng.Group("/api")

	api.GET("/anime/:id", getAnimeFromID)
}
