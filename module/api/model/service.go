package model

import "time"

type ServiceModel struct {
	ID         int32      `gorm:"-"`
	Title      string     `gorm:"column:title"`
	SubTitle   string     `gorm:"column:sub_title"`
	Body       string     `gorm:"column:body"`
	Image      string     `gorm:"column:image"`
	CategoryID int32      `gorm:"column:category_id"`
	Active     int32      `gorm:"column:active"`
	CreateTime *time.Time `gorm:"-"`
	UpdateTime *time.Time `gorm:"-"`
}

func (ServiceModel) TableName() string {
	return "grab_services"
}
