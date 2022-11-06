package goods

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	Name  string  `gorm:"name"`
	Price float64 `gorm:"price"`
	Color string  `gorm:"color"`
}
