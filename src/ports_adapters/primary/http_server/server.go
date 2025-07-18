package http_server

import (
	"github.com/go-playground/validator/v10"
	"svc-subman/src/application"
)

var validate = validator.New()

type Server struct {
	app *application.Application
}

func NewServer(app *application.Application) *Server {
	return &Server{app: app}
}
