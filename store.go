package userstore

import (
	"context"
	"database/sql"

	"github.com/dracory/neat"
	contractsschema "github.com/dracory/neat/contracts/database/schema"
)

// == TYPE ====================================================================

type storeImplementation struct {
	userTableName      string
	db                 *neat.Database
	automigrateEnabled bool
	debugEnabled       bool
}

// == INTERFACE ===============================================================

var _ StoreInterface = (*storeImplementation)(nil) // verify it extends the interface

// PUBLIC METHODS ============================================================

// AutoMigrate auto migrate (deprecated - use MigrateUp)
func (store *storeImplementation) AutoMigrate() error {
	return store.MigrateUp(context.Background())
}

// MigrateUp creates the user table
func (store *storeImplementation) MigrateUp(ctx context.Context, tx ...*sql.Tx) error {
	if store.db.Schema().HasTable(store.userTableName) {
		return nil
	}

	err := store.db.Schema().Create(store.userTableName, func(table contractsschema.Blueprint) {
		table.String(COLUMN_ID, 21)
		table.Primary(COLUMN_ID)
		table.String(COLUMN_STATUS, 40)
		table.String(COLUMN_FIRST_NAME, 50)
		table.String(COLUMN_MIDDLE_NAMES, 50)
		table.String(COLUMN_LAST_NAME, 50)
		table.String(COLUMN_BUSINESS_NAME, 100)
		table.String(COLUMN_PHONE, 20)
		table.String(COLUMN_EMAIL, 100)
		table.String(COLUMN_PASSWORD, 255)
		table.String(COLUMN_ROLE, 50)
		table.String(COLUMN_COUNTRY, 2)
		table.String(COLUMN_TIMEZONE, 40)
		table.String(COLUMN_PROFILE_IMAGE_URL, 255)
		table.Text(COLUMN_METAS)
		table.Text(COLUMN_MEMO)
		table.DateTime(COLUMN_CREATED_AT)
		table.DateTime(COLUMN_UPDATED_AT)
		table.DateTime(COLUMN_SOFT_DELETED_AT)
	})

	if err != nil {
		return err
	}

	return nil
}

// MigrateDown drops the user table
func (store *storeImplementation) MigrateDown(ctx context.Context, tx ...*sql.Tx) error {
	if !store.db.Schema().HasTable(store.userTableName) {
		return nil
	}

	err := store.db.Schema().Drop(store.userTableName)
	if err != nil {
		return err
	}
	return nil
}

// GetDB - returns the database
func (store *storeImplementation) GetDB() *sql.DB {
	db, _ := store.db.DB()
	return db
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
	if debug {
		st.db.EnableDebug()
	} else {
		st.db.DisableDebug()
	}
}
