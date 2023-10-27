package errors

type ErrWrongCompanyType struct{}

func (ErrWrongCompanyType) Error() string {
	return "wrong Company Type, use following option: (Corporations | NonProfit | Cooperative | Sole Proprietorship)"
}
