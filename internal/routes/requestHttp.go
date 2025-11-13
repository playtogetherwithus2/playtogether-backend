package routes

import (
	"net/http"
	"play-together/internal/model"
	"play-together/internal/service"

	"github.com/gin-gonic/gin"
)

func AddRequestRoutes(router *gin.RouterGroup, requestService *service.RequestService) {
	router.POST("/request", createRequestHandler(requestService))
	router.GET("/request", getAllRequestsHandler(requestService))
	router.PATCH("/request/:id", updateRequestHandler(requestService))
	router.GET("/request/:id", getRequestByIDHandler(requestService))
	router.DELETE("/request/:id", deleteRequestByIDHandler(requestService))
}

func createRequestHandler(requestService *service.RequestService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.Request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		requestID, err := requestService.CreateRequest(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Request created", "request_id": requestID})
	}
}

func updateRequestHandler(requestService *service.RequestService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var updateData map[string]interface{}

		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := requestService.UpdateRequest(c.Request.Context(), id, updateData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Request updated successfully"})
	}
}

func getAllRequestsHandler(requestService *service.RequestService) gin.HandlerFunc {
	return func(c *gin.Context) {
		senderID := c.Query("senders_id")
		receiverID := c.Query("receivers_id")
		includeUserData := c.Query("user_data") == "true"

		requests, err := requestService.GetAllRequests(c.Request.Context(), senderID, receiverID, includeUserData)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch requests", "details": err.Error()})
			return
		}

		c.JSON(200, gin.H{"requests": requests})
	}
}

func getRequestByIDHandler(requestService *service.RequestService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		includeUserData := c.Query("user_data") == "true"

		request, err := requestService.GetRequestByID(c.Request.Context(), id, includeUserData)
		if err != nil {
			c.JSON(404, gin.H{
				"error":   "Post not found",
				"details": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"request": request,
		})
	}
}

func deleteRequestByIDHandler(requestService *service.RequestService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := requestService.DeleteRequestByID(c.Request.Context(), id); err != nil {
			c.JSON(404, gin.H{
				"error":   "Failed to delete request",
				"details": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Request deleted successfully",
			"id":      id,
		})
	}
}
