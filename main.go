package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jungmu/go-web/api"
)

func main() {
	r := gin.Default()

	api.Init(r)

	r.Run(":6666")
}
