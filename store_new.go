package userstore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dracory/neat"
)

// NewStoreOptions define the options for creating a new user store
type NewStoreOptions struct {
	UserTableName      string
	RoleTableName      string
	DB                 *sql.DB
	AutomigrateEnabled bool
	RolesEnabled       bool
	DebugEnabled       bool
}

// NewStore creates a new user store
func NewStore(opts NewStoreOptions) (StoreInterface, error) {
	if opts.UserTableName == "" {
		return nil, errors.New("user store: UserTableName is required")
	}

	if opts.DB == nil {
		return nil, errors.New("user store: DB is required")
	}

	if opts.RolesEnabled && opts.RoleTableName == "" {
		return nil, errors.New("user store: RoleTableName is required when RolesEnabled is true")
	}

	neatDB, err := neat.NewFromSQLDB(opts.DB)
	if err != nil {
		return nil, err
	}

	store := &storeImplementation{
		userTableName:      opts.UserTableName,
		roleTableName:      opts.RoleTableName,
		db:                 neatDB,
		automigrateEnabled: opts.AutomigrateEnabled,
		rolesEnabled:       opts.RolesEnabled,
		debugEnabled:       opts.DebugEnabled,
	}

	if store.automigrateEnabled {
		err := store.MigrateUp(context.Background())

		if err != nil {
			return nil, err
		}
	}

	return store, nil
}
