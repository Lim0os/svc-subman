package dto

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"time"
)

// CreateSubscriptionRequest - DTO для создания подписки
type CreateSubscriptionRequest struct {
	ServiceName string    `json:"service_name" validate:"required"`
	Price       int       `json:"price" validate:"required,min=1"`
	UserID      string    `json:"user_id" validate:"required"`
	StartDate   MonthYear `json:"start_date" validate:"required"`
	EndDate     MonthYear `json:"end_date"`
}

type GetSubscriberAgree struct {
	UserID      *string   `json:"user_id"`
	ServiceName *string   `json:"service_name"`
	StartDate   MonthYear `json:"start_date" validate:"required"`
	EndDate     MonthYear `json:"end_date" validate:"required"`
}

type GetUserSubscribe struct {
	UserID      *string `json:"user_id"`
	ServiceName *string `json:"service_name"`
	Offset      int     `json:"offset" validate:"required"`
	Limit       int     `json:"limit" validate:"required"`
}

type DeleteSubscribe struct {
	Ids []uuid.UUID `json:"ids" validate:"required"`
}

// UpdateSubscriptionRequest - DTO для обновления подписки
type UpdateSubscriptionRequest struct {
	ServiceName *string   `json:"service_name,omitempty"`
	Price       *int      `json:"price,omitempty" binding:"min=1"`
	StartDate   MonthYear `json:"start_date,omitempty"` // Формат: "MM-YYYY"
	EndDate     MonthYear `json:"end_date,omitempty"`   // Формат: "MM-YYYY" или null для удаления даты
}

type MonthYear struct {
	time.Time
}

const MonthYearFormat = "01-2006"

func (my *MonthYear) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t, err := time.Parse(MonthYearFormat, s)
	if err != nil {
		return fmt.Errorf("invalid month-year format, expected MM-YYYY: %w", err)
	}

	my.Time = t
	return nil
}

func (my *MonthYear) MarshalJSON() ([]byte, error) {
	return json.Marshal(my.Format(MonthYearFormat))
}
