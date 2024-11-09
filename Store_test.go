package userstore

import (
	"database/sql"
	"os"
	"strings"
	"testing"

	"github.com/gouniverse/sb"
	"github.com/gouniverse/utils"
	_ "modernc.org/sqlite"
)

func initDB(filepath string) *sql.DB {
	if filepath != ":memory:" && utils.FileExists(filepath) {
		err := os.Remove(filepath) // remove database

		if err != nil {
			panic(err)
		}
	}

	dsn := filepath + "?parseTime=true"
	db, err := sql.Open("sqlite", dsn)

	if err != nil {
		panic(err)
	}

	return db
}

func TestStoreUserCreate(t *testing.T) {
	db := initDB(":memory:")

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

	err = store.UserCreate(user)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func TestStoreUserFindByEmail(t *testing.T) {
	db := initDB(":memory:")

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

	err = store.UserCreate(user)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	userFound, errFind := store.UserFindByEmail(user.Email())

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
	db := initDB(":memory:")

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

	err = store.UserCreate(user)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	userFound, errFind := store.UserFindByID(user.ID())

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
	db := initDB(":memory:")

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

	err = store.UserCreate(user)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.UserSoftDeleteByID(user.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if user.SoftDeletedAt() != sb.MAX_DATETIME {
		t.Fatal("User MUST NOT be soft deleted")
	}

	userFound, errFind := store.UserFindByID(user.ID())

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

	userFindWithDeleted, err := store.UserList(query)

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
