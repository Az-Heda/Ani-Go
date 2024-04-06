package api

import (
	"github.com/gin-gonic/gin"
)

func Init(eng *gin.Engine) {
	api := eng.Group("/api")
	// SetGenreEndpoints(api)
	// SetStudioEndpoints(api)
	// SetTypeEndpoints(api)
	SetStatusEndpoints(api)
	// SetSeasonEndpoints(api)
	SetImageEndpoints(api)
	// SetEpisodeEndpoints(api)
	// SetDescriptionEndpoints(api)
	// SetCharacterEndpoints(api)
	SetAnimeEndpoints(api)
}
