// Code generated by ent, DO NOT EDIT.

package token

import (
	"time"
)

const (
	// Label holds the string label denoting the token type in the database.
	Label = "token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUser holds the string denoting the user field in the database.
	FieldUser = "user"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldPermissions holds the string denoting the permissions field in the database.
	FieldPermissions = "permissions"
	// FieldGroup holds the string denoting the group field in the database.
	FieldGroup = "group"
	// Table holds the table name of the token in the database.
	Table = "tokens"
)

// Columns holds all SQL columns for token fields.
var Columns = []string{
	FieldID,
	FieldUser,
	FieldToken,
	FieldCreated,
	FieldPermissions,
	FieldGroup,
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

var (
	// DefaultCreated holds the default value on creation for the "created" field.
	DefaultCreated time.Time
)
