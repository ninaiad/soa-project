package handler

import (
	"log"
	"net/http"

	s "gateway/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service s.IService
}

func NewHandler(service s.IService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SetupRouter() *gin.Engine {
	router := gin.New()

	// Healthcheck
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello!")
	})

	user := router.Group("/user")
	{
		user.POST("/sign-up", h.signUp)
		user.POST("/sign-in", h.signIn)

		user.PUT("/", h.userIdentity, h.updateUser)
		user.DELETE("/", h.userIdentity, h.deleteUser)

		user.GET("/posts", h.userIdentity, h.getPageOfPosts)
	}

	post := router.Group("/post", h.userIdentity)
	{
		post.POST("/", h.createPost)
		post.PUT("/", h.updatePost)
		post.DELETE("/", h.deletePost)
		post.GET("/", h.getPost)

		post.POST("/view", h.viewPost)
		post.POST("/like", h.likePost)
		post.GET("/statistics", h.getPostStatistics)
	}

	statistics := router.Group("/posts/statistics", h.userIdentity)
	{
		statistics.GET("/users", h.getTopKUsers)
		statistics.GET("/posts", h.getTopKPosts)
	}

	log.Println("router set up")
	return router
}
