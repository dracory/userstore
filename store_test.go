package userstore

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

func initDB(filepath string) (*sql.DB, error) {
	if filepath != ":memory:" && fileExists(filepath) {
		err := os.Remove(filepath) // remove database

		if err != nil {
			return nil, err
		}
	}

	dsn := filepath + "?parseTime=true"
	db, err := sql.Open("sqlite", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func initStore(filepath string) (StoreInterface, error) {
	db, err := initDB(filepath)

	if err != nil {
		return nil, err
	}

	store, err := NewStore(NewStoreOptions{
		DB:                 db,
		UserTableName:      "user_table",
		RoleTableName:      "role_table",
		UserRoleTableName:  "user_role_table",
		GroupTableName:     "group_table",
		UserGroupTableName: "user_group_table",
		RolesEnabled:       true,
		GroupsEnabled:      true,
		AutomigrateEnabled: true,
	})

	if err != nil {
		return nil, err
	}

	if store == nil {
		return nil, errors.New("unexpected nil store")
	}

	return store, nil
}

func TestStoreMigrateUp_IDColumnLength(t *testing.T) {
	s, err := initStore(":memory:")
	if err != nil {
		t.Fatalf("failed to init store: %v", err)
	}

	// Verify insertion of 40-character IDs across all ID fields
	id40 := "1234567890123456789012345678901234567890"

	// 1. User Table
	user := NewUser()
	user.SetID(id40)
	err = s.UserCreate(context.Background(), user)
	if err != nil {
		t.Fatalf("failed to insert user with 40-char ID: %v", err)
	}

	// 2. Role Table
	role := NewRole()
	role.SetID(id40)
	err = s.RoleCreate(context.Background(), role)
	if err != nil {
		t.Fatalf("failed to insert role with 40-char ID: %v", err)
	}

	// 3. UserRole Table
	userRole := NewUserRole()
	userRole.SetID(id40)
	userRole.SetUserID(id40)
	userRole.SetRoleID(id40)
	err = s.UserRoleCreate(context.Background(), userRole)
	if err != nil {
		t.Fatalf("failed to insert user_role with 40-char IDs: %v", err)
	}

	// 4. Group Table
	group := NewGroup()
	group.SetID(id40)
	err = s.GroupCreate(context.Background(), group)
	if err != nil {
		t.Fatalf("failed to insert group with 40-char ID: %v", err)
	}

	// 5. UserGroup Table
	userGroup := NewUserGroup()
	userGroup.SetID(id40)
	userGroup.SetUserID(id40)
	userGroup.SetGroupID(id40)
	err = s.UserGroupCreate(context.Background(), userGroup)
	if err != nil {
		t.Fatalf("failed to insert user_group with 40-char IDs: %v", err)
	}
}
