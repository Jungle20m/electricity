package model

import "time"

type Order struct {
	ID            int        `json:"id" gorm:"column:id"`
	WorkingSiteID int        `json:"working_site_id" gorm:"column:working_site_id"`
	CustomerID    int        `json:"customer_id" gorm:"column:customer_id"`
	BrandName     string     `json:"brand_name" gorm:"column:brand_name"`
	Code          string     `json:"code" gorm:"column:code"`
	Status        int        `json:"status" gorm:"column:status"`
	OrderAtTime   *time.Time `json:"order_at_time" gorm:"column:order_at_time"`
	CreateTime    *time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime    *time.Time `json:"update_time" gorm:"column:update_time"`
}

func (Order) TableName() string {
	return "orders"
}
