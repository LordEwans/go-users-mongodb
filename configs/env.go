package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func checkErr() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
}

func EnvMongoURI() string {
	checkErr()

	return os.Getenv("MONGOURI")
}

func EnvPort() string {
	checkErr()

	return os.Getenv("PORT")
}
