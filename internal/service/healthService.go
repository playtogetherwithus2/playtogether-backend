package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthService struct{}

func NewHealthService() *HealthService {
	return &HealthService{}
}

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (s *HealthService) HealthCheck(c *gin.Context) {
	response := HealthResponse{
		Status:  "success",
		Message: "Service is running smoothly",
	}
	c.JSON(http.StatusOK, response)
}
