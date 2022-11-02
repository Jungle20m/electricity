package storage

import (
	"context"
	"github.com/Jungle20m/electricity/modules/api/model"
)

func (s *Storage) InsertService(ctx context.Context, service model.ServiceModel) error {
	result := s.db.Create(&service)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) InsertCategory(ctx context.Context, category model.CategoryModel) error {
	result := s.db.Create(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
