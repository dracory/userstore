package main

// Groups usage of the userstore package: enable groups, create groups,
// add users via the user_group junction table, and check
// memberships with UserGroups / UserHasGroups.
//
// Run with: go run ./examples/groups

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
		GroupTableName:     "group",
		UserGroupTableName: "user_group",
		GroupsEnabled:      true, // required to create the group and user_group tables
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

	// Create groups
	staff, err := store.GroupFindByHandleOrCreate(ctx, "staff", userstore.GROUP_STATUS_ACTIVE)
	if err != nil {
		return err
	}
	vip, err := store.GroupFindByHandleOrCreate(ctx, "vip", userstore.GROUP_STATUS_ACTIVE)
	if err != nil {
		return err
	}

	// Add the user to the groups
	_, err = store.UserGroupFindByUserIDAndGroupIDOrCreate(ctx, user.GetID(), staff.GetID())
	if err != nil {
		return err
	}
	_, err = store.UserGroupFindByUserIDAndGroupIDOrCreate(ctx, user.GetID(), vip.GetID())
	if err != nil {
		return err
	}

	// List the user's groups
	groups, err := store.UserGroups(ctx, user.GetID())
	if err != nil {
		return err
	}
	for _, group := range groups {
		fmt.Println("user is member of group:", group.GetHandle())
	}

	// Check single or multiple groups at once
	isMember, err := store.UserHasGroups(ctx, user.GetID(), []string{staff.GetID(), vip.GetID()})
	if err != nil {
		return err
	}
	fmt.Println("member of staff+vip:", isMember)

	return nil
}
