package service

import (
	"context"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/idekpas/xm-company/internal/xmcompany"
	e "github.com/idekpas/xm-company/pkg/errors"
)

type UpdateParams struct {
	CompanyID       int                    `valid:"required"`
	ID              *uuid.UUID             `valid:"optional"`
	Name            *string                `valid:"optional,stringlength(1|15)"`
	Description     *string                `valid:"optional,stringlength(0|3000)"`
	EmployeesAmount *int                   `valid:"optional"`
	Registered      *bool                  `valid:"optional"`
	Type            *xmcompany.CompanyType `valid:"optional"`
}

func (s Service) Update(ctx context.Context, params UpdateParams) (xmcompany.Company, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return xmcompany.Company{}, e.ErrArgument{Wrapped: err}
	}

	c, err := s.Get(ctx, params.CompanyID)
	if err != nil {
		return xmcompany.Company{}, err
	}

	err = s.updateCompany(params, &c)
	if err != nil {
		return xmcompany.Company{}, err
	}

	tx, err := s.repository.Db.BeginTxx(ctx, nil)
	if err != nil {
		return xmcompany.Company{}, err
	}
	defer tx.Rollback()

	err = s.repository.Update(ctx, c)
	if err != nil {
		return xmcompany.Company{}, err
	}

	err = tx.Commit()
	return c, err
}

func (s Service) updateCompany(params UpdateParams, c *xmcompany.Company) error {
	if params.ID != nil {
		c.ID = *params.ID
	}
	if params.Name != nil {
		c.Name = *params.Name
	}
	if params.Description != nil {
		c.Description = *params.Description
	}
	if params.EmployeesAmount != nil {
		c.EmployeesAmount = *params.EmployeesAmount
	}
	if params.Registered != nil {
		c.Registered = *params.Registered
	}
	if params.Type != nil {
		if !params.Type.IsValid() {
			return e.ErrArgument{Wrapped: e.ErrWrongCompanyType{}}
		}
		c.Type = *params.Type
	}
	return nil
}
