package model

import "time"

type CategoryModel struct {
	ID          int32      `gorm:"id"`
	Name        string     `gorm:"name"`
	Description string     `gorm:"description"`
	Active      int32      `gorm:"active"`
	CreateTime  *time.Time `gorm:"create_time"`
	UpdateTime  *time.Time `gorm:"update_time"`
}

func (CategoryModel) TableName() string {
	return "grab_supply_categories"
}
