package domain

type GetTransactionResponse struct {
	Status      *StatusCode  `json:"status"`
	Transaction []Transaction `json:"transactions"`
}

type Transaction struct {
	TransactionId     string `json:"transaction_id"`
	TransactionName   string `json:"transaction_name"`
	TransactionImage  string `json:"transaction_image"`
	TransactionIsBank bool   `json:"transaction_is_bank"`
}
