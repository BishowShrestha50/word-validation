package controller

import (
	"errors"
	"net/http"
	"strings"
	"word-validation/model"
	"word-validation/utils"
)

func (s *Controller) WordsHandler(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query().Get("word")
	if word == "" {
		utils.ERROR(w, http.StatusBadRequest, errors.New("required word"))
		return
	}
	valid := s.Words[strings.ToLower(word)]
	utils.JSON(w, http.StatusOK, model.Response{
		Valid: valid,
	})
}

func (s *Controller) Health(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, http.StatusOK, "OK")
}
