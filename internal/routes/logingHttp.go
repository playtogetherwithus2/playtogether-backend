package routes

import (
	"play-together/internal/model"
	"play-together/internal/service"

	"github.com/gin-gonic/gin"
)

func AddLoginRoutes(router *gin.RouterGroup, loginService *service.LoginService) {
	router.POST("/login", loginHandler(loginService))
	router.POST("/signup", signupHandler(loginService))
}

func loginHandler(loginService *service.LoginService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest model.LoginRequest

		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(400, gin.H{
				"error":   "Invalid request payload",
				"details": err.Error(),
			})
			return
		}

		token, err := loginService.Login(c.Request.Context(), loginRequest.Email, loginRequest.Password)
		if err != nil {
			c.JSON(401, gin.H{
				"error":   "Login failed",
				"details": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"token":   token,
			"message": "Login successful",
		})
	}
}

func signupHandler(loginService *service.LoginService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var signupRequest model.SignupRequest

		if err := c.ShouldBindJSON(&signupRequest); err != nil {
			c.JSON(400, gin.H{
				"error":   "Invalid request payload",
				"details": err.Error(),
			})
			return
		}

		user, err := loginService.Signup(c.Request.Context(), signupRequest.Email, signupRequest.Password)
		if err != nil {
			c.JSON(401, gin.H{
				"error":   "Signup failed",
				"details": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Signup successful",
			"user_id": user.UID,
		})
	}
}