package test

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReqPost(uri string, body interface{}) (req *http.Request) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req, _ = http.NewRequest("POST", uri, bytes.NewBuffer(jsonBody))
	req.Header.Add("Content-Type", gin.MIMEJSON)
	return
}

func ReqPut(uri string, body interface{}) (req *http.Request) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req, _ = http.NewRequest("PUT", uri, bytes.NewBuffer(jsonBody))
	req.Header.Add("Content-Type", gin.MIMEJSON)
	return
}

func ReqGet(uri string) (req *http.Request) {
	req, _ = http.NewRequest("GET", uri, nil)
	req.Header.Add("Content-Type", gin.MIMEJSON)
	return
}

func ReqDelete(uri string) (req *http.Request) {
	req, _ = http.NewRequest("DELETE", uri, nil)
	req.Header.Add("Content-Type", gin.MIMEJSON)
	return
}
