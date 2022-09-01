package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jungmu/go-web/api/user"
)

func Init(r *gin.Engine) {
	r.GET("/user/:id", user.Get)
}
