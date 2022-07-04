package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Bank holds the schema definition for the Bank entity.
type Bank struct {
	ent.Schema
}

// Fields of the Bank.
func (Bank) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name").Unique(),
		field.String("short_name").Unique().Optional(),
		field.String("name_th").Optional(),
		field.String("short_name_th").Optional(),
		field.String("bank_account_name").Optional(),
		field.String("logo").Optional(),
		field.String("bank_id").Optional(),
		field.Enum("status").Values("on", "off").Default("off"),
		field.Time("created_at").Default(time.Now).GoType(time.Time{}),
		field.Time("updated_at").Default(time.Now).GoType(time.Time{}),
	}
}

// Edges of the Bank.
func (Bank) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", BankAccount.Type),
	}
}
