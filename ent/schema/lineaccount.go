package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// LineAccount holds the schema definition for the à¸LineAccount entity.
type LineAccount struct {
	ent.Schema
}

// Fields of the LineAccount.
func (LineAccount) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("line_id").Default("N/A"),
		field.String("line_client_id").Optional(),
		field.Time("created_at").Default(time.Now).GoType(time.Time{}),
		field.Time("updated_at").Default(time.Now).GoType(time.Time{}),
	}
}

// Edges of the LineAccount.
func (LineAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("line").Unique(),
	}
}
