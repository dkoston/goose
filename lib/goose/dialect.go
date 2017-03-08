package goose

import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
	"bytes"
	"fmt"
)

// SqlDialect abstracts the details of specific SQL dialects
// for goose's few SQL specific statements
type SqlDialect interface {
	createVersionTableSql(tablePrefix string) string // sql string to create the goose_db_version table
	insertVersionSql(tablePrefix string) string      // sql string to insert the initial version table row
	dbVersionQuery(db *sql.DB, tablePrefix string) (*sql.Rows, error)
}

// drivers that we don't know about can ask for a dialect by name
func dialectByName(d string) SqlDialect {
	switch d {
	case "postgres":
		return &PostgresDialect{}
	case "mysql":
		return &MySqlDialect{}
	case "sqlite3":
		return &Sqlite3Dialect{}
	}

	return nil
}

// get the tablename used for migrations
func getVersionTableName(tablePrefix string) string {
	var buffer bytes.Buffer
	buffer.WriteString(tablePrefix)
	buffer.WriteString("_db_version")
	return buffer.String()
}


////////////////////////////
// Postgres
////////////////////////////

type PostgresDialect struct{}

func (pg PostgresDialect) createVersionTableSql(tablePrefix string) string {
	tableName := getVersionTableName(tablePrefix)
	return fmt.Sprintf(`CREATE TABLE %s (
            	id serial NOT NULL,
                version_id bigint NOT NULL,
                is_applied boolean NOT NULL,
                tstamp timestamp NULL default now(),
                PRIMARY KEY(id)
            );`, tableName)
}

func (pg PostgresDialect) insertVersionSql(tablePrefix string) string {
	tableName := getVersionTableName(tablePrefix)
	return fmt.Sprintf("INSERT INTO %s (version_id, is_applied) VALUES ($1, $2);", tableName)
}

func (pg PostgresDialect) dbVersionQuery(db *sql.DB, tablePrefix string) (*sql.Rows, error) {
	tableName := getVersionTableName(tablePrefix)
	statement := fmt.Sprintf("SELECT version_id, is_applied from %s ORDER BY id DESC", tableName)
	rows, err := db.Query(statement)

	// XXX: check for postgres specific error indicating the table doesn't exist.
	// for now, assume any error is because the table doesn't exist,
	// in which case we'll try to create it.
	if err != nil {
		return nil, ErrTableDoesNotExist
	}

	return rows, err
}

////////////////////////////
// MySQL
////////////////////////////

type MySqlDialect struct{}

func (m MySqlDialect) createVersionTableSql(tablePrefix string) string {
	tableName := getVersionTableName(tablePrefix)
	return fmt.Sprintf(`CREATE TABLE %s (
                id serial NOT NULL,
                version_id bigint NOT NULL,
                is_applied boolean NOT NULL,
                tstamp timestamp NULL default now(),
                PRIMARY KEY(id)
            );`, tableName)
}

func (m MySqlDialect) insertVersionSql(tablePrefix string) string {
	tableName := getVersionTableName(tablePrefix)
	return fmt.Sprintf("INSERT INTO %s (version_id, is_applied) VALUES (?, ?);", tableName)
}

func (m MySqlDialect) dbVersionQuery(db *sql.DB, tablePrefix string) (*sql.Rows, error) {
	tableName := getVersionTableName(tablePrefix)
	statement := fmt.Sprintf("SELECT version_id, is_applied from %s ORDER BY id DESC", tableName)
	rows, err := db.Query(statement)

	// XXX: check for mysql specific error indicating the table doesn't exist.
	// for now, assume any error is because the table doesn't exist,
	// in which case we'll try to create it.
	if err != nil {
		return nil, ErrTableDoesNotExist
	}

	return rows, err
}

////////////////////////////
// sqlite3
////////////////////////////

type Sqlite3Dialect struct{}

func (m Sqlite3Dialect) createVersionTableSql(tablePrefix string) string {
	tableName := getVersionTableName(tablePrefix)
	return fmt.Sprintf(`CREATE TABLE %s (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                version_id INTEGER NOT NULL,
                is_applied INTEGER NOT NULL,
                tstamp TIMESTAMP DEFAULT (datetime('now'))
            );`, tableName)
}

func (m Sqlite3Dialect) insertVersionSql(tablePrefix string) string {
	tableName := getVersionTableName(tablePrefix)
	return fmt.Sprintf("INSERT INTO %s (version_id, is_applied) VALUES (?, ?);", tableName)
}

func (m Sqlite3Dialect) dbVersionQuery(db *sql.DB, tablePrefix string) (*sql.Rows, error) {
	tableName := getVersionTableName(tablePrefix)
	statement := fmt.Sprintf("SELECT version_id, is_applied from %s ORDER BY id DESC", tableName)
	rows, err := db.Query(statement)

	switch err.(type) {
	case sqlite3.Error:
		return nil, ErrTableDoesNotExist
	}
	return rows, err
}
