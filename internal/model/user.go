package model

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type User struct {
    ID                uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
    Name              string         `gorm:"type:varchar(255);not null" json:"name"`
    CPFCNPJ           string         `gorm:"type:varchar(20);uniqueIndex;not null" json:"cpfcnpj"`
    Phone             string         `gorm:"type:varchar(20)" json:"phone"`
    Email             string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
    EmailVerifiedAt   *time.Time	 `json:"email_verified_at"`
    Password          string         `gorm:"type:varchar(255);not null" json:"-"`
    CreatedAt         time.Time		 `json:"created_at"`
    UpdatedAt         time.Time		 `json:"updated_at"`
    DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
