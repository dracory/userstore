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

// AutoMigrate auto migrate (deprecated - use MigrateUp)
func (store *storeImplementation) AutoMigrate() error {
	return store.MigrateUp()
}

// MigrateUp creates the user table
func (store *storeImplementation) MigrateUp(tx ...*sql.Tx) error {
	var txToUse *sql.Tx
	if len(tx) > 0 {
		txToUse = tx[0]
	}

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

	var errExec error
	if txToUse != nil {
		_, errExec = txToUse.Exec(sqlStr)
	} else {
		_, errExec = store.db.Exec(sqlStr)
	}

	if errExec != nil {
		return errExec
	}

	return nil
}

// MigrateDown drops the user table
func (store *storeImplementation) MigrateDown(tx ...*sql.Tx) error {
	var txToUse *sql.Tx
	if len(tx) > 0 {
		txToUse = tx[0]
	}

	sqlStr, err := store.sqlUserTableDrop()
	if err != nil {
		return err
	}

	if sqlStr == "" {
		return errors.New("user table drop sql is empty")
	}

	if store.db == nil {
		return errors.New("userstore: database is nil")
	}

	var errExec error
	if txToUse != nil {
		_, errExec = txToUse.Exec(sqlStr)
	} else {
		_, errExec = store.db.Exec(sqlStr)
	}

	if errExec != nil {
		return errExec
	}

	return nil
}

// GetDB - returns the database
func (store *storeImplementation) GetDB() *sql.DB {
	return store.db
}

// GetUserTableName returns the user table name
func (store *storeImplementation) GetUserTableName() string {
	return store.userTableName
}

// SetUserTableName sets the user table name
func (store *storeImplementation) SetUserTableName(tableName string) {
	store.userTableName = tableName
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
