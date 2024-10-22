package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *Repository[T] {
	return &Repository[T]{db: db}
}

func (r *Repository[T]) Create(model *T) error {
	if err := r.db.Create(model).Error; err != nil {
		return fmt.Errorf("failed to create model: %w", err)
	}
	return nil
}

func (r *Repository[T]) RetrieveOne(id int) (*T, error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to retrieve model: %w", err)
	}
	return &entity, nil
}

func (r *Repository[T]) RetrieveMany(limit int, orderBy string) ([]T, error) {
	var entities []T
	if err := r.db.Order(orderBy).Limit(limit).Find(&entities).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve models: %w", err)
	}
	return entities, nil
}

func (r *Repository[T]) Delete(model *T) error {
	if err := r.db.Delete(model).Error; err != nil {
		return fmt.Errorf("failed to delete model: %w", err)
	}
	return nil
}
