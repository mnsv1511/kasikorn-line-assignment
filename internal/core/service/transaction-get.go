package service

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mnsv1511/kasikorn-line-assignment/constant"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/core/service/domain"
)

func (s *ServiceImpl) GetTransaction(c echo.Context, userId string) (*domain.GetTransactionResponse, error) {
	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		return &domain.GetTransactionResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.INVALID_REQUEST_CODE),
				Description: string(constant.INVALID_REQUEST_DES),
			},
		}, err
	}

	transactionData, err := s.repository.GetTransaction(intUserId)
	if err != nil {
		return &domain.GetTransactionResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.TRANSACTION_NOT_FOUND_CODE),
				Description: string(constant.TRANSACTION_NOT_FOUND_DES),
			},
		}, err
	}

	transactionResp := make([]domain.Transaction, len(transactionData))
	for i, transaction := range transactionData {
		transactionResp[i] = domain.Transaction{
			TransactionId:     fmt.Sprintf("%d", transaction.ID),
			TransactionName:   transaction.Name,
			TransactionImage:  transaction.Image,
			TransactionIsBank: transaction.IsBank,
		}
	}

	return &domain.GetTransactionResponse{
		Transaction: transactionResp,
	}, nil
}
