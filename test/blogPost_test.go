package test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/jungmu/go-web/api"
)

type reqPost struct {
	Title string `uri:"title"`
	reqPostBody
}

type reqPostBody struct {
	SubTitle string `json:"subTitle"`
	Tags     string `json:"tags"`
	Content  string `json:"content"`
	PostKey  string `json:"postKey"`
}

func TestPost(t *testing.T) {
	r := gin.Default()

	api.Init(r)

	res := httptest.NewRecorder()

	reqStruct := reqPost{
		"qaasw",
		reqPostBody{
			"kkk1",
			"tag1, tag2",
			"# bbbbb\n ## C123\n ```co123123123de```\n ``` cc123123123c\n ```\n con12312312tent",
			"jungmu254!",
		},
	}

	req := ReqPost(fmt.Sprintf("/blog/article/%s", reqStruct.Title), reqStruct)

	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
}
