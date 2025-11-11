package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID  `gorm:"type:uuid;primaryKey"`
	UserName         string     `gorm:"uniqueIndex;not null"`
	Email            string     `gorm:"uniqueIndex;not null"`
	PasswordHash     *string    `gorm:""`
	AuthProviderType string     `gorm:"not null;default:'local'"`
	RoleID           uint       `gorm:"not null"`
	Role             Role       `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PhoneNumber      *string    `gorm:""`
	Verified         bool       `gorm:"default:false"`
	IsBlocked        bool       `gorm:"default:false"`
	CreatedAt        *time.Time `gorm:"autoCreateTime"`
	UpdatedAt        *time.Time `gorm:"autoUpdateTime"`
	DeletedAt        time.Time  `gorm:"index"`
}

type UserProfile struct {
	FirstName string  `gorm:"not null"`
	LastName  string  `gorm:"not null"`
	Address   string  ""
	DOB       *string `gorm:""`
	ImageUrl  *string
}

type UserAddress struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	AddressLine string `gorm:"not null"`
	City        string `gorm:"not null"`
	State       string
	Country     string `gorm:"not null"`
	ZipCode     string
}

type Role struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}
