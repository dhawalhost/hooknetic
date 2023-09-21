package repository

import (
	"time"

	"github.com/dhawalhost/hooknetic/models"
	"github.com/dhawalhost/hooknetic/utility"
)

func CreateWebhook(req models.WHRequest) (response models.WebhookData, err error) {
	response.Action = req.Action
	response.Description = req.Description
	response.CreatedDate = time.Now()
	response.Active = true
	response.HookID, err = utility.GenerateUUID()
	if err != nil {
		return models.WebhookData{}, err
	}
	models.WebhookCacheList[response.HookID] = response.Action
	return
}
