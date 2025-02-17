// Code generated by ent, DO NOT EDIT.

package accountdetails

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the accountdetails type in the database.
	Label = "account_details"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// FieldIsMainAccount holds the string denoting the is_main_account field in the database.
	FieldIsMainAccount = "is_main_account"
	// FieldProgress holds the string denoting the progress field in the database.
	FieldProgress = "progress"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeAccounts holds the string denoting the accounts edge name in mutations.
	EdgeAccounts = "accounts"
	// Table holds the table name of the accountdetails in the database.
	Table = "account_details"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "account_details"
	// UsersInverseTable is the table name for the Users entity.
	// It exists in this package in order to avoid circular dependency with the "users" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_id"
	// AccountsTable is the table that holds the accounts relation/edge.
	AccountsTable = "account_details"
	// AccountsInverseTable is the table name for the Accounts entity.
	// It exists in this package in order to avoid circular dependency with the "accounts" package.
	AccountsInverseTable = "accounts"
	// AccountsColumn is the table column denoting the accounts relation/edge.
	AccountsColumn = "account_id"
)

// Columns holds all SQL columns for accountdetails fields.
var Columns = []string{
	FieldID,
	FieldAccountID,
	FieldUserID,
	FieldColor,
	FieldIsMainAccount,
	FieldProgress,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the AccountDetails queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAccountID orders the results by the account_id field.
func ByAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByColor orders the results by the color field.
func ByColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldColor, opts...).ToFunc()
}

// ByIsMainAccount orders the results by the is_main_account field.
func ByIsMainAccount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsMainAccount, opts...).ToFunc()
}

// ByProgress orders the results by the progress field.
func ByProgress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProgress, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByUsersField orders the results by users field.
func ByUsersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), sql.OrderByField(field, opts...))
	}
}

// ByAccountsField orders the results by accounts field.
func ByAccountsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccountsStep(), sql.OrderByField(field, opts...))
	}
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
	)
}
func newAccountsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccountsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AccountsTable, AccountsColumn),
	)
}
