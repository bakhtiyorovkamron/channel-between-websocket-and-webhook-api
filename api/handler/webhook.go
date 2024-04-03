package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *handlerV1) ForwardToWebhook(c *gin.Context) {
	// This function will forward the data
	req := make(map[string]interface{})
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"status": "error", "message": "Invalid request"})
		return
	}

	myChannel <- req

	c.JSON(200, gin.H{"status": "success", "message": "Data forwarded to webhook successfully"})
}
