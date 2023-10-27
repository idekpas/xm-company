package handlers

import (
	"github.com/google/uuid"
	"github.com/idekpas/xm-company/internal/xmcompany"
	cs "github.com/idekpas/xm-company/internal/xmcompany/service"
	"net/http"
)

type createRequest struct {
	ID              uuid.UUID             `json:"id"`
	Name            string                `json:"name"`
	Description     string                `json:"description"`
	EmployeesAmount int                   `json:"employees-amount"`
	Registered      bool                  `json:"registered"`
	Type            xmcompany.CompanyType `json:"type"`
}

type createResponse struct {
	CompanyID int `json:"company-id"`
}

func (s service) Create() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := createRequest{}

		err := s.decode(r, &req)
		if err != nil {
			s.respond(w, err, 0)
			return
		}

		companyID, err := s.companyService.Create(r.Context(), cs.CreateParams{
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
		s.respond(w, createResponse{CompanyID: companyID}, http.StatusOK)
	}
}
