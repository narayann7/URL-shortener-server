package services

import (
	"log"

	"github.com/spf13/viper"
)

func Getenv(key string) string {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		//todo log fatal f means
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
