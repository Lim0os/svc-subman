package http_server

import (
	"errors"
	"net/http"
	"svc-subman/src/ports_adapters/primary/http_server/dto"
)

func (s Server) GetSubscribes(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	offset, err := getIntPtrStrict(query.Get("offset"))
	if err != nil {
		response := Error(errors.New("failed to get offset"), http.StatusBadRequest)
		writeResponse(w, response)
		return
	}
	limit, err := getIntPtrStrict(query.Get("limit"))
	if err != nil {
		response := Error(errors.New("failed to get limit"), http.StatusBadRequest)
		writeResponse(w, response)
		return
	}

	request := dto.GetUserSubscribe{
		UserID:      getStringPtr(query.Get("user_id")),
		ServiceName: getStringPtr(query.Get("service_name")),
		Offset:      offset,
		Limit:       limit,
	}

	resp, err := s.app.Commands.GetSubscribes.Handle(r.Context(), request)
	if err != nil {
		response := Error(err, http.StatusInternalServerError)
		writeResponse(w, response)
		return
	}
	response := Success(resp, http.StatusOK)
	writeResponse(w, response)
	return
}
