package products

import "gorm.io/gorm"

type ProductHandler struct {
	Db *gorm.DB
}
