package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	payload := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}

	if err := writeJSON(w, http.StatusOK, payload); err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
