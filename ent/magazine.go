// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/ts1106/favomic-api/ent/magazine"
)

// Magazine is the model entity for the Magazine schema.
type Magazine struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MagazineQuery when eager-loading is set.
	Edges MagazineEdges `json:"edges"`
}

// MagazineEdges holds the relations/edges for other nodes in the graph.
type MagazineEdges struct {
	// Comics holds the value of the comics edge.
	Comics []*Comic `json:"comics,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ComicsOrErr returns the Comics value or an error if the edge
// was not loaded in eager-loading.
func (e MagazineEdges) ComicsOrErr() ([]*Comic, error) {
	if e.loadedTypes[0] {
		return e.Comics, nil
	}
	return nil, &NotLoadedError{edge: "comics"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Magazine) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case magazine.FieldName:
			values[i] = new(sql.NullString)
		case magazine.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Magazine", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Magazine fields.
func (m *Magazine) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case magazine.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case magazine.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		}
	}
	return nil
}

// QueryComics queries the "comics" edge of the Magazine entity.
func (m *Magazine) QueryComics() *ComicQuery {
	return (&MagazineClient{config: m.config}).QueryComics(m)
}

// Update returns a builder for updating this Magazine.
// Note that you need to call Magazine.Unwrap() before calling this method if this Magazine
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Magazine) Update() *MagazineUpdateOne {
	return (&MagazineClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Magazine entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Magazine) Unwrap() *Magazine {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Magazine is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Magazine) String() string {
	var builder strings.Builder
	builder.WriteString("Magazine(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Magazines is a parsable slice of Magazine.
type Magazines []*Magazine

func (m Magazines) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
