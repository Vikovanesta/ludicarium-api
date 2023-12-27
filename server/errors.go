package server

import (
	"log"
	"net/http"
)

func serverError(w http.ResponseWriter, err error) {
	log.Println(err)
	errorResponse(w, http.StatusInternalServerError, "internal error")
}

func errorResponse(w http.ResponseWriter, code int, errs interface{}) {
	writeJson(w, code, M{"errors": errs})
}
