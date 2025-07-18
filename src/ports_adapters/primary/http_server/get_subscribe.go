package http_server

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

func (s Server) GetSubscribe(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		response := Error(errors.New("params id is requed"), http.StatusBadRequest)
		writeResponse(w, response)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		response := Error(errors.New("params id is not valid uuid"), http.StatusBadRequest)
		writeResponse(w, response)
		return
	}
	resp, err := s.app.Commands.GetSubscribe.Handle(r.Context(), id)
	if err != nil {
		response := Error(err, http.StatusInternalServerError)
		writeResponse(w, response)
		return
	}
	response := Success(resp, http.StatusOK)
	writeResponse(w, response)
	return
}
