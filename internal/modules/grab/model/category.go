package model

import "time"

type CategoryModel struct {
	ID          int32      `gorm:"-"`
	Name        string     `gorm:"column:name"`
	Description string     `gorm:"column:description"`
	Active      int32      `gorm:"column:active"`
	CreateTime  *time.Time `gorm:"-"`
	UpdateTime  *time.Time `gorm:"-"`
}

func (CategoryModel) TableName() string {
	return "grab_supply_categories"
}
