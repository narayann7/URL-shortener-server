package services

import (
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/narayann7/gourl/database"
)

func Getenv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func GetNewHash() string {
	id := uuid.New()
	hash := strings.Split(id.String(), "-")[0]
	hash = "testing"

	rdb0 := database.DatabaseInit(0)
	rdb0.Get(database.Ctx, "testing")

	return hash
}
