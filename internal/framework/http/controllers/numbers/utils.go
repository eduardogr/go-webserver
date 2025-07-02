package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func readFromMuxVars(r *http.Request, key string) (int, error) {
	vars := mux.Vars(r)
	return strconv.Atoi(vars[key])
}
