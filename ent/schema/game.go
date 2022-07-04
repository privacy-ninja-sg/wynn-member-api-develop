package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name"),
		field.String("banner").Optional(),
		field.Text("desc").Optional(),
		field.Enum("status").Values("on", "off").Default("off"),
		field.Time("created_at").Default(time.Now).GoType(time.Time{}),
		field.Time("updated_at").Default(time.Now).GoType(time.Time{}),
	}
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", GameAccount.Type),
		edge.To("transfers", TransferTransaction.Type),
	}
}
