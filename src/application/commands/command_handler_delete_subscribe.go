package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"strings"
	"svc-subman/src/common/decorator"
	"svc-subman/src/domain"
	"svc-subman/src/ports_adapters/primary/http_server/dto"
)

type deleteSubscribeCommand struct {
	repo   domain.ISubscribeRepository
	logger *slog.Logger
}

type DeleteSubscribeCommand decorator.CommandHandlerDecorator[dto.DeleteSubscribe, *string]

func NewDeleteSubscribeCommand(repo domain.ISubscribeRepository, logger *slog.Logger) decorator.CommandHandlerDecorator[dto.DeleteSubscribe, *string] {

	return decorator.ApplyCommandDecorator[dto.DeleteSubscribe, *string](deleteSubscribeCommand{repo: repo, logger: logger}, logger)
}

func (cr deleteSubscribeCommand) Handle(ctx context.Context, cmd dto.DeleteSubscribe) (*string, error) {
	var errorS []error
	var failedIDs []uuid.UUID

	for _, v := range cmd.Ids {
		err := cr.repo.Delete(ctx, v)
		if err != nil {
			cr.logger.Error("failed delete subscribe", err)
			errorS = append(errorS, err)
			failedIDs = append(failedIDs, v)
		}
	}

	if len(errorS) > 0 {
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("failed to delete %d subscriptions: [", len(failedIDs)))
		for i, id := range failedIDs {
			sb.WriteString(fmt.Sprintf("%d (%v)", id, errorS[i]))
			if i < len(failedIDs)-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString("]")
		return nil, errors.New(sb.String())
	}

	return nil, nil
}
