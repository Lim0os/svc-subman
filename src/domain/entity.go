package domain

import (
	"github.com/google/uuid"
	"time"
)

// Subscription - сущность подписки пользователя
type Subscription struct {
	ID          uuid.UUID  `json:"id" `
	ServiceName string     `json:"service_name" `
	Price       int        `json:"price" `
	UserID      string     `json:"user_id" `
	StartDate   time.Time  `json:"start_date" `
	EndDate     *time.Time `json:"end_date,omitempty" `
}
