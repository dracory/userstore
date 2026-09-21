package userstore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dracory/neat"
)

// NewStoreOptions define the options for creating a new user store
type NewStoreOptions struct {
	// Table names
	UserTableName      string
	RolesEnabled       bool
	RoleTableName      string
	UserRoleTableName  string
	GroupsEnabled      bool
	GroupTableName     string
	UserGroupTableName string

	// Database
	DB *sql.DB

	// Options
	AutomigrateEnabled bool
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

	if opts.RolesEnabled && opts.UserRoleTableName == "" {
		return nil, errors.New("user store: UserRoleTableName is required when RolesEnabled is true")
	}

	if opts.GroupsEnabled && opts.GroupTableName == "" {
		return nil, errors.New("user store: GroupTableName is required when GroupsEnabled is true")
	}

	if opts.GroupsEnabled && opts.UserGroupTableName == "" {
		return nil, errors.New("user store: UserGroupTableName is required when GroupsEnabled is true")
	}

	neatDB, err := neat.NewFromSQLDB(opts.DB)
	if err != nil {
		return nil, err
	}

	store := &storeImplementation{
		userTableName:      opts.UserTableName,
		roleTableName:      opts.RoleTableName,
		userRoleTableName:  opts.UserRoleTableName,
		groupTableName:     opts.GroupTableName,
		userGroupTableName: opts.UserGroupTableName,
		db:                 neatDB,
		automigrateEnabled: opts.AutomigrateEnabled,
		rolesEnabled:       opts.RolesEnabled,
		groupsEnabled:      opts.GroupsEnabled,
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
