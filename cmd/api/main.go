package main

import (
	"fmt"
	"net/http"

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
	http.ListenAndServe(port, r)
}
