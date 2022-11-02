package storage

import "context"

func (s *storage) WithTransaction(ctx context.Context, fn func(*storage) error) error {
	tx := s.DB.Begin()
	defer tx.Rollback()

	err := fn(&storage{DB: tx})
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}
