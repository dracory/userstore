package userstore

import (
	"context"
	"strings"
	"testing"
)

func TestStoreUserGroupCount(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	count, err := store.UserGroupCount(context.Background(), NewUserGroupQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 0 {
		t.Fatal("unexpected count:", count)
	}

	err = store.UserGroupCreate(context.Background(), NewUserGroup().
		SetUserID("user_1").
		SetGroupID("group_1"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	count, err = store.UserGroupCount(context.Background(), NewUserGroupQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 1 {
		t.Fatal("unexpected count:", count)
	}
}

func TestStoreUserGroupDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	userGroup := NewUserGroup().
		SetUserID("user_1").
		SetGroupID("group_1")

	err = store.UserGroupCreate(context.Background(), userGroup)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.UserGroupDeleteByID(context.Background(), userGroup.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	userGroupFound, err := store.UserGroupFindByID(context.Background(), userGroup.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if userGroupFound != nil {
		t.Fatal("UserGroup MUST be nil")
	}
}

func TestStoreUserGroupFindByUserIDAndGroupID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	userGroup := NewUserGroup().
		SetUserID("user_1").
		SetGroupID("group_1")

	err = store.UserGroupCreate(context.Background(), userGroup)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	userGroupFound, err := store.UserGroupFindByUserIDAndGroupID(context.Background(), "user_1", "group_1")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if userGroupFound == nil {
		t.Fatal("UserGroup MUST NOT be nil")
	}

	if userGroupFound.GetUserID() != "user_1" {
		t.Fatal("User IDs do not match")
	}

	if userGroupFound.GetGroupID() != "group_1" {
		t.Fatal("Group IDs do not match")
	}
}

func TestStoreUserGroupFindByUserIDAndGroupIDOrCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	userGroup, err := store.UserGroupFindByUserIDAndGroupIDOrCreate(context.Background(), "user_1", "group_1")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if userGroup == nil {
		t.Fatal("UserGroup MUST NOT be nil")
	}

	userGroupFound, err := store.UserGroupFindByUserIDAndGroupIDOrCreate(context.Background(), "user_1", "group_1")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if userGroupFound.GetID() != userGroup.GetID() {
		t.Fatal("UserGroup MUST be the existing one")
	}
}

func TestStoreUserGroupSoftDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	userGroup := NewUserGroup().
		SetUserID("user_1").
		SetGroupID("group_1")

	err = store.UserGroupCreate(context.Background(), userGroup)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.UserGroupSoftDeleteByID(context.Background(), userGroup.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	userGroupFound, err := store.UserGroupFindByID(context.Background(), userGroup.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if userGroupFound != nil {
		t.Fatal("UserGroup MUST be nil")
	}

	userGroupFindWithDeleted, err := store.UserGroupList(context.Background(), NewUserGroupQuery().
		SetSoftDeletedIncluded(true).
		SetID(userGroup.GetID()).
		SetLimit(1))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(userGroupFindWithDeleted) == 0 {
		t.Fatal("UserGroup MUST be soft deleted")
	}

	if strings.Contains(userGroupFindWithDeleted[0].GetSoftDeletedAt(), MAX_DATETIME) {
		t.Fatal("UserGroup MUST be soft deleted", userGroup.GetSoftDeletedAt())
	}

	if !userGroupFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("UserGroup MUST be soft deleted")
	}
}

func TestStoreUserHasGroups(t *testing.T) {
	store, err := initStore(":memory:")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	group1 := NewGroup().SetStatus(GROUP_STATUS_ACTIVE).SetHandle("staff").SetName("Staff")
	group2 := NewGroup().SetStatus(GROUP_STATUS_ACTIVE).SetHandle("alumni").SetName("Alumni")
	group3 := NewGroup().SetStatus(GROUP_STATUS_ACTIVE).SetHandle("vip").SetName("VIP")

	for _, group := range []GroupInterface{group1, group2, group3} {
		err = store.GroupCreate(context.Background(), group)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
	}

	err = store.UserGroupCreate(context.Background(), NewUserGroup().SetUserID("user_1").SetGroupID(group1.GetID()))
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	err = store.UserGroupCreate(context.Background(), NewUserGroup().SetUserID("user_1").SetGroupID(group2.GetID()))
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	hasGroups, err := store.UserHasGroups(context.Background(), "user_1", []string{group1.GetID(), group2.GetID()})
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if !hasGroups {
		t.Fatal("UserHasGroups should return true for assigned groups")
	}

	hasGroups, err = store.UserHasGroups(context.Background(), "user_1", []string{group1.GetID(), group3.GetID()})
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if hasGroups {
		t.Fatal("UserHasGroups should return false when a group is missing")
	}
}

func TestStoreUserGroups(t *testing.T) {
	store, err := initStore(":memory:")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	group1 := NewGroup().SetStatus(GROUP_STATUS_ACTIVE).SetHandle("staff").SetName("Staff")
	group2 := NewGroup().SetStatus(GROUP_STATUS_ACTIVE).SetHandle("alumni").SetName("Alumni")

	err = store.GroupCreate(context.Background(), group1)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	err = store.GroupCreate(context.Background(), group2)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.UserGroupCreate(context.Background(), NewUserGroup().SetUserID("user_1").SetGroupID(group1.GetID()))
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	err = store.UserGroupCreate(context.Background(), NewUserGroup().SetUserID("user_1").SetGroupID(group2.GetID()))
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	err = store.UserGroupCreate(context.Background(), NewUserGroup().SetUserID("user_2").SetGroupID(group1.GetID()))
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	groups, err := store.UserGroups(context.Background(), "user_1")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(groups) != 2 {
		t.Fatalf("Expected 2 groups for user_1, got %d", len(groups))
	}

	groups, err = store.UserGroups(context.Background(), "user_2")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(groups) != 1 {
		t.Fatalf("Expected 1 group for user_2, got %d", len(groups))
	}

	if groups[0].GetHandle() != "staff" {
		t.Fatalf("Expected handle 'staff', got '%s'", groups[0].GetHandle())
	}
}
