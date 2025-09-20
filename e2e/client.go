package e2e

import (
	"os"

	"github.com/NdoleStudio/neero-go"
	// Auto load .env file
	_ "github.com/joho/godotenv/autoload"
)

var client = neero.New(neero.WithSecretKey(os.Getenv("NEERO_SECRET_KEY")))
