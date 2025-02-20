package entities

import (
	"time"
)

type Book struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(64);unique;not null;"`
	AdminID     *string   `gorm:"type:varchar(64);"`
	Description string    `gorm:"type:varchar(128);not null;"`
	Picture     string    `gorm:"type:varchar(256);not null;"`
	Price       uint      `gorm:"not null;"`
	IsArchive   bool      `gorm:"not null;default:false;"`
	CreatedAt   time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt   time.Time `gorm:"not null;autoUpdateTime;"`
}
