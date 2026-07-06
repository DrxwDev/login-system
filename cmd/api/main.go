package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"

	"github.com/DrxwDev/login-system/internal/app"
)

func init() {
	if os.Getenv("RENDER") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("No .env file found")
		}
	}
}

func main() {
	fx.New(
		app.Module,
	).Run()
}
