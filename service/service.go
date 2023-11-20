package service

import (
	"carservice/infra/config"
	"carservice/infra/logger"
	"carservice/service/chain"
	"carservice/service/middleware"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Service struct {
	srv         *http.Server
	chainClient *chain.ChainClient
	eg          *gin.Engine
}

func NewService(eg *gin.Engine, chainClient *chain.ChainClient) *Service {

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GetServerConfig().Port),
		Handler: eg.Handler(),
	}
	return &Service{srv: srv, chainClient: chainClient, eg: eg}
}

func (s *Service) Run() {
	// cors
	apiGroup := s.eg.Group("api")
	apiGroup.Use(middleware.Cors())

	{
		// login
		apiGroup.POST("/user/login", s.Login)
		// register
		apiGroup.POST("/user/register", s.Register)
		// user info
		apiGroup.GET("/user/info", middleware.JWTAuth(), s.GetUserInfo)
		// set user info
		apiGroup.POST("/user/info", middleware.JWTAuth(), s.UpdateUserInfo)
		// charge score
		apiGroup.POST("/user/charge", middleware.JWTAuth(), s.ChargeScore)

		// high way
		apiGroup.POST("/highway/preview", middleware.JWTAuth(), s.PreviewHighWay)
		apiGroup.POST("/highway/start", middleware.JWTAuth(), s.StartHighWay)
		apiGroup.POST("/highway/end", middleware.JWTAuth(), s.EndHighWay)

		// charge
		apiGroup.POST("/charge/preview", middleware.JWTAuth(), s.PreviewCharge)
		apiGroup.POST("/charge/start", middleware.JWTAuth(), s.StartCharge)
		apiGroup.POST("/charge/end", middleware.JWTAuth(), s.EndCharge)

		// park
		apiGroup.POST("/park/preview", middleware.JWTAuth(), s.PreviewPark)
		apiGroup.POST("/park/start", middleware.JWTAuth(), s.StartPark)
		apiGroup.POST("/park/end", middleware.JWTAuth(), s.EndPark)

		// order
		apiGroup.POST("/order/pay", middleware.JWTAuth(), s.PayOrder)
		apiGroup.GET("/highway/orders", middleware.JWTAuth(), s.GetHighWayOrders)
		apiGroup.GET("/charge/orders", middleware.JWTAuth(), s.GetChargeOrders)
		apiGroup.GET("/park/orders", middleware.JWTAuth(), s.GetParkOrders)
		apiGroup.GET("/order/list", middleware.JWTAuth(), s.GetOrderList)

	}
	go func() {
		// 服务连接
		if err := s.srv.ListenAndServe(); err != nil {
			logger.Error("close server", err)
		}
	}()
}

func (s *Service) Shutdown() {
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
