package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

func TestReadFromVars(t *testing.T) {
	expected := 9
	url := fmt.Sprintf("/api/v1/numbers/%d", expected)
	r, _ := http.NewRequest("GET", url, nil)

	// Hack to try to fake gorilla/mux vars
	// Got from: https://stackoverflow.com/questions/34435185/unit-testing-for-functions-that-use-gorilla-mux-url-parameters
	vars := map[string]string{
		"id": fmt.Sprintf("%d", expected),
	}
	r = mux.SetURLVars(r, vars)

	id, err := readFromMuxVars(r, "id")
	if err != nil {
		t.Error(err)
		return
	}

	if id != expected {
		t.Errorf("readFromVars is not working. Expected: %d. Got: %d", expected, id)
		return
	}
}
