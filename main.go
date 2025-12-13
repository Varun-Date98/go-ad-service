package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Varun-Date98/go-ad-service/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type dbAPI struct {
	DB *database.Queries
	Redis *redis.Client
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

	redisUrl := os.Getenv("REDIS_URL")
	rdb := redis.NewClient(&redis.Options{
		Addr: redisUrl,
		DB: 0,
	})

	_, err = rdb.Ping(context.Background()).Result()

	if err != nil {
		rdb = nil
		log.Println("Could not connect to redis, ad capping is turned off")
	}

	db := dbAPI{
		DB: database.New(conn),
		Redis: rdb,
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
