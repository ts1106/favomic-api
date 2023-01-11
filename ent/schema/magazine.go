package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
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
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Annotations(
				entproto.Field(1),
			),
		field.String("name").
			Unique().
			NotEmpty().
			Annotations(
				entproto.Field(2),
			),
	}
}

// Edges of the Magazine.
func (Magazine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comics", Comic.Type).Annotations(
			entsql.Annotation{
				OnDelete: entsql.Cascade,
			},
			entproto.Field(3),
		),
	}
}

func (Magazine) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// entproto.Message(),
		// entproto.Service(),
	}
}
