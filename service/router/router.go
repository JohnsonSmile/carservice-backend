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

		// set user info
		apiGroup.POST("/user/info", middleware.JWTAuth(), handler.UpdateUserInfo)

		// high way
		apiGroup.POST("/highway/preview", middleware.JWTAuth(), handler.PreviewHighWay)
		apiGroup.POST("/highway/start", middleware.JWTAuth(), handler.StartHighWay)
		apiGroup.POST("/highway/end", middleware.JWTAuth(), handler.EndHighWay)

		// charge
		apiGroup.POST("/charge/preview", middleware.JWTAuth(), handler.PreviewCharge)
		apiGroup.POST("/charge/start", middleware.JWTAuth(), handler.StartCharge)
		apiGroup.POST("/charge/end", middleware.JWTAuth(), handler.EndCharge)

		// park
		apiGroup.POST("/park/preview", middleware.JWTAuth(), handler.PreviewPark)
		apiGroup.POST("/park/start", middleware.JWTAuth(), handler.StartPark)
		apiGroup.POST("/park/end", middleware.JWTAuth(), handler.EndPark)

		// order
		apiGroup.POST("/order/pay", middleware.JWTAuth(), handler.PayOrder)
		apiGroup.GET("/highway/orders", middleware.JWTAuth(), handler.GetHighWayOrders)
		apiGroup.GET("/charge/orders", middleware.JWTAuth(), handler.GetChargeOrders)

	}

	return eg
}
