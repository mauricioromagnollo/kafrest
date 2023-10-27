package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mauricioromagnollo/kafrest/external/config"
	"github.com/segmentio/kafka-go"
)

// PublishTopicMessagesController is the struct that will handle the /topics/{topicName}/messages endpoint.
type PublishTopicMessagesController struct {
	Environment *config.Environment
	Logger      *config.Logger
}

type publishTopicsRequest struct {
	Messages []map[string]interface{} `json:"messages"`
}

type publishTopicsResponse struct {
	Topic             string                   `json:"topic"`
	PublishedMessages []map[string]interface{} `json:"published_messages"`
}

// NewPublishTopicMessagesController is the constructor for the PublishTopicMessages struct.
func NewPublishTopicMessagesController(env *config.Environment, logger *config.Logger) *PublishTopicMessagesController {
	return &PublishTopicMessagesController{
		Environment: env,
		Logger:      logger,
	}
}

// Handle is the function that will be called when a request is made to the /publish endpoint.
func (pc PublishTopicMessagesController) Handle(w http.ResponseWriter, r *http.Request) {
	var requestData publishTopicsRequest
	var responseData publishTopicsResponse

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		return
	}

	topicName := chi.URLParam(r, "topicName")

	kw := kafka.Writer{
		Addr:     kafka.TCP(pc.Environment.KafkaHost),
		Topic:    topicName,
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

		responseData.PublishedMessages = append(responseData.PublishedMessages, message)
	}

	responseData.Topic = topicName
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
