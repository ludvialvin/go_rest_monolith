package models

import (
	"go_rest_monolith/config/database"
	"go_rest_monolith/types"

	"gorm.io/gorm"
)

func FindUserAuth(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.Gorm2.Model(&types.User{}).Take(dest, conds...)
}

func FindUserByEmail(dest interface{}, email string) *gorm.DB {
	return FindUserAuth(dest, "email = ?", email)
}