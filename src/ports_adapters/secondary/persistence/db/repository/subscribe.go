package repository

import (
	"context"
	"fmt"
	"svc-subman/src/ports_adapters/secondary/persistence/db/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"svc-subman/src/domain"
)

type SubscribeRepo struct {
	db *gorm.DB
}

var _ domain.ISubscribeRepository = &SubscribeRepo{}

func NewSubscribeRepo(db *gorm.DB) *SubscribeRepo {
	return &SubscribeRepo{
		db: db,
	}
}

func (r *SubscribeRepo) Create(ctx context.Context, sub *models.Subscription) error {
	if sub.ID == uuid.Nil {
		sub.ID = uuid.New()
	}
	return r.db.WithContext(ctx).Create(sub).Error
}

func (r *SubscribeRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.Subscription, error) {
	var sub models.Subscription
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&sub).Error
	if err != nil {
		return nil, fmt.Errorf("subscription not found: %w", err)
	}
	return &sub, nil
}

func (r *SubscribeRepo) Update(ctx context.Context, sub *models.Subscription) error {
	return r.db.WithContext(ctx).Save(sub).Error
}

func (r *SubscribeRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Subscription{}).Error
}

func (r *SubscribeRepo) List(ctx context.Context, userID *string, serviceName *string, offset, limit int) ([]*models.Subscription, error) {
	var subs []*models.Subscription
	query := r.db.WithContext(ctx).Model(&models.Subscription{})

	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}

	if serviceName != nil {
		query = query.Where("service_name ILIKE ?", fmt.Sprintf("%%%s%%", *serviceName))
	}

	err := query.Offset(offset).Limit(limit).Find(&subs).Error
	return subs, err
}

func (r *SubscribeRepo) GetTotalCost(ctx context.Context, userID *string, serviceName *string, startDate, endDate time.Time) (int, error) {
	var total struct {
		Sum int
	}

	query := r.db.WithContext(ctx).Model(&models.Subscription{}).
		Select("COALESCE(SUM(price), 0) as sum").
		Where("start_date <= ? AND (end_date IS NULL OR end_date >= ?)", endDate, startDate)

	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}

	if serviceName != nil {
		query = query.Where("service_name ILIKE ?", fmt.Sprintf("%%%s%%", *serviceName))
	}

	err := query.Scan(&total).Error
	return total.Sum, err
}

func (r *SubscribeRepo) GetActiveSubscriptions(ctx context.Context, date time.Time) ([]*models.Subscription, error) {
	var subs []*models.Subscription
	err := r.db.WithContext(ctx).Where("start_date <= ? AND (end_date IS NULL OR end_date >= ?)", date, date).Find(&subs).Error
	return subs, err
}
