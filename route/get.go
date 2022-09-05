package route

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jungmu/go-web/db"
	"github.com/jungmu/go-web/ent/blog"
	"github.com/jungmu/go-web/markdown"
)

type reqGet struct {
	Title string `uri:"title"`
}

func Get(c *gin.Context) {
	r := reqGet{}
	err := c.BindUri(&r)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "bad params",
		})
		return
	}

	b, err := db.Client().Blog.Query().
		Where(blog.Title(r.Title)).
		Only(c)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": err,
		})
	}

	c.HTML(http.StatusOK, "blog-post.tmpl", gin.H{
		"title":    b.Title,
		"create":   b.CreateDatetime.Format("2006-01-02 15:04:05"),
		"update":   b.UpdateDatetime.Format("2006-01-02 15:04:05"),
		"markdown": markdown.MdToHtml(b.Content),
	})
}
