package repository

import (
	"context"
	"fmt"
	"github.com/idekpas/xm-company/internal/xmcompany"
	"github.com/idekpas/xm-company/pkg/db"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{Db: db}
}

func (r Repository) Create(ctx context.Context, ce *xmcompany.Company) error {
	q := `INSERT INTO public.company (id, name, description, employees_amount, registered, type) 
		VALUES (:id, :name, :description, :employees_amount, :registered, :type) RETURNING companyid;`
	rows, err := r.Db.NamedQueryContext(ctx, q, ce)
	if err != nil {
		return db.Handle(err)
	}

	for rows.Next() {
		err = rows.StructScan(ce)
		if err != nil {
			return db.Handle(err)
		}
	}
	return db.Handle(err)
}

func (r Repository) Update(ctx context.Context, ce xmcompany.Company) error {
	q := `UPDATE company SET 
                	id = :id, 
                	name = :name, 
    		  		description = :description, 
    		  	    employees_amount = :employees_amount, 
    		  	    registered = :registered, 
    		  	    type = :type
				WHERE companyid = :companyid;`
	_, err := r.Db.NamedExecContext(ctx, q, ce)
	return db.Handle(err)
}

func (r Repository) Delete(ctx context.Context, companyID int) error {
	q := `DELETE FROM company WHERE companyid = $1;`
	res, err := r.Db.ExecContext(ctx, q, companyID)
	if err != nil {
		return db.Handle(err)
	}

	c, err := res.RowsAffected()
	if err != nil {
		return db.Handle(err)
	}
	if c != 1 {
		return db.ErrCompanyNotFound{}
	}
	return db.Handle(err)
}

func (r Repository) Find(ctx context.Context, companyID int) (xmcompany.Company, error) {
	ce := xmcompany.Company{}
	q := fmt.Sprintf(
		"SELECT * FROM company WHERE companyid = $1;",
	)
	err := r.Db.GetContext(ctx, &ce, q, companyID)
	return ce, db.Handle(err)
}
