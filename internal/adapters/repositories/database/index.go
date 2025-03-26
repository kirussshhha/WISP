package database

import (
	"WISP/internal/adapters/repositories/database/models"
	"WISP/internal/config"
	"WISP/internal/interfaces"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository interface {
	interfaces.UserServiceInterface
	interfaces.TeamMembersServiceInterface
	interfaces.TeamServiceInterface
}

type Database struct {
	*gorm.DB
}

func New(config *config.DBConfig) (Repository, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	err = db.AutoMigrate(&dbm.User{}, &dbm.Team{}, &dbm.TeamMember{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &Database{
		db,
	}, nil
}
