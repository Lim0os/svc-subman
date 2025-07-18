package commands

import (
	"svc-subman/src/domain"
	"svc-subman/src/ports_adapters/secondary/persistence/db/models"
)

func convertToDomain(subscribe *models.Subscription) (domain.Subscription, error) {
	return domain.Subscription{
		ID:          subscribe.ID,
		ServiceName: subscribe.ServiceName,
		StartDate:   subscribe.StartDate,
		EndDate:     subscribe.EndDate,
		Price:       subscribe.Price,
		UserID:      subscribe.UserID,
	}, nil
}
