// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountBalancesColumns holds the columns for the "account_balances" table.
	AccountBalancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "amount", Type: field.TypeFloat64},
		{Name: "account_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// AccountBalancesTable holds the schema information for the "account_balances" table.
	AccountBalancesTable = &schema.Table{
		Name:       "account_balances",
		Columns:    AccountBalancesColumns,
		PrimaryKey: []*schema.Column{AccountBalancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "account_balances_accounts_account_balances",
				Columns:    []*schema.Column{AccountBalancesColumns[2]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "account_balances_users_account_balances",
				Columns:    []*schema.Column{AccountBalancesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AccountDetailsColumns holds the columns for the "account_details" table.
	AccountDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "color", Type: field.TypeString},
		{Name: "is_main_account", Type: field.TypeBool},
		{Name: "progress", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
		{Name: "account_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// AccountDetailsTable holds the schema information for the "account_details" table.
	AccountDetailsTable = &schema.Table{
		Name:       "account_details",
		Columns:    AccountDetailsColumns,
		PrimaryKey: []*schema.Column{AccountDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "account_details_accounts_account_details",
				Columns:    []*schema.Column{AccountDetailsColumns[5]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "account_details_users_account_details",
				Columns:    []*schema.Column{AccountDetailsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AccountFlagsColumns holds the columns for the "account_flags" table.
	AccountFlagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "flag_type", Type: field.TypeString},
		{Name: "flag_value", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "account_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// AccountFlagsTable holds the schema information for the "account_flags" table.
	AccountFlagsTable = &schema.Table{
		Name:       "account_flags",
		Columns:    AccountFlagsColumns,
		PrimaryKey: []*schema.Column{AccountFlagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "account_flags_accounts_account_flags",
				Columns:    []*schema.Column{AccountFlagsColumns[5]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "account_flags_users_account_flags",
				Columns:    []*schema.Column{AccountFlagsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString},
		{Name: "currency", Type: field.TypeString},
		{Name: "account_number", Type: field.TypeString},
		{Name: "issuer", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounts_users_accounts",
				Columns:    []*schema.Column{AccountsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BannersColumns holds the columns for the "banners" table.
	BannersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "image", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// BannersTable holds the schema information for the "banners" table.
	BannersTable = &schema.Table{
		Name:       "banners",
		Columns:    BannersColumns,
		PrimaryKey: []*schema.Column{BannersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "banners_users_banners",
				Columns:    []*schema.Column{BannersColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DebitCardDesignsColumns holds the columns for the "debit_card_designs" table.
	DebitCardDesignsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "color", Type: field.TypeString},
		{Name: "border_color", Type: field.TypeString},
		{Name: "card_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// DebitCardDesignsTable holds the schema information for the "debit_card_designs" table.
	DebitCardDesignsTable = &schema.Table{
		Name:       "debit_card_designs",
		Columns:    DebitCardDesignsColumns,
		PrimaryKey: []*schema.Column{DebitCardDesignsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "debit_card_designs_debit_cards_debit_card_design",
				Columns:    []*schema.Column{DebitCardDesignsColumns[3]},
				RefColumns: []*schema.Column{DebitCardsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "debit_card_designs_users_debit_card_design",
				Columns:    []*schema.Column{DebitCardDesignsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DebitCardDetailsColumns holds the columns for the "debit_card_details" table.
	DebitCardDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "issuer", Type: field.TypeString},
		{Name: "number", Type: field.TypeString},
		{Name: "card_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// DebitCardDetailsTable holds the schema information for the "debit_card_details" table.
	DebitCardDetailsTable = &schema.Table{
		Name:       "debit_card_details",
		Columns:    DebitCardDetailsColumns,
		PrimaryKey: []*schema.Column{DebitCardDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "debit_card_details_debit_cards_debit_card_details",
				Columns:    []*schema.Column{DebitCardDetailsColumns[3]},
				RefColumns: []*schema.Column{DebitCardsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "debit_card_details_users_debit_card_details",
				Columns:    []*schema.Column{DebitCardDetailsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DebitCardStatusColumns holds the columns for the "debit_card_status" table.
	DebitCardStatusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeString},
		{Name: "card_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// DebitCardStatusTable holds the schema information for the "debit_card_status" table.
	DebitCardStatusTable = &schema.Table{
		Name:       "debit_card_status",
		Columns:    DebitCardStatusColumns,
		PrimaryKey: []*schema.Column{DebitCardStatusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "debit_card_status_debit_cards_debit_card_status",
				Columns:    []*schema.Column{DebitCardStatusColumns[2]},
				RefColumns: []*schema.Column{DebitCardsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "debit_card_status_users_debit_card_status",
				Columns:    []*schema.Column{DebitCardStatusColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DebitCardsColumns holds the columns for the "debit_cards" table.
	DebitCardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// DebitCardsTable holds the schema information for the "debit_cards" table.
	DebitCardsTable = &schema.Table{
		Name:       "debit_cards",
		Columns:    DebitCardsColumns,
		PrimaryKey: []*schema.Column{DebitCardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "debit_cards_users_debit_cards",
				Columns:    []*schema.Column{DebitCardsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "image", Type: field.TypeString},
		{Name: "is_bank", Type: field.TypeBool},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transactions_users_transactions",
				Columns:    []*schema.Column{TransactionsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserGreetingsColumns holds the columns for the "user_greetings" table.
	UserGreetingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "greeting", Type: field.TypeString, Size: 2147483647},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// UserGreetingsTable holds the schema information for the "user_greetings" table.
	UserGreetingsTable = &schema.Table{
		Name:       "user_greetings",
		Columns:    UserGreetingsColumns,
		PrimaryKey: []*schema.Column{UserGreetingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_greetings_users_user_greetings",
				Columns:    []*schema.Column{UserGreetingsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "image_url", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountBalancesTable,
		AccountDetailsTable,
		AccountFlagsTable,
		AccountsTable,
		BannersTable,
		DebitCardDesignsTable,
		DebitCardDetailsTable,
		DebitCardStatusTable,
		DebitCardsTable,
		TransactionsTable,
		UserGreetingsTable,
		UsersTable,
	}
)

func init() {
	AccountBalancesTable.ForeignKeys[0].RefTable = AccountsTable
	AccountBalancesTable.ForeignKeys[1].RefTable = UsersTable
	AccountDetailsTable.ForeignKeys[0].RefTable = AccountsTable
	AccountDetailsTable.ForeignKeys[1].RefTable = UsersTable
	AccountFlagsTable.ForeignKeys[0].RefTable = AccountsTable
	AccountFlagsTable.ForeignKeys[1].RefTable = UsersTable
	AccountsTable.ForeignKeys[0].RefTable = UsersTable
	BannersTable.ForeignKeys[0].RefTable = UsersTable
	DebitCardDesignsTable.ForeignKeys[0].RefTable = DebitCardsTable
	DebitCardDesignsTable.ForeignKeys[1].RefTable = UsersTable
	DebitCardDetailsTable.ForeignKeys[0].RefTable = DebitCardsTable
	DebitCardDetailsTable.ForeignKeys[1].RefTable = UsersTable
	DebitCardStatusTable.ForeignKeys[0].RefTable = DebitCardsTable
	DebitCardStatusTable.ForeignKeys[1].RefTable = UsersTable
	DebitCardsTable.ForeignKeys[0].RefTable = UsersTable
	TransactionsTable.ForeignKeys[0].RefTable = UsersTable
	UserGreetingsTable.ForeignKeys[0].RefTable = UsersTable
}
