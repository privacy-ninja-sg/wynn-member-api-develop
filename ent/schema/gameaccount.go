package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// GameAccount holds the schema definition for the GameAccount entity.
type GameAccount struct {
	ent.Schema
}

// Fields of the GameAccount.
func (GameAccount) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Time("created_at").Default(time.Now).GoType(time.Time{}),
		field.Time("updated_at").Default(time.Now).GoType(time.Time{}),
	}
}

// Edges of the GameAccount.
func (GameAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("games").Unique(),
		edge.From("game", Game.Type).Ref("accounts").Unique(),
		edge.To("pgslot", PgSlotAccount.Type),
		edge.To("pretty", PrettyGameAccount.Type),
		edge.To("sagame", SAGameAccount.Type),
	}
}
