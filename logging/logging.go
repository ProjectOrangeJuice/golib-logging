package logging

import (
	"log"
	"net/http"
)

func Log(function string, msg string, err error) {
	log.Printf("func='%s' message='%s' error='%s'", function, msg, err)
}

func LogHTTP(handler, msg string, err error, w http.ResponseWriter) {
	Log(handler, msg, err)
	http.Error(w, msg, http.StatusInternalServerError)
}
