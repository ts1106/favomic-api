package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("title").NotEmpty(),
		field.UUID("author_id", uuid.UUID{}),
		field.UUID("magazine_id", uuid.UUID{}),
	}
}

// Edges of the Comic.
func (Comic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", Author.Type).Ref("comics").Unique().Field("author_id").Required(),
		edge.From("magazine", Magazine.Type).Ref("comics").Unique().Field("magazine_id").Required(),
		edge.To("tags", Tag.Type).Annotations(
			entsql.Annotation{
				OnDelete: entsql.Cascade,
			},
		),
		edge.To("episodes", Episode.Type).Annotations(
			entsql.Annotation{
				OnDelete: entsql.Cascade,
			},
		),
	}
}

func (Comic) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title").Edges("author").Unique(),
	}
}
