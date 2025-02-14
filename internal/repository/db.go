package repository

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mnsv1511/kasikorn-line-assignment/ent"
)

func InitDatabase() (*ent.Client, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	clientDB, err := ent.Open("postgres", os.Getenv("PSQL_DB_CONNECT"))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
		return nil, err
	}

	if err := clientDB.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return nil, err
	}

	return clientDB, nil
}

func CheckDataExists(client *ent.Client) (bool, error) {

	userCount, err := client.Users.Query().Count(context.Background())
	if err != nil {
		return false, err
	}
	if userCount > 0 {
		return true, nil
	}
	return false, nil
}

func MockUp() {
	ctx := context.Background()

	client, err := InitDatabase()
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// user greeting
	userGreetingTester := client.UserGreetings.Create().
		SetGreeting("Have a nice day").
		SaveX(ctx)

	// transaction
	transactionTester1 := client.Transactions.Create().
		SetName("Emily").
		SetImage("emily_img.png").
		SetIsBank(true).
		SaveX(ctx)

	transactionTester2 := client.Transactions.Create().
		SetName("AbcdEfghiJKlmN").
		SetImage("AbcdEfghiJKlmN_img.png").
		SetIsBank(true).
		SaveX(ctx)

	transactionTester3 := client.Transactions.Create().
		SetName("Jone Kiersten").
		SetImage("Jone_Kiersten_img.png").
		SetIsBank(true).
		SaveX(ctx)

	transactionTester4 := client.Transactions.Create().
		SetName("Emily").
		SetImage("emily_img.png").
		SetIsBank(true).
		SaveX(ctx)

	transactionTester5 := client.Transactions.Create().
		SetName("Emily").
		SetImage("emily_img.png").
		SetIsBank(true).
		SaveX(ctx)

	transactionTester6 := client.Transactions.Create().
		SetName("MarkYu Gonzales").
		SetImage("MarkYu_Gonzales_img.png").
		SetIsBank(true).
		SaveX(ctx)

	// debit card design
	debitCardDesignTester1 := client.DebitCardDesign.Create().
		SetColor("#00BFFF").
		SetBorderColor("#00BFFF").
		SaveX(ctx)

	debitCardDesignTester2 := client.DebitCardDesign.Create().
		SetColor("#FF8200").
		SetBorderColor("#FF8200").
		SaveX(ctx)

	debitCardDesignTester3 := client.DebitCardDesign.Create().
		SetColor("#FFFFFF").
		SetBorderColor("#FFFFFF").
		SaveX(ctx)

	debitCardDesignTester4 := client.DebitCardDesign.Create().
		SetColor("#87CEFA").
		SetBorderColor("#87CEFA").
		SaveX(ctx)

	// debit card detail
	debitCardDetailTester1 := client.DebitCardDetails.Create().
		SetIssuer("TestLab").
		SetNumber("ABCD0123456789").
		SaveX(ctx)

	debitCardDetailTester2 := client.DebitCardDetails.Create().
		SetIssuer("TestLab").
		SetNumber("ABCD0123456789").
		SaveX(ctx)

	debitCardDetailTester3 := client.DebitCardDetails.Create().
		SetIssuer("TestLab").
		SetNumber("ABCD0123456789").
		SaveX(ctx)

	debitCardDetailTester4 := client.DebitCardDetails.Create().
		SetIssuer("TestLab").
		SetNumber("ABCD0123456789").
		SaveX(ctx)

	// debit card status
	debitCardStatusTester1 := client.DebitCardStatus.Create().
		SetStatus("In progress").
		SaveX(ctx)

	debitCardStatusTester2 := client.DebitCardStatus.Create().
		SetStatus("In progress").
		SaveX(ctx)

	debitCardStatusTester3 := client.DebitCardStatus.Create().
		SetStatus("In progress").
		SaveX(ctx)

	debitCardStatusTester4 := client.DebitCardStatus.Create().
		SetStatus("In progress").
		SaveX(ctx)

	// debit card
	debitCardTester1 := client.DebitCards.Create().
		SetName("For my dream").
		AddDebitCardDesignIDs(debitCardDesignTester1.ID).
		AddDebitCardDetailIDs(debitCardDesignTester1.ID).
		AddDebitCardStatuIDs(debitCardStatusTester1.ID).
		SaveX(ctx)

	debitCardTester2 := client.DebitCards.Create().
		SetName("For my dream").
		AddDebitCardDesignIDs(debitCardDesignTester2.ID).
		AddDebitCardDetailIDs(debitCardDesignTester2.ID).
		AddDebitCardStatuIDs(debitCardStatusTester2.ID).
		SaveX(ctx)

	debitCardTester3 := client.DebitCards.Create().
		SetName("For my dream").
		AddDebitCardDesignIDs(debitCardDesignTester3.ID).
		AddDebitCardDetailIDs(debitCardDesignTester3.ID).
		AddDebitCardStatuIDs(debitCardStatusTester3.ID).
		SaveX(ctx)

	debitCardTester4 := client.DebitCards.Create().
		SetName("For my dream").
		AddDebitCardDesignIDs(debitCardDesignTester4.ID).
		AddDebitCardDetailIDs(debitCardDesignTester4.ID).
		AddDebitCardStatuIDs(debitCardStatusTester4.ID).
		SaveX(ctx)

	// account balances
	accountBalanceTester1 := client.AccountBalances.Create().
		SetAmount(8837999.00).
		SaveX(ctx)

	accountBalanceTester2 := client.AccountBalances.Create().
		SetAmount(300.100).
		SaveX(ctx)

	accountBalanceTester3 := client.AccountBalances.Create().
		SetAmount(30000.00).
		SaveX(ctx)

	accountBalanceTester4 := client.AccountBalances.Create().
		SetAmount(30000.00).
		SaveX(ctx)

	// account details
	accountDetailTester1 := client.AccountDetails.Create().
		SetColor("#46D2D2").
		SetIsMainAccount(true).
		SetProgress(0).
		SetName("Saving Account").
		SaveX(ctx)

	accountDetailTester2 := client.AccountDetails.Create().
		SetColor("#FF7F50").
		SetIsMainAccount(false).
		SetProgress(0).
		SetName("Credit Loan").
		SaveX(ctx)

	accountDetailTester3 := client.AccountDetails.Create().
		SetColor("#A390EE").
		SetIsMainAccount(false).
		SetProgress(24).
		SetName("Travel New York").
		SaveX(ctx)

	accountDetailTester4 := client.AccountDetails.Create().
		SetColor("#6482FF").
		SetIsMainAccount(false).
		SetProgress(0).
		SetName("Need to repay").
		SaveX(ctx)

	// account flags
	accountFlagTester1 := client.AccountFlags.Create().
		SetFlagType("main").
		SetFlagValue("Disbursement").
		SaveX(ctx)

	accountFlagTester2 := client.AccountFlags.Create().
		SetFlagType("no main").
		SetFlagValue("Overdue").
		SaveX(ctx)

	// accounts
	accountTester1 := client.Accounts.Create().
		SetType("Smart account").
		SetCurrency("").
		SetAccountNumber("568-2-817-40-9").
		SetIssuer("TestLab").
		AddAccountBalanceIDs(accountBalanceTester1.ID).
		AddAccountDetailIDs(accountDetailTester1.ID).
		SaveX(ctx)

	accountTester2 := client.Accounts.Create().
		SetType("Credit Loan").
		SetCurrency("").
		SetAccountNumber("568-2-817-40-9").
		SetIssuer("").
		AddAccountBalanceIDs(accountBalanceTester2.ID).
		AddAccountDetailIDs(accountDetailTester2.ID).
		SaveX(ctx)

	accountTester3 := client.Accounts.Create().
		SetType("Goal driven savings").
		SetCurrency("").
		SetAccountNumber("568-2-817-40-9").
		SetIssuer("TestLab").
		AddAccountBalanceIDs(accountBalanceTester3.ID).
		AddAccountDetailIDs(accountDetailTester3.ID).
		SaveX(ctx)

	accountTester4 := client.Accounts.Create().
		SetType("Credit Loan").
		SetCurrency("").
		SetAccountNumber("568-2-817-40-9").
		SetIssuer("").
		AddAccountBalanceIDs(accountBalanceTester4.ID).
		AddAccountDetailIDs(accountDetailTester4.ID).
		AddAccountFlagIDs(accountFlagTester1.ID, accountFlagTester2.ID).
		SaveX(ctx)

	// user
	userTester := client.Users.Create().
		SetName("Clare").
		SetImageURL("clare_img.png").
		AddUserGreetingIDs(userGreetingTester.ID).
		AddTransactionIDs(transactionTester1.ID, transactionTester2.ID, transactionTester3.ID, transactionTester4.ID, transactionTester5.ID, transactionTester6.ID).
		AddDebitCardDesignIDs(debitCardDesignTester1.ID, debitCardDesignTester2.ID, debitCardDesignTester3.ID, debitCardDesignTester4.ID).
		AddDebitCardDetailIDs(debitCardDetailTester1.ID, debitCardDetailTester2.ID, debitCardDetailTester3.ID, debitCardDetailTester4.ID).
		AddDebitCardStatuIDs(debitCardStatusTester1.ID, debitCardStatusTester2.ID, debitCardStatusTester3.ID, debitCardStatusTester4.ID).
		AddDebitCardIDs(debitCardTester1.ID, debitCardTester2.ID, debitCardTester3.ID, debitCardTester4.ID).
		AddAccountBalanceIDs(accountBalanceTester1.ID, accountBalanceTester2.ID, accountBalanceTester3.ID, accountBalanceTester4.ID).
		AddAccountDetailIDs(accountDetailTester1.ID, accountDetailTester2.ID, accountDetailTester3.ID, accountDetailTester4.ID).
		AddAccountFlagIDs(accountFlagTester1.ID, accountFlagTester2.ID).
		AddAccountIDs(accountTester1.ID, accountTester2.ID, accountTester3.ID, accountTester4.ID).
		SaveX(ctx)

	log.Println("Create UserGreetings ", userGreetingTester)
	log.Println("Create Transactions ", transactionTester1, transactionTester2, transactionTester3, transactionTester4, transactionTester5, transactionTester6)
	log.Println("Create DebitCardDesign ", debitCardDesignTester1, debitCardDesignTester2, debitCardDesignTester3, debitCardDesignTester4)
	log.Println("Create DebitCardStatus ", debitCardStatusTester1, debitCardStatusTester2, debitCardStatusTester3, debitCardStatusTester4)
	log.Println("Create DebitCardDetails ", debitCardDetailTester1, debitCardDetailTester2, debitCardDetailTester3, debitCardDetailTester4)
	log.Println("Create DebitCards ", debitCardTester1, debitCardTester2, debitCardTester3, debitCardTester4)
	log.Println("Create AccountBalances ", accountBalanceTester1, accountBalanceTester2, accountBalanceTester3, accountBalanceTester4)
	log.Println("Create AccountDetails ", accountDetailTester1, accountDetailTester2, accountDetailTester3, accountDetailTester4)
	log.Println("Create AccountFlags ", accountFlagTester1, accountFlagTester2)
	log.Println("Create Accounts ", accountTester1, accountTester2, accountTester3, accountTester4)
	log.Println("Create User ", userTester)
}
