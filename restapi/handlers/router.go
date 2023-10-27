package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Register(r *mux.Router, lg *logrus.Logger, db *sqlx.DB) {
	handler := newHandler(lg, db)
	// todo: health api
	//r.HandleFunc("/health", handler.Health())
	r.HandleFunc("/company", handler.Create()).Methods(http.MethodPost)
	r.HandleFunc("/company/{companyid}", handler.Get()).Methods(http.MethodGet)
	r.HandleFunc("/company/{companyid}", handler.Update()).Methods(http.MethodPut)
	r.HandleFunc("/company/{companyid}", handler.Delete()).Methods(http.MethodDelete)
}
