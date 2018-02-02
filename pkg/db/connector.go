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

const driverName = "postgres"

var dbConnStr = cli.StringFlag{
	Name:     "db-conn-str",
	EnvKey:   "DB_CONN_STR",
	DefValue: "postgres://postgres:@localhost:5432/postgres",
}

var dbMaxLife = cli.DurationFlag{
	Name:     "db-max-life",
	EnvKey:   "DB_MAX_LIFE",
	DefValue: 1 * time.Minute,
}

var dbMaxIdle = cli.IntFlag{
	Name:     "db-max-idle",
	EnvKey:   "DB_MAX_IDLE",
	DefValue: 0,
}

var dbMaxOpen = cli.IntFlag{
	Name:     "db-max-open",
	EnvKey:   "DB_MAX_OPEN",
	DefValue: 36,
}

func Flags() []cli.Flag {
	return []cli.Flag{
		dbMaxOpen,
		dbMaxIdle,
		dbConnStr,
		dbMaxLife,
	}
}

// Connector models a connector to backend database
type Connector struct {
	db sq.DBProxyBeginner
}

// NewConnector creates a new database connector
func NewConnector() (connector *Connector, err error) {
	if !cli.IsSet(dbConnStr.GetName()) {
		return nil, fmt.Errorf("no %v set, abort connection", dbConnStr.GetName())
	}
	var db *sql.DB
	if db, err = sql.Open(driverName, cli.GetString(dbConnStr.GetName())); err != nil {
		return nil, err
	}
	if cli.IsSet(dbMaxOpen.GetName()) {
		db.SetMaxOpenConns(cli.GetInt(dbMaxOpen.Name))
	} else {
		db.SetMaxOpenConns(36)
	}
	if cli.IsSet(dbMaxIdle.GetName()) {
		db.SetMaxIdleConns(cli.GetInt(dbMaxIdle.GetName()))
	} else {
		db.SetMaxIdleConns(0)
	}
	if cli.IsSet(dbMaxLife.GetName()) {
		db.SetConnMaxLifetime(cli.GetDuration(dbMaxLife.GetName()))
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
