package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv()  {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("an error occured : %s", err.Error())
		return 
	}
}
