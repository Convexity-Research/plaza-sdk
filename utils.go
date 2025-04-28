package main

import (
	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}
}
