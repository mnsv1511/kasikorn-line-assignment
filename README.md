# kasikorn-line-assignment
This project involves building an API for a Backend Developer Assignment, specifically designed to provide data for the UI used in the test.

Note: This project does not include authentication; therefore, the user ID is mocked for data retrieval in the handler file. To test the API, you can refer to the provided instructions for guidance.

## Assignment API Spec
Notion: [Assignment API Spec](https://held-bosworth-7d6.notion.site/1901e5ab5e4e808ea515d5a64dd1c865?v=1901e5ab5e4e80c895ae000c4108c1e6&pvs=4)

## Tools
- Go Language
- PostgreSQL
- Docker

## How to run API
1. Clone this project
2. Install PostgreSQL with Docker:
```bash
docker compose up -d
```
3. Run Program:
```bash
go run main.go 
```

## Example API 
### GET User detail for login
- GET `localhost:8080/user/login`
- Response : 200 OK
```json
{
  "user_name": "Clare",
  "user_image_url": "clare_img.png"
}
```

### GET User Greeting
- GET `localhost:8080/user/greeting`
- Response : 200 OK
```json
{
  "user_name": "Clare",
  "user_greeting": "Have a nice day"
}
```

### GET Main account by user id
- GET `localhost:8080/account/main`
- Response : 200 OK
```json
{
  "account_id": "1",
  "account_name": "Saving Account",
  "account_amount": 62000.00,
  "account_type": "Smart account",
  "account_number": "568-2-817-40-9",
  "account_issuer": "TestLab",
}
```

### GET List transaction by user id
- GET `localhost:8080/transaction`
- Response : 200 OK
```json
{
  "transactions": [
    {
      "transaction_id": "1",
      "transaction_name": "Emily",
      "transaction_image": "emily_img.png",
      "transaction_is_bank": true,
    },
    {
      "transaction_id": "2",
      "transaction_name": "AbcdEfghiJKlmN",
      "transaction_image": "AbcdEfghiJKlmN_img.png",
      "transaction_is_bank": true,
    },
    {
      "transaction_id": "3",
      "transaction_name": "Jone Kiersten",
      "transaction_image": "Jone_Kiersten_img.png",
      "transaction_is_bank": true,
    },
    {
      "transaction_id": "4",
      "transaction_name": "Emily",
      "transaction_image": "emily_img.png",
      "transaction_is_bank": true,
    },
    {
      "transaction_id": "5",
      "transaction_name": "Emily",
      "transaction_image": "emily_img.png",
      "transaction_is_bank": true,
    },
    {
      "transaction_id": "6",
      "transaction_name": "MarkYu Gonzales",
      "transaction_image": "MarkYu_Gonzales_img.png",
      "transaction_is_bank": true,
    },
  ]
}
```

### GET List debit card example by user id
- GET `localhost:8080/debit-card/example`
- Response : 200 OK
```json
{
  "debit_cards": [
    {
      "debit_card_id": "1",
      "debit_card_name": "For my dream",
      "debit_card_status": "In progress",
      "debit_card_issuer": "TestLab",
      "debit_card_number": "ABCD0123456789",
      "debit_card_color": "#00BFFF",
      "debit_card_border_color": "#00BFFF",
    },
    {
      "debit_card_id": "2",
      "debit_card_name": "For my dream",
      "debit_card_status": "In progress",
      "debit_card_issuer": "TestLab",
      "debit_card_number": "ABCD0123456789",
      "debit_card_color": "#FF8200",
      "debit_card_border_color": "#FF8200",
    },
    {
      "debit_card_id": "3",
      "debit_card_name": "For my dream",
      "debit_card_status": "In progress",
      "debit_card_issuer": "TestLab",
      "debit_card_number": "ABCD0123456789",
      "debit_card_color": "#FFFFFF",
      "debit_card_border_color": "#FFFFFF",
    },
    {
      "debit_card_id": "4",
      "debit_card_name": "For my dream",
      "debit_card_status": "In progress",
      "debit_card_issuer": "TestLab",
      "debit_card_number": "ABCD0123456789",
      "debit_card_color": "#87CEFA",
      "debit_card_border_color": "#87CEFA",
    },
  ]
}
```

### GET List account by user id
- GET `localhost:8080/account`
- Response : 200 OK
```json
{
  "accounts": [
    {
      "account_id": "1",
      "account_name": "Saving Account",
      "account_amount": 8837999.00,
      "account_type": "Smart account",
      "account_number": "568-2-817-40-9",
      "account_issuer": "TestLab",
      "account_color": "#46D2D2",
      "account_progress": 0,
      "account_flags": [],
    },
    {
      "account_id": "2",
      "account_name": "Credit Loan",
      "account_amount": 300.100,
      "account_type": "Credit Loan",
      "account_number": "568-2-817-40-9",
      "account_issuer": "",
      "account_color": "#FF7F50",
      "account_progress": 0,
      "account_flags": [],
    },
    {
      "account_id": "3",
      "account_name": "Travel New York",
      "account_amount": 30000.00,
      "account_type": "Goal driven savings",
      "account_number": "568-2-817-40-9",
      "account_issuer": "TestLab",
      "account_color": "#A390EE",
      "account_progress": 24,
      "account_flags": [],
    },
    {
      "account_id": "4",
      "account_name": "Need to repay",
      "account_amount": 30000.00,
      "account_type": "Credit Loan",
      "account_number": "568-2-817-40-9",
      "account_issuer": "",
      "account_color": "#6482FF",
      "account_progress": 0,
      "account_flags": [
        {
          "account_flag_id": "1",
          "account_flag_type": "main",
          "account_flag_value": "Disbursement",
        },
        {
          "account_flag_id": "2",
          "account_flag_type": "no main",
          "account_flag_value": "Overdue",
        },
      ],
    },
    
  ]
}
```