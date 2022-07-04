// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"wynn-member-api/ent/channel"
	"wynn-member-api/ent/user"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// Tel holds the value of the "tel" field.
	Tel string `json:"tel,omitempty"`
	// Picture holds the value of the "picture" field.
	Picture string `json:"picture,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Status holds the value of the "status" field.
	Status user.Status `json:"status,omitempty"`
	// Bonus holds the value of the "bonus" field.
	Bonus user.Bonus `json:"bonus,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	channel_user *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Games holds the value of the games edge.
	Games []*GameAccount `json:"games,omitempty"`
	// Transfers holds the value of the transfers edge.
	Transfers []*TransferTransaction `json:"transfers,omitempty"`
	// Banks holds the value of the banks edge.
	Banks []*BankAccount `json:"banks,omitempty"`
	// AccessToken holds the value of the access_token edge.
	AccessToken []*AccessToken `json:"access_token,omitempty"`
	// Line holds the value of the line edge.
	Line []*LineAccount `json:"line,omitempty"`
	// Wallet holds the value of the wallet edge.
	Wallet []*MasterWalletTransaction `json:"wallet,omitempty"`
	// Channel holds the value of the channel edge.
	Channel *Channel `json:"channel,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// GamesOrErr returns the Games value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GamesOrErr() ([]*GameAccount, error) {
	if e.loadedTypes[0] {
		return e.Games, nil
	}
	return nil, &NotLoadedError{edge: "games"}
}

// TransfersOrErr returns the Transfers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TransfersOrErr() ([]*TransferTransaction, error) {
	if e.loadedTypes[1] {
		return e.Transfers, nil
	}
	return nil, &NotLoadedError{edge: "transfers"}
}

// BanksOrErr returns the Banks value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) BanksOrErr() ([]*BankAccount, error) {
	if e.loadedTypes[2] {
		return e.Banks, nil
	}
	return nil, &NotLoadedError{edge: "banks"}
}

// AccessTokenOrErr returns the AccessToken value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AccessTokenOrErr() ([]*AccessToken, error) {
	if e.loadedTypes[3] {
		return e.AccessToken, nil
	}
	return nil, &NotLoadedError{edge: "access_token"}
}

// LineOrErr returns the Line value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LineOrErr() ([]*LineAccount, error) {
	if e.loadedTypes[4] {
		return e.Line, nil
	}
	return nil, &NotLoadedError{edge: "line"}
}

// WalletOrErr returns the Wallet value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) WalletOrErr() ([]*MasterWalletTransaction, error) {
	if e.loadedTypes[5] {
		return e.Wallet, nil
	}
	return nil, &NotLoadedError{edge: "wallet"}
}

// ChannelOrErr returns the Channel value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ChannelOrErr() (*Channel, error) {
	if e.loadedTypes[6] {
		if e.Channel == nil {
			// The edge channel was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: channel.Label}
		}
		return e.Channel, nil
	}
	return nil, &NotLoadedError{edge: "channel"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldTel, user.FieldPicture, user.FieldUsername, user.FieldPassword, user.FieldStatus, user.FieldBonus:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case user.FieldUUID:
			values[i] = new(uuid.UUID)
		case user.ForeignKeys[0]: // channel_user
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				u.UUID = *value
			}
		case user.FieldTel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tel", values[i])
			} else if value.Valid {
				u.Tel = value.String
			}
		case user.FieldPicture:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field picture", values[i])
			} else if value.Valid {
				u.Picture = value.String
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = user.Status(value.String)
			}
		case user.FieldBonus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bonus", values[i])
			} else if value.Valid {
				u.Bonus = user.Bonus(value.String)
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field channel_user", value)
			} else if value.Valid {
				u.channel_user = new(int)
				*u.channel_user = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryGames queries the "games" edge of the User entity.
func (u *User) QueryGames() *GameAccountQuery {
	return (&UserClient{config: u.config}).QueryGames(u)
}

// QueryTransfers queries the "transfers" edge of the User entity.
func (u *User) QueryTransfers() *TransferTransactionQuery {
	return (&UserClient{config: u.config}).QueryTransfers(u)
}

// QueryBanks queries the "banks" edge of the User entity.
func (u *User) QueryBanks() *BankAccountQuery {
	return (&UserClient{config: u.config}).QueryBanks(u)
}

// QueryAccessToken queries the "access_token" edge of the User entity.
func (u *User) QueryAccessToken() *AccessTokenQuery {
	return (&UserClient{config: u.config}).QueryAccessToken(u)
}

// QueryLine queries the "line" edge of the User entity.
func (u *User) QueryLine() *LineAccountQuery {
	return (&UserClient{config: u.config}).QueryLine(u)
}

// QueryWallet queries the "wallet" edge of the User entity.
func (u *User) QueryWallet() *MasterWalletTransactionQuery {
	return (&UserClient{config: u.config}).QueryWallet(u)
}

// QueryChannel queries the "channel" edge of the User entity.
func (u *User) QueryChannel() *ChannelQuery {
	return (&UserClient{config: u.config}).QueryChannel(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(fmt.Sprintf("%v", u.UUID))
	builder.WriteString(", tel=")
	builder.WriteString(u.Tel)
	builder.WriteString(", picture=")
	builder.WriteString(u.Picture)
	builder.WriteString(", username=")
	builder.WriteString(u.Username)
	builder.WriteString(", password=<sensitive>")
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteString(", bonus=")
	builder.WriteString(fmt.Sprintf("%v", u.Bonus))
	builder.WriteString(", created_at=")
	builder.WriteString(fmt.Sprintf("%v", u.CreatedAt))
	builder.WriteString(", updated_at=")
	builder.WriteString(fmt.Sprintf("%v", u.UpdatedAt))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}