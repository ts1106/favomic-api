package schema

import (
	"time"

	"entgo.io/ent"
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
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("title").NotEmpty(),
		field.String("url").Unique().NotEmpty(),
		field.String("Thumbnail").NotEmpty(),
		field.Time("updated_at").Default(time.Now),
		field.UUID("comic_id", uuid.UUID{}),
	}
}

// Edges of the Episode.
func (Episode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("comic", Comic.Type).Ref("episodes").Unique().Field("comic_id").Required(),
	}
}
