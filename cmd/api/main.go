package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/mauricioromagnollo/kafrest/external/config"
	"github.com/mauricioromagnollo/kafrest/external/controllers"
)

func main() {
	env := config.NewEnvironment()

	statusController := controllers.NewStatusController(env)
	publishController := controllers.NewPublishController(env)

	r := chi.NewRouter()

	r.Get("/status", statusController.Handle)
	r.Post("/messages", publishController.Handle)

	port := fmt.Sprintf(":%v", env.AppPort)

	server := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
