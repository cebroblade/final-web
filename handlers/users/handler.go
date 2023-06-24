package users

import "gorm.io/gorm"

type UserHandler struct {
	Db *gorm.DB
}
