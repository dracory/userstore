package userstore

import (
	"context"
	"database/sql"
	"os"
	"strings"
	"testing"

	"github.com/gouniverse/base/database"
	"github.com/gouniverse/sb"
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

	return store, nil
}

func TestStoreUserCreate(t *testing.T) {
	db, err := initDB(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	store, err := NewStore(NewStoreOptions{
		DB:                 db,
		UserTableName:      "user_table_create",
		AutomigrateEnabled: true,
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if store == nil {
		t.Fatal("unexpected nil store")
	}

	user := NewUser().
		SetStatus(USER_STATUS_UNVERIFIED).
		SetFirstName("John").
		SetMiddleNames("").
		SetLastName("Doe").
		SetEmail("test@test.com").
		SetPassword("").
		SetProfileImageUrl("http://test.com/profile.png")

	err = store.UserCreate(context.Background(), user)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func TestStoreUserFindByEmail(t *testing.T) {
	db, err := initDB(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	store, err := NewStore(NewStoreOptions{
		DB:                 db,
		UserTableName:      "user_table_find_by_email",
		AutomigrateEnabled: true,
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	user := NewUser().
		SetStatus(USER_STATUS_UNVERIFIED).
		SetFirstName("John").
		SetMiddleNames("").
		SetLastName("Doe").
		SetEmail("test@test.com").
		SetPassword("").
		SetProfileImageUrl("http://test.com/profile.png")

	err = user.SetMetas(map[string]string{
		"education_1": "Education 1",
		"education_2": "Education 2",
		"education_3": "Education 3",
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.UserCreate(database.Context(context.Background(), db), user)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	userFound, errFind := store.UserFindByEmail(database.Context(context.Background(), db), user.Email())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if userFound == nil {
		t.Fatal("User MUST NOT be nil")
	}

	if userFound.ID() != user.ID() {
		t.Fatal("IDs do not match")
	}

	if userFound.Email() != user.Email() {
		t.Fatal("Emails do not match")
	}

	if userFound.FirstName() != user.FirstName() {
		t.Fatal("First names do not match")
	}

	if userFound.MiddleNames() != user.MiddleNames() {
		t.Fatal("Middle names do not match")
	}

	if userFound.LastName() != user.LastName() {
		t.Fatal("Last names do not match")
	}

	if userFound.ProfileImageUrl() != user.ProfileImageUrl() {
		t.Fatal("Profile image URLs do not match")
	}

	if userFound.Status() != user.Status() {
		t.Fatal("Statuses do not match")
	}

	if userFound.Role() != user.Role() {
		t.Fatal("Roles do not match")
	}

	if userFound.Role() != USER_ROLE_USER {
		t.Fatal("Roles MUST be:", USER_ROLE_USER, `found:`, userFound.Role())
	}

	if userFound.Meta("education_1") != user.Meta("education_1") {
		t.Fatal("Metas do not match")
	}

	if userFound.Meta("education_2") != user.Meta("education_2") {
		t.Fatal("Metas do not match")
	}

	if userFound.Meta("education_3") != user.Meta("education_3") {
		t.Fatal("Metas do not match")
	}
}

func TestStoreUserFindByID(t *testing.T) {
	db, err := initDB(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	store, err := NewStore(NewStoreOptions{
		DB:                 db,
		UserTableName:      "user_table_find_by_id",
		AutomigrateEnabled: true,
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	user := NewUser().
		SetStatus(USER_STATUS_UNVERIFIED).
		SetFirstName("John").
		SetMiddleNames("").
		SetLastName("Doe").
		SetEmail("test@test.com").
		SetPassword("").
		SetProfileImageUrl("http://test.com/profile.png")

	err = user.SetMetas(map[string]string{
		"education_1": "Education 1",
		"education_2": "Education 2",
		"education_3": "Education 3",
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.UserCreate(database.Context(context.Background(), db), user)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	userFound, errFind := store.UserFindByID(database.Context(context.Background(), db), user.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if userFound == nil {
		t.Fatal("User MUST NOT be nil")
	}

	if userFound.ID() != user.ID() {
		t.Fatal("IDs do not match")
	}

	if userFound.Email() != user.Email() {
		t.Fatal("Emails do not match")
	}

	if userFound.FirstName() != user.FirstName() {
		t.Fatal("First names do not match")
	}

	if userFound.MiddleNames() != user.MiddleNames() {
		t.Fatal("Middle names do not match")
	}

	if userFound.LastName() != user.LastName() {
		t.Fatal("Last names do not match")
	}

	if userFound.ProfileImageUrl() != user.ProfileImageUrl() {
		t.Fatal("Profile image URLs do not match")
	}

	if userFound.Status() != user.Status() {
		t.Fatal("Statuses do not match")
	}

	if userFound.Role() != user.Role() {
		t.Fatal("Roles do not match")
	}

	if userFound.Role() != USER_ROLE_USER {
		t.Fatal("Roles MUST be:", USER_ROLE_USER, `found:`, userFound.Role())
	}

	if userFound.Meta("education_1") != user.Meta("education_1") {
		t.Fatal("Metas do not match")
	}

	if userFound.Meta("education_2") != user.Meta("education_2") {
		t.Fatal("Metas do not match")
	}

	if userFound.Meta("education_3") != user.Meta("education_3") {
		t.Fatal("Metas do not match")
	}
}

func TestStoreUserSoftDelete(t *testing.T) {
	db, err := initDB(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	store, err := NewStore(NewStoreOptions{
		DB:                 db,
		UserTableName:      "user_table_find_by_id",
		AutomigrateEnabled: true,
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if store == nil {
		t.Fatal("unexpected nil store")
	}

	user := NewUser().
		SetStatus(USER_STATUS_UNVERIFIED).
		SetFirstName("John").
		SetMiddleNames("").
		SetLastName("Doe").
		SetEmail("test@test.com").
		SetPassword("").
		SetProfileImageUrl("http://test.com/profile.png")

	err = store.UserCreate(database.Context(context.Background(), db), user)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.UserSoftDeleteByID(database.Context(context.Background(), db), user.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if user.SoftDeletedAt() != sb.MAX_DATETIME {
		t.Fatal("User MUST NOT be soft deleted")
	}

	userFound, errFind := store.UserFindByID(database.Context(context.Background(), db), user.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if userFound != nil {
		t.Fatal("User MUST be nil")
	}
	query := NewUserQuery().SetWithSoftDeleted(true)

	query, err = query.SetID(user.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	query, err = query.SetLimit(1)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	userFindWithDeleted, err := store.UserList(database.Context(context.Background(), db), query)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(userFindWithDeleted) == 0 {
		t.Fatal("Exam MUST be soft deleted")
	}

	if strings.Contains(userFindWithDeleted[0].SoftDeletedAt(), sb.MAX_DATETIME) {
		t.Fatal("User MUST be soft deleted", user.SoftDeletedAt())
	}

	if !userFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("Page MUST be soft deleted")
	}

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
