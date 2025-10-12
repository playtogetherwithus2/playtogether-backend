package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	response := HealthResponse{
		Status:  "success",
		Message: "Service is running smoothly",
	}
	c.JSON(http.StatusOK, response)
}
