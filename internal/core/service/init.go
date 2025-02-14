package service

import (
	"github.com/labstack/echo/v4"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/core/service/domain"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/repository"
)

type Service interface {
	GetAccountList(c echo.Context, userId string) (*domain.GetAccountListResponse, error)
	GetAccountMain(c echo.Context, userId string) (*domain.GetAccountMainResponse, error)
	GetDebitCardExample(c echo.Context, userId string) (*domain.GetDebitCardExampleResponse, error)
	GetTransaction(c echo.Context, userId string) (*domain.GetTransactionResponse, error)
	GetUserGreeting(c echo.Context, userId string) (*domain.GetUserGreetingResponse, error)
	GetUserLoginDetail(c echo.Context, userId string) (*domain.GetUserLoginDetailResponse, error)
}

type ServiceImpl struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) (Service, error) {
	return &ServiceImpl{
		repository: repo,
	}, nil
}
