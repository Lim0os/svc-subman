package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Subscription struct {
	gorm.Model
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ServiceName string     `gorm:"not null"`
	Price       int        `gorm:"not null;check:price > 0"`
	UserID      string     `gorm:"type:string;not null;index"`
	StartDate   time.Time  `gorm:"not null"`
	EndDate     *time.Time `gorm:"null"`
}

func (Subscription) TableName() string {
	return "subscription_service.subscriptions"
}
