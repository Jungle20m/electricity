package storage

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	orderModel "github.com/Jungle20m/electricity/internal/modules/order/model"
)

func (s *Storage) GetOrderByID(ctx context.Context, id int) (*orderModel.Order, error) {
	var record orderModel.Order
	sql := `
		SELECT id, working_site_id, customer_id, brand_name, code, status, order_at_time, create_time, update_time
		FROM orders
		WHERE id=?
		LIMIT 1
	`
	err := s.db.Raw(sql, id).First(&record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("record not found")
	}
	return &record, nil
}
