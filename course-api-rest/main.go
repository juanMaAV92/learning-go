package main

import (
	"context"
	"log"
	"os"

	"api.com/go/rest-ws/handlers"
	"api.com/go/rest-ws/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	server, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DataBaseURL: DATABASE_URL,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = server.Start(BindRoutes)
	if err != nil {
		log.Fatal(err)
	}

}

func BindRoutes(s server.Server, r *mux.Router) error {
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods("GET")
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods("POST")
	return nil
}
