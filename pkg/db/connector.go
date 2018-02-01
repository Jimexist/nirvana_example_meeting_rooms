package db

import (
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/caicloud/nirvana"
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
func NewConnector(config nirvana.Config) (connector *Connector, err error) {
	var db *sql.DB
	if connStr, ok := config.Config(connStr).(string); !ok {
		return nil, fmt.Errorf("cannot get `%v` from config", connStr)
	} else if db, err = sql.Open(driverName, connStr); err != nil {
		return nil, err
	}
	if dbMaxOpen, ok := config.Config(dbMaxOpen).(int); !ok {
		db.SetMaxOpenConns(36)
	} else {
		db.SetMaxOpenConns(dbMaxOpen)
	}
	if dbMaxIdle, ok := config.Config(dbMaxIdle).(int); !ok {
		db.SetMaxIdleConns(0)
	} else {
		db.SetMaxIdleConns(dbMaxIdle)
	}
	if dbMaxLife, ok := config.Config(dbMaxLife).(time.Duration); !ok {
		db.SetConnMaxLifetime(1 * time.Minute)
	} else {
		db.SetConnMaxLifetime(dbMaxLife)
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
