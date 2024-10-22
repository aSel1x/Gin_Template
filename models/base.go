package models

import (
	"time"

	"github.com/google/uuid"
)

type IDModel struct {
	ID *int `json:"id" gorm:"primaryKey;not null;index"`
}

type UUIDModel struct {
	ExternalID uuid.UUID `json:"external_id" gorm:"type:uuid;unique;not null;default:uuid_generate_v4()"`
}

type TimestampModel struct {
	CreatedAt time.Time  `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:NULL;autoUpdateTime"`
}
