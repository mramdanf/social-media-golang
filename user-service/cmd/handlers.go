package main

import (
	"fmt"
	"net/http"
)

func (app *Config) testGet(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("yuhu"),
	}
	app.writeJSON(w, http.StatusAccepted, payload)
}