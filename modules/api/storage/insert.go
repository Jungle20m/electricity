package storage

import (
	"context"
	"github.com/Jungle20m/electricity/modules/api/model"
)

func (s *storage) InsertService(ctx context.Context, service model.ServiceModel) error {
	result := s.DB.Create(&service)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *storage) InsertCategory(ctx context.Context, category model.CategoryModel) error {
	result := s.DB.Create(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
