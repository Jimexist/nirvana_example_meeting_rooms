package db

import (
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/caicloud/nirvana/cli"

	// import pq because we use Postgres driver
	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
	connStr    = "connStr"
	dbMaxOpen  = "dbMaxOpen"
	dbMaxIdle  = "dbMaxIdle"
	dbMaxLife  = "dbMaxLife"
)

// Connector models a connector to backend database
type Connector struct {
	db sq.DBProxyBeginner
}

// NewConnector creates a new database connector
func NewConnector() (connector *Connector, err error) {
	if !cli.IsSet(connStr) {
		return nil, fmt.Errorf("no %v set, abort connection", connStr)
	}
	var db *sql.DB
	if db, err = sql.Open(driverName, cli.GetString(connStr)); err != nil {
		return nil, err
	}
	if cli.IsSet(dbMaxOpen) {
		db.SetMaxOpenConns(cli.GetInt(dbMaxOpen))
	} else {
		db.SetMaxOpenConns(36)
	}
	if cli.IsSet(dbMaxIdle) {
		db.SetMaxIdleConns(cli.GetInt(dbMaxIdle))
	} else {
		db.SetMaxIdleConns(0)
	}
	if cli.IsSet(dbMaxLife) {
		db.SetConnMaxLifetime(cli.GetDuration(dbMaxLife))
	} else {
		db.SetConnMaxLifetime(1 * time.Minute)
	}
	return &Connector{
		db: sq.NewStmtCacheProxy(db),
	}, nil
}

// NewStatement creates a statement builder to be used in queries.
func (c *Connector) NewStatement() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(c.db)

}

// NewTx creates a statement builder as well as the underlying sql.Tx
// object to be used in date updates.
func (c *Connector) NewTx() (*sql.Tx, sq.StatementBuilderType, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return nil, sq.StatementBuilderType{}, err
	}
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(tx)
	return tx, builder, nil
}
