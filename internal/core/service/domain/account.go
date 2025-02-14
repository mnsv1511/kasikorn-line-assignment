package domain

type GetAccountListResponse struct {
	Status  *StatusCode `json:"status"`
	Account []Account    `json:"accounts"`
}

type Account struct {
	AccountId       string      `json:"account_id"`
	AccountName     string      `json:"account_name"`
	AccountAmount   float64     `json:"account_amount"`
	AccountType     string      `json:"account_type"`
	AccountNumber   string      `json:"account_number"`
	AccountIssuer   string      `json:"account_issuer"`
	AccountColor    string      `json:"account_color"`
	AccountProgress int         `json:"account_progress"`
	AccountFlag     []AccountFlag `json:"account_flags"`
}

type AccountFlag struct {
	AccountFlagId    string     `json:"account_flag_id"`
	AccountFlagType  string     `json:"account_flag_type"`
	AccountFlagValue string     `json:"account_flag_value"`
}

type GetAccountMainResponse struct {
	Status        *StatusCode `json:"status"`
	AccountId     string     `json:"account_id"`
	AccountName   string     `json:"account_name"`
	AccountAmount float64    `json:"account_amount"`
	AccountType   string     `json:"account_type"`
	AccountNumber string     `json:"account_number"`
	AccountIssuer string     `json:"account_issuer"`
}


