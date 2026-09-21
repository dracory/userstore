package main

// Roles usage of the userstore package: enable roles, create roles,
// assign them to a user via the user_role junction table, and check
// memberships with UserRoles / UserHasRoles.
//
// Run with: go run ./examples/roles

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/dracory/userstore"
	_ "modernc.org/sqlite"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite", ":memory:?parseTime=true")
	if err != nil {
		return err
	}
	defer db.Close()

	store, err := userstore.NewStore(userstore.NewStoreOptions{
		DB:                 db,
		UserTableName:      "user",
		RoleTableName:      "role",
		UserRoleTableName:  "user_role",
		RolesEnabled:       true, // required to create the role and user_role tables
		AutomigrateEnabled: true,
	})
	if err != nil {
		return err
	}

	// Create a user
	user := userstore.NewUser().
		SetStatus(userstore.USER_STATUS_ACTIVE).
		SetEmail("john@example.com")

	if err := store.UserCreate(ctx, user); err != nil {
		return err
	}

	// Create roles
	admin, err := store.RoleFindByHandleOrCreate(ctx, "administrator", userstore.ROLE_STATUS_ACTIVE)
	if err != nil {
		return err
	}
	manager, err := store.RoleFindByHandleOrCreate(ctx, "manager", userstore.ROLE_STATUS_ACTIVE)
	if err != nil {
		return err
	}

	// Assign roles to the user
	_, err = store.UserRoleFindByUserIDAndRoleIDOrCreate(ctx, user.GetID(), admin.GetID())
	if err != nil {
		return err
	}
	_, err = store.UserRoleFindByUserIDAndRoleIDOrCreate(ctx, user.GetID(), manager.GetID())
	if err != nil {
		return err
	}

	// List the user's roles
	roles, err := store.UserRoles(ctx, user.GetID())
	if err != nil {
		return err
	}
	for _, role := range roles {
		fmt.Println("user has role:", role.GetHandle())
	}

	// Check single or multiple roles at once
	hasBoth, err := store.UserHasRoles(ctx, user.GetID(), []string{admin.GetID(), manager.GetID()})
	if err != nil {
		return err
	}
	fmt.Println("has admin+manager:", hasBoth)

	return nil
}
