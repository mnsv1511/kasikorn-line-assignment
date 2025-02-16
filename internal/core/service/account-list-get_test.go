package service

import (
	"errors"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mnsv1511/kasikorn-line-assignment/constant"
	"github.com/mnsv1511/kasikorn-line-assignment/ent"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/core/service/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountList(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		mockRepo MockRepository
		want     *domain.GetAccountListResponse
		wantErr  error
	}{
		{
			name:  "success get account list",
			input: "1",
			mockRepo: MockRepository{
				MockGetAccountsListByUserId: func(userId int) ([]*ent.Accounts, error) {
					return []*ent.Accounts{
						{
							ID:            1,
							UserID:        1,
							Type:          "Smart account",
							Currency:      "",
							AccountNumber: "568-2-817-40-9",
							Issuer:        "TestLab",
						},
						{
							ID:            2,
							UserID:        1,
							Type:          "Credit Loan",
							Currency:      "",
							AccountNumber: "568-2-817-40-9",
							Issuer:        "",
						},
						{
							ID:            3,
							UserID:        1,
							Type:          "Goal driven savings",
							Currency:      "",
							AccountNumber: "568-2-817-40-9",
							Issuer:        "TestLab",
						},
						{
							ID:            4,
							UserID:        1,
							Type:          "Credit Loan",
							Currency:      "",
							AccountNumber: "568-2-817-40-9",
							Issuer:        "",
						},
					}, nil
				},
				MockGetAccountBalancesListByUserId: func(userId int) ([]*ent.AccountBalances, error) {
					return []*ent.AccountBalances{
						{
							ID:        1,
							AccountID: 1,
							UserID:    1,
							Amount:    8837999.00,
						},
						{
							ID:        2,
							AccountID: 2,
							UserID:    1,
							Amount:    300.100,
						},
						{
							ID:        3,
							AccountID: 3,
							UserID:    1,
							Amount:    30000.00,
						},
						{
							ID:        4,
							AccountID: 4,
							UserID:    1,
							Amount:    30000.00,
						},
					}, nil
				},
				MockGetAccountDetailsListByUserId: func(userId int) ([]*ent.AccountDetails, error) {
					return []*ent.AccountDetails{
						{
							ID:            1,
							AccountID:     1,
							UserID:        1,
							Color:         "#46D2D2",
							IsMainAccount: true,
							Progress:      0,
							Name:          "Saving Account",
						},
						{
							ID:            2,
							AccountID:     2,
							UserID:        1,
							Color:         "#FF7F50",
							IsMainAccount: false,
							Progress:      0,
							Name:          "Credit Loan",
						},
						{
							ID:            3,
							AccountID:     3,
							UserID:        1,
							Color:         "#A390EE",
							IsMainAccount: false,
							Progress:      24,
							Name:          "Travel New York",
						},
						{
							ID:            4,
							AccountID:     4,
							UserID:        1,
							Color:         "#6482FF",
							IsMainAccount: false,
							Progress:      0,
							Name:          "Need to repay",
						},
					}, nil
				},
				MockGetAccountFlagsListByUserId: func(userId int) ([]*ent.AccountFlags, error) {
					return []*ent.AccountFlags{
						{
							ID:        1,
							AccountID: 4,
							UserID:    1,
							FlagType:  "main",
							FlagValue: "Disbursement",
						},
						{
							ID:        2,
							AccountID: 4,
							UserID:    1,
							FlagType:  "no main",
							FlagValue: "Overdue",
						},
					}, nil
				},
			},
			want: &domain.GetAccountListResponse{
				Account: []domain.Account{
					{
						AccountId:       "1",
						AccountName:     "Saving Account",
						AccountAmount:   8837999.00,
						AccountType:     "Smart account",
						AccountNumber:   "568-2-817-40-9",
						AccountIssuer:   "TestLab",
						AccountColor:    "#46D2D2",
						AccountProgress: 0,
						AccountFlag:     nil,
					},
					{
						AccountId:       "2",
						AccountName:     "Credit Loan",
						AccountAmount:   300.100,
						AccountType:     "Credit Loan",
						AccountNumber:   "568-2-817-40-9",
						AccountIssuer:   "",
						AccountColor:    "#FF7F50",
						AccountProgress: 0,
						AccountFlag:     nil,
					},
					{
						AccountId:       "3",
						AccountName:     "Travel New York",
						AccountAmount:   30000.00,
						AccountType:     "Goal driven savings",
						AccountNumber:   "568-2-817-40-9",
						AccountIssuer:   "TestLab",
						AccountColor:    "#A390EE",
						AccountProgress: 24,
						AccountFlag:     nil,
					},
					{
						AccountId:       "4",
						AccountName:     "Need to repay",
						AccountAmount:   30000.00,
						AccountType:     "Credit Loan",
						AccountNumber:   "568-2-817-40-9",
						AccountIssuer:   "",
						AccountColor:    "#6482FF",
						AccountProgress: 0,
						AccountFlag: []domain.AccountFlag{
							{
								AccountFlagId:    "1",
								AccountFlagType:  "main",
								AccountFlagValue: "Disbursement",
							},
							{
								AccountFlagId:    "2",
								AccountFlagType:  "no main",
								AccountFlagValue: "Overdue",
							},
						},
					},
				},
			},
		},
		{
			name:  "error req get account list",
			input: "a",
			want: &domain.GetAccountListResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.INVALID_REQUEST_CODE),
					Description: string(constant.INVALID_REQUEST_DES),
				},
			},
			wantErr: errors.New("error req get account list"),
		},
		{
			name:  "error repo get accounts list by user id",
			input: "2",
			mockRepo: MockRepository{
				MockGetAccountsListByUserId: func(userId int) ([]*ent.Accounts, error) {
					return nil, errors.New("account not found")
				},
			},
			want: &domain.GetAccountListResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.ACCOUNT_NOT_FOUND_CODE),
					Description: string(constant.ACCOUNT_NOT_FOUND_DES),
				},
			},
			wantErr: errors.New("account not found"),
		},
		{
			name:  "error repo get account balances list by user id",
			input: "1",
			mockRepo: MockRepository{
				MockGetAccountsListByUserId: func(userId int) ([]*ent.Accounts, error) {
					return []*ent.Accounts{
						{
							ID:            1,
							UserID:        1,
							Type:          "Smart account",
							Currency:      "",
							AccountNumber: "568-2-817-40-9",
							Issuer:        "TestLab",
						},
						{
							ID:            2,
							UserID:        1,
							Type:          "Credit Loan",
							Currency:      "",
							AccountNumber: "568-2-817-40-9",
							Issuer:        "",
						},
						{
							ID:            3,
							UserID:        1,
							Type:          "Goal driven savings",
							Currency:      "",
							AccountNumber: "568-2-817-40-9",
							Issuer:        "TestLab",
						},
						{
							ID:            4,
							UserID:        1,
							Type:          "Credit Loan",
							Currency:      "",
							AccountNumber: "568-2-817-40-9",
							Issuer:        "",
						},
					}, nil
				},
				MockGetAccountBalancesListByUserId: func(userId int) ([]*ent.AccountBalances, error) {
					return nil, errors.New("account not found")
				},
			},
			want: &domain.GetAccountListResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.ACCOUNT_NOT_FOUND_CODE),
					Description: string(constant.ACCOUNT_NOT_FOUND_DES),
				},
			},
			wantErr: errors.New("account not found"),
		},
		{
			name:  "error repo get account details list by user id",
			input: "1",
			mockRepo: MockRepository{
				MockGetAccountsListByUserId: func(userId int) ([]*ent.Accounts, error) {
					return []*ent.Accounts{
						{
							ID:            1,
							UserID:        1,
							Type:          "Smart account",
							Currency:      "",
							AccountNumber: "568-2-817-40-9",
							Issuer:        "TestLab",
						},
						{
							ID:            2,
							UserID:        1,
							Type:          "Credit Loan",
							Currency:      "",
							AccountNumber: "568-2-817-40-9",
							Issuer:        "",
						},
						{
							ID:            3,
							UserID:        1,
							Type:          "Goal driven savings",
							Currency:      "",
							AccountNumber: "568-2-817-40-9",
							Issuer:        "TestLab",
						},
						{
							ID:            4,
							UserID:        1,
							Type:          "Credit Loan",
							Currency:      "",
							AccountNumber: "568-2-817-40-9",
							Issuer:        "",
						},
					}, nil
				},
				MockGetAccountBalancesListByUserId: func(userId int) ([]*ent.AccountBalances, error) {
					return []*ent.AccountBalances{
						{
							ID:        1,
							AccountID: 1,
							UserID:    1,
							Amount:    8837999.00,
						},
						{
							ID:        2,
							AccountID: 2,
							UserID:    1,
							Amount:    300.100,
						},
						{
							ID:        3,
							AccountID: 3,
							UserID:    1,
							Amount:    30000.00,
						},
						{
							ID:        4,
							AccountID: 4,
							UserID:    1,
							Amount:    30000.00,
						},
					}, nil
				},
				MockGetAccountDetailsListByUserId: func(userId int) ([]*ent.AccountDetails, error) {
					return nil, errors.New("account not found")
				},
			},
			want: &domain.GetAccountListResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.ACCOUNT_NOT_FOUND_CODE),
					Description: string(constant.ACCOUNT_NOT_FOUND_DES),
				},
			},
			wantErr: errors.New("account not found"),
		},
		{
			name:  "success get account list",
			input: "1",
			mockRepo: MockRepository{
				MockGetAccountsListByUserId: func(userId int) ([]*ent.Accounts, error) {
					return []*ent.Accounts{
						{
							ID:            1,
							UserID:        1,
							Type:          "Smart account",
							Currency:      "",
							AccountNumber: "568-2-817-40-9",
							Issuer:        "TestLab",
						},
						{
							ID:            2,
							UserID:        1,
							Type:          "Credit Loan",
							Currency:      "",
							AccountNumber: "568-2-817-40-9",
							Issuer:        "",
						},
						{
							ID:            3,
							UserID:        1,
							Type:          "Goal driven savings",
							Currency:      "",
							AccountNumber: "568-2-817-40-9",
							Issuer:        "TestLab",
						},
						{
							ID:            4,
							UserID:        1,
							Type:          "Credit Loan",
							Currency:      "",
							AccountNumber: "568-2-817-40-9",
							Issuer:        "",
						},
					}, nil
				},
				MockGetAccountBalancesListByUserId: func(userId int) ([]*ent.AccountBalances, error) {
					return []*ent.AccountBalances{
						{
							ID:        1,
							AccountID: 1,
							UserID:    1,
							Amount:    8837999.00,
						},
						{
							ID:        2,
							AccountID: 2,
							UserID:    1,
							Amount:    300.100,
						},
						{
							ID:        3,
							AccountID: 3,
							UserID:    1,
							Amount:    30000.00,
						},
						{
							ID:        4,
							AccountID: 4,
							UserID:    1,
							Amount:    30000.00,
						},
					}, nil
				},
				MockGetAccountDetailsListByUserId: func(userId int) ([]*ent.AccountDetails, error) {
					return []*ent.AccountDetails{
						{
							ID:            1,
							AccountID:     1,
							UserID:        1,
							Color:         "#46D2D2",
							IsMainAccount: true,
							Progress:      0,
							Name:          "Saving Account",
						},
						{
							ID:            2,
							AccountID:     2,
							UserID:        1,
							Color:         "#FF7F50",
							IsMainAccount: false,
							Progress:      0,
							Name:          "Credit Loan",
						},
						{
							ID:            3,
							AccountID:     3,
							UserID:        1,
							Color:         "#A390EE",
							IsMainAccount: false,
							Progress:      24,
							Name:          "Travel New York",
						},
						{
							ID:            4,
							AccountID:     4,
							UserID:        1,
							Color:         "#6482FF",
							IsMainAccount: false,
							Progress:      0,
							Name:          "Need to repay",
						},
					}, nil
				},
				MockGetAccountFlagsListByUserId: func(userId int) ([]*ent.AccountFlags, error) {
					return nil, errors.New("account not found")
				},
			},
			want: &domain.GetAccountListResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.ACCOUNT_NOT_FOUND_CODE),
					Description: string(constant.ACCOUNT_NOT_FOUND_DES),
				},
			},
			wantErr: errors.New("account not found"),
		},
	}
	for _, testCase := range testCases {
		c := echo.New().AcquireContext()
		t.Run(testCase.name, func(t *testing.T) {
			s := MockService(&testCase.mockRepo)
			got, err := s.GetAccountList(c, testCase.input)
			if err != nil {
				assert.Equal(t, testCase.want, got)
			} else {
				assert.Equal(t, testCase.wantErr, err)
			}
		})
	}
}
