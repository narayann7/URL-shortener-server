package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func DatabaseInit(dbId int) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("rdb_url"),
		Password: "",
		DB:       dbId,
	})
	return rdb

}
