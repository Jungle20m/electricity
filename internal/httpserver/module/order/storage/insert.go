package storage

import (
	"context"
	model2 "github.com/Jungle20m/electricity/internal/httpserver/module/order/model"
)

func (s *Storage) InsertService(ctx context.Context, service model2.ServiceModel) error {
	result := s.db.Create(&service)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) InsertCategory(ctx context.Context, category model2.CategoryModel) error {
	result := s.db.Create(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
