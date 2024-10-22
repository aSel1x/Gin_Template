package adapters

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/aSel1x/Gin_Template/models"
	"github.com/aSel1x/Gin_Template/repositories"
)

type Postgres struct {
	*gorm.DB
	UserRepo *repositories.UserRepo
}

func NewPostgres(postgresDsn string) (*Postgres, error) {
	db, err := gorm.Open(postgres.Open(postgresDsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres database: %w", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate users table: %w", err)
	}

	userRepo := repositories.NewUserRepo(db)

	return &Postgres{
		DB:       db,
		UserRepo: userRepo,
	}, nil
}
