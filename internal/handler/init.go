package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/core/service"
)

type Handler interface {
	TestPing(c echo.Context) error
}

type HandlerImpl struct {
	svc service.Service
}

func NewHadler(svc service.Service) (Handler, error) {
	return &HandlerImpl{
		svc: svc,
	}, nil
}
