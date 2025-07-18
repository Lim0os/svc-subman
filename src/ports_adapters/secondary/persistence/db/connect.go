package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"svc-subman/src/common/config"
	"svc-subman/src/domain"
	"svc-subman/src/ports_adapters/secondary/persistence/db/models"
	"svc-subman/src/ports_adapters/secondary/persistence/db/repository"
)

type Repositories struct {
	SubscribeRepo domain.ISubscribeRepository
}

func Migrate(db *gorm.DB, schema string) error {

	if err := db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", "subscription_service")).Error; err != nil {
		return fmt.Errorf("failed to create schema: %v", err)
	}

	if err := db.AutoMigrate(&models.Subscription{}); err != nil {
		return fmt.Errorf("failed to migrate tables: %v", err)
	}

	log.Println("Migration completed successfully!")
	return nil
}

func NewPostgresDB(cfg config.Repo) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DbHost, cfg.DbUser, cfg.DbPass, cfg.DbName, cfg.DbPort, cfg.DbSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	db.Exec(fmt.Sprintf("SET search_path TO %s", cfg.DbShema))

	return db, nil
}

func NewRepositories(cfg config.Repo) (*Repositories, error) {

	db, err := NewPostgresDB(cfg)
	if err != nil {
		return nil, err
	}

	if err := Migrate(db, cfg.DbShema); err != nil {
		return nil, err
	}

	subscribeRepo := repository.NewSubscribeRepo(db)

	return &Repositories{
		SubscribeRepo: subscribeRepo,
	}, nil
}
