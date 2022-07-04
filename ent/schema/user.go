package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("tel").MaxLen(10).Unique(),
		field.String("picture").Optional(),
		field.String("username").Unique(),
		field.String("password").Sensitive().NotEmpty(),
		field.Enum("status").Values("inactive", "active", "pending").Default("inactive"),
		field.Enum("bonus").Values("accepted", "rejected").Default("rejected"),
		field.Time("created_at").Default(time.Now).GoType(time.Time{}),
		field.Time("updated_at").Default(time.Now).GoType(time.Time{}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("games", GameAccount.Type),
		edge.To("transfers", TransferTransaction.Type),
		edge.To("banks", BankAccount.Type),
		edge.To("access_token", AccessToken.Type),
		edge.To("line", LineAccount.Type),
		edge.To("wallet", MasterWalletTransaction.Type),
		edge.From("channel", Channel.Type).Ref("user").Unique(),
	}
}
