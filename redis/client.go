package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

const (
	CHAT_KEY = "chat:simple"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func InsertChat(chat string) error {
	return rdb.LPush(context.TODO(), CHAT_KEY, chat).Err()
}

func GetLastChatList() (chatList []string, err error) {
	chatList, err = rdb.LRange(context.TODO(), CHAT_KEY, 0, 100).Result()
	if err != nil {
		log.Println(err)
	}
	for i, j := 0, len(chatList)-1; i < j; i, j = i+1, j-1 {
		chatList[i], chatList[j] = chatList[j], chatList[i]
	}
	return
}

func Set(key string, value string) error {
	return rdb.Set(context.TODO(), key, value, 0).Err()
}

func Get(key string) (string, error) {
	return rdb.Get(context.TODO(), "key").Result()
}

func Del(keys []string) (int64, error) {
	return rdb.Del(context.TODO(), keys...).Result()
}
