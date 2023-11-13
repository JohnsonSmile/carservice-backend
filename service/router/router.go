package router

import (
	"carservice/service/handler"
	"carservice/service/middleware"

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

	// cors
	apiGroup := eg.Group("api")
	apiGroup.Use(middleware.Cors())

	{
		// login
		apiGroup.POST("/user/login", handler.Login)
		// register
		apiGroup.POST("/user/register", handler.Register)

		// user info
		apiGroup.GET("/user/info", middleware.JWTAuth(), handler.GetUserInfo)

		// order
		apiGroup.POST("/order/create", middleware.JWTAuth(), handler.CreateOrder)

	}

	return eg
}
