// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthorsColumns holds the columns for the "authors" table.
	AuthorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// AuthorsTable holds the schema information for the "authors" table.
	AuthorsTable = &schema.Table{
		Name:       "authors",
		Columns:    AuthorsColumns,
		PrimaryKey: []*schema.Column{AuthorsColumns[0]},
	}
	// ComicsColumns holds the columns for the "comics" table.
	ComicsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "author_id", Type: field.TypeUUID},
		{Name: "magazine_id", Type: field.TypeUUID},
	}
	// ComicsTable holds the schema information for the "comics" table.
	ComicsTable = &schema.Table{
		Name:       "comics",
		Columns:    ComicsColumns,
		PrimaryKey: []*schema.Column{ComicsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comics_authors_comics",
				Columns:    []*schema.Column{ComicsColumns[2]},
				RefColumns: []*schema.Column{AuthorsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "comics_magazines_comics",
				Columns:    []*schema.Column{ComicsColumns[3]},
				RefColumns: []*schema.Column{MagazinesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "comic_title_author_id",
				Unique:  true,
				Columns: []*schema.Column{ComicsColumns[1], ComicsColumns[2]},
			},
		},
	}
	// EpisodesColumns holds the columns for the "episodes" table.
	EpisodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "url", Type: field.TypeString, Unique: true},
		{Name: "thumbnail", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "comic_id", Type: field.TypeUUID},
	}
	// EpisodesTable holds the schema information for the "episodes" table.
	EpisodesTable = &schema.Table{
		Name:       "episodes",
		Columns:    EpisodesColumns,
		PrimaryKey: []*schema.Column{EpisodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "episodes_comics_episodes",
				Columns:    []*schema.Column{EpisodesColumns[5]},
				RefColumns: []*schema.Column{ComicsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// MagazinesColumns holds the columns for the "magazines" table.
	MagazinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// MagazinesTable holds the schema information for the "magazines" table.
	MagazinesTable = &schema.Table{
		Name:       "magazines",
		Columns:    MagazinesColumns,
		PrimaryKey: []*schema.Column{MagazinesColumns[0]},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// ComicTagsColumns holds the columns for the "comic_tags" table.
	ComicTagsColumns = []*schema.Column{
		{Name: "comic_id", Type: field.TypeUUID},
		{Name: "tag_id", Type: field.TypeUUID},
	}
	// ComicTagsTable holds the schema information for the "comic_tags" table.
	ComicTagsTable = &schema.Table{
		Name:       "comic_tags",
		Columns:    ComicTagsColumns,
		PrimaryKey: []*schema.Column{ComicTagsColumns[0], ComicTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comic_tags_comic_id",
				Columns:    []*schema.Column{ComicTagsColumns[0]},
				RefColumns: []*schema.Column{ComicsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "comic_tags_tag_id",
				Columns:    []*schema.Column{ComicTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthorsTable,
		ComicsTable,
		EpisodesTable,
		MagazinesTable,
		TagsTable,
		ComicTagsTable,
	}
)

func init() {
	ComicsTable.ForeignKeys[0].RefTable = AuthorsTable
	ComicsTable.ForeignKeys[1].RefTable = MagazinesTable
	EpisodesTable.ForeignKeys[0].RefTable = ComicsTable
	ComicTagsTable.ForeignKeys[0].RefTable = ComicsTable
	ComicTagsTable.ForeignKeys[1].RefTable = TagsTable
}