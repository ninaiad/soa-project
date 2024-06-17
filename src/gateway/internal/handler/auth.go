package handler

import (
	"log"
	"net/http"

	post_pb "gateway/internal/service/posts"
	stat_pb "gateway/internal/service/statistics"
	"gateway/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input user.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	userId, err := h.service.CreateUser(input)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" { // violation of unique constraint
				newErrorResponse(c, http.StatusBadRequest, "invalid input body: login exists")
				return
			}
		}

		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, _, err := h.service.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Println("new successful sign-up")

	c.JSON(http.StatusOK, map[string]interface{}{
		"token":   token,
		"user_id": userId,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, userId, err := h.service.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Println("new successful sign-in")

	c.JSON(http.StatusOK, map[string]interface{}{
		"token":   token,
		"user_id": userId,
	})
}

func (h *Handler) updateUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input user.UserPublic
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	updatedData, err := h.service.UpdateUser(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("successful updateUser request")
	c.JSON(http.StatusOK, updatedData)
}

func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.Authorization.DeleteUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	_, err = h.service.PostsServerClient.DeleteUser(&gin.Context{}, &post_pb.UserId{Id: userId})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = h.service.StatisticsServiceClient.DeleteUser(&gin.Context{}, &stat_pb.UserId{Id: userId})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("successful updateUser request")
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
