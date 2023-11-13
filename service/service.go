package service

import (
	"carservice/infra/config"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Service struct {
	srv *http.Server
}

func NewService(eg *gin.Engine) *Service {

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GetServerConfig().Port),
		Handler: eg.Handler(),
	}
	return &Service{srv: srv}
}

func (s *Service) Run() {
	go func() {
		// 服务连接
		if err := s.srv.ListenAndServe(); err != nil {
			panic(err)
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
