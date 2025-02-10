package service

import (
	"github.com/labstack/echo/v4"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/repository"
)

type Service interface {
	TestPing(c echo.Context) (string, error)
}

type ServiceImpl struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) (Service, error) {
	return &ServiceImpl{
		repository: repo,
	}, nil
}
