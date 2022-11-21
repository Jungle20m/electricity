package storage

import (
	"context"
	grabModel "github.com/Jungle20m/electricity/internal/module/grab/model"
)

func (s *Storage) InsertService(ctx context.Context, service grabModel.ServiceCreate) error {
	result := s.db.Create(&service)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) InsertCategory(ctx context.Context, category grabModel.CategoryModel) error {
	result := s.db.Create(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
