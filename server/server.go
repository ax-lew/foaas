package server

import (
	"github.com/ax-lew/foaas/domain/handler"
	"github.com/ax-lew/foaas/domain/service"
	"github.com/ax-lew/foaas/logger"
	"github.com/ax-lew/foaas/ratelimiter"
	"github.com/gin-gonic/gin"
)

type Server struct {
	rateLimiter ratelimiter.RateLimiter
	service     service.Service
}

func NewServer(rateLimiter ratelimiter.RateLimiter, service service.Service) *Server {
	return &Server{rateLimiter: rateLimiter, service: service}
}

func (s *Server) Run() {
	handler := handler.NewHandler(s.rateLimiter, s.service)

	engine := gin.Default()
	engine.GET("/message", handler.Handle)
	if err := engine.Run(); err != nil {
		logger.Logger.Errorf("Server stopped: %s", err.Error())
	}
}
