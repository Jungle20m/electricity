package storage

import "context"

func (s *Storage) WithTransaction(ctx context.Context, fn func(*Storage) error) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	err := fn(&Storage{db: tx})
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}
