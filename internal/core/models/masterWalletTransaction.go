package models

import (
	"github.com/google/uuid"
	"time"
	"wynn-member-api/ent/masterwallettransaction"
)

type MasterWallet struct {
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// Debit holds the value of the "debit" field.
	Debit float32 `json:"debit,omitempty"`
	// Credit holds the value of the "credit" field.
	Credit float32 `json:"credit,omitempty"`
	// Balance holds the value of the "balance" field.
	Balance float32 `json:"balance,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
	// TxnType holds the value of the "txn_type" field.
	TxnType masterwallettransaction.TxnType `json:"txn_type,omitempty"`
	// Status holds the value of the "status" field.
	Status masterwallettransaction.Status `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
