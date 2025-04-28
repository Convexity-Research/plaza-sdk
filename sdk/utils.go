package sdk

import (
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if _, err := os.Stat(".env"); err == nil {
		godotenv.Load()
	}
}
