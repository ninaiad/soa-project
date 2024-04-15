package handler

import (
	"context"
	"log"
	"net/http"
	"soa/posts"
	"strconv"
	"time"

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
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	postId, err := h.services.PostsServerClient.CreatePost(context.Background(), &posts.CreateRequest{AuthorId: int32(userId), Text: input.Text})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("successful createPost request")
	c.JSON(http.StatusOK, postIdMsg{PostId: int(postId.PostId)})
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
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	_, err = h.services.PostsServerClient.UpdatePost(context.Background(), &posts.UpdateRequest{AuthorId: int32(userId), PostId: int32(postId), Text: input.Text})
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

	_, err = h.services.PostsServerClient.DeletePost(context.Background(), &posts.PostIdRequest{AuthorId: int32(userId), PostId: int32(postId)})
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

	post, err := h.services.PostsServerClient.GetPost(context.Background(), &posts.PostIdRequest{AuthorId: int32(userId), PostId: int32(postId)})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("successful getPost request")
	c.JSON(http.StatusOK, postResponse{Text: post.Text, TimeUpdated: post.TimeUpdated.AsTime().Format(time.RFC3339)})
}

func (h *Handler) getPageOfPosts(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
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

	posts, err := h.services.PostsServerClient.GetPageOfPosts(context.Background(), &posts.GetPageOfPostsRequest{AuthorId: int32(userId), PageNum: int32(pageNum), PageSize: int32(pageSize)})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	postsData := make([]postResponse, posts.PageSize)
	for i := range (*posts).Posts {
		postsData[i] = postResponse{Text: (*posts).Posts[i].Text, TimeUpdated: (*posts).Posts[i].TimeUpdated.AsTime().Format(time.RFC3339)}
	}

	log.Println("successful getPageOfPosts request")
	c.JSON(http.StatusOK, postsByPageOutput{PageNum: int(posts.PageNum), PageSize: int(posts.PageSize), Posts: postsData})
}
