package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mauricioromagnollo/kafkarest/external/config"
	"github.com/segmentio/kafka-go"
)

type PublishController struct {
	Environment *config.Environment
}

type PublishRequest struct {
	Topic    string                   `json:"topic"`
	Messages []map[string]interface{} `json:"messages"`
}

type PublishResponse struct {
	MessagesPublished []map[string]interface{} `json:"messages_published"`
}

func NewPublishController(env *config.Environment) *PublishController {
	return &PublishController{
		Environment: env,
	}
}

func (pc PublishController) Handle(w http.ResponseWriter, r *http.Request) {
	var requestData PublishRequest
	var responseData PublishResponse

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		return
	}

	kw := kafka.Writer{
		Addr:     kafka.TCP(pc.Environment.KafkaHost),
		Topic:    requestData.Topic,
		Balancer: &kafka.LeastBytes{},
	}

	for _, message := range requestData.Messages {
		parsedMessage, err := json.Marshal(message)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if err := kw.WriteMessages(context.Background(), kafka.Message{Value: parsedMessage}); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		responseData.MessagesPublished = append(responseData.MessagesPublished, message)
	}

	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)
}
