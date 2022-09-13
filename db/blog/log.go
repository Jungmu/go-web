package blog

import (
	"context"
	"log"

	"github.com/goccy/go-json"
	"github.com/jungmu/go-web/db"
	"github.com/jungmu/go-web/ent/bloglog"
	"github.com/jungmu/go-web/xconst"
)

func Log(blogID int64, url string, reason string, clientIP string, detail interface{}) {
	go syncLog(blogID, url, reason, clientIP, detail)
}

func syncLog(blogID int64, url string, reason string, clientIP string, detail interface{}) {
	jsonDetail, err := json.Marshal(detail)
	if err != nil {
		log.Printf("blog log detail errror : %s", err.Error())
	}

	_, err = db.Client().BlogLog.Create().
		SetBlogID(blogID).
		SetURL(url).
		SetReason(reason).
		SetDetail(string(jsonDetail)).
		SetClientIP(clientIP).
		Save(context.TODO())

	if err != nil {
		log.Printf("blog log create errror : %s", err.Error())
	}
}

func GetViewCount(id int64) (count int, err error) {
	count, err = db.Client().BlogLog.Query().
		Where(
			bloglog.BlogID(id),
			bloglog.Reason(xconst.ReasonView),
		).
		Count(context.TODO())

	return
}

func GetTotalViewCount() (count int, err error) {
	count, err = db.Client().BlogLog.Query().
		Where(
			bloglog.Reason(xconst.ReasonView),
		).
		Count(context.TODO())

	return
}
