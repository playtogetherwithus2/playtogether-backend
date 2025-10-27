package routes

import (
	"net/http"
	"play-together/internal/model"
	"play-together/internal/service"

	"github.com/gin-gonic/gin"
)

func AddChatRoutes(router *gin.RouterGroup, chatService *service.ChatService) {
	router.POST("/groups", createGroupHandler(chatService))
	router.GET("/groups", getAllGroupsHandler(chatService))
	router.GET("/groups/:groupId/messages", getMessagesHandler(chatService))
	router.POST("/groups/:groupId/messages", sendMessageHandler(chatService))
	router.POST("/group/:id/add", addMemberHandler(chatService))
	router.POST("/group/:id/remove", removeMemberHandler(chatService))
	router.GET("/group/:id/details", getGroupDetailsHandler(chatService))
}

func getAllGroupsHandler(chatService *service.ChatService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		memberID := c.Query("member_id")

		groups, err := chatService.GetAllGroups(ctx, memberID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to fetch groups",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"groups": groups,
		})
	}
}

func createGroupHandler(chatService *service.ChatService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.CreateGroupRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		groupID, err := chatService.CreateGroup(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Group created", "group_id": groupID})
	}
}

func getMessagesHandler(chatService *service.ChatService) gin.HandlerFunc {
	return func(c *gin.Context) {
		groupID := c.Param("groupId")

		messages, err := chatService.GetMessages(c.Request.Context(), groupID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"messages": messages})
	}
}

func sendMessageHandler(chatService *service.ChatService) gin.HandlerFunc {
	return func(c *gin.Context) {
		groupID := c.Param("groupId")
		var req model.SendMessageRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := chatService.SendMessage(c.Request.Context(), groupID, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
	}
}

func addMemberHandler(chatService *service.ChatService) gin.HandlerFunc {
	return func(c *gin.Context) {
		groupID := c.Param("id")
		var req model.ModifyMemberRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := chatService.AddMember(c.Request.Context(), groupID, req); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Member added successfully"})
	}
}

func removeMemberHandler(chatService *service.ChatService) gin.HandlerFunc {
	return func(c *gin.Context) {
		groupID := c.Param("id")
		var req model.ModifyMemberRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := chatService.RemoveMember(c.Request.Context(), groupID, req); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Member removed successfully"})
	}
}

func getGroupDetailsHandler(chatService *service.ChatService) gin.HandlerFunc {
	return func(c *gin.Context) {
		groupID := c.Param("id")

		group, err := chatService.GetGroupDetails(c.Request.Context(), groupID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"group": group})
	}
}
