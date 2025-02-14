package repository

import (
	"context"
	"fmt"

	"github.com/mnsv1511/kasikorn-line-assignment/ent"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/debitcarddesign"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/debitcarddetails"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/debitcards"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/debitcardstatus"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/users"
)

func (r *RepositoryImpl) GetDebitCardsListByUserId(userId int) ([]*ent.DebitCards, error) {
	debitCardData, err := r.clientDB.DebitCards.Query().
		WithUsers().
		Where(debitcards.HasUsersWith(users.ID(userId))).
		All(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return debitCardData, nil
}

func (r *RepositoryImpl) GetDebitCardDesignListByUserId(userId int) ([]*ent.DebitCardDesign, error) {
	debitCardData, err := r.clientDB.DebitCardDesign.Query().
		Where(debitcarddesign.UserIDEQ(userId)).
		WithUsers().
		WithDebitCards().
		All(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return debitCardData, nil
}

func (r *RepositoryImpl) GetDebitCardDetailListByUserId(userId int) ([]*ent.DebitCardDetails, error) {
	debitCardData, err := r.clientDB.DebitCardDetails.Query().
		Where(debitcarddetails.UserIDEQ(userId)).
		WithUsers().
		WithDebitCards().
		All(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return debitCardData, nil
}

func (r *RepositoryImpl) GetDebitCardStatusListByUserId(userId int) ([]*ent.DebitCardStatus, error) {
	debitCardData, err := r.clientDB.DebitCardStatus.Query().
		Where(debitcardstatus.UserIDEQ(userId)).
		WithUsers().
		WithDebitCards().
		All(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return debitCardData, nil
}