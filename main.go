package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jungmu/go-web/api"
)

func main() {
	r := gin.Default()
	api.Init(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
