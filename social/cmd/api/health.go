package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	//w.Header().Add("content-type", "application/json")
	//_, err := w.Write([]byte(`{"status":"ok"}`))
	//if err != nil {
	//	return
	//}
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}
	if err := writeJSON(w, http.StatusOK, data); err != nil {
		err := writeJSONError(w, http.StatusInternalServerError, err.Error())
		if err != nil {
			return
		}
	}
}
