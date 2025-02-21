package entities

import (
	"time"
)

type PurchaseHistory struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	CustomerID  string    `gorm:"type:varchar(64);not null;index"`
	Customer    Customer  `gorm:"foreignKey:CustomerID;references:ID"`
	BookID      uint64    `gorm:"not null;index"`
	Book        Book      `gorm:"foreignKey:BookID;references:ID"`
	Quantity    uint      `gorm:"not null"`
	IsPurchased bool      `gorm:"not null;default:true"`
	CreatedAt   time.Time `gorm:"not null;autoCreateTime"`
}
