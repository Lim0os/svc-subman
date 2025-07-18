package commands

import (
	"context"

	"log/slog"
	"svc-subman/src/common/decorator"
	"svc-subman/src/domain"
	"svc-subman/src/ports_adapters/primary/http_server/dto"
)

type getSubscribesQuvery struct {
	repo   domain.ISubscribeRepository
	logger *slog.Logger
}

type GetSubscribesQuvery decorator.CommandHandlerDecorator[dto.GetUserSubscribe, []domain.Subscription]

func NewGetSubscribesQuvery(repo domain.ISubscribeRepository, logger *slog.Logger) decorator.CommandHandlerDecorator[dto.GetUserSubscribe, []domain.Subscription] {

	return decorator.ApplyCommandDecorator[dto.GetUserSubscribe, []domain.Subscription](getSubscribesQuvery{repo: repo, logger: logger}, logger)
}

func (cr getSubscribesQuvery) Handle(ctx context.Context, cmd dto.GetUserSubscribe) ([]domain.Subscription, error) {
	resp, err := cr.repo.List(ctx, cmd.UserID, cmd.ServiceName, cmd.Offset, cmd.Limit)
	if err != nil {
		return nil, err
	}
	response := make([]domain.Subscription, 0, len(resp))
	for _, sub := range resp {
		r, err := convertToDomain(sub)
		if err != nil {
			return nil, err
		}
		response = append(response, r)
	}
	return response, nil

}
