package decorator

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
)

type CommandHandlerDecorator[C any, R any] interface {
	Handle(ctx context.Context, cmd C) (R, error)
}

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}

func ApplyCommandDecorator[C any, R any](
	handler CommandHandlerDecorator[C, R],
	logger *slog.Logger,
) CommandHandlerDecorator[C, R] {
	return CommandLoggingDecorator[C, R]{
		logger: logger,
		base:   handler,
	}
}
