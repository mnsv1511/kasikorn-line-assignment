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

func TestGetAccountMain(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		mockRepo MockRepository
		want     *domain.GetAccountMainResponse
		wantErr  error
	}{
		{
			name:  "success get account main",
			input: "1",
			mockRepo: MockRepository{
				MockGetAccountDetailsMainByUserId: func(userId int) (*ent.AccountDetails, error) {
					return &ent.AccountDetails{
						ID:            1,
						AccountID:     1,
						UserID:        1,
						Color:         "#46D2D2",
						IsMainAccount: true,
						Progress:      0,
						Name:          "Saving Account",
					}, nil
				},
				MockGetAccountsByAccountId: func(accountId int) (*ent.Accounts, error) {
					return &ent.Accounts{
						ID:            1,
						UserID:        1,
						Type:          "Smart account",
						Currency:      "",
						AccountNumber: "568-2-817-40-9",
						Issuer:        "TestLab",
					}, nil
				},
				MockGetAccountBalancesByAccountId: func(accountId int) (*ent.AccountBalances, error) {
					return &ent.AccountBalances{
						ID:        1,
						AccountID: 1,
						UserID:    1,
						Amount:    8837999.00,
					}, nil
				},
			},
			want: &domain.GetAccountMainResponse{
				AccountId:     "1",
				AccountName:   "Saving Account",
				AccountAmount: 8837999.00,
				AccountType:   "Smart account",
				AccountNumber: "568-2-817-40-9",
				AccountIssuer: "TestLab",
			},
		},
		{
			name:  "error req get account main",
			input: "a",
			want: &domain.GetAccountMainResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.INVALID_REQUEST_CODE),
					Description: string(constant.INVALID_REQUEST_DES),
				},
			},
			wantErr: errors.New("error req get account list"),
		},
		{
			name:  "error get account details main by user id",
			input: "1",
			mockRepo: MockRepository{
				MockGetAccountDetailsMainByUserId: func(userId int) (*ent.AccountDetails, error) {
					return nil, errors.New("account not found")
				},
			},
			want: &domain.GetAccountMainResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.ACCOUNT_NOT_FOUND_CODE),
					Description: string(constant.ACCOUNT_NOT_FOUND_DES),
				},
			},
			wantErr: errors.New("account not found"),
		},
		{
			name:  "error repo get account by account id",
			input: "1",
			mockRepo: MockRepository{
				MockGetAccountDetailsMainByUserId: func(userId int) (*ent.AccountDetails, error) {
					return &ent.AccountDetails{
						ID:            1,
						AccountID:     1,
						UserID:        1,
						Color:         "#46D2D2",
						IsMainAccount: true,
						Progress:      0,
						Name:          "Saving Account",
					}, nil
				},
				MockGetAccountsByAccountId: func(accountId int) (*ent.Accounts, error) {
					return nil, errors.New("account not found")
				},
			},
			want: &domain.GetAccountMainResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.ACCOUNT_NOT_FOUND_CODE),
					Description: string(constant.ACCOUNT_NOT_FOUND_DES),
				},
			},
			wantErr: errors.New("account not found"),
		},
		{
			name:  "success get account main",
			input: "1",
			mockRepo: MockRepository{
				MockGetAccountDetailsMainByUserId: func(userId int) (*ent.AccountDetails, error) {
					return &ent.AccountDetails{
						ID:            1,
						AccountID:     1,
						UserID:        1,
						Color:         "#46D2D2",
						IsMainAccount: true,
						Progress:      0,
						Name:          "Saving Account",
					}, nil
				},
				MockGetAccountsByAccountId: func(accountId int) (*ent.Accounts, error) {
					return &ent.Accounts{
						ID:            1,
						UserID:        1,
						Type:          "Smart account",
						Currency:      "",
						AccountNumber: "568-2-817-40-9",
						Issuer:        "TestLab",
					}, nil
				},
				MockGetAccountBalancesByAccountId: func(accountId int) (*ent.AccountBalances, error) {
					return nil, errors.New("account not found")
				},
			},
			want: &domain.GetAccountMainResponse{
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
			got, err := s.GetAccountMain(c, testCase.input)
			if err != nil {
				assert.Equal(t, testCase.want, got)
			} else {
				assert.Equal(t, testCase.wantErr, err)
			}
		})
	}
}
