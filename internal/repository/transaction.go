package repository

import (
	"context"
	"fmt"

	"github.com/mnsv1511/kasikorn-line-assignment/ent"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/transactions"
)

func (r *RepositoryImpl) GetTransaction(userId int) ([]*ent.Transactions, error) {
	transactionData, err := r.clientDB.Transactions.Query().
		Where(transactions.UserIDEQ(userId)).
		WithUsers().
		All(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return transactionData, nil
}
