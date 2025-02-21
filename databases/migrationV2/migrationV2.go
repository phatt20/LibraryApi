package main

import (
	"github.com/phatt20/LibraryApi/config"
	"github.com/phatt20/LibraryApi/databases"
	"github.com/phatt20/LibraryApi/entities"
	enum "github.com/phatt20/LibraryApi/enum"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(&conf.Database)
	tx := db.Connect().Begin()
	itemsAdding(tx)

	tx.Commit()
	if tx.Error != nil {
		tx.Rollback()
		panic(tx.Error)
	}
}

func itemsAdding(tx *gorm.DB) {
	books := []entities.Book{
		{
			Name:        "คัมภีร์เวทลี้ลับ",
			Description: "ตำราศักดิ์สิทธิ์ที่บรรจุคาถาโบราณและความรู้เกี่ยวกับเวทมนตร์ที่สูญหายไป",
			Price:       250,
			Picture:     "https://i.pinimg.com/564x/9a/45/b2/9a45b2d056fc5e47bf79c7e3d3c1b3a4.jpg",
			Status:      enum.StatusAvailable,
		},
		{
			Name:        "พงศาวดารแห่งอาณาจักรที่สาบสูญ",
			Description: "เรื่องราวมหากาพย์เกี่ยวกับการเกิดและล่มสลายของอาณาจักรเวทมนตร์",
			Price:       180,
			Picture:     "https://i.pinimg.com/564x/2c/3e/7d/2c3e7d9e2cd81d15b6e81623e01b13f6.jpg",
			Status:      enum.StatusAvailable,
		},
		{
			Name:        "ศิลปะแห่งการเล่นแร่แปรธาตุ",
			Description: "คู่มือการเล่นแร่แปรธาตุและความลับที่ถูกเก็บซ่อนไว้",
			Price:       220,
			Picture:     "https://i.pinimg.com/564x/6a/98/23/6a9823f49a23a9295b8f78c556634d3f.jpg",
			Status:      enum.StatusAvailable,
		},
		{
			Name:        "สารานุกรมสัตว์วิเศษ",
			Description: "หนังสือรวบรวมข้อมูลและภาพประกอบของสัตว์ในตำนานและที่อยู่อาศัยของพวกมัน",
			Price:       160,
			Picture:     "https://i.pinimg.com/564x/4d/2a/12/4d2a12a68b2c7b89c3b575b1bfe2e73b.jpg",
			Status:      enum.StatusAvailable,
		},
		{
			Name:        "บันทึกที่สาบสูญของนักเดินทางข้ามกาลเวลา",
			Description: "ไดอารี่ปริศนาที่บันทึกเรื่องราวของการเดินทางผ่านห้วงเวลาและมิติต่าง ๆ",
			Price:       300,
			Picture:     "https://i.pinimg.com/564x/77/3e/4d/773e4d3d4aa7d2cba56b2f4e933a328b.jpg",
			Status:      enum.StatusAvailable,
		},
		{
			Name:        "ยอดนักสืบจิ๋วโคนัน",
			Description: "เรื่องราวของนักสืบวัยเด็กที่ไขปริศนาอาชญากรรมสุดซับซ้อน",
			Price:       95,
			Picture:     "https://i.pinimg.com/564x/25/b6/8d/25b68dfc83e4f6165b7a9c4c13a6a9b7.jpg",
			Status:      enum.StatusAvailable,
		},
		{
			Name:        "วันพีซ",
			Description: "การผจญภัยของลูฟี่และพรรคพวกในการตามหาสมบัติที่ยิ่งใหญ่ที่สุด",
			Price:       99,
			Picture:     "https://i.pinimg.com/564x/81/a7/30/81a7301b4c4a7db964e8f6f4f54a9f3b.jpg",
			Status:      enum.StatusAvailable,
		},
		{
			Name:        "นารูโตะ",
			Description: "เรื่องราวของนินจาหนุ่มที่ต้องการเป็นโฮคาเงะและปกป้องหมู่บ้านของเขา",
			Price:       90,
			Picture:     "https://i.pinimg.com/564x/8b/7d/42/8b7d42e66e3f2cbf3d842b3bbab1de45.jpg",
			Status:      enum.StatusAvailable,
		},
		{
			Name:        "ดาบพิฆาตอสูร",
			Description: "เด็กหนุ่มที่ต้องการล้างแค้นอสูรและช่วยเหลือน้องสาวที่กลายเป็นอสูร",
			Price:       120,
			Picture:     "https://i.pinimg.com/564x/0f/33/2e/0f332eb8d2e6795e3eec5c9e9e59a52d.jpg",
			Status:      enum.StatusAvailable,
		},
		{
			Name:        "ไฮคิว!! คู่ตบฟ้าประทาน",
			Description: "เรื่องราวของทีมวอลเลย์บอลที่ต้องการขึ้นเป็นสุดยอดของประเทศ",
			Price:       110,
			Picture:     "https://i.pinimg.com/564x/35/b7/ea/35b7ea8a8cc51cba5f3b0deccca5a6b4.jpg",
			Status:      enum.StatusAvailable,
		},
	}

	tx.CreateInBatches(books, len(books))
}
