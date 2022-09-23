package route

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jungmu/go-web/redis"
)

func AIChat(c *gin.Context) {
	c.HTML(http.StatusOK, "ai-chat.tmpl", gin.H{
		"title": "openAI chat demo",
	})
}

func Chat(c *gin.Context) {
	chatList, err := redis.GetLastChatList()
	if err != nil {
		chatList = []string{}
	}

	c.HTML(http.StatusOK, "chat.tmpl", gin.H{
		"title":      "websocket chat",
		"ip":         c.ClientIP(),
		"chat":       chatList,
		"randomName": fmt.Sprintf("guest-%d", time.Now().UnixMilli()),
	})
}
