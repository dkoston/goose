
package main

import (
    "database/sql"
    "fmt"
    "log"
)

func Up_20130106222315(txn *sql.Tx) {
    stmt, err := txn.Prepare("INSERT INTO post (id, title, body) VALUES(2, 'test pos 2t', 'this is a test post 2');")
    if err != nil {
        txn.Rollback()
    }
    defer stmt.Close()
    err = txn.Commit()
    if err != nil {
        log.Fatal(err)
    }
}

func Down_20130106222315(txn *sql.Tx) {
    fmt.Println("Hello from migration 20130106222315 Down!")
}
