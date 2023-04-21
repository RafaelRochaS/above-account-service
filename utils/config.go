package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var PORT int
var HOST string

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	HOST = os.Getenv(("HOST"))
	PORT, err = strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		PORT = 50000
	}
}
