package service

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mnsv1511/kasikorn-line-assignment/constant"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/core/service/domain"
)

func (s *ServiceImpl) GetAccountList(c echo.Context, userId string) (*domain.GetAccountListResponse, error) {
	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		return &domain.GetAccountListResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.INVALID_REQUEST_CODE),
				Description: string(constant.INVALID_REQUEST_DES),
			},
		}, err
	}
	accountListData, err := s.repository.GetAccountsListByUserId(intUserId)
	if err != nil {
		return &domain.GetAccountListResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.ACCOUNT_NOT_FOUND_CODE),
				Description: string(constant.ACCOUNT_NOT_FOUND_DES),
			},
		}, err
	}
	accountBalanceListData, err := s.repository.GetAccountBalancesListByUserId(intUserId)
	if err != nil {
		return &domain.GetAccountListResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.ACCOUNT_NOT_FOUND_CODE),
				Description: string(constant.ACCOUNT_NOT_FOUND_DES),
			},
		}, err
	}
	accountDetailListData, err := s.repository.GetAccountDetailsListByUserId(intUserId)
	if err != nil {
		return &domain.GetAccountListResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.ACCOUNT_NOT_FOUND_CODE),
				Description: string(constant.ACCOUNT_NOT_FOUND_DES),
			},
		}, err
	}
	accountFlagListData, err := s.repository.GetAccountFlagsListByUserId(intUserId)
	if err != nil {
		return &domain.GetAccountListResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.ACCOUNT_NOT_FOUND_CODE),
				Description: string(constant.ACCOUNT_NOT_FOUND_DES),
			},
		}, err
	}
	
	accountListMap := make(map[int]domain.Account)
	for _, account := range accountListData {
		accountListMap[account.ID] = domain.Account{
			AccountId:     fmt.Sprintf("%d", account.ID),
			AccountType:   account.Type,
			AccountNumber: account.AccountNumber,
			AccountIssuer: account.Issuer,
		}
	}


	for _, accountBalance := range accountBalanceListData {
		if account, ok := accountListMap[accountBalance.AccountID]; ok {
			account.AccountAmount = accountBalance.Amount
			accountListMap[accountBalance.AccountID] = account
		}

	}

	for _, accountDetail := range accountDetailListData {
		if account, ok := accountListMap[accountDetail.AccountID]; ok {
			account.AccountColor = accountDetail.Color
			account.AccountProgress = accountDetail.Progress
			account.AccountName = accountDetail.Name
			accountListMap[accountDetail.AccountID] = account
		}

	}

	//accountFlagListmap := make(map[int]domain.Account)
	for _, accountFlag := range accountFlagListData {
		if account, ok := accountListMap[accountFlag.AccountID]; ok {
			accountFlagDetail := domain.AccountFlag{
				AccountFlagId:    fmt.Sprintf("%d", accountFlag.ID),
				AccountFlagType:  accountFlag.FlagType,
				AccountFlagValue: accountFlag.FlagValue,
			}
			account.AccountFlag = append(account.AccountFlag, accountFlagDetail)
			accountListMap[accountFlag.AccountID] = account
		}
	}

	var accountResp []domain.Account
	for _, account := range accountListData {
		accountResp = append(accountResp, accountListMap[account.ID])
	}

	return &domain.GetAccountListResponse{
		Account: accountResp,
	}, nil
}
