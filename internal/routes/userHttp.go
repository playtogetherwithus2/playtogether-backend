package routes

import (
	"net/http"
	"play-together/internal/model"
	"play-together/internal/service"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(router *gin.RouterGroup, userService *service.UserService) {
	router.GET("/users", getUsersHandler(userService))
	router.GET("/users/:id", getUserByIDHandler(userService))
	router.POST("/users/id", getUsersByIDsHandler(userService))
}

func getUsersHandler(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		users, err := userService.GetUsers(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}

func getUserByIDHandler(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		userID := c.Param("id")

		user, err := userService.GetUserByID(ctx, userID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}

func getUsersByIDsHandler(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.UserIDsRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid request payload",
				"details": err.Error(),
			})
			return
		}

		if len(req.UserIDs) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id field cannot be empty",
			})
			return
		}

		ctx := c.Request.Context()
		users, err := userService.GetUsersByIDs(ctx, req.UserIDs)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}
