// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"wynn-member-api/ent/game"
	"wynn-member-api/ent/transfertransaction"
	"wynn-member-api/ent/user"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// TransferTransaction is the model entity for the TransferTransaction schema.
type TransferTransaction struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount float32 `json:"amount,omitempty"`
	// Status holds the value of the "status" field.
	Status transfertransaction.Status `json:"status,omitempty"`
	// TxnType holds the value of the "txn_type" field.
	TxnType transfertransaction.TxnType `json:"txn_type,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransferTransactionQuery when eager-loading is set.
	Edges          TransferTransactionEdges `json:"edges"`
	game_transfers *int
	user_transfers *int
}

// TransferTransactionEdges holds the relations/edges for other nodes in the graph.
type TransferTransactionEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Game holds the value of the game edge.
	Game *Game `json:"game,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransferTransactionEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// GameOrErr returns the Game value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransferTransactionEdges) GameOrErr() (*Game, error) {
	if e.loadedTypes[1] {
		if e.Game == nil {
			// The edge game was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: game.Label}
		}
		return e.Game, nil
	}
	return nil, &NotLoadedError{edge: "game"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TransferTransaction) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case transfertransaction.FieldAmount:
			values[i] = new(sql.NullFloat64)
		case transfertransaction.FieldID:
			values[i] = new(sql.NullInt64)
		case transfertransaction.FieldStatus, transfertransaction.FieldTxnType:
			values[i] = new(sql.NullString)
		case transfertransaction.FieldCreatedAt, transfertransaction.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case transfertransaction.FieldUUID:
			values[i] = new(uuid.UUID)
		case transfertransaction.ForeignKeys[0]: // game_transfers
			values[i] = new(sql.NullInt64)
		case transfertransaction.ForeignKeys[1]: // user_transfers
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TransferTransaction", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TransferTransaction fields.
func (tt *TransferTransaction) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transfertransaction.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tt.ID = int(value.Int64)
		case transfertransaction.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				tt.UUID = *value
			}
		case transfertransaction.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				tt.Amount = float32(value.Float64)
			}
		case transfertransaction.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				tt.Status = transfertransaction.Status(value.String)
			}
		case transfertransaction.FieldTxnType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field txn_type", values[i])
			} else if value.Valid {
				tt.TxnType = transfertransaction.TxnType(value.String)
			}
		case transfertransaction.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tt.CreatedAt = value.Time
			}
		case transfertransaction.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tt.UpdatedAt = value.Time
			}
		case transfertransaction.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field game_transfers", value)
			} else if value.Valid {
				tt.game_transfers = new(int)
				*tt.game_transfers = int(value.Int64)
			}
		case transfertransaction.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_transfers", value)
			} else if value.Valid {
				tt.user_transfers = new(int)
				*tt.user_transfers = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the TransferTransaction entity.
func (tt *TransferTransaction) QueryOwner() *UserQuery {
	return (&TransferTransactionClient{config: tt.config}).QueryOwner(tt)
}

// QueryGame queries the "game" edge of the TransferTransaction entity.
func (tt *TransferTransaction) QueryGame() *GameQuery {
	return (&TransferTransactionClient{config: tt.config}).QueryGame(tt)
}

// Update returns a builder for updating this TransferTransaction.
// Note that you need to call TransferTransaction.Unwrap() before calling this method if this TransferTransaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (tt *TransferTransaction) Update() *TransferTransactionUpdateOne {
	return (&TransferTransactionClient{config: tt.config}).UpdateOne(tt)
}

// Unwrap unwraps the TransferTransaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tt *TransferTransaction) Unwrap() *TransferTransaction {
	tx, ok := tt.config.driver.(*txDriver)
	if !ok {
		panic("ent: TransferTransaction is not a transactional entity")
	}
	tt.config.driver = tx.drv
	return tt
}

// String implements the fmt.Stringer.
func (tt *TransferTransaction) String() string {
	var builder strings.Builder
	builder.WriteString("TransferTransaction(")
	builder.WriteString(fmt.Sprintf("id=%v", tt.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(fmt.Sprintf("%v", tt.UUID))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", tt.Amount))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", tt.Status))
	builder.WriteString(", txn_type=")
	builder.WriteString(fmt.Sprintf("%v", tt.TxnType))
	builder.WriteString(", created_at=")
	builder.WriteString(fmt.Sprintf("%v", tt.CreatedAt))
	builder.WriteString(", updated_at=")
	builder.WriteString(fmt.Sprintf("%v", tt.UpdatedAt))
	builder.WriteByte(')')
	return builder.String()
}

// TransferTransactions is a parsable slice of TransferTransaction.
type TransferTransactions []*TransferTransaction

func (tt TransferTransactions) config(cfg config) {
	for _i := range tt {
		tt[_i].config = cfg
	}
}
