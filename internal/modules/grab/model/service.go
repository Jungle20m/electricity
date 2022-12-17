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

type ServiceCreate struct {
	Title      string `json:"title" gorm:"title"`
	SubTitle   string `json:"sub_title" gorm:"sub_title"`
	Body       string `json:"body" gorm:"body"`
	Image      string `json:"image" gorm:"image"`
	CategoryID int32  `json:"category_id" gorm:"category_id"`
	Active     int32  `json:"-" gorm:"active"`
}

func (ServiceCreate) TableName() string {
	return ServiceModel{}.TableName()
}
