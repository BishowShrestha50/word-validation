package controller

import (
	"encoding/json"
	"net/http"
	"strings"
	"word-validation/model"
)

func (s *Controller) WordsHandler(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query().Get("word")
	valid := s.Words[strings.ToLower(word)]
	response := model.Response{
		Valid: valid,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Controller) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("ok")
}
