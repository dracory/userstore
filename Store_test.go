package userstore

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"testing"

	"github.com/gouniverse/base/database"
	"github.com/gouniverse/utils"
	_ "modernc.org/sqlite"
)

func initDB(filepath string) (*sql.DB, error) {
	if filepath != ":memory:" && utils.FileExists(filepath) {
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

func TestStoreWithTx(t *testing.T) {
	store, err := initStore("test_store_with_tx.db")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if store == nil {
		t.Fatal("unexpected nil store")
	}

	db := store.DB()

	if db == nil {
		t.Fatal("unexpected nil db")
	}

	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	tx, err := db.Begin()

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if tx == nil {
		t.Fatal("unexpected nil tx")
	}

	txCtx := database.Context(context.Background(), tx)

	// create user
	user := NewUser().
		SetStatus(USER_STATUS_UNVERIFIED).
		SetFirstName("John").
		SetMiddleNames("").
		SetLastName("Doe").
		SetEmail("test@test.com").
		SetPassword("").
		SetProfileImageUrl("http://test.com/profile.png")

	err = store.UserCreate(txCtx, user)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// update user
	user.SetFirstName("John 2")
	err = store.UserUpdate(txCtx, user)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// check user
	userFound, errFind := store.UserFindByID(database.Context(context.Background(), db), user.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if userFound != nil {
		t.Fatal("User MUST be nil, as transaction not committed")
	}

	if err := tx.Commit(); err != nil {
		t.Fatal("unexpected error:", err)
	}

	// check user
	userFound, errFind = store.UserFindByID(database.Context(context.Background(), db), user.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if userFound == nil {
		t.Fatal("User MUST be not nil, as transaction committed")
	}

	if userFound.FirstName() != "John 2" {
		t.Fatal("User MUST be John 2, as transaction committed")
	}
}
