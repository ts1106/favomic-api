// Code generated by ent, DO NOT EDIT.

package tag

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeComics holds the string denoting the comics edge name in mutations.
	EdgeComics = "comics"
	// Table holds the table name of the tag in the database.
	Table = "tags"
	// ComicsTable is the table that holds the comics relation/edge. The primary key declared below.
	ComicsTable = "comic_tags"
	// ComicsInverseTable is the table name for the Comic entity.
	// It exists in this package in order to avoid circular dependency with the "comic" package.
	ComicsInverseTable = "comics"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// ComicsPrimaryKey and ComicsColumn2 are the table columns denoting the
	// primary key for the comics relation (M2M).
	ComicsPrimaryKey = []string{"comic_id", "tag_id"}
)

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
