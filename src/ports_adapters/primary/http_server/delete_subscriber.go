package http_server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"svc-subman/src/ports_adapters/primary/http_server/dto"
)

func (s Server) DeleteSubscriber(w http.ResponseWriter, r *http.Request) {
	var req dto.DeleteSubscribe
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := Error(err, http.StatusBadRequest)
		writeResponse(w, response)
		return
	}
	if err := validate.Struct(req); err != nil {
		writeResponse(w, Error(err, http.StatusBadRequest))
		return
	}
	if len(req.Ids) == 0 {
		response := Error(fmt.Errorf("list ids is null"), http.StatusBadRequest)
		writeResponse(w, response)
		return
	}

	_, err := s.app.Commands.DeleteSubscribe.Handle(r.Context(), req)
	if err != nil {
		response := Error(err, http.StatusBadRequest)
		writeResponse(w, response)
		return
	}
	response := Success("All subscribers have been deleted")
	writeResponse(w, response)
	return
}
