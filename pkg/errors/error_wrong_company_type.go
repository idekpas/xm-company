package errors

type ErrWrongCompanyType struct {
}

func (e ErrWrongCompanyType) Error() string {
	return "wrong company type"
}
