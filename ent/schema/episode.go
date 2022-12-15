package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Episode holds the schema definition for the Episode entity.
type Episode struct {
	ent.Schema
}

// Fields of the Episode.
func (Episode) Fields() []ent.Field {
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
		field.String("url").
			Unique().
			NotEmpty().
			Annotations(
				entproto.Field(3),
			),
		field.String("Thumbnail").
			NotEmpty().
			Annotations(
				entproto.Field(4),
			),
		field.Time("updated_at").
			Default(time.Now).
			Annotations(
				entproto.Field(5),
			),
		field.UUID("comic_id", uuid.UUID{}).
			Annotations(
				entproto.Field(6),
			),
	}
}

// Edges of the Episode.
func (Episode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("comic", Comic.Type).
			Ref("episodes").
			Unique().
			Field("comic_id").
			Required().
			Annotations(
				entproto.Field(7),
			),
	}
}

func (Episode) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
