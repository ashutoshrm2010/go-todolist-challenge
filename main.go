package main

import (
	"net/http"
	"os"

	"github.com/PeakActivity/go-todolist-challenge/app"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// Loads the env variables
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	// Prints the address of our api to the console
	logrus.Info("Starting Server on http://localhost:", os.Getenv("PORT"))

	// Server router on given port
	server := app.NewServer()
	logrus.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), server))
}
