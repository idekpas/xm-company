package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s service) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		companyID, err := strconv.Atoi(vars["companyID"])
		if err != nil {
			s.respond(w, err, 0)
			return
		}

		err = s.companyService.Delete(r.Context(), companyID)
		if err != nil {
			s.respond(w, err, 0)
			return
		}
		s.respond(w, nil, http.StatusOK)
	}
}
