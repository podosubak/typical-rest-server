package dbtxn

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/typical-go/typical-go/pkg/errkit"
)

// ContextKey to get transaction
const ContextKey key = iota

type (
	key int
	// Context of transaction
	Context struct {
		TxMap map[*sql.DB]Tx
		Err   error
	}
	// CommitFn is commit function to close the transaction
	CommitFn func() error
	// UseHandler responsible to handle transaction
	UseHandler struct {
		*Context
		sq.BaseRunner
	}
	// Tx is interface for *db.Tx
	Tx interface {
		sq.BaseRunner
		Rollback() error
		Commit() error
	}
)

// NewContext return new instance of Context
func NewContext() *Context {
	return &Context{TxMap: make(map[*sql.DB]Tx)}
}

// Begin transaction
func Begin(parent *context.Context) *Context {
	c := NewContext()
	*parent = context.WithValue(*parent, ContextKey, c)
	return c
}

// Use transaction if possible
func Use(ctx context.Context, db *sql.DB) (*UseHandler, error) {
	if ctx == nil {
		return nil, errors.New("dbtxn: missing context.Context")
	}

	c := Find(ctx)
	if c == nil { // NOTE: not transactional
		return &UseHandler{BaseRunner: db}, nil
	}

	tx, err := c.Begin(ctx, db)
	if err != nil {
		return nil, err
	}

	return &UseHandler{BaseRunner: tx, Context: c}, nil
}

// Find transaction context
func Find(ctx context.Context) *Context {
	if ctx == nil {
		return nil
	}
	c, _ := ctx.Value(ContextKey).(*Context)
	return c
}

// Error of transaction
func Error(ctx context.Context) error {
	if c := Find(ctx); c != nil {
		return c.Err
	}
	return nil
}

//
// Context
//

// Begin transaction
func (c *Context) Begin(ctx context.Context, db *sql.DB) (sq.BaseRunner, error) {
	tx, ok := c.TxMap[db]
	if ok {
		return tx, nil
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		c.Err = fmt.Errorf("dbtxn: %w", err)
		return nil, c.Err
	}
	c.TxMap[db] = tx
	return tx, nil
}

// Commit if no error
func (c *Context) Commit() error {
	var errs errkit.Errors
	if c.Err != nil {
		for _, tx := range c.TxMap {
			errs.Append(tx.Rollback())
		}
	} else {
		for _, tx := range c.TxMap {
			errs.Append(tx.Commit())
		}
	}

	return errs.Unwrap()
}

// SetError to set error to txn context
func (c *Context) SetError(err error) bool {
	if c != nil && err != nil {
		c.Err = err
		return true
	}
	return false
}
