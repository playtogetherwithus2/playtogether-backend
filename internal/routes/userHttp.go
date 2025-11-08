package routes

import (
	"fmt"
	"net/http"
	"play-together/internal/model"
	"play-together/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(router *gin.RouterGroup, userService *service.UserService) {
	router.GET("/users", getUsersHandler(userService))
	router.GET("/users/:id", getUserByIDHandler(userService))
	router.POST("/users/id", getUsersByIDsHandler(userService))
	router.PATCH("/users/:id", updateUserHandler(userService))
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

func updateUserHandler(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id") 

		if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10 MB limit
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
			return
		}

		var req model.UpdateUserRequest

		req.UserName = c.PostForm("user_name")
		req.Name = c.PostForm("name")
		req.Gender = c.PostForm("gender")
		req.Bio = c.PostForm("bio")
		req.City = c.PostForm("city")
		req.PreferredTime = c.PostForm("preferred_time")

		if ageStr := c.PostForm("age"); ageStr != "" {
			var age int
			if _, err := fmt.Sscanf(ageStr, "%d", &age); err == nil {
				req.Age = age
			}
		}

		if sports := c.PostForm("sports_interested"); sports != "" {
			req.SportsInterested = strings.Split(sports, ",")
		}

		if days := c.PostForm("availability_days"); days != "" {
			req.AvailabilityDays = strings.Split(days, ",")
		}

		if locations := c.PostForm("preferred_locations"); locations != "" {
			req.PreferredLocations = strings.Split(locations, ",")
		}

		_, header, err := c.Request.FormFile("profile_photo")
		if err == nil && header != nil {
			filePath := fmt.Sprintf("/tmp/%s", header.Filename)
			if err := c.SaveUploadedFile(header, filePath); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save uploaded image"})
				return
			}
			req.ProfilePhotoURL = filePath
		}

		ctx := c.Request.Context()
		if err := userService.UpdateUser(ctx, id, req); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":   "User updated successfully",
			"id":        id,
			"user_name": req.UserName,
		})
	}
}
