package service

import (
	"context"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/idekpas/xm-company/internal/xmcompany"
	e "github.com/idekpas/xm-company/pkg/errors"
)

type CreateParams struct {
	ID              uuid.UUID             `valid:"required"`
	Name            string                `valid:"required,stringlength(1|15)"`
	Description     string                `valid:"optional,stringlength(0|3000)"`
	EmployeesAmount int                   `valid:"required"`
	Registered      bool                  `valid:"required"`
	Type            xmcompany.CompanyType `valid:"required"`
}

func (s Service) Create(ctx context.Context, params CreateParams) (int, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return 0, e.ErrArgument{Wrapped: err}
	}

	if !params.Type.IsValid() {
		return 0, e.ErrArgument{Wrapped: e.ErrWrongCompanyType{}}
	}

	tx, err := s.repository.Db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	ce := xmcompany.Company{
		ID:              params.ID,
		Name:            params.Name,
		Description:     params.Description,
		EmployeesAmount: params.EmployeesAmount,
		Registered:      params.Registered,
		Type:            params.Type,
	}
	err = s.repository.Create(ctx, &ce)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	return ce.CompanyID, err
}
