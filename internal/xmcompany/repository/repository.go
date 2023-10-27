package repository

import (
	"context"
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
	q := `INSERT INTO company (ID, Name, Description, Employees_Amount, Registered, Type) 
		VALUES (:id, :name, :description, :employees_amount, registered, type) RETURNING companyID;`
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
                	ID = :id, 
                	Name = :name, 
    		  		Description = :description, 
    		  	    Employees_Amount = :employees_amount, 
    		  	    Registered = :registered, 
    		  	    Type = :type
				WHERE companyID = :companyID;`
	_, err := r.Db.NamedExecContext(ctx, q, ce)
	return db.Handle(err)
}

func (r Repository) Delete(ctx context.Context, companyID int) error {
	q := `DELETE FROM company WHERE companyID = (?);`
	res, err := r.Db.ExecContext(ctx, q, companyID)
	if err != nil {
		return db.Handle(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return db.Handle(err)
	}
	if count != 1 {
		return db.ErrCompanyNotFound{}
	}
	return db.Handle(err)
}
