package route

import (
	"io/fs"
	"net/http"
	"os"
	"sort"

	"github.com/gin-gonic/gin"
)

func readDir(dirname string) ([]fs.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}

func Mobile(c *gin.Context) {
	targetDir := "./assets/gallary"
	fileNames := make([]string, 0, 30)
	files, _ := readDir(targetDir)
	for _, file := range files {
		fileNames = append(fileNames, targetDir[1:]+"/"+file.Name())
	}

	c.HTML(http.StatusOK, "weding-mobile.tmpl", gin.H{
		"title": "정종선 ❤️ 김지안 결혼합니다.",
		"files": fileNames,
	})
}
