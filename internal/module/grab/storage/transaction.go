package storage

import "context"

func (s *Storage) WithTx(ctx context.Context, fn func(storage *Storage) error) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	err := fn(&Storage{db: tx})
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

//
//func (s *Storage) ExtractTx(ctx)
