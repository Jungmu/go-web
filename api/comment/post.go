package comment

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jungmu/go-web/db"
	blogDB "github.com/jungmu/go-web/db/blog"
	"github.com/jungmu/go-web/ent/blog"
	"github.com/jungmu/go-web/xconst"
)

type reqPost struct {
	Title string `uri:"title"`
	reqPostBody
}

type reqPostBody struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	Content   string `json:"content"`
	CommentID int64  `json:"commentID"`
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
	hash.Write([]byte(r.Password))
	password := hex.EncodeToString(hash.Sum(nil))

	result, err := db.Client().Comment.Create().
		SetName(r.Name).
		SetPassword(password).
		SetContent(r.Content).
		SetCommentID(r.CommentID).
		Save(c)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": err,
		})
		return
	}

	err = db.Client().Blog.Update().
		Where(
			blog.Title(r.Title),
		).
		AddComments(result).
		Exec(c)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": err,
		})
		return
	}

	blogDB.Log(result.ID, c.Request.RequestURI, xconst.ReasonPostComment, c.ClientIP(), result)

	c.Status(http.StatusOK)
}
