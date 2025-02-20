package entities

import "time"

type CustomerCoin struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement;"`
	CustomerID string    `gorm:"type:varchar(64);not null"`
	Amount     int64     `gorm:"not null;"`
	CreateAt   time.Time `gorm:"not null;autoCreateTime;"`
}
