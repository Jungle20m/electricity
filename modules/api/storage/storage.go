package storage

import "gorm.io/gorm"

type storage struct {
	DB *gorm.DB
}

func NewStorage(db *gorm.DB) *storage {
	return &storage{}
}
