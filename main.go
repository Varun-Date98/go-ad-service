package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Varun-Date98/go-ad-service/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type dbAPI struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT not found in environment")
	}

	dbURL := os.Getenv("DB_URL")
	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Could not connect to postgres DB,", err)
	}

	db := dbAPI{
		DB: database.New(conn),
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
	v1Router.Get("/ad", db.adHandler)
	v1Router.Get("/err", errorHandler)
	v1Router.Get("/healthz", readinessHandler)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Println("Starting ad server on port:", portString)
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal("Server error occurred:", err)
	}
}
