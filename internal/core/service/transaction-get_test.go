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

func TestGetTransaction(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		mockRepo MockRepository
		want     *domain.GetTransactionResponse
		wantErr  error
	}{
		{
			name:  "success get transaction",
			input: "1",
			mockRepo: MockRepository{
				MockGetTransaction: func(userId int) ([]*ent.Transactions, error) {
					return []*ent.Transactions{
						{
							ID:     1,
							UserID: 1,
							Name:   "Emily",
							Image:  "emily_img.png",
							IsBank: true,
						},
						{
							ID:     2,
							UserID: 1,
							Name:   "AbcdEfghiJKlmN",
							Image:  "AbcdEfghiJKlmN_img.png",
							IsBank: true,
						},
						{
							ID:     3,
							UserID: 1,
							Name:   "Jone Kiersten",
							Image:  "Jone_Kiersten_img.png",
							IsBank: true,
						},
						{
							ID:     4,
							UserID: 1,
							Name:   "Emily",
							Image:  "emily_img.png",
							IsBank: true,
						},
						{
							ID:     5,
							UserID: 1,
							Name:   "Emily",
							Image:  "emily_img.png",
							IsBank: true,
						},
						{
							ID:     6,
							UserID: 1,
							Name:   "MarkYu Gonzales",
							Image:  "MarkYu_Gonzales_img.png",
							IsBank: true,
						},
					}, nil
				},
			},
			want: &domain.GetTransactionResponse{
				Transaction: []domain.Transaction{
					{
						TransactionId:     "1",
						TransactionName:   "Emily",
						TransactionImage:  "emily_img.png",
						TransactionIsBank: true,
					},
					{
						TransactionId:     "2",
						TransactionName:   "AbcdEfghiJKlmN",
						TransactionImage:  "AbcdEfghiJKlmN_img.png",
						TransactionIsBank: true,
					},
					{
						TransactionId:     "3",
						TransactionName:   "Jone Kiersten",
						TransactionImage:  "Jone_Kiersten_img.png",
						TransactionIsBank: true,
					},
					{
						TransactionId:     "4",
						TransactionName:   "Emily",
						TransactionImage:  "emily_img.png",
						TransactionIsBank: true,
					},
					{
						TransactionId:     "5",
						TransactionName:   "Emily",
						TransactionImage:  "emily_img.png",
						TransactionIsBank: true,
					},
					{
						TransactionId:     "6",
						TransactionName:   "MarkYu Gonzales",
						TransactionImage:  "MarkYu_Gonzales_img.png",
						TransactionIsBank: true,
					},
				},
			},
		},
		{
			name:  "error req get transaction",
			input: "a",
			want: &domain.GetTransactionResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.INVALID_REQUEST_CODE),
					Description: string(constant.INVALID_REQUEST_DES),
				},
			},
			wantErr: errors.New("error req get transaction"),
		},
		{
			name:  "error repo get transaction",
			input: "1",
			mockRepo: MockRepository{
				MockGetTransaction: func(userId int) ([]*ent.Transactions, error) {
					return nil, errors.New("error repo get transaction")
				},
			},
			want: &domain.GetTransactionResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.TRANSACTION_NOT_FOUND_CODE),
					Description: string(constant.TRANSACTION_NOT_FOUND_DES),
				},
			},
			wantErr: errors.New("error repo get transaction"),
		},
	}
	for _, testCase := range testCases {
		c := echo.New().AcquireContext()
		t.Run(testCase.name, func(t *testing.T) {
			s := MockService(&testCase.mockRepo)
			got, err := s.GetTransaction(c, testCase.input)
			if err != nil {
				assert.Equal(t, testCase.want, got)
			} else {
				assert.Equal(t, testCase.wantErr, err)
			}
		})
	}
}
