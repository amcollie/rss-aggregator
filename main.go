package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
        log.Fatal("PORT not set in the environment")
    }

	router := chi.NewRouter()
	router.Use(corsMiddleware)

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerError)


	srv := &http.Server {
		Addr: ":" + portString, 
		Handler: router,
	}

	router.Mount("/v1", v1Router)

	log.Printf("Server is running on port %s\n", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
    }
}