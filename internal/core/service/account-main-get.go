package service

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mnsv1511/kasikorn-line-assignment/constant"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/core/service/domain"
)

func (s *ServiceImpl) GetAccountMain(c echo.Context, userId string) (*domain.GetAccountMainResponse, error) {
	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		return &domain.GetAccountMainResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.INVALID_REQUEST_CODE),
				Description: string(constant.INVALID_REQUEST_DES),
			},
		}, err
	}
	accountDetailMain, err := s.repository.GetAccountDetailsMainByUserId(intUserId)
	if err != nil {
		return &domain.GetAccountMainResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.ACCOUNT_NOT_FOUND_CODE),
				Description: string(constant.ACCOUNT_NOT_FOUND_DES),
			},
		}, err
	}
	accountMain, err := s.repository.GetAccountsByAccountId(accountDetailMain.ID)
	if err != nil {
		return &domain.GetAccountMainResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.ACCOUNT_NOT_FOUND_CODE),
				Description: string(constant.ACCOUNT_NOT_FOUND_DES),
			},
		}, err
	}
	accountBalancesMain, err := s.repository.GetAccountBalancesByAccountId(accountDetailMain.ID)
	if err != nil {
		return &domain.GetAccountMainResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.ACCOUNT_NOT_FOUND_CODE),
				Description: string(constant.ACCOUNT_NOT_FOUND_DES),
			},
		}, err
	}

	return &domain.GetAccountMainResponse{
		AccountId:     fmt.Sprintf("%d", accountDetailMain.ID),
		AccountName:   accountDetailMain.Name,
		AccountAmount: accountBalancesMain.Amount,
		AccountType:   accountMain.Type,
		AccountNumber: accountMain.AccountNumber,
		AccountIssuer: accountMain.Issuer,
	}, nil
}
