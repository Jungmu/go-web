package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jungmu/go-web/api/blog"
	"github.com/jungmu/go-web/route"
)

func Init(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")

	blogGroup := r.Group("blog")
	{
		blogGroup.GET("", route.Index)
		blogGroup.GET("post", route.Post)
		blogGroup.GET("article/:title", route.Get)
		blogGroup.POST("article/:title", blog.Post)
	}
}
