package commands

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"svc-subman/src/common/decorator"
	"svc-subman/src/domain"
)

type getSubscriberQuvery struct {
	repo   domain.ISubscribeRepository
	logger *slog.Logger
}

type GetSubscriberQuvery decorator.CommandHandlerDecorator[uuid.UUID, domain.Subscription]

func NewGetSubscriberQuvery(repo domain.ISubscribeRepository, logger *slog.Logger) decorator.CommandHandlerDecorator[uuid.UUID, domain.Subscription] {

	return decorator.ApplyCommandDecorator[uuid.UUID, domain.Subscription](getSubscriberQuvery{repo: repo, logger: logger}, logger)
}

func (cr getSubscriberQuvery) Handle(ctx context.Context, cmd uuid.UUID) (domain.Subscription, error) {
	model, err := cr.repo.GetByID(ctx, cmd)
	if err != nil {
		return domain.Subscription{}, err
	}
	response, _ := convertToDomain(model)
	return response, nil

}
