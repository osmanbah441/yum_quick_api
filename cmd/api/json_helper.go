package main

import (
	"encoding/json"
	"net/http"
)

type envelop map[string]any

func (app *application) writeJSON(w http.ResponseWriter, data envelop, status int, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	return err

}
