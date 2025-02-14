package dto

import (
	"github.com/mnsv1511/kasikorn-line-assignment/internal/core/service/domain"
)

type GetTransactionResponse struct {
	Status      *StatusCode    `json:"status,omitempty"`
	Transaction []Transaction `json:"transactions"`
}

type Transaction struct {
	TransactionId     string `json:"transaction_id"`
	TransactionName   string `json:"transaction_name"`
	TransactionImage  string `json:"transaction_image"`
	TransactionIsBank bool   `json:"transaction_is_bank"`
}

func (d GetTransactionResponse) FromDomain(dm *domain.GetTransactionResponse) *GetTransactionResponse {
	transactions := make([]Transaction, len(dm.Transaction))
	for i, transaction := range dm.Transaction {
		transactions[i] = Transaction{
			TransactionId:     transaction.TransactionId,
			TransactionName:   transaction.TransactionName,
			TransactionImage:  transaction.TransactionImage,
			TransactionIsBank: transaction.TransactionIsBank,
		}
	}

	return &GetTransactionResponse{
		Transaction: transactions,
	}
}
