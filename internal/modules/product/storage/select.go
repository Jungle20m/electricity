package storage

import (
	"context"
	"errors"
	"fmt"
	productModel "github.com/Jungle20m/electricity/internal/modules/product/model"
	"gorm.io/gorm"
)

func (s *Storage) GetProductByID(ctx context.Context, id int) (*productModel.Product, error) {
	var record productModel.Product
	sql := `
		SELECT id, working_site_id, order_id, category_id, code, quantity, price, name, commission, status, point_back, point_back_status, order_at_time, create_time, update_time
		FROM product
		WHERE id=?
		LIMIT 1
	`
	err := s.db.Raw(sql, id).First(&record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("record not found")
	}
	return &record, nil
}
