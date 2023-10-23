package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err!= nil {
		log.Printf("Fail to marshal payload: %v\n", payload)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	if statusCode > 499 {
		log.Println("Responding with 5XX error:", message)
	}

	type ErrorResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(w, statusCode, ErrorResponse{
		Error: message,
	})
}