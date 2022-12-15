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

// Author holds the schema definition for the Author entity.
type Author struct {
	ent.Schema
}

// Fields of the Author.
func (Author) Fields() []ent.Field {
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

// Edges of the Author.
func (Author) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comics", Comic.Type).Annotations(
			entsql.Annotation{
				OnDelete: entsql.Cascade,
			},
			entproto.Field(3),
		),
	}
}

func (Author) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
