package service

import "github.com/idekpas/xm-company/internal/xmcompany/repository"

type Service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repository: r,
	}
}
