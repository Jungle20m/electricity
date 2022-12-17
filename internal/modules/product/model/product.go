package model

import "time"

type Product struct {
	ID              int        `json:"id" gorm:"column:id"`
	WorkingSiteID   int        `json:"working_site_id" gorm:"column:working_site_id"`
	OrderID         int        `json:"order_id" gorm:"column:order_id"`
	CategoryID      int        `json:"category_id" gorm:"column:category_id"`
	Code            string     `json:"code" gorm:"column:code"`
	Quantity        int        `json:"quantity" gorm:"column:quantity"`
	Price           float64    `json:"price" gorm:"column:price"`
	Name            string     `json:"name" gorm:"column:name"`
	Commission      float64    `json:"commission" gorm:"column:commission"`
	Status          string     `json:"status" gorm:"column:status"`
	PointBack       int        `json:"point_back" gorm:"column:point_back"`
	PointBackStatus int        `json:"point_back_status" gorm:"column:point_back_status"`
	CreateTime      *time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime      *time.Time `json:"update_time" gorm:"column:update_time"`
}

func (Product) TableName() string {
	return "products"
}
