package model

import (
	"gorm.io/gorm"
	"time"
)

type TimestampModel struct {
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type SoftDeleteModel struct {
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type BaseModel struct {
	TimestampModel
	ID uint64 `gorm:"primaryKey;autoIncrement"`
}
