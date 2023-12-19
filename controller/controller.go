package controller

import (
	"fmt"
	"net/http"
	"os"
	"word-validation/service"

	"github.com/sirupsen/logrus"
)

type Controller struct {
	Words   map[string]bool
	service service.IWord
}

func NewController() *Controller {
	ctl := &Controller{}
	ctl.service = service.NewService()

	filePath := os.Getenv("WORD_FILE_PATH")
	if filePath == "" {
		filePath = "words.txt"
	}
	words, err := ctl.service.GetAllWords(filePath)
	if err != nil {
		logrus.Fatalf("unable to map words %v ", err)
	}
	ctl.Words = words
	ctl.Routes()
	return ctl
}

func (ctl *Controller) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logrus.Infof("server start on port %s ", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return fmt.Errorf("unable to run server on port %s :: %v ", port, err)
	}
	return nil
}
