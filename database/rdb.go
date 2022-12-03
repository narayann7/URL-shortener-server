package database

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var Ctx = context.Background()

func DatabaseInit(dbId int) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("RDB_URL"),
		Password: "",
		DB:       dbId,
	})
	return rdb

}
func GetNewHash() string {
	hash := ""
	var continueLoop bool = true
	rdb0 := DatabaseInit(0)
	defer rdb0.Close()
	//loop will continue till it find a hash which is not used
	for continueLoop {
		id := uuid.New()
		hash = strings.Split(id.String(), "-")[0]
		_, err := rdb0.Get(Ctx, hash).Result()
		if err != nil {
			//if err is not nil which means the hash has no value to it
			//and is not used for url
			continueLoop = false
		}
	}
	return hash
}

func SetUrlInDb(url string, urlHash string, expiry time.Duration) bool {
	rdb0 := DatabaseInit(0)
	defer rdb0.Close()

	status, err := rdb0.Set(Ctx, urlHash, url, expiry).Result()

	if status == "OK" && err != nil {
		return true
	} else {
		return false
	}

}
