package mysql

import "gorm.io/gorm"

type MysqlOption func(*gorm.DB)

func WithMaxIdleConns() MysqlOption {
	return func(db *gorm.DB) {

	}
}
