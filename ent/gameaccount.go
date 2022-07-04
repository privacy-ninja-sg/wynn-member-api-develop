// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"wynn-member-api/ent/game"
	"wynn-member-api/ent/gameaccount"
	"wynn-member-api/ent/user"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// GameAccount is the model entity for the GameAccount schema.
type GameAccount struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GameAccountQuery when eager-loading is set.
	Edges         GameAccountEdges `json:"edges"`
	game_accounts *int
	user_games    *int
}

// GameAccountEdges holds the relations/edges for other nodes in the graph.
type GameAccountEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Game holds the value of the game edge.
	Game *Game `json:"game,omitempty"`
	// Pgslot holds the value of the pgslot edge.
	Pgslot []*PgSlotAccount `json:"pgslot,omitempty"`
	// Pretty holds the value of the pretty edge.
	Pretty []*PrettyGameAccount `json:"pretty,omitempty"`
	// Sagame holds the value of the sagame edge.
	Sagame []*SAGameAccount `json:"sagame,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GameAccountEdges) OwnerOrErr() (*User, error) {
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
func (e GameAccountEdges) GameOrErr() (*Game, error) {
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

// PgslotOrErr returns the Pgslot value or an error if the edge
// was not loaded in eager-loading.
func (e GameAccountEdges) PgslotOrErr() ([]*PgSlotAccount, error) {
	if e.loadedTypes[2] {
		return e.Pgslot, nil
	}
	return nil, &NotLoadedError{edge: "pgslot"}
}

// PrettyOrErr returns the Pretty value or an error if the edge
// was not loaded in eager-loading.
func (e GameAccountEdges) PrettyOrErr() ([]*PrettyGameAccount, error) {
	if e.loadedTypes[3] {
		return e.Pretty, nil
	}
	return nil, &NotLoadedError{edge: "pretty"}
}

// SagameOrErr returns the Sagame value or an error if the edge
// was not loaded in eager-loading.
func (e GameAccountEdges) SagameOrErr() ([]*SAGameAccount, error) {
	if e.loadedTypes[4] {
		return e.Sagame, nil
	}
	return nil, &NotLoadedError{edge: "sagame"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GameAccount) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case gameaccount.FieldID:
			values[i] = new(sql.NullInt64)
		case gameaccount.FieldCreatedAt, gameaccount.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case gameaccount.FieldUUID:
			values[i] = new(uuid.UUID)
		case gameaccount.ForeignKeys[0]: // game_accounts
			values[i] = new(sql.NullInt64)
		case gameaccount.ForeignKeys[1]: // user_games
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GameAccount", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GameAccount fields.
func (ga *GameAccount) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case gameaccount.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ga.ID = int(value.Int64)
		case gameaccount.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				ga.UUID = *value
			}
		case gameaccount.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ga.CreatedAt = value.Time
			}
		case gameaccount.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ga.UpdatedAt = value.Time
			}
		case gameaccount.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field game_accounts", value)
			} else if value.Valid {
				ga.game_accounts = new(int)
				*ga.game_accounts = int(value.Int64)
			}
		case gameaccount.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_games", value)
			} else if value.Valid {
				ga.user_games = new(int)
				*ga.user_games = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the GameAccount entity.
func (ga *GameAccount) QueryOwner() *UserQuery {
	return (&GameAccountClient{config: ga.config}).QueryOwner(ga)
}

// QueryGame queries the "game" edge of the GameAccount entity.
func (ga *GameAccount) QueryGame() *GameQuery {
	return (&GameAccountClient{config: ga.config}).QueryGame(ga)
}

// QueryPgslot queries the "pgslot" edge of the GameAccount entity.
func (ga *GameAccount) QueryPgslot() *PgSlotAccountQuery {
	return (&GameAccountClient{config: ga.config}).QueryPgslot(ga)
}

// QueryPretty queries the "pretty" edge of the GameAccount entity.
func (ga *GameAccount) QueryPretty() *PrettyGameAccountQuery {
	return (&GameAccountClient{config: ga.config}).QueryPretty(ga)
}

// QuerySagame queries the "sagame" edge of the GameAccount entity.
func (ga *GameAccount) QuerySagame() *SAGameAccountQuery {
	return (&GameAccountClient{config: ga.config}).QuerySagame(ga)
}

// Update returns a builder for updating this GameAccount.
// Note that you need to call GameAccount.Unwrap() before calling this method if this GameAccount
// was returned from a transaction, and the transaction was committed or rolled back.
func (ga *GameAccount) Update() *GameAccountUpdateOne {
	return (&GameAccountClient{config: ga.config}).UpdateOne(ga)
}

// Unwrap unwraps the GameAccount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ga *GameAccount) Unwrap() *GameAccount {
	tx, ok := ga.config.driver.(*txDriver)
	if !ok {
		panic("ent: GameAccount is not a transactional entity")
	}
	ga.config.driver = tx.drv
	return ga
}

// String implements the fmt.Stringer.
func (ga *GameAccount) String() string {
	var builder strings.Builder
	builder.WriteString("GameAccount(")
	builder.WriteString(fmt.Sprintf("id=%v", ga.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(fmt.Sprintf("%v", ga.UUID))
	builder.WriteString(", created_at=")
	builder.WriteString(fmt.Sprintf("%v", ga.CreatedAt))
	builder.WriteString(", updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ga.UpdatedAt))
	builder.WriteByte(')')
	return builder.String()
}

// GameAccounts is a parsable slice of GameAccount.
type GameAccounts []*GameAccount

func (ga GameAccounts) config(cfg config) {
	for _i := range ga {
		ga[_i].config = cfg
	}
}