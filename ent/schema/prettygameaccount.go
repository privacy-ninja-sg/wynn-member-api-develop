package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"regexp"
	"time"
)

// PrettyGameAccount holds the schema definition for the PrettyGameAccount entity.
type PrettyGameAccount struct {
	ent.Schema
} 

// Fields of the PrettyGameAccount.
func (PrettyGameAccount) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("username").Optional(),
		field.String("password").Sensitive().Optional(),
		field.String("desktop_uri").
			Match(regexp.MustCompile("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()!@:%_\\+.~#?&\\/\\/=]*)")).
			Optional(),
		field.String("mobile_uri").
			Match(regexp.MustCompile("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()!@:%_\\+.~#?&\\/\\/=]*)")).
			Optional(),
		field.Time("created_at").Default(time.Now).GoType(time.Time{}),
		field.Time("updated_at").Default(time.Now).GoType(time.Time{}),
		field.Text("raw_data").Optional(),
	}
}

// Edges of the PrettyGameAccount.
func (PrettyGameAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", GameAccount.Type).Ref("pretty").Unique(),
	}
}
