package storage

import (
	"context"
	grabModel "github.com/Jungle20m/electricity/internal/modules/grab/model"
)

func (s *Storage) InsertService(ctx context.Context, service grabModel.ServiceCreate) error {
	db := s.GetConnection(ctx)
	result := db.Create(&service)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) InsertCategory(ctx context.Context, service grabModel.ServiceCreate) error {
	result := s.db.Create(&service)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
