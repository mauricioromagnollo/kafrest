package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mauricioromagnollo/kafrest/external/config"
	"github.com/segmentio/kafka-go"
)

// PublishController is the struct that will handle the /publish endpoint.
type PublishController struct {
	Environment *config.Environment
}

type publishRequest struct {
	Topic    string                   `json:"topic"`
	Messages []map[string]interface{} `json:"messages"`
}

type publishResponse struct {
	MessagesPublished []map[string]interface{} `json:"messages_published"`
}

// NewPublishController is the constructor for the PublishController struct.
func NewPublishController(env *config.Environment) *PublishController {
	return &PublishController{
		Environment: env,
	}
}

// Handle is the function that will be called when a request is made to the /publish endpoint.
func (pc PublishController) Handle(w http.ResponseWriter, r *http.Request) {
	var requestData publishRequest
	var responseData publishResponse

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
			_, err := w.Write([]byte(err.Error()))
			if err != nil {
				panic(err)
			}

			return
		}

		if err := kw.WriteMessages(context.Background(), kafka.Message{Value: parsedMessage}); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte(err.Error()))
			if err != nil {
				panic(err)
			}

			return
		}

		responseData.MessagesPublished = append(responseData.MessagesPublished, message)
	}

	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			panic(err)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, werr := w.Write(responseJSON)
	if err != nil {
		panic(werr)
	}
}
