package handlers

import (
	"github.com/gorilla/mux"
	r "github.com/idekpas/xm-company/internal/xmcompany/repository"
	s "github.com/idekpas/xm-company/internal/xmcompany/service"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type service struct {
	logger         *logrus.Logger
	router         *mux.Router
	companyService s.Service
}

func newHandler(lg *logrus.Logger, db *sqlx.DB) service {
	return service{
		logger:         lg,
		companyService: s.NewService(r.NewRepository(db)),
	}
}
