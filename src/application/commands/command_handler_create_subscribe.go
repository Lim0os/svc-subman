package commands

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"svc-subman/src/common/decorator"
	"svc-subman/src/domain"
	"svc-subman/src/ports_adapters/primary/http_server/dto"
	"svc-subman/src/ports_adapters/secondary/persistence/db/models"
)

type createSubscribeCommand struct {
	repo   domain.ISubscribeRepository
	logger *slog.Logger
}

type CreateSubscribeCommand decorator.CommandHandlerDecorator[dto.CreateSubscriptionRequest, string]

func NewCreateSubscribeCommand(repo domain.ISubscribeRepository, logger *slog.Logger) decorator.CommandHandlerDecorator[dto.CreateSubscriptionRequest, string] {

	return decorator.ApplyCommandDecorator[dto.CreateSubscriptionRequest, string](createSubscribeCommand{repo: repo, logger: logger}, logger)
}

func (cr createSubscribeCommand) Handle(ctx context.Context, cmd dto.CreateSubscriptionRequest) (string, error) {
	reposModels := &models.Subscription{
		ID:          uuid.New(),
		ServiceName: cmd.ServiceName,
		Price:       cmd.Price,
		UserID:      cmd.UserID,
		StartDate:   cmd.StartDate.Time,
		EndDate:     &cmd.EndDate.Time,
	}

	err := cr.repo.Create(ctx, reposModels)
	if err != nil {
		cr.logger.Error("failed create subscribe", err)
		return "", err
	}
	return reposModels.ID.String(), nil
}
