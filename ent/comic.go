// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/ts1106/favomic-api/ent/author"
	"github.com/ts1106/favomic-api/ent/comic"
	"github.com/ts1106/favomic-api/ent/magazine"
)

// Comic is the model entity for the Comic schema.
type Comic struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// AuthorID holds the value of the "author_id" field.
	AuthorID uuid.UUID `json:"author_id,omitempty"`
	// MagazineID holds the value of the "magazine_id" field.
	MagazineID uuid.UUID `json:"magazine_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ComicQuery when eager-loading is set.
	Edges ComicEdges `json:"edges"`
}

// ComicEdges holds the relations/edges for other nodes in the graph.
type ComicEdges struct {
	// Author holds the value of the author edge.
	Author *Author `json:"author,omitempty"`
	// Magazine holds the value of the magazine edge.
	Magazine *Magazine `json:"magazine,omitempty"`
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// Episodes holds the value of the episodes edge.
	Episodes []*Episode `json:"episodes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ComicEdges) AuthorOrErr() (*Author, error) {
	if e.loadedTypes[0] {
		if e.Author == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: author.Label}
		}
		return e.Author, nil
	}
	return nil, &NotLoadedError{edge: "author"}
}

// MagazineOrErr returns the Magazine value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ComicEdges) MagazineOrErr() (*Magazine, error) {
	if e.loadedTypes[1] {
		if e.Magazine == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: magazine.Label}
		}
		return e.Magazine, nil
	}
	return nil, &NotLoadedError{edge: "magazine"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e ComicEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[2] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// EpisodesOrErr returns the Episodes value or an error if the edge
// was not loaded in eager-loading.
func (e ComicEdges) EpisodesOrErr() ([]*Episode, error) {
	if e.loadedTypes[3] {
		return e.Episodes, nil
	}
	return nil, &NotLoadedError{edge: "episodes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comic) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case comic.FieldTitle:
			values[i] = new(sql.NullString)
		case comic.FieldID, comic.FieldAuthorID, comic.FieldMagazineID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Comic", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comic fields.
func (c *Comic) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comic.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case comic.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				c.Title = value.String
			}
		case comic.FieldAuthorID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field author_id", values[i])
			} else if value != nil {
				c.AuthorID = *value
			}
		case comic.FieldMagazineID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field magazine_id", values[i])
			} else if value != nil {
				c.MagazineID = *value
			}
		}
	}
	return nil
}

// QueryAuthor queries the "author" edge of the Comic entity.
func (c *Comic) QueryAuthor() *AuthorQuery {
	return (&ComicClient{config: c.config}).QueryAuthor(c)
}

// QueryMagazine queries the "magazine" edge of the Comic entity.
func (c *Comic) QueryMagazine() *MagazineQuery {
	return (&ComicClient{config: c.config}).QueryMagazine(c)
}

// QueryTags queries the "tags" edge of the Comic entity.
func (c *Comic) QueryTags() *TagQuery {
	return (&ComicClient{config: c.config}).QueryTags(c)
}

// QueryEpisodes queries the "episodes" edge of the Comic entity.
func (c *Comic) QueryEpisodes() *EpisodeQuery {
	return (&ComicClient{config: c.config}).QueryEpisodes(c)
}

// Update returns a builder for updating this Comic.
// Note that you need to call Comic.Unwrap() before calling this method if this Comic
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comic) Update() *ComicUpdateOne {
	return (&ComicClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Comic entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comic) Unwrap() *Comic {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comic is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comic) String() string {
	var builder strings.Builder
	builder.WriteString("Comic(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("title=")
	builder.WriteString(c.Title)
	builder.WriteString(", ")
	builder.WriteString("author_id=")
	builder.WriteString(fmt.Sprintf("%v", c.AuthorID))
	builder.WriteString(", ")
	builder.WriteString("magazine_id=")
	builder.WriteString(fmt.Sprintf("%v", c.MagazineID))
	builder.WriteByte(')')
	return builder.String()
}

// Comics is a parsable slice of Comic.
type Comics []*Comic

func (c Comics) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
