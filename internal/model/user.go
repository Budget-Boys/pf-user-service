package model

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type User struct {
    ID                uuid.UUID      `gorm:"type:char(36);primaryKey"`
    Name              string         `gorm:"type:varchar(255);not null"`
    CPFCNPJ           string         `gorm:"type:varchar(20);uniqueIndex;not null"`
    Phone             string         `gorm:"type:varchar(20)"`
    Email             string         `gorm:"type:varchar(255);uniqueIndex;not null"`
    EmailVerifiedAt   *time.Time
    Password          string         `gorm:"type:varchar(255);not null"`
    CreatedAt         time.Time
    UpdatedAt         time.Time
    DeletedAt         gorm.DeletedAt `gorm:"index"`
}
