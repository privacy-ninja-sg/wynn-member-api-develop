package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// TransferTransaction holds the schema definition for the TransferTransaction entity.
type TransferTransaction struct {
	ent.Schema
}

// Fields of the TransferTransaction.
func (TransferTransaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Float32("amount"),
		field.Enum("status").Values("waiting", "processing", "successfully", "rejected").Default("waiting"),
		field.Enum("txn_type").Values("deposit", "withdraw").Default("deposit"),
		field.Time("created_at").Default(time.Now).GoType(time.Time{}),
		field.Time("updated_at").Default(time.Now).GoType(time.Time{}),
	}
}

// Edges of the TransferTransaction.
func (TransferTransaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("transfers").Unique(),
		edge.From("game", Game.Type).Ref("transfers").Unique(),
	}
}
