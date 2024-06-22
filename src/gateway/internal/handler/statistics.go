package handler

import (
	"log"
	"net/http"
	"strconv"

	"gateway/internal/service"

	"github.com/gin-gonic/gin"
)

type postStatistics struct {
	Id             int64  `json:"id"`
	AuthorId       int64  `json:"author_id"`
	AuthorUsername string `json:"author_username"`
	NumLikes       uint64 `json:"num_likes"`
	NumViews       uint64 `json:"num_views"`
}

type topPosts struct {
	Posts []postStatistics `json:"posts"`
}

type userStatistics struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	NumLikes uint64 `json:"num_likes"`
	NumViews uint64 `json:"num_views"`
}

type topUsers struct {
	Users []userStatistics `json:"users"`
}

func (h *Handler) viewPost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var authorId int64
	authorIdS, ok := c.GetQuery("author_id")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "author_id parameter is not set")
		return
	}

	authorId, err = strconv.ParseInt(authorIdS, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "author_id parameter is not a number")
		return
	}

	postIdS, ok := c.GetQuery("id")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no id parameter for post")
		return
	}
	postId, err := strconv.ParseInt(postIdS, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id parameter not a number")
		return
	}

	if err = h.service.AddEvent(postId, authorId, userId, service.View); err != nil {
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
		newErrorResponse(c, http.StatusBadRequest, "author_id parameter is not set")
		return
	}

	authorId, err = strconv.ParseInt(authorIdS, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "author_id parameter is not a number")
		return
	}

	postIdS, ok := c.GetQuery("id")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no id parameter for post")
		return
	}
	postId, err := strconv.ParseInt(postIdS, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id parameter not a number")
		return
	}

	if err = h.service.AddEvent(postId, authorId, userId, service.Like); err != nil {
		log.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "error adding like")
		return
	}

	log.Println("successful likePost request")
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) getPostStatistics(c *gin.Context) {
	postIdS, ok := c.GetQuery("id")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no id parameter for post")
		return
	}

	postId, err := strconv.ParseInt(postIdS, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id parameter not a number")
		return
	}

	s, err := h.service.GetPostStatistics(postId)
	if err != nil {
		log.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "error getting post statistics")
		return
	}

	username, err := h.service.GetUsername(s.GetAuthorId())
	if err != nil {
		log.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "error getting posts statistics")
		return
	}

	log.Println("successful getPostStatistics request")
	c.JSON(http.StatusOK,
		postStatistics{
			Id:             postId,
			AuthorId:       s.GetAuthorId(),
			AuthorUsername: username,
			NumLikes:       s.GetNumLikes(),
			NumViews:       s.GetNumViews(),
		})
}

func (h *Handler) getTopKPosts(c *gin.Context) {
	eventTypeS, ok := c.GetQuery("event_type")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no event_type parameter for get top k posts")
		return
	}

	kS, ok := c.GetQuery("k")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no k parameter for get top k posts")
		return
	}

	k, err := strconv.ParseUint(kS, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "k parameter not a number")
		return
	}

	res, err := h.service.GetTopKPosts(service.EventType(eventTypeS), k)
	if err != nil {
		log.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "error getting post statistics")
		return
	}

	posts := []postStatistics{}
	for _, p := range res.Posts {
		username, err := h.service.GetUsername(p.GetAuthorId())
		if err != nil {
			log.Println(err.Error())
			newErrorResponse(c, http.StatusInternalServerError, "error getting posts statistics")
			return
		}

		posts = append(posts,
			postStatistics{
				Id:             p.GetPostId(),
				AuthorId:       p.GetAuthorId(),
				AuthorUsername: username,
				NumLikes:       p.GetNumLikes(),
				NumViews:       p.GetNumViews(),
			})
	}

	log.Println("successful getTopKPosts request")
	c.JSON(http.StatusOK, topPosts{Posts: posts})
}

func (h *Handler) getTopKUsers(c *gin.Context) {
	eventTypeS, ok := c.GetQuery("event_type")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no event_type parameter for get top k posts")
		return
	}

	kS, ok := c.GetQuery("k")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no k parameter for get top k users")
		return
	}

	k, err := strconv.ParseUint(kS, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "k parameter not a number")
		return
	}

	res, err := h.service.GetTopKUsers(service.EventType(eventTypeS), k)
	if err != nil {
		log.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "error getting users statistics")
		return
	}

	users := []userStatistics{}
	for _, u := range res.Users {
		username, err := h.service.GetUsername(u.Id)
		if err != nil {
			log.Println(err.Error())
			newErrorResponse(c, http.StatusInternalServerError, "error getting users statistics")
			return
		}

		users = append(users,
			userStatistics{
				Id:       u.GetId(),
				Username: username,
				NumLikes: u.GetNumLikes(),
				NumViews: u.GetNumViews(),
			})
	}

	log.Println("successful getTopKUsers request")
	c.JSON(http.StatusOK, topUsers{Users: users})
}
