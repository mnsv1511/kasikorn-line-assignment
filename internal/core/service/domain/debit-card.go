package domain

type GetDebitCardExampleResponse struct {
	Status    *StatusCode `json:"status"`
	DebitCard []DebitCard  `json:"debit_cards"`
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
