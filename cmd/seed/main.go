package main

import (
	"log"
	"paw-me-back/internal/db"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Couldn't load .env file")
	}
	conn, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	db.Seed(conn)
}
