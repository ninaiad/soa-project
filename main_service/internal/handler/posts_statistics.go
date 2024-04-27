package handler

import (
	"log"
	"net/http"
	"soa-main/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) viewPost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var authorId int64
	authorIdS, ok := c.GetQuery("author_id")
	if !ok {
		authorId = int64(userId)
	} else {
		authorId, err = strconv.ParseInt(authorIdS, 10, 64)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "author_id parameter is not a number")
			return
		}
	}

	postIdS, ok := c.GetQuery("post_id")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no id parameter for post")
		return
	}
	postId, err := strconv.ParseInt(postIdS, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id parameter not a number")
		return
	}

	if err = h.services.AddEvent(postId, authorId, service.View); err != nil {
		log.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "error adding view")
		return
	}

	log.Println("successful viewPost request")
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) likePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var authorId int64
	authorIdS, ok := c.GetQuery("author_id")
	if !ok {
		authorId = int64(userId)
	} else {
		authorId, err = strconv.ParseInt(authorIdS, 10, 64)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "author_id parameter is not a number")
			return
		}
	}

	postIdS, ok := c.GetQuery("post_id")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no id parameter for post")
		return
	}
	postId, err := strconv.ParseInt(postIdS, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id parameter not a number")
		return
	}

	if err = h.services.AddEvent(postId, authorId, service.Like); err != nil {
		log.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "error adding like")
		return
	}

	log.Println("successful likePost request")
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
