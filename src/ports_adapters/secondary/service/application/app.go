package application

import (
	"log/slog"
	"svc-subman/src/application"
	"svc-subman/src/application/commands"
	"svc-subman/src/domain"
)

func InitApp(db domain.ISubscribeRepository, logger *slog.Logger) (*application.Application, error) {
	return &application.Application{
		Commands: application.Command{
			CreateSubscribe: commands.NewCreateSubscribeCommand(db, logger),
			DeleteSubscribe: commands.NewDeleteSubscribeCommand(db, logger),
			AgreeSubscribe:  commands.NewAgreeSubscriberQuvery(db, logger),
			GetSubscribes:   commands.NewGetSubscribesQuvery(db, logger),
			GetSubscribe:    commands.NewGetSubscriberQuvery(db, logger)},
	}, nil
}
