// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"wynn-member-api/ent/gameaccount"
	"wynn-member-api/ent/sagameaccount"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// SAGameAccount is the model entity for the SAGameAccount schema.
type SAGameAccount struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// DesktopURI holds the value of the "desktop_uri" field.
	DesktopURI string `json:"desktop_uri,omitempty"`
	// MobileURI holds the value of the "mobile_uri" field.
	MobileURI string `json:"mobile_uri,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// RawData holds the value of the "raw_data" field.
	RawData string `json:"raw_data,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SAGameAccountQuery when eager-loading is set.
	Edges               SAGameAccountEdges `json:"edges"`
	game_account_sagame *int
}

// SAGameAccountEdges holds the relations/edges for other nodes in the graph.
type SAGameAccountEdges struct {
	// Owner holds the value of the owner edge.
	Owner *GameAccount `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SAGameAccountEdges) OwnerOrErr() (*GameAccount, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: gameaccount.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SAGameAccount) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sagameaccount.FieldID:
			values[i] = new(sql.NullInt64)
		case sagameaccount.FieldUsername, sagameaccount.FieldPassword, sagameaccount.FieldDesktopURI, sagameaccount.FieldMobileURI, sagameaccount.FieldRawData:
			values[i] = new(sql.NullString)
		case sagameaccount.FieldCreatedAt, sagameaccount.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case sagameaccount.FieldUUID:
			values[i] = new(uuid.UUID)
		case sagameaccount.ForeignKeys[0]: // game_account_sagame
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SAGameAccount", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SAGameAccount fields.
func (sga *SAGameAccount) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sagameaccount.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sga.ID = int(value.Int64)
		case sagameaccount.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				sga.UUID = *value
			}
		case sagameaccount.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				sga.Username = value.String
			}
		case sagameaccount.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				sga.Password = value.String
			}
		case sagameaccount.FieldDesktopURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desktop_uri", values[i])
			} else if value.Valid {
				sga.DesktopURI = value.String
			}
		case sagameaccount.FieldMobileURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mobile_uri", values[i])
			} else if value.Valid {
				sga.MobileURI = value.String
			}
		case sagameaccount.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sga.CreatedAt = value.Time
			}
		case sagameaccount.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sga.UpdatedAt = value.Time
			}
		case sagameaccount.FieldRawData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field raw_data", values[i])
			} else if value.Valid {
				sga.RawData = value.String
			}
		case sagameaccount.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field game_account_sagame", value)
			} else if value.Valid {
				sga.game_account_sagame = new(int)
				*sga.game_account_sagame = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the SAGameAccount entity.
func (sga *SAGameAccount) QueryOwner() *GameAccountQuery {
	return (&SAGameAccountClient{config: sga.config}).QueryOwner(sga)
}

// Update returns a builder for updating this SAGameAccount.
// Note that you need to call SAGameAccount.Unwrap() before calling this method if this SAGameAccount
// was returned from a transaction, and the transaction was committed or rolled back.
func (sga *SAGameAccount) Update() *SAGameAccountUpdateOne {
	return (&SAGameAccountClient{config: sga.config}).UpdateOne(sga)
}

// Unwrap unwraps the SAGameAccount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sga *SAGameAccount) Unwrap() *SAGameAccount {
	tx, ok := sga.config.driver.(*txDriver)
	if !ok {
		panic("ent: SAGameAccount is not a transactional entity")
	}
	sga.config.driver = tx.drv
	return sga
}

// String implements the fmt.Stringer.
func (sga *SAGameAccount) String() string {
	var builder strings.Builder
	builder.WriteString("SAGameAccount(")
	builder.WriteString(fmt.Sprintf("id=%v", sga.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(fmt.Sprintf("%v", sga.UUID))
	builder.WriteString(", username=")
	builder.WriteString(sga.Username)
	builder.WriteString(", password=<sensitive>")
	builder.WriteString(", desktop_uri=")
	builder.WriteString(sga.DesktopURI)
	builder.WriteString(", mobile_uri=")
	builder.WriteString(sga.MobileURI)
	builder.WriteString(", created_at=")
	builder.WriteString(fmt.Sprintf("%v", sga.CreatedAt))
	builder.WriteString(", updated_at=")
	builder.WriteString(fmt.Sprintf("%v", sga.UpdatedAt))
	builder.WriteString(", raw_data=")
	builder.WriteString(sga.RawData)
	builder.WriteByte(')')
	return builder.String()
}

// SAGameAccounts is a parsable slice of SAGameAccount.
type SAGameAccounts []*SAGameAccount

func (sga SAGameAccounts) config(cfg config) {
	for _i := range sga {
		sga[_i].config = cfg
	}
}