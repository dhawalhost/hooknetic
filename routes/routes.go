package routes

import (
	"github.com/dhawalhost/hooknetic/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/webhooks/create", handlers.CreateWebhooks())
	router.POST("/webhooks/:token", handlers.WebhookListener())
}
