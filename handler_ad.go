package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)


func adHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	ageStr := q.Get("age")
	age := 0
	
	if ageStr != "" {
		parsed, err := strconv.Atoi(ageStr)
		
		if err == nil {
			age = parsed
		} else {
			log.Printf("Could not get age from query, %v", err)
		}
	}

	user := UserContext{
		UserID:  q.Get("userId"),
		Age:     age,
		Country: q.Get("country"),
	}

	placement := PlacementContext{
		PlacementID: q.Get("placementId"),
	}

	decision := selectAd(user, placement)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")

	if decision == nil {
		// No eligible campaign for this user/context
		w.WriteHeader(http.StatusNoContent) // 204
		return
	}

	if err := json.NewEncoder(w).Encode(decision); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}
