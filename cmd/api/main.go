package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/mauricioromagnollo/kafrest/external/config"
	"github.com/mauricioromagnollo/kafrest/external/controllers"
)

func main() {
	env := config.NewEnvironment()

	logger, err := config.NewLogger(env.AppEnv != "test")
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	statusController := controllers.NewStatusController(env)
	publishTopicMessagesController := controllers.NewPublishTopicMessagesController(env, logger)

	r := chi.NewRouter()

	r.Get("/status", statusController.Handle)

	r.Route("/topics", func(r chi.Router) {
		r.Route("/{topicName}", func(r chi.Router) {
			r.Post("/messages", publishTopicMessagesController.Handle)
		})
	})

	port := fmt.Sprintf(":%v", env.AppPort)

	server := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	logger.Info("Starting server")
	if err := server.ListenAndServe(); err != nil {
		logger.Error(fmt.Sprintf("Failed to start server %v", err))
		panic(err)
	}
}
