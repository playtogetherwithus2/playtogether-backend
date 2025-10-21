package routes

import (
	"play-together/internal/model"
	"play-together/internal/service"

	"github.com/gin-gonic/gin"
)

func AddPostRoutes(router *gin.RouterGroup, postService *service.PostService) {
	router.POST("/posts", createPostHandler(postService))
	router.GET("/posts", getAllPostsHandler(postService))
	router.GET("/posts/:id", getPostByIDHandler(postService))
}

func createPostHandler(postService *service.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var post model.GamePost
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request payload", "details": err.Error()})
			return
		}

		id, err := postService.CreatePost(c.Request.Context(), &post)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create post", "details": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "Post created successfully",
			"post_id": id,
		})
	}
}

func getAllPostsHandler(postService *service.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		posts, err := postService.GetAllPosts(c.Request.Context())
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch posts", "details": err.Error()})
			return
		}
		c.JSON(200, gin.H{"posts": posts})
	}
}

func getPostByIDHandler(postService *service.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		post, err := postService.GetPostByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(404, gin.H{
				"error":   "Post not found",
				"details": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"post": post,
		})
	}
}
