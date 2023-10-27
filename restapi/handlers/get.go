package handlers

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/idekpas/xm-company/internal/xmcompany"
	e "github.com/idekpas/xm-company/pkg/errors"
	"net/http"
	"strconv"
)

type getResponse struct {
	CompanyID       int                   `json:"company-id"`
	ID              uuid.UUID             `json:"id"`
	Name            string                `json:"name"`
	Description     string                `json:"description"`
	EmployeesAmount int                   `json:"employees-amount"`
	Registered      bool                  `json:"registered"`
	Type            xmcompany.CompanyType `json:"type"`
}

func (s service) Get() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		companyID, err := strconv.Atoi(vars["companyid"])
		if err != nil {
			s.respond(w, e.ErrArgument{
				Wrapped: errors.New("invalid company ID in url"),
			}, 0)
			return
		}

		c, err := s.companyService.Get(r.Context(), companyID)
		if err != nil {
			s.respond(w, err, 0)
			return
		}
		s.respond(w, getResponse{
			CompanyID:       c.CompanyID,
			ID:              c.ID,
			Name:            c.Name,
			Description:     c.Description,
			EmployeesAmount: c.EmployeesAmount,
			Registered:      c.Registered,
			Type:            c.Type,
		}, http.StatusOK)
	}
}
