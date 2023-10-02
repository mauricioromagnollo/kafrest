package controllers

import (
	"net/http"

	"github.com/mauricioromagnollo/kafkarest/external/config"
)

type StatusController struct {
	Environment *config.Environment
}

func NewStatusController(env *config.Environment) *StatusController {
	return &StatusController{
		Environment: env,
	}
}

func (sc StatusController) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Published successfully"))
}
