package controller

import "net/http"

func (ctl Controller) Routes() {
	http.HandleFunc("/v1/validate", ctl.WordsHandler)
	http.HandleFunc("/health", ctl.Health)

}
