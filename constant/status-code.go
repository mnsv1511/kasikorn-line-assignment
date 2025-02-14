package constant

type StatusCode string

const (
	INVALID_REQUEST_CODE         StatusCode = "1001"
	USER_NOT_FOUND_CODE          StatusCode = "2001"
	USER_GREETING_NOT_FOUND_CODE StatusCode = "2002"
	TRANSACTION_NOT_FOUND_CODE   StatusCode = "2003"
	ACCOUNT_NOT_FOUND_CODE       StatusCode = "2004"
	DEBIT_CARD_NOT_FOUND_CODE    StatusCode = "2005"
)

type StatusDescription string

const (
	INVALID_REQUEST_DES         StatusDescription = "Invalid request"
	USER_NOT_FOUND_DES          StatusDescription = "User not found"
	USER_GREETING_NOT_FOUND_DES StatusDescription = "User greeting not found"
	TRANSACTION_NOT_FOUND_DES   StatusDescription = "Transaction not found"
	ACCOUNT_NOT_FOUND_DES       StatusDescription = "Account not found"
	DEBIT_CARD_NOT_FOUND_DES    StatusDescription = "Debit card not found"
)
