package enum

type BookStatus string

const (
	StatusAvailable BookStatus = "available" // พร้อมให้ยืมหรือขาย
	StatusBorrowed  BookStatus = "borrowed"  // กำลังถูกยืม
	StatusArchived  BookStatus = "archived"  // เก็บถาวร
	StatusSold      BookStatus = "sold"      // ถูกขายไปแล้ว
)
