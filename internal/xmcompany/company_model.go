package xmcompany

import "github.com/google/uuid"

type CompanyType string

const (
	Corporations       CompanyType = "Corporations"
	NonProfit          CompanyType = "NonProfit"
	Cooperative        CompanyType = "Cooperative"
	SoleProprietorship CompanyType = "Sole Proprietorship"
)

func (ct CompanyType) IsValid() bool {
	switch ct {
	case Corporations, NonProfit, Cooperative, SoleProprietorship:
		return true
	}
	return false
}

type Company struct {
	CompanyID       int         `db:"companyid"`
	ID              uuid.UUID   `db:"id"`
	Name            string      `db:"name"`
	Description     string      `db:"description"`
	EmployeesAmount int         `db:"employees_amount"`
	Registered      bool        `db:"registered"`
	Type            CompanyType `db:"type"`
}
