package route

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jungmu/go-web/db"
	"github.com/jungmu/go-web/ent/blog"
)

type reqEdit struct {
	Title string `uri:"title"`
}

func Edit(c *gin.Context) {
	r := reqEdit{}
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

	c.HTML(http.StatusOK, "edit.tmpl", gin.H{
		"title":    b.Title,
		"subTitle": b.SubTitle,
		"tags":     b.Tags,
		"content":  b.Content,
	})
}
