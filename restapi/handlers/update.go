package handlers

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/idekpas/xm-company/internal/xmcompany"
	cs "github.com/idekpas/xm-company/internal/xmcompany/service"
	e "github.com/idekpas/xm-company/pkg/errors"
	"net/http"
	"strconv"
)

type updateRequest struct {
	ID              *uuid.UUID             `json:"id"`
	Name            *string                `json:"name"`
	Description     *string                `json:"description"`
	EmployeesAmount *int                   `json:"employees-amount"`
	Registered      *bool                  `json:"registered"`
	Type            *xmcompany.CompanyType `json:"type"`
}

type updateResponse struct {
	CompanyID       int                   `json:"company-id"`
	ID              uuid.UUID             `json:"id"`
	Name            string                `json:"name"`
	Description     string                `json:"description"`
	EmployeesAmount int                   `json:"employees-amount"`
	Registered      bool                  `json:"registered"`
	Type            xmcompany.CompanyType `json:"type"`
}

func (s service) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		companyID, err := strconv.Atoi(vars["companyid"])
		if err != nil {
			s.respond(w, e.ErrArgument{
				Wrapped: errors.New("invalid company ID in url"),
			}, 0)
			return
		}

		req := updateRequest{}
		err = s.decode(r, &req)
		if err != nil {
			s.respond(w, err, 0)
			return
		}

		c, err := s.companyService.Update(r.Context(), cs.UpdateParams{
			CompanyID:       companyID,
			ID:              req.ID,
			Name:            req.Name,
			Description:     req.Description,
			EmployeesAmount: req.EmployeesAmount,
			Registered:      req.Registered,
			Type:            req.Type,
		})
		if err != nil {
			s.respond(w, err, 0)
			return
		}
		s.respond(w, updateResponse{
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
