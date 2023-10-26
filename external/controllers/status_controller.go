package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mauricioromagnollo/kafrest/external/config"
)

// StatusController is the struct that will handle the /status endpoint.
type StatusController struct {
	Environment *config.Environment
}

type statusResponse struct {
	Message string `json:"message"`
	Broker  string `json:"broker"`
}

// NewStatusController is the constructor for the StatusController struct.
func NewStatusController(env *config.Environment) *StatusController {
	return &StatusController{
		Environment: env,
	}
}

// Handle is the function that will be called when a request is made to the /status endpoint.
func (sc StatusController) Handle(w http.ResponseWriter, _ *http.Request) {
	responseData := statusResponse{
		Message: "OK",
		Broker:  sc.Environment.KafkaHost,
	}

	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			panic(err)
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	_, werr := w.Write(responseJSON)
	if werr != nil {
		panic(werr)
	}
}
