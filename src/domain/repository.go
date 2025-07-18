package domain

import (
	"context"
	"github.com/google/uuid"
	"svc-subman/src/ports_adapters/secondary/persistence/db/models"
	"time"
)

type ISubscribeRepository interface {
	GetActiveSubscriptions(ctx context.Context, date time.Time) ([]*models.Subscription, error)
	GetTotalCost(ctx context.Context, userID *string, serviceName *string, startDate, endDate time.Time) (int, error)
	List(ctx context.Context, userID *string, serviceName *string, offset, limit int) ([]*models.Subscription, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, sub *models.Subscription) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Subscription, error)
	Create(ctx context.Context, sub *models.Subscription) error
}
