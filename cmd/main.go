package main

import (
	"GetSongsDataService/internal/server"
	logrusCustom "GetSongsDataService/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrusCustom.InitLogger()

	err := godotenv.Load()
	if err != nil {
		logrusCustom.Logger.Fatalf("Error loading .env file")
	}

	logrusCustom.LogWithLocation(logrus.InfoLevel, "Successfully loaded environment variables")

	appGRPC := server.NewApp()

	err = appGRPC.Run(os.Getenv("APP_PORT"))
	if err != nil {
		logrusCustom.Logger.Fatalf("Error running server")
	}
}
