package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net/http"
	"workshop/internal/config"

	"github.com/go-chi/chi"

	"workshop/internal/handler"
)

func main() {
	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml" , &cfg)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler()

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	log.Print("starting server")
	err = http.ListenAndServe(":8080", r)
	log.Fatal(err)

	log.Print("shutting server down")
}
