package handler

import (
	"github.com/ax-lew/foaas/domain/service"
	"github.com/ax-lew/foaas/logger"
	"github.com/ax-lew/foaas/ratelimiter"
	"github.com/gin-gonic/gin"
	"net/http"
)

const userIDHeader = "User-id"

type Handler struct {
	rateLimiter ratelimiter.RateLimiter
	service     service.Service
}

func NewHandler(rateLimiter ratelimiter.RateLimiter, service service.Service) *Handler {
	return &Handler{
		rateLimiter: rateLimiter,
		service:     service,
	}
}

func (h *Handler) Handle(ctx *gin.Context) {
	userID := ctx.GetHeader(userIDHeader)
	if len(userID) == 0 {
		logger.Logger.Errorf("Bad request: missing user ID")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Must send user ID"})
		return
	}

	if !h.rateLimiter.Allowed(userID) {
		logger.Logger.Warnf("Max requests allowed exceeded for user %s", userID)
		ctx.JSON(http.StatusTooManyRequests, gin.H{
			"error": "max requests allowed exceeded",
		})
		return
	}

	response, err := h.service.GetFuckOff(userID)
	if err != nil {
		logger.Logger.Errorf("Error getting fuck off message: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
	return
}
