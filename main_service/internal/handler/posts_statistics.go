package handler

import (
	"log"
	"net/http"
	"soa-main/internal/service"
	"strconv"

	pb "soa-main/internal/statistics_proto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) viewPost(c *gin.Context) {
	var authorId int64
	authorIdS, ok := c.GetQuery("author_id")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "author_id parameter is not set")
		return
	}

	authorId, err := strconv.ParseInt(authorIdS, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "author_id parameter is not a number")
		return
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
	var authorId int64
	authorIdS, ok := c.GetQuery("author_id")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "author_id parameter is not set")
		return
	}

	authorId, err := strconv.ParseInt(authorIdS, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "author_id parameter is not a number")
		return
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

type postStatistics struct {
	Id          int32  `json:"ID"`
	AuthorId    int32  `json:"author_id"`
	AuthorLogin string `json:"author_login"`
	NumLikes    uint64 `json:"num_likes"`
	NumViews    uint64 `json:"num_views"`
}

func (h *Handler) getPostStatistics(c *gin.Context) {
	postIdS, ok := c.GetQuery("post_id")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no id parameter for post")
		return
	}

	postId, err := strconv.Atoi(postIdS)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id parameter not a number")
		return
	}

	statistics, err := h.services.GetPostStatistics(c.Request.Context(), &pb.PostId{PostId: int32(postId)})
	if err != nil {
		log.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "error getting post statistics")
		return
	}

	userData, err := h.services.GetUserLogin(int(statistics.GetAuthorId()))
	if err != nil {
		log.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "error getting posts statistics")
		return
	}

	log.Println("successful getPostStatistics request")
	c.JSON(http.StatusOK,
		postStatistics{
			Id:          int32(postId),
			AuthorId:    statistics.GetAuthorId(),
			AuthorLogin: userData.Username,
			NumLikes:    statistics.GetNumLikes(),
			NumViews:    statistics.GetNumViews(),
		})
}

type topPosts struct {
	Posts []postStatistics `json:"posts"`
}

func (h *Handler) getTopKPosts(c *gin.Context) {
	eventTypeS, ok := c.GetQuery("event_type")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no event_type parameter for get top k posts")
		return
	}

	var eventType pb.EventType
	if eventTypeS == "like" {
		eventType = pb.EventType_LIKE
	} else if eventTypeS == "view" {
		eventType = pb.EventType_VIEW
	} else {
		newErrorResponse(c, http.StatusBadRequest, "unknown event_type value for get top k posts")
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

	res, err := h.services.GetTopKPosts(c.Request.Context(), &pb.TopKRequest{K: k, Event: eventType})
	if err != nil {
		log.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "error getting post statistics")
		return
	}

	posts := []postStatistics{}
	for _, p := range res.Posts {
		userData, err := h.services.GetUserLogin(int(p.GetAuthorId()))
		if err != nil {
			log.Println(err.Error())
			newErrorResponse(c, http.StatusInternalServerError, "error getting posts statistics")
			return
		}

		posts = append(posts,
			postStatistics{
				Id:          p.GetPostId(),
				AuthorId:    p.GetAuthorId(),
				AuthorLogin: userData.Username,
				NumLikes:    p.GetNumLikes(),
				NumViews:    p.GetNumViews(),
			})
	}

	log.Println("successful getTopKPosts request")
	c.JSON(http.StatusOK, topPosts{Posts: posts})
}

type userStatistics struct {
	Id       int32  `json:"ID"`
	Login    string `json:"login"`
	NumLikes uint64 `json:"num_likes"`
	NumViews uint64 `json:"num_views"`
}

type topUsers struct {
	Users []userStatistics `json:"users"`
}

func (h *Handler) getTopKUsers(c *gin.Context) {
	eventTypeS, ok := c.GetQuery("event_type")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no event_type parameter for get top k posts")
		return
	}

	var eventType pb.EventType
	if eventTypeS == "like" {
		eventType = pb.EventType_LIKE
	} else if eventTypeS == "view" {
		eventType = pb.EventType_VIEW
	} else {
		newErrorResponse(c, http.StatusBadRequest, "unknown event_type value for get top k users")
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

	res, err := h.services.GetTopKUsers(c.Request.Context(), &pb.TopKRequest{K: k, Event: eventType})
	if err != nil {
		log.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "error getting users statistics")
		return
	}

	users := []userStatistics{}
	for _, u := range res.Users {
		userData, err := h.services.GetUserLogin(int(u.GetAuthorId()))
		if err != nil {
			log.Println(err.Error())
			newErrorResponse(c, http.StatusInternalServerError, "error getting users statistics")
			return
		}

		users = append(users,
			userStatistics{
				Id:       u.GetAuthorId(),
				Login:    userData.Username,
				NumLikes: u.GetNumLikes(),
				NumViews: u.GetNumViews(),
			})
	}

	log.Println("successful getTopKUsers request")
	c.JSON(http.StatusOK, topUsers{Users: users})
}
