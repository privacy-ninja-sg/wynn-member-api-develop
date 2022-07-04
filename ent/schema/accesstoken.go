package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// AccessToken holds the schema definition for the AccessToken entity.
type AccessToken struct {
	ent.Schema
}

// Fields of the AccessToken.
func (AccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("line_token").Optional(),
		field.String("access_token").NotEmpty(),
		field.Time("token_expire").GoType(time.Time{}),
		field.String("ip").Optional(),
		field.Time("created_at").Default(time.Now).GoType(time.Time{}),
		field.Time("updated_at").Default(time.Now).GoType(time.Time{}),
	}
}

// Edges of the AccessToken.
func (AccessToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("access_token").Unique().Required(),
	}
}
