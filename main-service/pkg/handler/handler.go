package handler

import (
	"log"
	"mainservice/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) SetupRouter() *gin.Engine {
	router := gin.New()
	router.POST("/sign-up", h.signUp)
	router.POST("/sign-in", h.signIn)
	router.PUT("/user", h.userIdentity, h.updateUser)

	log.Println("router set up")
	return router
}
