package database

import (
	"database/sql"
	"log"
)

type DbLike interface {
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
}

func RollbackTx(tx *sql.Tx) {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		log.Fatalf("update drivers: unable to rollback: %v", rollbackErr)
	}
}
