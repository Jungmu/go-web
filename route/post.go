package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	c.HTML(http.StatusOK, "post.tmpl", gin.H{
		"title": "Write Article",
	})
}
