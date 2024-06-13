package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"gateway/internal/service/posts_proto"

	"github.com/gin-gonic/gin"
)

type postTextMsg struct {
	Text string `json:"text"`
}

type postIdMsg struct {
	PostId int `json:"post_id"`
}

type postResponse struct {
	Text        string `json:"text"`
	TimeUpdated string `json:"time_updated"`
}

type postsByPageOutput struct {
	PageNum  int            `json:"page_num"`
	PageSize int            `json:"page_size"`
	Posts    []postResponse `json:"posts"`
}

func (h *Handler) createPost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input postTextMsg
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	resp, err := h.service.CreatePost(
		context.Background(),
		&posts_proto.CreateRequest{AuthorId: int32(userId), Text: input.Text})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("successful createPost request")
	c.JSON(http.StatusOK, postIdMsg{PostId: int(resp.PostId)})
}

func (h *Handler) updatePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	postIdS, ok := c.GetQuery("id")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no id parameter for updated post")
		return
	}

	postId, err := strconv.Atoi(postIdS)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id parameter not a number")
		return
	}

	var input postTextMsg
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	_, err = h.service.UpdatePost(
		context.Background(),
		&posts_proto.UpdateRequest{
			AuthorId: int32(userId),
			PostId:   int32(postId),
			Text:     input.Text,
		})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("successful updatePost request")
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) deletePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	postIdS, ok := c.GetQuery("id")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no id parameter for deleted post")
		return
	}

	postId, err := strconv.Atoi(postIdS)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id parameter not a number")
		return
	}

	_, err = h.service.DeletePost(
		context.Background(),
		&posts_proto.PostIdRequest{
			AuthorId: int32(userId),
			PostId:   int32(postId),
		})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("successful deletePost request")
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) getPost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var authorId int
	authorIdS, ok := c.GetQuery("author_id")
	if !ok {
		authorId = userId
	} else {
		authorId, err = strconv.Atoi(authorIdS)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "author_id parameter is not a number")
			return
		}
	}

	postIdS, ok := c.GetQuery("id")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no id parameter for post")
		return
	}

	postId, err := strconv.Atoi(postIdS)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id parameter not a number")
		return
	}

	post, err := h.service.GetPost(
		context.Background(),
		&posts_proto.PostIdRequest{
			AuthorId: int32(authorId),
			PostId:   int32(postId),
		})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("successful getPost request")
	c.JSON(http.StatusOK,
		postResponse{
			Text:        post.Text,
			TimeUpdated: post.TimeUpdated.AsTime().Format(time.RFC3339),
		})
}

func (h *Handler) getPageOfPosts(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var authorId int
	authorIdS, ok := c.GetQuery("author_id")
	if !ok {
		authorId = userId
	} else {
		authorId, err = strconv.Atoi(authorIdS)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "author_id parameter is not a number")
			return
		}
	}

	pageNumS, ok := c.GetQuery("page_num")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no page_num parameter")
		return
	}

	pageNum, err := strconv.Atoi(pageNumS)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "page_num parameter not a number")
		return
	}

	pageSizeS, ok := c.GetQuery("page_size")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "no page_size parameter")
		return
	}

	pageSize, err := strconv.Atoi(pageSizeS)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "page_size parameter not a number")
		return
	}

	posts, err := h.service.GetPageOfPosts(
		context.Background(),
		&posts_proto.GetPageOfPostsRequest{
			AuthorId: int32(authorId),
			PageNum:  int32(pageNum),
			PageSize: int32(pageSize),
		})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	postsData := make([]postResponse, posts.PageSize)
	for i := range (*posts).Posts {
		postsData[i] = postResponse{
			Text:        (*posts).Posts[i].Text,
			TimeUpdated: (*posts).Posts[i].TimeUpdated.AsTime().Format(time.RFC3339),
		}
	}

	log.Println("successful getPageOfPosts request")
	c.JSON(http.StatusOK,
		postsByPageOutput{
			PageNum:  int(posts.PageNum),
			PageSize: int(posts.PageSize),
			Posts:    postsData,
		})
}
