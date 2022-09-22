package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Chat(c *gin.Context) {
	c.HTML(http.StatusOK, "ai-chat.tmpl", gin.H{
		"title": "openAI chat demo",
	})
}
