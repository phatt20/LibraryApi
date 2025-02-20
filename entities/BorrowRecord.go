package entities

import "time"

type BorrowRecord struct {
	ID          uint64     `gorm:"primaryKey;autoIncrement"`
	BookID      uint64     `gorm:"not null;index;"`
	Book        Book       `gorm:"foreignKey:BookID;references:ID"`
	CustomerID  string     `gorm:"type:varchar(64);not null;index;"`
	Customer    Customer   `gorm:"foreignKey:CustomerID;references:ID"`
	BorrowedAt  time.Time  `gorm:"not null;autoCreateTime;"`
	ReturnDueAt time.Time  `gorm:"not null;"`
	ReturnedAt  *time.Time `gorm:""`
	IsPurchased bool       `gorm:"not null;default:false;"` // true = ซื้อขาด, false = ยืม
}
