package databases

import "gorm.io/gorm"

type Databases interface {
	Connect() *gorm.DB
}
