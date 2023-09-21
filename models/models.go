package models

import (
	"time"
)

type WebhookData struct {
	HookID      string    `json:"hookId"`
	Action      string    `json:"action"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	CreatedDate time.Time `json:"createdDate"`
}

type WebhookPayload struct {
	Payload map[string]interface{} `json:"payload"`
}

var WebhookCacheList = make(map[string]string)

type WHRequest struct {
	Action      string `json:"action"`
	Description string `json:"description,omitempty"`
}
