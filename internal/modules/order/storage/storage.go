package storage

import "gorm.io/gorm"

type Storage struct {
	db *gorm.DB
}

func NewMysqlStorage(db *gorm.DB) *Storage {
	return &Storage{db: db}
}
