package http_server

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"
	"svc-subman/src/ports_adapters/primary/http_server/dto"
)

func (s Server) AgreeSubscribe(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	startDate, err := parseMonthYear(query.Get("start_date"))
	if err != nil {
		response := Error(errors.New("failed to get start_date"), http.StatusBadRequest)
		writeResponse(w, response)
		return
	}
	endDate, err := parseMonthYear(query.Get("end_date"))
	if err != nil {
		response := Error(errors.New("failed to get ent_date"), http.StatusBadRequest)
		writeResponse(w, response)
		return
	}
	request := dto.GetSubscriberAgree{
		UserID:      getStringPtr(query.Get("user_id")),
		ServiceName: getStringPtr(query.Get("service_name")),
		StartDate:   *startDate,
		EndDate:     *endDate,
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		response := Error(err, http.StatusBadRequest)
		writeResponse(w, response)
		return
	}

	resp, err := s.app.Commands.AgreeSubscribe.Handle(r.Context(), request)
	if err != nil {
		response := Error(err, http.StatusInternalServerError)
		writeResponse(w, response)
		return
	}
	response := Success(resp)
	writeResponse(w, response)
}
