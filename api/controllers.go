package api

import (
	"encoding/json"
	"net/http"
)

type HandlerFuncE func(http.ResponseWriter, *http.Request) error

func WriteJSON(writer http.ResponseWriter, statusCode int, v any) error {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	return json.NewEncoder(writer).Encode(v)
}

func HandleFunc(fn HandlerFuncE) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := fn(writer, request); err != nil {
			WriteJSON(writer, http.StatusInternalServerError, err.Error())
		}
	}
}
