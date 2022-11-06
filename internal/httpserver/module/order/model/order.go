package model

type OrderModel struct {
	ID                    int    `gorm:"column:id"`
	OrderCode             string `gorm:"column:order_code"`
	GrabServiceID         int    `gorm:"column:grab_service_id"`
	CustomerWorkingSiteID int    `gorm:"column:customer_working_site_id"`
}
