package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize(isDebug bool) *gin.Engine {
	// gin debug mode
	if isDebug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	eg := gin.Default()
	return eg
}
