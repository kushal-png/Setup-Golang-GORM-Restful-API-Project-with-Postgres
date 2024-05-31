package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	Id uuid.UUID `gorm:"type:uuid;primary_kay"`
	Name string  `gorm:"type:varchar(255);not null"`
	Email string `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
