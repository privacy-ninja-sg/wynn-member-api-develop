package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// MasterWalletTransaction holds the schema definition for the MasterWalletTransaction entity.
type MasterWalletTransaction struct {
	ent.Schema
}

// Fields of the MasterWalletTransaction.
func (MasterWalletTransaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Float32("debit").Default(0),   // deposit
		field.Float32("credit").Default(0),  // withdraw
		field.Float32("balance").Default(0), // ending balance
		field.String("remark").Optional(),
		field.Enum("txn_type").Values("deposit", "withdraw", "transfer"),
		field.Enum("status").Values("waiting", "pending", "successfully", "rejected").Default("waiting"),
		field.Time("created_at").Default(time.Now).GoType(time.Time{}),
		field.Time("updated_at").Default(time.Now).GoType(time.Time{}),
	}
}

// Edges of the MasterWalletTransaction.
func (MasterWalletTransaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("wallet").Unique(),
	}
}
