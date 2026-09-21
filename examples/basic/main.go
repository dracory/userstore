package main

// Basic usage of the userstore package: create the store,
// create a user, find it, update it, and soft delete it.
//
// Run with: go run ./examples/basic

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
		AutomigrateEnabled: true,
	})
	if err != nil {
		return err
	}

	// Create a user
	user := userstore.NewUser().
		SetStatus(userstore.USER_STATUS_ACTIVE).
		SetFirstName("John").
		SetLastName("Doe").
		SetEmail("john@example.com")

	if err := store.UserCreate(ctx, user); err != nil {
		return err
	}
	fmt.Println("created user:", user.GetID(), user.GetEmail())

	// Find by email
	found, err := store.UserFindByEmail(ctx, "john@example.com")
	if err != nil {
		return err
	}
	fmt.Println("found user:", found.GetFirstName(), found.GetLastName())

	// Update
	found.SetFirstName("Jane")
	if err := store.UserUpdate(ctx, found); err != nil {
		return err
	}
	fmt.Println("updated first name:", found.GetFirstName())

	// Count
	count, err := store.UserCount(ctx, userstore.NewUserQuery().SetStatus(userstore.USER_STATUS_ACTIVE))
	if err != nil {
		return err
	}
	fmt.Println("active users:", count)

	// Soft delete
	if err := store.UserSoftDelete(ctx, user); err != nil {
		return err
	}
	fmt.Println("soft deleted user:", user.GetID())

	return nil
}
