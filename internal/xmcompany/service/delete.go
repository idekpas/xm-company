package service

import (
	"context"
)

func (s Service) Delete(ctx context.Context, id int) error {
	_, err := s.Get(ctx, id)
	if err != nil {
		return err
	}

	err = s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
