package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dhawalhost/hooknetic/models"
	"github.com/dhawalhost/hooknetic/repository"
	"github.com/gin-gonic/gin"
)

func CreateWebhooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := &models.WHRequest{}
		err := c.BindJSON(request)
		if err != nil {
			log.Println("Error: ", err)
			c.JSON(http.StatusBadRequest, ErrorResponseMessage("Invalid request"))
			return
		}

		response, _ := repository.CreateWebhook(*request)
		c.JSON(http.StatusOK, ResponseMessage("Webhook Created Successfully", response))
	}
}

func WebhookListener() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request models.WebhookPayload
		HookID := c.Param("token")
		action, ok := models.WebhookCacheList[HookID]
		if !ok {
			c.JSON(http.StatusNotFound, ErrorResponseMessage("Webhook not found"))
			return
		}

		err := c.BindJSON(&request)
		if err != nil {
			log.Println("Error: ", err)
			c.JSON(http.StatusBadRequest, ErrorResponseMessage("Invalid payload"))
			return
		}
		c.JSON(http.StatusOK, ResponseMessage(fmt.Sprintf("Payload for the action:%v has been received successfully", action), request.Payload))
	}
}

func ErrorResponseMessage(msg interface{}) gin.H {
	return gin.H{"error": msg}
}

func ResponseMessage(msg interface{}, data ...any) gin.H {
	response := gin.H{}
	response["message"] = msg
	for _, v := range data {
		response["data"] = v
	}
	return response
}
