package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mauricioromagnollo/kafrest/external/config"
)

type StatusController struct {
	Environment *config.Environment
}

type StatusResponse struct {
	Message string `json:"message"`
	Broker  string `json:"broker"`
}

func NewStatusController(env *config.Environment) *StatusController {
	return &StatusController{
		Environment: env,
	}
}

func (sc StatusController) Handle(w http.ResponseWriter, _ *http.Request) {
	responseData := StatusResponse{
		Message: "OK",
		Broker:  sc.Environment.KafkaHost,
	}

	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
