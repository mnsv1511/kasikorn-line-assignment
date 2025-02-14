package repository

import (
	"log"

	"github.com/mnsv1511/kasikorn-line-assignment/ent"
)

type Repository interface {
	// Account
	GetAccountsListByUserId(userId int) ([]*ent.Accounts, error)
	GetAccountsByAccountId(accountId int) (*ent.Accounts, error)
	GetAccountBalancesListByUserId(userId int) ([]*ent.AccountBalances, error)
	GetAccountBalancesByAccountId(accountId int) (*ent.AccountBalances, error)
	GetAccountDetailsListByUserId(userId int) ([]*ent.AccountDetails, error)
	GetAccountDetailsMainByUserId(userId int) (*ent.AccountDetails, error)
	GetAccountFlagsListByUserId(userId int) ([]*ent.AccountFlags, error)

	// Debit Cards
	GetDebitCardsListByUserId(userId int) ([]*ent.DebitCards, error)
	GetDebitCardDesignListByUserId(debitCardId int) ([]*ent.DebitCardDesign, error)
	GetDebitCardDetailListByUserId(debitCardId int) ([]*ent.DebitCardDetails, error)
	GetDebitCardStatusListByUserId(debitCardId int) ([]*ent.DebitCardStatus, error)

	// Transaction
	GetTransaction(userId int) ([]*ent.Transactions, error)

	// User
	GetUser(userId int) (*ent.Users, error)
	GetUserGreeting(userId int) (*ent.UserGreetings, error)
}

type RepositoryImpl struct {
	clientDB *ent.Client
}

func NewRepository() (Repository, error) {
	clientDB, err := InitDatabase()
	if err != nil {
		return nil, err
	}

	dataExists, err := CheckDataExists(clientDB)
	if err != nil {
		log.Fatalf("failed checking data exists: %v", err)
	}

	if dataExists == false {
		log.Println("Mocking up data...")
		MockUp()
		log.Println("Mock up data successfully.")
	}

	return &RepositoryImpl{
		clientDB: clientDB,
	}, nil
}
