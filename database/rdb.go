package database

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
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
		data, err := rdb0.Get(Ctx, hash).Result()
		if len(data) == 0 && err != nil {
			//if err is not nil which means the hash has no value to it
			//and is not used for any url
			//length of data == 0 means empty string came
			continueLoop = false
		}
	}
	return hash
}

func SetUrlInDb(url string, urlHash string, expiry time.Duration) bool {

	rdb0 := DatabaseInit(0)
	defer rdb0.Close()

	status, err := rdb0.Set(Ctx, urlHash, url, expiry).Result()
	if status == "OK" && err == nil {
		return true
	} else {
		return false
	}

}
func CheckIsHashUnique(hash string) bool {

	rdb0 := DatabaseInit(0)
	defer rdb0.Close()

	data, err := rdb0.Get(Ctx, hash).Result()
	fmt.Println(data)
	fmt.Println(err)
	//if err is not nil which means the hash has no value to it
	//and is not used for any url
	//length of data == 0 means empty string came
	if len(data) == 0 && err != nil {
		return true
	} else {
		return false
	}

}
func CounterReducer(ip string) error {
	rdb1 := DatabaseInit(1)
	defer rdb1.Close()

	quota, err := rdb1.Get(Ctx, ip).Result()
	if err != nil {
		if err == redis.Nil {
			return errors.New("rate limit exceeded.reload the limit")
		} else {
			return err
		}
	}
	timeRemaining, _ := rdb1.Get(Ctx, ip).Time()
	quotaInInt, err := strconv.Atoi(quota)
	if err != nil {
		return err
	}
	if quotaInInt <= 0 {
		return errors.New("rate limit exceeded, please try after :" + timeRemaining.String())
	}
	_, err = rdb1.Decr(Ctx, ip).Result()
	if err != nil {
		return err
	}
	return nil
}
