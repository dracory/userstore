package userstore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dracory/database"
)

// == TYPE ====================================================================

type store struct {
	// roleTableName      string
	userTableName      string
	db                 *sql.DB
	dbDriverName       string
	automigrateEnabled bool
	debugEnabled       bool
}

// == INTERFACE ===============================================================

var _ StoreInterface = (*store)(nil) // verify it extends the interface

// PUBLIC METHODS ============================================================

// AutoMigrate auto migrate
func (store *store) AutoMigrate() error {
	sqlStr := store.sqlUserTableCreate()

	if sqlStr == "" {
		return errors.New("user table create sql is empty")
	}

	if store.db == nil {
		return errors.New("userstore: database is nil")
	}

	_, err := store.db.Exec(sqlStr)

	if err != nil {
		return err
	}

	return nil
}

// DB - returns the database
func (store *store) DB() *sql.DB {
	return store.db
}

// EnableDebug - enables the debug option
func (st *store) EnableDebug(debug bool) {
	st.debugEnabled = debug
}

func (store *store) toQuerableContext(ctx context.Context) database.QueryableContext {
	if database.IsQueryableContext(ctx) {
		return ctx.(database.QueryableContext)
	}

	return database.Context(ctx, store.db)
}
