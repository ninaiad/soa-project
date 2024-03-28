package handler

import (
	"context"
	"log"
	"net/http"
	"soa/posts"
	"time"

	"github.com/gin-gonic/gin"
)

type postTextMsg struct {
	Text string `json:"text"`
}

type postIdMsg struct {
	PostId int `json:"post_id"`
}

type updatePostInput struct {
	postIdMsg
	postTextMsg
}

type postResponse struct {
	Text        string `json:"text"`
	TimeUpdated string `json:"time_updated"`
}

type postsByPageInput struct {
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
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

	var input updatePostInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	_, err = h.services.PostsServerClient.UpdatePost(context.Background(), &posts.UpdateRequest{AuthorId: int32(userId), PostId: int32(input.PostId), Text: input.Text})
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

	var input postIdMsg
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	_, err = h.services.PostsServerClient.DeletePost(context.Background(), &posts.PostIdRequest{AuthorId: int32(userId), PostId: int32(input.PostId)})
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

	var input postIdMsg
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	post, err := h.services.PostsServerClient.GetPost(context.Background(), &posts.PostIdRequest{AuthorId: int32(userId), PostId: int32(input.PostId)})
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

	var input postsByPageInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	posts, err := h.services.PostsServerClient.GetPageOfPosts(context.Background(), &posts.GetPageOfPostsRequest{AuthorId: int32(userId), PageNum: int32(input.PageNum), PageSize: int32(input.PageSize)})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	postsData := make([]postResponse, posts.PageSize)
	for i := range (*posts).Posts {
		postsData[i] = postResponse{Text: (*posts).Posts[i].Text, TimeUpdated: (*posts).Posts[i].TimeUpdated.AsTime().Format(time.RFC3339)}
	}

	log.Println("successful getPostsbyPage request")
	c.JSON(http.StatusOK, postsByPageOutput{PageNum: int(posts.PageNum), PageSize: int(posts.PageSize), Posts: postsData})
}
