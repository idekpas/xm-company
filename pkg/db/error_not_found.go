package db

import (
	"database/sql"
	"errors"
	"fmt"
)

type ErrCompanyNotFound struct{}

func (ErrCompanyNotFound) Error() string {
	return "Company not found!"
}
func (ErrCompanyNotFound) Unwrap() error {
	return fmt.Errorf("object not found")
}

func Handle(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrCompanyNotFound{}
	}
	return err
}
