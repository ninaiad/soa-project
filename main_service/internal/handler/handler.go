package handler

import (
	"log"
	"soa-main/internal/service"

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

	router.POST("/post", h.userIdentity, h.createPost)
	router.PUT("/post", h.userIdentity, h.updatePost)
	router.DELETE("/post", h.userIdentity, h.deletePost)
	router.GET("/post", h.userIdentity, h.getPost)
	router.GET("/posts", h.userIdentity, h.getPageOfPosts)

	router.POST("/view", h.userIdentity, h.viewPost)
	router.POST("/like", h.userIdentity, h.likePost)

	log.Println("router set up")
	return router
}
