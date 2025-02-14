package dto

import "github.com/mnsv1511/kasikorn-line-assignment/internal/core/service/domain"

type GetDebitCardExampleResponse struct {
	Status    *StatusCode  `json:"status,omitempty"`
	DebitCard []DebitCard `json:"debit_cards"`
}

type DebitCard struct {
	DebitCardId          string `json:"debit_card_id"`
	DebitCardName        string `json:"debit_card_name"`
	DebitCardStatus      string `json:"debit_card_status"`
	DebitCardIssuer      string `json:"debit_card_issuer"`
	DebitCardNumber      string `json:"debit_card_number"`
	DebitCardColor       string `json:"debit_card_color"`
	DebitCardBorderColor string `json:"debit_card_border_color"`
}

func (d GetDebitCardExampleResponse) FromDomain(dm *domain.GetDebitCardExampleResponse) *GetDebitCardExampleResponse {
	debitCards := make([]DebitCard, len(dm.DebitCard))
	for i, debitCard := range dm.DebitCard {
		debitCards[i] = DebitCard{
			DebitCardId:          debitCard.DebitCardId,
			DebitCardName:        debitCard.DebitCardName,
			DebitCardStatus:      debitCard.DebitCardStatus,
			DebitCardIssuer:      debitCard.DebitCardIssuer,
			DebitCardNumber:      debitCard.DebitCardNumber,
			DebitCardColor:       debitCard.DebitCardColor,
			DebitCardBorderColor: debitCard.DebitCardBorderColor,
		}
	}
	return &GetDebitCardExampleResponse{
		DebitCard: debitCards,
	}
}
