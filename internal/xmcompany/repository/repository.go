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
	q := `INSERT INTO company (ID, Name, Description, EmployeesAmount, Registered, Type) 
		VALUES (:id, :name, :desc, :empAmount, registered, type) RETURNING companyID;`
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
