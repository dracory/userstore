package userstore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dracory/database"
)

// == TYPE ====================================================================

type storeImplementation struct {
	// roleTableName      string
	userTableName      string
	db                 *sql.DB
	dbDriverName       string
	automigrateEnabled bool
	debugEnabled       bool
}

// == INTERFACE ===============================================================

var _ StoreInterface = (*storeImplementation)(nil) // verify it extends the interface

// PUBLIC METHODS ============================================================

// AutoMigrate auto migrate
func (store *storeImplementation) AutoMigrate() error {
	sqlStr, err := store.sqlUserTableCreate()
	if err != nil {
		return err
	}

	if sqlStr == "" {
		return errors.New("user table create sql is empty")
	}

	if store.db == nil {
		return errors.New("userstore: database is nil")
	}

	_, err = store.db.Exec(sqlStr)

	if err != nil {
		return err
	}

	return nil
}

// DB - returns the database
func (store *storeImplementation) DB() *sql.DB {
	return store.db
}

// EnableDebug - enables the debug option
func (st *storeImplementation) EnableDebug(debug bool) {
	st.debugEnabled = debug
}

func (store *storeImplementation) toQuerableContext(ctx context.Context) database.QueryableContext {
	if database.IsQueryableContext(ctx) {
		return ctx.(database.QueryableContext)
	}

	return database.Context(ctx, store.db)
}
