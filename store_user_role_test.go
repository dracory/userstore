package userstore

import (
	"context"
	"strings"
	"testing"
)

func TestStoreUserRoleCount(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	count, err := store.UserRoleCount(context.Background(), NewUserRoleQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 0 {
		t.Fatal("unexpected count:", count)
	}

	err = store.UserRoleCreate(context.Background(), NewUserRole().
		SetUserID("user_1").
		SetRoleID("role_1"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	count, err = store.UserRoleCount(context.Background(), NewUserRoleQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 1 {
		t.Fatal("unexpected count:", count)
	}
}

func TestStoreUserRoleCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	userRole := NewUserRole().
		SetUserID("user_1").
		SetRoleID("role_1")

	err = store.UserRoleCreate(context.Background(), userRole)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func TestStoreUserRoleDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	userRole := NewUserRole().
		SetUserID("user_1").
		SetRoleID("role_1")

	err = store.UserRoleCreate(context.Background(), userRole)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.UserRoleDeleteByID(context.Background(), userRole.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	userRoleFound, err := store.UserRoleFindByID(context.Background(), userRole.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if userRoleFound != nil {
		t.Fatal("UserRole MUST be nil")
	}
}

func TestStoreUserRoleFindByUserIDAndRoleID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	userRole := NewUserRole().
		SetUserID("user_1").
		SetRoleID("role_1")

	err = store.UserRoleCreate(context.Background(), userRole)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	userRoleFound, err := store.UserRoleFindByUserIDAndRoleID(context.Background(), "user_1", "role_1")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if userRoleFound == nil {
		t.Fatal("UserRole MUST NOT be nil")
	}

	if userRoleFound.GetID() != userRole.GetID() {
		t.Fatal("IDs do not match")
	}

	if userRoleFound.GetUserID() != "user_1" {
		t.Fatal("User IDs do not match")
	}

	if userRoleFound.GetRoleID() != "role_1" {
		t.Fatal("Role IDs do not match")
	}
}

func TestStoreUserRoleFindByUserIDAndRoleIDOrCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	userRole, err := store.UserRoleFindByUserIDAndRoleIDOrCreate(context.Background(), "user_1", "role_1")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if userRole == nil {
		t.Fatal("UserRole MUST NOT be nil")
	}

	userRoleFound, err := store.UserRoleFindByUserIDAndRoleIDOrCreate(context.Background(), "user_1", "role_1")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if userRoleFound.GetID() != userRole.GetID() {
		t.Fatal("UserRole MUST be the existing one")
	}
}

func TestStoreUserRoleList(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	userRoles := []UserRoleInterface{
		NewUserRole().SetUserID("user_1").SetRoleID("role_1"),
		NewUserRole().SetUserID("user_1").SetRoleID("role_2"),
		NewUserRole().SetUserID("user_2").SetRoleID("role_1"),
	}

	for _, userRole := range userRoles {
		err = store.UserRoleCreate(context.Background(), userRole)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	}

	listUser, err := store.UserRoleList(context.Background(), NewUserRoleQuery().SetUserID("user_1"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(listUser) != 2 {
		t.Fatal("unexpected list length:", len(listUser))
	}

	listRole, err := store.UserRoleList(context.Background(), NewUserRoleQuery().SetRoleID("role_1"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(listRole) != 2 {
		t.Fatal("unexpected list length:", len(listRole))
	}
}

func TestStoreUserRoleSoftDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	userRole := NewUserRole().
		SetUserID("user_1").
		SetRoleID("role_1")

	err = store.UserRoleCreate(context.Background(), userRole)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.UserRoleSoftDeleteByID(context.Background(), userRole.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	userRoleFound, err := store.UserRoleFindByID(context.Background(), userRole.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if userRoleFound != nil {
		t.Fatal("UserRole MUST be nil")
	}

	userRoleFindWithDeleted, err := store.UserRoleList(context.Background(), NewUserRoleQuery().
		SetSoftDeletedIncluded(true).
		SetID(userRole.GetID()).
		SetLimit(1))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(userRoleFindWithDeleted) == 0 {
		t.Fatal("UserRole MUST be soft deleted")
	}

	if strings.Contains(userRoleFindWithDeleted[0].GetSoftDeletedAt(), MAX_DATETIME) {
		t.Fatal("UserRole MUST be soft deleted", userRole.GetSoftDeletedAt())
	}

	if !userRoleFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("UserRole MUST be soft deleted")
	}
}

func TestStoreUserRoleUpdate(t *testing.T) {
	store, err := initStore(":memory:")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	userRole := NewUserRole().
		SetUserID("user_1").
		SetRoleID("role_1")

	err = store.UserRoleCreate(context.Background(), userRole)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	userRole.SetRoleID("role_2")

	err = store.UserRoleUpdate(context.Background(), userRole)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	userRoleFound, err := store.UserRoleFindByID(context.Background(), userRole.GetID())
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if userRoleFound == nil {
		t.Fatal("UserRole MUST NOT be nil")
	}

	if userRoleFound.GetRoleID() != "role_2" {
		t.Fatal("Role ID MUST be 'role_2', found:", userRoleFound.GetRoleID())
	}
}
