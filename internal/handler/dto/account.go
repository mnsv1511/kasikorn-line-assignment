package dto

import "github.com/mnsv1511/kasikorn-line-assignment/internal/core/service/domain"

type GetAccountListResponse struct {
	Status  *StatusCode `json:"status,omitempty"`
	Account []Account  `json:"accounts"`
}

type Account struct {
	AccountId       string        `json:"account_id"`
	AccountName     string        `json:"account_name"`
	AccountAmount   float64       `json:"account_amount"`
	AccountType     string        `json:"account_type"`
	AccountNumber   string        `json:"account_number"`
	AccountIssuer   string        `json:"account_issuer"`
	AccountColor    string        `json:"account_color"`
	AccountProgress int           `json:"account_progress"`
	AccountFlag     []AccountFlag `json:"account_flags"`
}

type AccountFlag struct {
	AccountFlagId    string     `json:"account_flag_id"`
	AccountFlagType  string     `json:"account_flag_type"`
	AccountFlagValue string     `json:"account_flag_value"`
}

type GetAccountMainResponse struct {
	Status        *StatusCode `json:"status,omitempty"`
	AccountId     string     `json:"account_id"`
	AccountName   string     `json:"account_name"`
	AccountAmount float64    `json:"account_amount"`
	AccountType   string     `json:"account_type"`
	AccountNumber string     `json:"account_number"`
	AccountIssuer string     `json:"account_issuer"`
}

func (d GetAccountListResponse) FromDomain(dm *domain.GetAccountListResponse) *GetAccountListResponse {
	accounts := make([]Account, len(dm.Account))
	for i, account := range dm.Account {
		accountFlag := make([]AccountFlag, len(account.AccountFlag))
		for j, flag := range account.AccountFlag {
			accountFlag[j] = AccountFlag{
				AccountFlagId:    flag.AccountFlagId,
				AccountFlagType:  flag.AccountFlagType,
				AccountFlagValue: flag.AccountFlagValue,
			}
		}
		accounts[i] = Account{
			AccountId:       account.AccountId,
			AccountName:     account.AccountName,
			AccountAmount:   account.AccountAmount,
			AccountType:     account.AccountType,
			AccountNumber:   account.AccountNumber,
			AccountIssuer:   account.AccountIssuer,
			AccountColor:    account.AccountColor,
			AccountProgress: account.AccountProgress,
			AccountFlag:     accountFlag,
		}
	}

	return &GetAccountListResponse{
		Account: accounts,
	}
}

func (d GetAccountMainResponse) FromDomain(dm *domain.GetAccountMainResponse) *GetAccountMainResponse {
	return &GetAccountMainResponse{
		AccountId:     dm.AccountId,
		AccountName:   dm.AccountName,
		AccountAmount: dm.AccountAmount,
		AccountType:   dm.AccountType,
		AccountNumber: dm.AccountNumber,
		AccountIssuer: dm.AccountIssuer,
	}
}
