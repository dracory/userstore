package userstore

import (
	"context"
	"strings"
	"testing"

	"github.com/dracory/database"
	"github.com/dracory/sb"
)

func TestStoreUserCount(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	count, err := store.UserCount(context.Background(), NewUserQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 0 {
		t.Fatal("unexpected count:", count)
	}

	err = store.UserCreate(context.Background(), NewUser().
		SetStatus(USER_STATUS_UNVERIFIED).
		SetFirstName("John").
		SetMiddleNames("").
		SetLastName("Doe").
		SetEmail("test@test.com").
		SetPassword("").
		SetProfileImageUrl("http://test.com/profile.png"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	count, err = store.UserCount(context.Background(), NewUserQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 1 {
		t.Fatal("unexpected count:", count)
	}

	err = store.UserCreate(context.Background(), NewUser().
		SetStatus(USER_STATUS_UNVERIFIED).
		SetFirstName("John").
		SetMiddleNames("").
		SetLastName("Doe").
		SetEmail("test2@test.com").
		SetPassword("").
		SetProfileImageUrl("http://test.com/profile.png"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	count, err = store.UserCount(context.Background(), NewUserQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 2 {
		t.Fatal("unexpected count:", count)
	}
}

func TestStoreUserCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

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

func TestStoreUserDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

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

	err = store.UserDelete(context.Background(), user)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	userFound, err := store.UserFindByID(context.Background(), user.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if userFound != nil {
		t.Fatal("User MUST be nil")
	}

	userFindWithDeleted, err := store.UserList(context.Background(), NewUserQuery().
		SetID(user.GetID()).
		SetSoftDeletedIncluded(true))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(userFindWithDeleted) != 0 {
		t.Fatal("User MUST be nil")
	}
}

func TestStoreUserDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

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

	err = store.UserDeleteByID(context.Background(), user.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	userFound, err := store.UserFindByID(context.Background(), user.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if userFound != nil {
		t.Fatal("User MUST be nil")
	}

	userFindWithDeleted, err := store.UserList(context.Background(), NewUserQuery().
		SetID(user.GetID()).
		SetSoftDeletedIncluded(true))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(userFindWithDeleted) != 0 {
		t.Fatal("User MUST NOT be found")
	}
}

func TestStoreUserFindByEmail(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

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

	err = store.UserCreate(database.Context(context.Background(), store.DB()), user)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	userFound, errFind := store.UserFindByEmail(database.Context(context.Background(), store.DB()), user.GetEmail())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if userFound == nil {
		t.Fatal("User MUST NOT be nil")
	}

	if userFound.GetID() != user.GetID() {
		t.Fatal("IDs do not match")
	}

	if userFound.GetEmail() != user.GetEmail() {
		t.Fatal("Emails do not match")
	}

	if userFound.GetFirstName() != user.GetFirstName() {
		t.Fatal("First names do not match")
	}

	if userFound.GetMiddleNames() != user.GetMiddleNames() {
		t.Fatal("Middle names do not match")
	}

	if userFound.GetLastName() != user.GetLastName() {
		t.Fatal("Last names do not match")
	}

	if userFound.GetProfileImageUrl() != user.GetProfileImageUrl() {
		t.Fatal("Profile image URLs do not match")
	}

	if userFound.GetStatus() != user.GetStatus() {
		t.Fatal("Statuses do not match")
	}

	if userFound.GetRole() != user.GetRole() {
		t.Fatal("Roles do not match")
	}

	if userFound.GetRole() != USER_ROLE_USER {
		t.Fatal("Roles MUST be:", USER_ROLE_USER, `found:`, userFound.GetRole())
	}

	if userFound.GetMeta("education_1") != user.GetMeta("education_1") {
		t.Fatal("Metas do not match")
	}

	if userFound.GetMeta("education_2") != user.GetMeta("education_2") {
		t.Fatal("Metas do not match")
	}

	if userFound.GetMeta("education_3") != user.GetMeta("education_3") {
		t.Fatal("Metas do not match")
	}
}

func TestStoreUserFindByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

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

	ctx := database.Context(context.Background(), store.DB())
	err = store.UserCreate(ctx, user)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	userFound, errFind := store.UserFindByID(ctx, user.GetID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if userFound == nil {
		t.Fatal("User MUST NOT be nil")
	}

	if userFound.GetID() != user.GetID() {
		t.Fatal("IDs do not match")
	}

	if userFound.GetEmail() != user.GetEmail() {
		t.Fatal("Emails do not match")
	}

	if userFound.GetFirstName() != user.GetFirstName() {
		t.Fatal("First names do not match")
	}

	if userFound.GetMiddleNames() != user.GetMiddleNames() {
		t.Fatal("Middle names do not match")
	}

	if userFound.GetLastName() != user.GetLastName() {
		t.Fatal("Last names do not match")
	}

	if userFound.GetProfileImageUrl() != user.GetProfileImageUrl() {
		t.Fatal("Profile image URLs do not match")
	}

	if userFound.GetStatus() != user.GetStatus() {
		t.Fatal("Statuses do not match")
	}

	if userFound.GetRole() != user.GetRole() {
		t.Fatal("Roles do not match")
	}

	if userFound.GetRole() != USER_ROLE_USER {
		t.Fatal("Roles MUST be:", USER_ROLE_USER, `found:`, userFound.GetRole())
	}

	if userFound.GetMeta("education_1") != user.GetMeta("education_1") {
		t.Fatal("Metas do not match")
	}

	if userFound.GetMeta("education_2") != user.GetMeta("education_2") {
		t.Fatal("Metas do not match")
	}

	if userFound.GetMeta("education_3") != user.GetMeta("education_3") {
		t.Fatal("Metas do not match")
	}
}

func TestStoreUserList(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	users := []UserInterface{
		NewUser().
			SetStatus(USER_STATUS_UNVERIFIED).
			SetFirstName("John").
			SetMiddleNames("").
			SetLastName("Doe").
			SetEmail("test@test.com").
			SetPassword("").
			SetProfileImageUrl("http://test.com/profile.png"),
		NewUser().
			SetStatus(USER_STATUS_ACTIVE).
			SetFirstName("Jane").
			SetMiddleNames("").
			SetLastName("Doe").
			SetEmail("test2@test.com").
			SetPassword("").
			SetProfileImageUrl("http://test.com/profile.png"),
	}

	for _, user := range users {
		err = store.UserCreate(context.Background(), user)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	}

	listActive, err := store.UserList(context.Background(), NewUserQuery().SetStatus(USER_STATUS_ACTIVE))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(listActive) != 1 {
		t.Fatal("unexpected list length:", len(listActive))
	}

	listEmail, err := store.UserList(context.Background(), NewUserQuery().SetEmail("test2@test.com"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(listEmail) != 1 {
		t.Fatal("unexpected list length:", len(listEmail))
	}
}

func TestStoreUserSoftDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

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

	err = store.UserSoftDelete(context.Background(), user)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if user.GetSoftDeletedAt() == sb.MAX_DATETIME {
		t.Fatal("User MUST be soft deleted")
	}

	userFound, errFind := store.UserFindByID(context.Background(), user.GetID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if userFound != nil {
		t.Fatal("User MUST be soft deleted, so MUST be nil")
	}

	userFindWithDeleted, err := store.UserList(context.Background(), NewUserQuery().
		SetSoftDeletedIncluded(true).
		SetID(user.GetID()).
		SetLimit(1))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(userFindWithDeleted) == 0 {
		t.Fatal("User MUST be soft deleted")
	}

	if strings.Contains(userFindWithDeleted[0].GetSoftDeletedAt(), sb.MAX_DATETIME) {
		t.Fatal("User MUST be soft deleted", user.GetSoftDeletedAt())
	}

	if !userFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("User MUST be soft deleted")
	}
}

func TestStoreUserSoftDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

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

	err = store.UserSoftDeleteByID(context.Background(), user.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if user.GetSoftDeletedAt() != sb.MAX_DATETIME {
		t.Fatal("User MUST NOT be soft deleted, as it was soft deleted by ID")
	}

	userFound, errFind := store.UserFindByID(context.Background(), user.GetID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if userFound != nil {
		t.Fatal("User MUST be nil")
	}
	query := NewUserQuery().
		SetSoftDeletedIncluded(true).
		SetID(user.GetID()).
		SetLimit(1)

	userFindWithDeleted, err := store.UserList(context.Background(), query)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(userFindWithDeleted) == 0 {
		t.Fatal("User MUST be soft deleted")
	}

	if strings.Contains(userFindWithDeleted[0].GetSoftDeletedAt(), sb.MAX_DATETIME) {
		t.Fatal("User MUST be soft deleted", user.GetSoftDeletedAt())
	}

	if !userFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("User MUST be soft deleted")
	}
}

func TestStoreUserMetaLike(t *testing.T) {
	store, err := initStore(":memory:")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	// Create a user with metadata containing the search term
	user := NewUser()

	err = user.SetMetas(map[string]string{"key": "searchTermValue"})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.UserCreate(context.Background(), user)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	user2 := NewUser()

	err = user2.SetMetas(map[string]string{"key": "searchTermValue2"})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.UserCreate(context.Background(), user2)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// Query for users matching the search term
	users, err := store.UserList(context.Background(), NewUserQuery().SetMetaLike(`"key":"searchTermValue"`))
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// Assertions
	if len(users) != 1 {
		t.Errorf("Expected one user to be found, but got %d", len(users))
	}
	if users[0].GetID() != user.GetID() {
		t.Errorf("Incorrect user returned, expected ID %s, but got %s", user.GetID(), users[0].GetID())
	}
}
