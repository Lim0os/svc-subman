package http_server

import (
	"encoding/json"
	_ "github.com/go-playground/validator/v10"
	"net/http"
	"svc-subman/src/ports_adapters/primary/http_server/dto"
)

func (s Server) CreateSubscribe(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateSubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := Error(err, http.StatusBadRequest)
		writeResponse(w, response)
		return
	}
	if err := validate.Struct(req); err != nil {
		writeResponse(w, Error(err, http.StatusBadRequest))
		return
	}

	id, err := s.app.Commands.CreateSubscribe.Handle(r.Context(), req)
	if err != nil {
		response := Error(err, http.StatusBadRequest)
		writeResponse(w, response)
		return
	}
	response := Success(id)
	writeResponse(w, response)
	return
}
