package repository

import (
	"context"
	"fmt"

	"github.com/mnsv1511/kasikorn-line-assignment/ent"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/accountbalances"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/accountdetails"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/accountflags"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/accounts"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/users"
)

func (r *RepositoryImpl) GetAccountsListByUserId(userId int) ([]*ent.Accounts, error) {
	accountsData, err := r.clientDB.Accounts.Query().
		WithUsers().
		Where(accounts.HasUsersWith(users.ID(userId))).
		All(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return accountsData, nil
}

func (r *RepositoryImpl) GetAccountsByAccountId(accountId int) (*ent.Accounts, error) {
	accountsData, err := r.clientDB.Accounts.Query().
		WithUsers().
		Where(accounts.IDEQ(accountId)).
		Only(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return accountsData, nil
}

func (r *RepositoryImpl) GetAccountBalancesListByUserId(userId  int) ([]*ent.AccountBalances, error) {
	accountsData, err := r.clientDB.AccountBalances.Query().
		Where(accountbalances.UserIDEQ(userId)).
		WithUsers().
		WithAccounts().
		All(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return accountsData, nil
}

func (r *RepositoryImpl) GetAccountBalancesByAccountId(accountId int) (*ent.AccountBalances, error) {
	accountsData, err := r.clientDB.AccountBalances.Query().
		Where(accountbalances.AccountIDEQ(accountId)).
		WithUsers().
		WithAccounts().
		Only(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return accountsData, nil
}

func (r *RepositoryImpl) GetAccountDetailsListByUserId(userId int) ([]*ent.AccountDetails, error) {
	accountsData, err := r.clientDB.AccountDetails.Query().
		Where(accountdetails.UserIDEQ(userId)).
		WithUsers().
		WithAccounts().
		All(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return accountsData, nil
}

func (r *RepositoryImpl) GetAccountDetailsMainByUserId(userId int) (*ent.AccountDetails, error) {
	accountsData, err := r.clientDB.AccountDetails.Query().
		Where(accountdetails.And(
			accountdetails.UserIDEQ(userId),
			accountdetails.IsMainAccountEQ(true),
			)).
		WithUsers().
		WithAccounts().
		Only(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return accountsData, nil
}

func (r *RepositoryImpl) GetAccountFlagsListByUserId(userId int) ([]*ent.AccountFlags, error) {
	accountsData, err := r.clientDB.AccountFlags.Query().
		Where(accountflags.UserIDEQ(userId)).
		WithUsers().
		WithAccounts().
		All(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return accountsData, nil
}
