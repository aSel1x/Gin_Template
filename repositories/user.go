package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/aSel1x/Gin_Template/models"
)

type UserRepo struct {
	*Repository[models.User]
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		Repository: NewRepository[models.User](db),
	}
}

func (r *UserRepo) RetrieveByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where(
		&models.User{UserBase: models.UserBase{Username: username}}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to retrieve user by username: %w", err)
	}
	return &user, nil
}
