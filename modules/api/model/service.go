package model

type ServiceModel struct {
	//ID         int32      `gorm:"id"`
	Title      string `gorm:"title"`
	SubTitle   string `gorm:"sub_title"`
	Body       string `gorm:"body"`
	Image      string `gorm:"image"`
	CategoryID int32  `gorm:"category_id"`
	Active     int32  `gorm:"active"`
	//CreateTime *time.Time `gorm:"create_time"`
	//UpdateTime *time.Time `gorm:"update_time"`
}

func (ServiceModel) TableName() string {
	return "grab_services"
}
