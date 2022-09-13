package route

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jungmu/go-web/db"
	blogDB "github.com/jungmu/go-web/db/blog"
	"github.com/jungmu/go-web/ent/blog"
	"github.com/jungmu/go-web/markdown"
	"github.com/jungmu/go-web/xconst"
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

	blogDB.Log(b.ID, c.Request.RequestURI, xconst.ReasonView, c.ClientIP(), r)

	viewCount, err := blogDB.GetViewCount(b.ID)
	if err != nil {
		log.Panicln(err)
	}

	visitCount, err := blogDB.GetVisitCount(b.ID)
	if err != nil {
		log.Panicln(err)
	}

	c.HTML(http.StatusOK, "blog-post.tmpl", gin.H{
		"title":      b.Title,
		"create":     b.CreateDatetime.Format("2006-01-02 15:04:05"),
		"update":     b.UpdateDatetime.Format("2006-01-02 15:04:05"),
		"markdown":   markdown.MdToHtml(b.Content),
		"viewCount":  viewCount,
		"visitCount": visitCount,
	})
}
