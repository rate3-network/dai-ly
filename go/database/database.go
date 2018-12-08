package database

import (
	"database/sql/driver"
	"time"

	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
	"github.com/rate-engineering/dai-ly/go/zap"
)

type Store struct {
	*dbr.Connection
	Logger *zap.Logger
}

const (
	TransactionsT = "transactions"

	TXHashC     = "tx_hash"
	SenderC     = "sender"
	ReceiverC   = "receiver"
	DelegateC   = "delegate"
	FeeAmountC  = "fee_amount"
	SendAmountC = "send_amount"
	TokenC      = "token"
	NonceC      = "nonce"

	LastCreatedC  = "last_created"
	LastModifiedC = "last_modified"
)

// pgTimeFormat is RFC3339Nano which has timezone information and accepted by pg
var PgTimeFormat = time.RFC3339Nano

type PGTimezone struct{}

func (d PGTimezone) QuoteIdent(s string) string {
	return dialect.PostgreSQL.QuoteIdent(s)
}

func (d PGTimezone) EncodeString(s string) string {
	return dialect.PostgreSQL.EncodeString(s)
}

func (d PGTimezone) EncodeBool(b bool) string {
	return dialect.PostgreSQL.EncodeBool(b)
}

func (d PGTimezone) EncodeTime(t time.Time) string {
	return `'` + t.UTC().Format(PgTimeFormat) + `'`
}

func (d PGTimezone) EncodeBytes(b []byte) string {
	return dialect.PostgreSQL.EncodeBytes(b)
}

func (d PGTimezone) Placeholder(n int) string {
	return dialect.PostgreSQL.Placeholder(n)
}

// sql now utility function with timezone
type betterNowSentinel struct{}

// NowTZ replaces dbr.Now
// NOTE: DO NOT USE dbr.Now unless db is timezone is in UTC
var NowTZ = betterNowSentinel{}

func (n betterNowSentinel) Value() (driver.Value, error) {
	now := time.Now().UTC().Format(PgTimeFormat)
	return now, nil
}
