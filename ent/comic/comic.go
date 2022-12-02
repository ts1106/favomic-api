// Code generated by ent, DO NOT EDIT.

package comic

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the comic type in the database.
	Label = "comic"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldAuthorID holds the string denoting the author_id field in the database.
	FieldAuthorID = "author_id"
	// FieldMagazineID holds the string denoting the magazine_id field in the database.
	FieldMagazineID = "magazine_id"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeMagazine holds the string denoting the magazine edge name in mutations.
	EdgeMagazine = "magazine"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeEpisodes holds the string denoting the episodes edge name in mutations.
	EdgeEpisodes = "episodes"
	// Table holds the table name of the comic in the database.
	Table = "comics"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "comics"
	// AuthorInverseTable is the table name for the Author entity.
	// It exists in this package in order to avoid circular dependency with the "author" package.
	AuthorInverseTable = "authors"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "author_id"
	// MagazineTable is the table that holds the magazine relation/edge.
	MagazineTable = "comics"
	// MagazineInverseTable is the table name for the Magazine entity.
	// It exists in this package in order to avoid circular dependency with the "magazine" package.
	MagazineInverseTable = "magazines"
	// MagazineColumn is the table column denoting the magazine relation/edge.
	MagazineColumn = "magazine_id"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "comic_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// EpisodesTable is the table that holds the episodes relation/edge.
	EpisodesTable = "episodes"
	// EpisodesInverseTable is the table name for the Episode entity.
	// It exists in this package in order to avoid circular dependency with the "episode" package.
	EpisodesInverseTable = "episodes"
	// EpisodesColumn is the table column denoting the episodes relation/edge.
	EpisodesColumn = "comic_id"
)

// Columns holds all SQL columns for comic fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldAuthorID,
	FieldMagazineID,
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"comic_id", "tag_id"}
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
