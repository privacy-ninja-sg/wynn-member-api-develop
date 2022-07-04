// Code generated by entc, DO NOT EDIT.

package gameaccount

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the gameaccount type in the database.
	Label = "game_account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeGame holds the string denoting the game edge name in mutations.
	EdgeGame = "game"
	// EdgePgslot holds the string denoting the pgslot edge name in mutations.
	EdgePgslot = "pgslot"
	// EdgePretty holds the string denoting the pretty edge name in mutations.
	EdgePretty = "pretty"
	// EdgeSagame holds the string denoting the sagame edge name in mutations.
	EdgeSagame = "sagame"
	// Table holds the table name of the gameaccount in the database.
	Table = "game_accounts"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "game_accounts"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_games"
	// GameTable is the table that holds the game relation/edge.
	GameTable = "game_accounts"
	// GameInverseTable is the table name for the Game entity.
	// It exists in this package in order to avoid circular dependency with the "game" package.
	GameInverseTable = "games"
	// GameColumn is the table column denoting the game relation/edge.
	GameColumn = "game_accounts"
	// PgslotTable is the table that holds the pgslot relation/edge.
	PgslotTable = "pg_slot_accounts"
	// PgslotInverseTable is the table name for the PgSlotAccount entity.
	// It exists in this package in order to avoid circular dependency with the "pgslotaccount" package.
	PgslotInverseTable = "pg_slot_accounts"
	// PgslotColumn is the table column denoting the pgslot relation/edge.
	PgslotColumn = "game_account_pgslot"
	// PrettyTable is the table that holds the pretty relation/edge.
	PrettyTable = "pretty_game_accounts"
	// PrettyInverseTable is the table name for the PrettyGameAccount entity.
	// It exists in this package in order to avoid circular dependency with the "prettygameaccount" package.
	PrettyInverseTable = "pretty_game_accounts"
	// PrettyColumn is the table column denoting the pretty relation/edge.
	PrettyColumn = "game_account_pretty"
	// SagameTable is the table that holds the sagame relation/edge.
	SagameTable = "sa_game_accounts"
	// SagameInverseTable is the table name for the SAGameAccount entity.
	// It exists in this package in order to avoid circular dependency with the "sagameaccount" package.
	SagameInverseTable = "sa_game_accounts"
	// SagameColumn is the table column denoting the sagame relation/edge.
	SagameColumn = "game_account_sagame"
)

// Columns holds all SQL columns for gameaccount fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "game_accounts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"game_accounts",
	"user_games",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)