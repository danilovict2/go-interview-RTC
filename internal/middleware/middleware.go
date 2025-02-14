package middleware

import "gorm.io/gorm"

type Middleware struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) *Middleware {
	return &Middleware{
		DB: DB,
	}
}
