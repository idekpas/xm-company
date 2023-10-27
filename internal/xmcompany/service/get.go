package service

import (
	"context"
	"errors"
	"github.com/idekpas/xm-company/internal/xmcompany"
	"github.com/idekpas/xm-company/pkg/db"
	e "github.com/idekpas/xm-company/pkg/errors"
)

func (s Service) Get(ctx context.Context, id int) (xmcompany.Company, error) {
	c, err := s.repository.Find(ctx, id)
	if err == nil {
		return c, nil
	}

	if errors.As(err, &db.ErrCompanyNotFound{}) {
		return xmcompany.Company{}, e.ErrArgument{Wrapped: db.ErrCompanyNotFound{}}
	}

	return xmcompany.Company{}, err
}
