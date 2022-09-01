package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jungmu/go-web/db"
)

type req struct {
	ID int64 `uri:"id"`
}

func Get(c *gin.Context) {
	r := req{}
	c.BindUri(&r)

	u, err := db.Client().User.Get(c, r.ID)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, u)
}
