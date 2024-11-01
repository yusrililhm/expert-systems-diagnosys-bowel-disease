package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv()  {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("an error occured : %s", err.Error())
		return 
	}
}
