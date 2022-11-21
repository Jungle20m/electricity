package model

import "time"

type OrderModel struct {
	ID                       int        `gorm:"column:id"`
	OrderCode                string     `gorm:"column:order_code"`
	GrabServiceID            int        `gorm:"column:grab_service_id"`
	CustomerWorkingSiteID    int        `gorm:"column:customer_working_site_id"`
	CustomerID               int        `gorm:"column:customer_id"`
	CustomerName             string     `gorm:"column:customer_name"`
	CustomerPhone            string     `gorm:"column:customer_phone"`
	CustomerCode             string     `gorm:"column:customer_code"`
	Latitude                 float64    `gorm:"column:latitude"`
	Longitude                float64    `gorm:"column:longitude"`
	Address                  string     `gorm:"column:address"`
	ExpectedSupportTime      *time.Time `gorm:"column:expected_support_time"`
	CustomerNode             string     `gorm:"column:customer_note"`
	CustomerCancelReason     string     `gorm:"column:customer_cancel_reason"`
	RatingNote               string     `gorm:"column:rating_note"`
	RatingStart              string     `gorm:"column:rating_star"`
	ElectricianWorkerSiteID  int        `gorm:"column:electrician_worker_site_id"`
	ElectricianWorkingSiteId int        `gorm:"column:electrician_working_site_id"`
	ElectricianNote          string     `gorm:"column:electrician_note"`
	PaymentRequestNote       string     `gorm:"column:payment_request_note"`
	ElectricianCancelReason  string     `gorm:"column:electrician_cancel_reason"`
	Status                   string     `gorm:"column:status"`
	CreateTime               *time.Time `gorm:"column:create_time"`
	UpdateTime               *time.Time `gorm:"column:update_time"`
}
