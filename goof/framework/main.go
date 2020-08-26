package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main()  {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("the root directory does not exist .env file")
	}
}