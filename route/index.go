package route

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jungmu/go-web/db"
	blogDB "github.com/jungmu/go-web/db/blog"
	"github.com/jungmu/go-web/ent"
	"github.com/jungmu/go-web/ent/blog"
	"github.com/jungmu/go-web/redis"
	"github.com/jungmu/go-web/xconst"
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
		Where(blog.TagsContains(strings.TrimSpace(r.Tag))).
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

	blogDB.Log(0, c.Request.RequestURI, xconst.ReasonIndex, c.ClientIP(), r)

	totalView, err := blogDB.GetTotalViewCount()
	if err != nil {
		log.Panicln(err)
	}

	totalVisit, err := blogDB.GetTotalVisitCount()
	if err != nil {
		log.Panicln(err)
	}

	chatList, err := redis.GetLastChatList()
	if err != nil {
		chatList = []string{}
	}

	data := gin.H{
		"title":      "article list",
		"Posts":      Posts,
		"totalView":  totalView,
		"totalVisit": totalVisit,
		"ip":         c.ClientIP(),
		"chat":       chatList,
		"randomName": fmt.Sprintf("guest-%d", time.Now().UnixMilli()),
	}
	c.HTML(http.StatusOK, "index.tmpl", data)
}
