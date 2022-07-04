package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// BankAccount holds the schema definition for the BankAccount entity.
type BankAccount struct {
	ent.Schema
}

// Fields of the BankAccount.
func (BankAccount) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("bank_account_id"),
		field.String("bank_account_id_last").MaxLen(4).Unique(),
		field.String("bank_account_name"),
		field.Enum("status").Values("pending", "rejected", "approved").Default("pending"),
		field.String("bank_code").Optional(),
		field.Time("created_at").Default(time.Now).GoType(time.Time{}),
		field.Time("updated_at").Default(time.Now).GoType(time.Time{}),
	}
}

// Edges of the BankAccount.
func (BankAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("banks").Unique(),
		edge.From("bank", Bank.Type).Ref("accounts").Unique(),
	}
}
