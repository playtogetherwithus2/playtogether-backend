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
	router.GET("/request/:id", getRequestByIDHandler(requestService))
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

func getAllRequestsHandler(requestService *service.RequestService) gin.HandlerFunc {
	return func(c *gin.Context) {
		requests, err := requestService.GetAllRequests(c.Request.Context())
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch posts", "details": err.Error()})
			return
		}
		c.JSON(200, gin.H{"requests": requests})
	}
}

func getRequestByIDHandler(requestService *service.RequestService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		request, err := requestService.GetRequestByID(c.Request.Context(), id)
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
