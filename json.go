package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface {}) {
	data, err := json.Marshal(payload)

	if err != nil {
		fmt.Printf("Failed to marshal json response %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error", msg)
	}

	type errorResp struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errorResp{msg})
}