package commands

import (
	"context"
	"log/slog"
	"svc-subman/src/common/decorator"
	"svc-subman/src/domain"
	"svc-subman/src/ports_adapters/primary/http_server/dto"
)

type agreeSubscriberQuvery struct {
	repo   domain.ISubscribeRepository
	logger *slog.Logger
}

type AgreeSubscriberQuvery decorator.CommandHandlerDecorator[dto.GetSubscriberAgree, int]

func NewAgreeSubscriberQuvery(repo domain.ISubscribeRepository, logger *slog.Logger) decorator.CommandHandlerDecorator[dto.GetSubscriberAgree, int] {

	return decorator.ApplyCommandDecorator[dto.GetSubscriberAgree, int](agreeSubscriberQuvery{repo: repo, logger: logger}, logger)
}

func (cr agreeSubscriberQuvery) Handle(ctx context.Context, cmd dto.GetSubscriberAgree) (int, error) {
	resp, err := cr.repo.GetTotalCost(ctx, cmd.UserID, cmd.ServiceName, cmd.StartDate.Time, cmd.EndDate.Time)
	if err != nil {
		return 0, err
	}
	return resp, nil

}
