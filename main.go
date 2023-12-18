package main

import (
	"log"
	"word-validation/controller"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load configuration")
	}
	logrus.Info("successfully load configuration")
	ctl := controller.NewController()
	ctl.Run()
}
