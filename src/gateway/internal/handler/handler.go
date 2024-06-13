package handler

import (
	"log"

	s "gateway/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *s.Service
}

func NewHandler(service *s.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SetupRouter() *gin.Engine {
	router := gin.New()

	user := router.Group("/user")
	{
		user.POST("/sign-up", h.signUp)
		user.POST("/sign-in", h.signIn)

		user.PUT("/", h.userIdentity, h.updateUser)

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

	posts := router.Group("/posts/statistics", h.userIdentity)
	{
		posts.GET("/users", h.getTopKUsers)
		posts.GET("/posts", h.getTopKPosts)
	}

	log.Println("router set up")
	return router
}
