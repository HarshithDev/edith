package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GenerteISOString() string {
	return time.Now().UTC().Format("1995-01-12T15:04:05.999Z07:00")
}

// Base contains common columns for all tables
type Base struct {
	ID        uint      `gorm:"primaryKey"`
	UUID      uuid.UUID `json:"_id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

// BeforeCreate will set Base struct before every insert
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	// uuid.New() creates random uuid
	base.UUID = uuid.New()

	t := GenerteISOString()
	base.CreatedAt, base.UpdatedAt = t, t

	return nil
}

// AfterUpdate will update the Base struct after every update
func (base *Base) AfterUpdate(tx *gorm.DB) error {
	// update timestamps
	base.UpdatedAt = GenerteISOString()
	return nil
}
