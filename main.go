package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	seedCampaigns()
	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT not found in environment")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/ad", adHandler)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Println("Starting ad server on port:", portString)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Server error occurred:", err)
	}
}
