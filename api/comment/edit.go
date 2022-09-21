package comment

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jungmu/go-web/db"
	blogDB "github.com/jungmu/go-web/db/blog"
	"github.com/jungmu/go-web/ent/comment"
	"github.com/jungmu/go-web/xconst"
)

type reqEdit struct {
	Title string `uri:"title"`
	reqEditBody
}

type reqEditBody struct {
	ID       int64  `json:"id"`
	Password string `json:"password"`
	Content  string `json:"content"`
}

func Edit(c *gin.Context) {
	r := reqEdit{}

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
	hash.Write([]byte(r.Password))
	password := hex.EncodeToString(hash.Sum(nil))

	result, err := db.Client().Comment.Update().
		Where(
			comment.ID(r.ID),
			comment.Password(password),
		).
		SetContent(r.Content).
		Save(c)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": err,
		})
		return
	}

	if result != 1 {
		err = fmt.Errorf("effect row count must be 1, count : %d", result)
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": err,
		})
		return
	}

	blogDB.Log(r.ID, c.Request.RequestURI, xconst.ReasonEditComment, c.ClientIP(), r.Content)

	c.Status(http.StatusOK)
}
