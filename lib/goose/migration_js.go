package goose

import (
	"encoding/gob"
	"log"
	"os"
	"os/exec"
	"database/sql"
)

func init() {
	gob.Register(PostgresDialect{})
	gob.Register(MySqlDialect{})
	gob.Register(Sqlite3Dialect{})
}

//
// Run a .js migration.
//
// This does not handle transactions. The JS file is run with `node filename.js`
// and is responsible for transaction control.
//
func runJSMigration(conf *DBConf, db *sql.DB, path string, version int64, direction bool, tablePrefix string) error {
	cmd := exec.Command("node", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("`node %v` failed: %v", path, err)
	}

	//Finalize migration if the node script worked
	if err := FinalizeMigration(conf, db, direction, version, tablePrefix); err != nil {
		log.Fatalf("error finalizing migration %s, quitting. (%v)", path, err)
	}

	return nil
}