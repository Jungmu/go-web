package route

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jungmu/go-web/db"
	"github.com/jungmu/go-web/ent"
	"github.com/jungmu/go-web/ent/blog"
)

type blogPost struct {
	Title          string   `json:"title"`
	SubTitle       string   `json:"subTitle"`
	Tags           []string `json:"tags"`
	UpdateDateTime string   `json:"updateDateTime"`
	CreateDateTime string   `json:"createDateTime"`
}

type req struct {
	Tag  string `form:"tag"`
	Page int    `form:"page"`
}

const pageSize = 20

func Index(c *gin.Context) {
	r := req{}
	err := c.Bind(&r)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "bad params",
		})
		return
	}

	b, err := db.Client().Blog.Query().
		Where(blog.TagsContains(r.Tag)).
		Limit(pageSize).
		Offset(pageSize * r.Page).
		Order(ent.Desc(blog.FieldCreateDatetime)).
		All(c)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": err,
		})
	}

	Posts := []blogPost{}

	for _, p := range b {
		Posts = append(Posts, blogPost{
			p.Title,
			p.SubTitle,
			strings.Split(p.Tags, ","),
			p.UpdateDatetime.Format("2006-01-02 15:04:05"),
			p.CreateDatetime.Format("2006-01-02 15:04:05"),
		})
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "article list",
		"Posts": Posts,
	})
}
