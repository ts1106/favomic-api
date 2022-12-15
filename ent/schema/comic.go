package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Comic holds the schema definition for the Comic entity.
type Comic struct {
	ent.Schema
}

// Fields of the Comic.
func (Comic) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Annotations(
				entproto.Field(1),
			),
		field.String("title").
			NotEmpty().
			Annotations(
				entproto.Field(2),
			),
		field.UUID("author_id", uuid.UUID{}).
			Annotations(
				entproto.Field(3),
			),
		field.UUID("magazine_id", uuid.UUID{}).
			Annotations(
				entproto.Field(4),
			),
	}
}

// Edges of the Comic.
func (Comic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", Author.Type).
			Ref("comics").
			Unique().
			Field("author_id").
			Required().
			Annotations(
				entproto.Field(5),
			),
		edge.From("magazine", Magazine.Type).
			Ref("comics").
			Unique().
			Field("magazine_id").
			Required().
			Annotations(
				entproto.Field(6),
			),
		edge.To("tags", Tag.Type).Annotations(
			entsql.Annotation{
				OnDelete: entsql.Cascade,
			},
			entproto.Field(7),
		),
		edge.To("episodes", Episode.Type).Annotations(
			entsql.Annotation{
				OnDelete: entsql.Cascade,
			},
			entproto.Field(8),
		),
	}
}

func (Comic) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title").
			Edges("author").
			Unique(),
	}
}

func (Comic) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
