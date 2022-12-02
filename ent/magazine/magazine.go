// Code generated by ent, DO NOT EDIT.

package magazine

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the magazine type in the database.
	Label = "magazine"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeComics holds the string denoting the comics edge name in mutations.
	EdgeComics = "comics"
	// Table holds the table name of the magazine in the database.
	Table = "magazines"
	// ComicsTable is the table that holds the comics relation/edge.
	ComicsTable = "comics"
	// ComicsInverseTable is the table name for the Comic entity.
	// It exists in this package in order to avoid circular dependency with the "comic" package.
	ComicsInverseTable = "comics"
	// ComicsColumn is the table column denoting the comics relation/edge.
	ComicsColumn = "magazine_id"
)

// Columns holds all SQL columns for magazine fields.
var Columns = []string{
	FieldID,
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

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
