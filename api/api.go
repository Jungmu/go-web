package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jungmu/go-web/api/ai"
	"github.com/jungmu/go-web/api/blog"
	"github.com/jungmu/go-web/api/comment"
	"github.com/jungmu/go-web/route"
)

func Init(r *gin.Engine) {
	r.Static("/assets", "./assets")
	r.StaticFile("/robots.txt", "./assets/robots.txt")
	r.StaticFile("/googlef7a4061446ba6717.html", "./assets/googlef7a4061446ba6717.html")
	r.StaticFile("/naver7150e5ff6ce115875d92e8b3c751ea7f.html", "./assets/naver7150e5ff6ce115875d92e8b3c751ea7f.html")
	r.LoadHTMLGlob("templates/*")
	r.GET("", route.Index)

	blogGroup := r.Group("blog")
	{
		blogGroup.GET("", route.Index)
		blogGroup.GET("post", route.Post)
		blogGroup.GET("article/:title", route.Get)
		blogGroup.GET("edit/:title", route.Edit)
		blogGroup.PUT("article/:title", blog.Edit)
		blogGroup.POST("article/:title", blog.Post)

		blogGroup.POST("comment/:title", comment.Post)
	}
	aiGroup := r.Group("ai")
	{
		aiGroup.GET("chat", route.Chat)
		aiGroup.POST("chat", ai.Chat)
	}
}
