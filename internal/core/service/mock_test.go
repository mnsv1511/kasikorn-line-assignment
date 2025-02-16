package service

import (
	"github.com/mnsv1511/kasikorn-line-assignment/ent"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/repository"
)

type MockRepository struct {
	MockGetAccountsListByUserId        func(userId int) ([]*ent.Accounts, error)
	MockGetAccountsByAccountId         func(accountId int) (*ent.Accounts, error)
	MockGetAccountBalancesListByUserId func(userId int) ([]*ent.AccountBalances, error)
	MockGetAccountBalancesByAccountId  func(accountId int) (*ent.AccountBalances, error)
	MockGetAccountDetailsListByUserId  func(userId int) ([]*ent.AccountDetails, error)
	MockGetAccountDetailsMainByUserId  func(userId int) (*ent.AccountDetails, error)
	MockGetAccountFlagsListByUserId    func(userId int) ([]*ent.AccountFlags, error)
	MockGetDebitCardsListByUserId      func(userId int) ([]*ent.DebitCards, error)
	MockGetDebitCardDesignListByUserId func(userId int) ([]*ent.DebitCardDesign, error)
	MockGetDebitCardDetailListByUserId func(userId int) ([]*ent.DebitCardDetails, error)
	MockGetDebitCardStatusListByUserId func(userId int) ([]*ent.DebitCardStatus, error)
	MockGetTransaction                 func(userId int) ([]*ent.Transactions, error)
	MockGetUser                        func(userId int) (*ent.Users, error)
	MockGetUserGreeting                func(userId int) (*ent.UserGreetings, error)
}

func (s *MockRepository) GetAccountsListByUserId(userId int) ([]*ent.Accounts, error) {
	return s.MockGetAccountsListByUserId(userId)
}

func (s *MockRepository) GetAccountsByAccountId(accountId int) (*ent.Accounts, error) {
	return s.MockGetAccountsByAccountId(accountId)
}

func (s *MockRepository) GetAccountBalancesListByUserId(userId int) ([]*ent.AccountBalances, error) {
	return s.MockGetAccountBalancesListByUserId(userId)
}

func (s *MockRepository) GetAccountBalancesByAccountId(accountId int) (*ent.AccountBalances, error) {
	return s.MockGetAccountBalancesByAccountId(accountId)
}

func (s *MockRepository) GetAccountDetailsListByUserId(userId int) ([]*ent.AccountDetails, error) {
	return s. MockGetAccountDetailsListByUserId(userId)
}

func (s *MockRepository) GetAccountDetailsMainByUserId(userId int) (*ent.AccountDetails, error) {
	return s.MockGetAccountDetailsMainByUserId(userId)
}

func (s *MockRepository) GetAccountFlagsListByUserId(userId int) ([]*ent.AccountFlags, error) {
	return s.MockGetAccountFlagsListByUserId(userId)
}

func (s *MockRepository) GetDebitCardsListByUserId(userId int) ([]*ent.DebitCards, error) {
	return s.MockGetDebitCardsListByUserId(userId)
}

func (s *MockRepository) GetDebitCardDesignListByUserId(debitCardId int) ([]*ent.DebitCardDesign, error) {
	return s.MockGetDebitCardDesignListByUserId(debitCardId)
}

func (s *MockRepository) GetDebitCardDetailListByUserId(debitCardId int) ([]*ent.DebitCardDetails, error) {
	return s.MockGetDebitCardDetailListByUserId(debitCardId)
}

func (s *MockRepository) GetDebitCardStatusListByUserId(debitCardId int) ([]*ent.DebitCardStatus, error) {
	return s.MockGetDebitCardStatusListByUserId(debitCardId)
}

func (s *MockRepository) GetTransaction(userId int) ([]*ent.Transactions, error) {
	return s.MockGetTransaction(userId)
}

func (s *MockRepository) GetUser(userId int) (*ent.Users, error) {
	return s.MockGetUser(userId)
}

func (s *MockRepository) GetUserGreeting(userId int) (*ent.UserGreetings, error) {
	return s.MockGetUserGreeting(userId)
}

func MockService(repo repository.Repository) Service {
	return &ServiceImpl{
		repository: repo,
	}
}
