package blog

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jungmu/go-web/db"
	blogDB "github.com/jungmu/go-web/db/blog"
	"github.com/jungmu/go-web/xconst"
)

type reqPost struct {
	Title string `uri:"title"`
	reqPostBody
}

type reqPostBody struct {
	SubTitle string `json:"subTitle"`
	Tags     string `json:"tags" example:"tag1,tag2"`
	Content  string `json:"content"`
	PostKey  string `json:"postKey"`
}

func Post(c *gin.Context) {
	r := reqPost{}

	err := c.BindUri(&r)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "bad uri",
		})
		return
	}

	err = c.Bind(&r)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "bad params",
		})
		return
	}

	hash := sha256.New()
	hash.Write([]byte(r.PostKey))
	key := hex.EncodeToString(hash.Sum(nil))

	if strings.Compare(key, os.Getenv("POST_KEY")) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "postKey miss match",
		})
		return
	}

	result, err := db.Client().Blog.Create().
		SetTitle(r.Title).
		SetSubTitle(r.SubTitle).
		SetTags(r.Tags).
		SetContent(r.Content).
		Save(c)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": err,
		})
		return
	}

	blogDB.Log(result.ID, c.Request.RequestURI, xconst.ReasonPost, c.ClientIP(), result)

	c.Status(http.StatusOK)
}
