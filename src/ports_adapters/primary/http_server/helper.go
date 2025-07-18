package http_server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"svc-subman/src/ports_adapters/primary/http_server/dto"
	"time"
)

const (
	StatusSuccess = "success"
	StatusError   = "error"
)

type IResponse interface {
	JSON() ([]byte, error)
	StatusCode() int
}

type SuccessResponse[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
	Meta   any    `json:"meta,omitempty"`
}

func (r *SuccessResponse[T]) JSON() ([]byte, error) {
	r.Status = StatusSuccess
	return json.Marshal(r)
}

func (r *SuccessResponse[T]) StatusCode() int {
	return 200
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Error   string `json:"error"`
	Details any    `json:"details,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func (r *ErrorResponse) JSON() ([]byte, error) {
	r.Status = StatusError
	return json.Marshal(r)
}

func (r *ErrorResponse) StatusCode() int {
	if r.Code >= 400 && r.Code < 600 {
		return r.Code
	}
	return 400
}

var responseGenerator = struct {
	Success func(data any, meta ...any) IResponse
	Error   func(err error, details any, code ...int) IResponse
}{
	Success: func(data any, meta ...any) IResponse {
		resp := &SuccessResponse[any]{
			Data: data,
		}
		if len(meta) > 0 {
			resp.Meta = meta[0]
		}
		return resp
	},
	Error: func(err error, details any, code ...int) IResponse {
		resp := &ErrorResponse{
			Error:   err.Error(),
			Details: details,
		}
		if len(code) > 0 {
			resp.Code = code[0]
		}
		return resp
	},
}

func Success(data any, meta ...any) IResponse {
	return responseGenerator.Success(data, meta...)
}

func Error(err error, details any, code ...int) IResponse {
	return responseGenerator.Error(err, details, code...)
}

func writeResponse(w http.ResponseWriter, response IResponse) {
	jsonData, err := response.JSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"status":"error","error":"failed to encode response"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode())
	w.Write(jsonData)
}

func getStringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func getIntPtrStrict(s string) (int, error) {
	if s == "" {
		return 0, errors.New("string is empty")
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %s", s)
	}
	return num, nil
}

func parseMonthYear(s string) (*dto.MonthYear, error) {
	if s == "" {
		return nil, errors.New("data is required")
	}
	t, err := time.Parse(dto.MonthYearFormat, s)
	if err != nil {
		return nil, errors.New("data is required")
	}
	return &dto.MonthYear{Time: t}, nil
}
