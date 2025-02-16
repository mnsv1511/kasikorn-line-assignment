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

func TestGetDebitCardExample(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		mockRepo MockRepository
		want     *domain.GetDebitCardExampleResponse
		wantErr  error
	}{
		{
			name:  "success get debit card example",
			input: "1",
			mockRepo: MockRepository{
				MockGetDebitCardsListByUserId: func(userId int) ([]*ent.DebitCards, error) {
					return []*ent.DebitCards{
						{
							ID:     1,
							UserID: 1,
							Name:   "For my dream",
						},
						{
							ID:     2,
							UserID: 1,
							Name:   "For my dream",
						},
						{
							ID:     3,
							UserID: 1,
							Name:   "For my dream",
						},
						{
							ID:     4,
							UserID: 1,
							Name:   "For my dream",
						},
					}, nil
				},
				MockGetDebitCardDesignListByUserId: func(userId int) ([]*ent.DebitCardDesign, error) {
					return []*ent.DebitCardDesign{
						{
							ID:          1,
							CardID:      1,
							UserID:      1,
							Color:       "#00BFFF",
							BorderColor: "#00BFFF",
						},
						{
							ID:          2,
							CardID:      2,
							UserID:      1,
							Color:       "#FF8200",
							BorderColor: "#FF8200",
						},
						{
							ID:          3,
							CardID:      3,
							UserID:      1,
							Color:       "#FFFFFF",
							BorderColor: "#FFFFFF",
						},
						{
							ID:          4,
							CardID:      4,
							UserID:      1,
							Color:       "#87CEFA",
							BorderColor: "#87CEFA",
						},
					}, nil
				},
				MockGetDebitCardDetailListByUserId: func(userId int) ([]*ent.DebitCardDetails, error) {
					return []*ent.DebitCardDetails{
						{
							ID:     1,
							CardID: 1,
							UserID: 1,
							Issuer: "TestLab",
							Number: "ABCD0123456789",
						},
						{
							ID:     2,
							CardID: 2,
							UserID: 1,
							Issuer: "TestLab",
							Number: "ABCD0123456789",
						},
						{
							ID:     3,
							CardID: 3,
							UserID: 1,
							Issuer: "TestLab",
							Number: "ABCD0123456789",
						},
						{
							ID:     4,
							CardID: 4,
							UserID: 1,
							Issuer: "TestLab",
							Number: "ABCD0123456789",
						},
					}, nil
				},
				MockGetDebitCardStatusListByUserId: func(userId int) ([]*ent.DebitCardStatus, error) {
					return []*ent.DebitCardStatus{
						{
							ID:     1,
							CardID: 1,
							UserID: 1,
							Status: "In progress",
						},
						{
							ID:     2,
							CardID: 2,
							UserID: 1,
							Status: "In progress",
						},
						{
							ID:     3,
							CardID: 3,
							UserID: 1,
							Status: "In progress",
						},
						{
							ID:     4,
							CardID: 4,
							UserID: 1,
							Status: "In progress",
						},
					}, nil
				},
			},
			want: &domain.GetDebitCardExampleResponse{
				DebitCard: []domain.DebitCard{
					{
						DebitCardId:          "1",
						DebitCardName:        "For my dream",
						DebitCardStatus:      "In progress",
						DebitCardIssuer:      "TestLab",
						DebitCardNumber:      "ABCD0123456789",
						DebitCardColor:       "#00BFFF",
						DebitCardBorderColor: "#00BFFF",
					},
					{
						DebitCardId:          "2",
						DebitCardName:        "For my dream",
						DebitCardStatus:      "In progress",
						DebitCardIssuer:      "TestLab",
						DebitCardNumber:      "ABCD0123456789",
						DebitCardColor:       "#FF8200",
						DebitCardBorderColor: "#FF8200",
					},
					{
						DebitCardId:          "3",
						DebitCardName:        "For my dream",
						DebitCardStatus:      "In progress",
						DebitCardIssuer:      "TestLab",
						DebitCardNumber:      "ABCD0123456789",
						DebitCardColor:       "#FFFFFF",
						DebitCardBorderColor: "#FFFFFF",
					},
					{
						DebitCardId:          "4",
						DebitCardName:        "For my dream",
						DebitCardStatus:      "In progress",
						DebitCardIssuer:      "TestLab",
						DebitCardNumber:      "ABCD0123456789",
						DebitCardColor:       "#87CEFA",
						DebitCardBorderColor: "#87CEFA",
					},
				},
			},
		},
		{
			name:  "error req get debit card example",
			input: "a",
			want: &domain.GetDebitCardExampleResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.INVALID_REQUEST_CODE),
					Description: string(constant.INVALID_REQUEST_DES),
				},
			},
			wantErr: errors.New("error req get debit card example"),
		},
		{
			name:  "error repo get debit card list by user id",
			input: "1",
			mockRepo: MockRepository{
				MockGetDebitCardsListByUserId: func(userId int) ([]*ent.DebitCards, error) {
					return nil, errors.New("debit card not found")
				},
			},
			want: &domain.GetDebitCardExampleResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.DEBIT_CARD_NOT_FOUND_CODE),
					Description: string(constant.DEBIT_CARD_NOT_FOUND_DES),
				},
			},
			wantErr: errors.New("debit card not found"),
		},
		{
			name:  "error repo get debit card design list by user id",
			input: "1",
			mockRepo: MockRepository{
				MockGetDebitCardsListByUserId: func(userId int) ([]*ent.DebitCards, error) {
					return []*ent.DebitCards{
						{
							ID:     1,
							UserID: 1,
							Name:   "For my dream",
						},
						{
							ID:     2,
							UserID: 1,
							Name:   "For my dream",
						},
						{
							ID:     3,
							UserID: 1,
							Name:   "For my dream",
						},
						{
							ID:     4,
							UserID: 1,
							Name:   "For my dream",
						},
					}, nil
				},
				MockGetDebitCardDesignListByUserId: func(userId int) ([]*ent.DebitCardDesign, error) {
					return nil, errors.New("debit card not found")
				},
			},
			want: &domain.GetDebitCardExampleResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.DEBIT_CARD_NOT_FOUND_CODE),
					Description: string(constant.DEBIT_CARD_NOT_FOUND_DES),
				},
			},
			wantErr: errors.New("debit card not found"),
		},
		{
			name:  "error get debit card detail list by user id",
			input: "1",
			mockRepo: MockRepository{
				MockGetDebitCardsListByUserId: func(userId int) ([]*ent.DebitCards, error) {
					return []*ent.DebitCards{
						{
							ID:     1,
							UserID: 1,
							Name:   "For my dream",
						},
						{
							ID:     2,
							UserID: 1,
							Name:   "For my dream",
						},
						{
							ID:     3,
							UserID: 1,
							Name:   "For my dream",
						},
						{
							ID:     4,
							UserID: 1,
							Name:   "For my dream",
						},
					}, nil
				},
				MockGetDebitCardDesignListByUserId: func(userId int) ([]*ent.DebitCardDesign, error) {
					return []*ent.DebitCardDesign{
						{
							ID:          1,
							CardID:      1,
							UserID:      1,
							Color:       "#00BFFF",
							BorderColor: "#00BFFF",
						},
						{
							ID:          2,
							CardID:      2,
							UserID:      1,
							Color:       "#FF8200",
							BorderColor: "#FF8200",
						},
						{
							ID:          3,
							CardID:      3,
							UserID:      1,
							Color:       "#FFFFFF",
							BorderColor: "#FFFFFF",
						},
						{
							ID:          4,
							CardID:      4,
							UserID:      1,
							Color:       "#87CEFA",
							BorderColor: "#87CEFA",
						},
					}, nil
				},
				MockGetDebitCardDetailListByUserId: func(userId int) ([]*ent.DebitCardDetails, error) {
					return nil, errors.New("debit card not found")
				},
			},
			want: &domain.GetDebitCardExampleResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.DEBIT_CARD_NOT_FOUND_CODE),
					Description: string(constant.DEBIT_CARD_NOT_FOUND_DES),
				},
			},
			wantErr: errors.New("debit card not found"),
		},
		{
			name:  "error get debit card status list by user id",
			input: "1",
			mockRepo: MockRepository{
				MockGetDebitCardsListByUserId: func(userId int) ([]*ent.DebitCards, error) {
					return []*ent.DebitCards{
						{
							ID:     1,
							UserID: 1,
							Name:   "For my dream",
						},
						{
							ID:     2,
							UserID: 1,
							Name:   "For my dream",
						},
						{
							ID:     3,
							UserID: 1,
							Name:   "For my dream",
						},
						{
							ID:     4,
							UserID: 1,
							Name:   "For my dream",
						},
					}, nil
				},
				MockGetDebitCardDesignListByUserId: func(userId int) ([]*ent.DebitCardDesign, error) {
					return []*ent.DebitCardDesign{
						{
							ID:          1,
							CardID:      1,
							UserID:      1,
							Color:       "#00BFFF",
							BorderColor: "#00BFFF",
						},
						{
							ID:          2,
							CardID:      2,
							UserID:      1,
							Color:       "#FF8200",
							BorderColor: "#FF8200",
						},
						{
							ID:          3,
							CardID:      3,
							UserID:      1,
							Color:       "#FFFFFF",
							BorderColor: "#FFFFFF",
						},
						{
							ID:          4,
							CardID:      4,
							UserID:      1,
							Color:       "#87CEFA",
							BorderColor: "#87CEFA",
						},
					}, nil
				},
				MockGetDebitCardDetailListByUserId: func(userId int) ([]*ent.DebitCardDetails, error) {
					return []*ent.DebitCardDetails{
						{
							ID:     1,
							CardID: 1,
							UserID: 1,
							Issuer: "TestLab",
							Number: "ABCD0123456789",
						},
						{
							ID:     2,
							CardID: 2,
							UserID: 1,
							Issuer: "TestLab",
							Number: "ABCD0123456789",
						},
						{
							ID:     3,
							CardID: 3,
							UserID: 1,
							Issuer: "TestLab",
							Number: "ABCD0123456789",
						},
						{
							ID:     4,
							CardID: 4,
							UserID: 1,
							Issuer: "TestLab",
							Number: "ABCD0123456789",
						},
					}, nil
				},
				MockGetDebitCardStatusListByUserId: func(userId int) ([]*ent.DebitCardStatus, error) {
					return nil, errors.New("debit card not found")
				},
			},
			want: &domain.GetDebitCardExampleResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.DEBIT_CARD_NOT_FOUND_CODE),
					Description: string(constant.DEBIT_CARD_NOT_FOUND_DES),
				},
			},
			wantErr: errors.New("debit card not found"),
		},
	}
	for _, testCase := range testCases {
		c := echo.New().AcquireContext()
		t.Run(testCase.name, func(t *testing.T) {
			s := MockService(&testCase.mockRepo)
			got, err := s.GetDebitCardExample(c, testCase.input)
			if err != nil {
				assert.Equal(t, testCase.want, got)
			} else {
				assert.Equal(t, testCase.wantErr, err)
			}
		})
	}
}
