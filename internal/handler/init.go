package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/core/service"
)

type Handler interface {
	GetAccountList(c echo.Context) error 
	GetAccountMain(c echo.Context) error 
	GetDebitCardExample(c echo.Context) error
	GetTransaction(c echo.Context) error
	GetUserGreeting(c echo.Context) error 
	GetUserLoginDetail(c echo.Context) error 
}

type HandlerImpl struct {
	svc service.Service
}

func NewHadler(svc service.Service) (Handler, error) {
	return &HandlerImpl{
		svc: svc,
	}, nil
}
