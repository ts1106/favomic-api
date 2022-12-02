package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Magazine holds the schema definition for the Magazine entity.
type Magazine struct {
	ent.Schema
}

// Fields of the Magazine.
func (Magazine) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").Unique().NotEmpty(),
	}
}

// Edges of the Magazine.
func (Magazine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comics", Comic.Type).Annotations(
			entsql.Annotation{
				OnDelete: entsql.Cascade,
			},
		),
	}
}
