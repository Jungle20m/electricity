package model

import "time"

type ServiceModel struct {
	ID         int32      `gorm:"-"`
	Title      string     `gorm:"title"`
	SubTitle   string     `gorm:"sub_title"`
	Body       string     `gorm:"body"`
	Image      string     `gorm:"image"`
	CategoryID int32      `gorm:"category_id"`
	Active     int32      `gorm:"active"`
	CreateTime *time.Time `gorm:"-"`
	UpdateTime *time.Time `gorm:"-"`
}

func (ServiceModel) TableName() string {
	return "grab_services"
}
