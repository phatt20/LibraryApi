package entities

import (
	"time"

	enum "github.com/phatt20/LibraryApi/enum"
)

type Book struct {
	ID          uint64          `gorm:"primaryKey;autoIncrement"`
	Name        string          `gorm:"type:varchar(64);unique;not null;"`
	AdminID     *string         `gorm:"type:varchar(64);"`
	Description string          `gorm:"type:varchar(128);not null;"`
	Picture     string          `gorm:"type:varchar(256);not null;"`
	Price       uint            `gorm:"not null;"`
	Status      enum.BookStatus `gorm:"type:varchar(20);not null;default:'available'"` // ใช้ Enum แทน bool
	CreatedAt   time.Time       `gorm:"not null;autoCreateTime;"`
	UpdatedAt   time.Time       `gorm:"not null;autoUpdateTime;"`
}
