package userstore

import (
	"context"
	"strings"
	"testing"
)

func TestStoreGroupCount(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	count, err := store.GroupCount(context.Background(), NewGroupQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 0 {
		t.Fatal("unexpected count:", count)
	}

	err = store.GroupCreate(context.Background(), NewGroup().
		SetStatus(GROUP_STATUS_ACTIVE).
		SetHandle("staff").
		SetName("Staff"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	count, err = store.GroupCount(context.Background(), NewGroupQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 1 {
		t.Fatal("unexpected count:", count)
	}
}

func TestStoreGroupCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	group := NewGroup().
		SetStatus(GROUP_STATUS_ACTIVE).
		SetHandle("staff").
		SetName("Staff")

	err = store.GroupCreate(context.Background(), group)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func TestStoreGroupDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	group := NewGroup().
		SetStatus(GROUP_STATUS_ACTIVE).
		SetHandle("staff").
		SetName("Staff")

	err = store.GroupCreate(context.Background(), group)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.GroupDeleteByID(context.Background(), group.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	groupFound, err := store.GroupFindByID(context.Background(), group.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if groupFound != nil {
		t.Fatal("Group MUST be nil")
	}
}

func TestStoreGroupFindByHandle(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	group := NewGroup().
		SetStatus(GROUP_STATUS_ACTIVE).
		SetHandle("staff").
		SetName("Staff")

	err = group.SetMetas(map[string]string{
		"department": "engineering",
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.GroupCreate(context.Background(), group)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	groupFound, errFind := store.GroupFindByHandle(context.Background(), group.GetHandle())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if groupFound == nil {
		t.Fatal("Group MUST NOT be nil")
	}

	if groupFound.GetID() != group.GetID() {
		t.Fatal("IDs do not match")
	}

	if groupFound.GetHandle() != group.GetHandle() {
		t.Fatal("Handles do not match")
	}

	if groupFound.GetName() != group.GetName() {
		t.Fatal("Names do not match")
	}

	if groupFound.GetMeta("department") != "engineering" {
		t.Fatal("Metas do not match")
	}
}

func TestStoreGroupFindByHandleOrCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	group, err := store.GroupFindByHandleOrCreate(context.Background(), "staff", GROUP_STATUS_ACTIVE)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if group == nil {
		t.Fatal("Group MUST NOT be nil")
	}

	if group.GetHandle() != "staff" {
		t.Fatal("Handles do not match")
	}

	groupFound, err := store.GroupFindByHandleOrCreate(context.Background(), "staff", GROUP_STATUS_ACTIVE)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if groupFound.GetID() != group.GetID() {
		t.Fatal("Group MUST be the existing one")
	}
}

func TestStoreGroupList(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	groups := []GroupInterface{
		NewGroup().SetStatus(GROUP_STATUS_ACTIVE).SetHandle("staff").SetName("Staff"),
		NewGroup().SetStatus(GROUP_STATUS_INACTIVE).SetHandle("alumni").SetName("Alumni"),
	}

	for _, group := range groups {
		err = store.GroupCreate(context.Background(), group)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	}

	listActive, err := store.GroupList(context.Background(), NewGroupQuery().SetStatus(GROUP_STATUS_ACTIVE))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(listActive) != 1 {
		t.Fatal("unexpected list length:", len(listActive))
	}
}

func TestStoreGroupSoftDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	group := NewGroup().
		SetStatus(GROUP_STATUS_ACTIVE).
		SetHandle("staff").
		SetName("Staff")

	err = store.GroupCreate(context.Background(), group)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.GroupSoftDeleteByID(context.Background(), group.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	groupFound, err := store.GroupFindByID(context.Background(), group.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if groupFound != nil {
		t.Fatal("Group MUST be nil")
	}

	groupFindWithDeleted, err := store.GroupList(context.Background(), NewGroupQuery().
		SetSoftDeletedIncluded(true).
		SetID(group.GetID()).
		SetLimit(1))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(groupFindWithDeleted) == 0 {
		t.Fatal("Group MUST be soft deleted")
	}

	if strings.Contains(groupFindWithDeleted[0].GetSoftDeletedAt(), MAX_DATETIME) {
		t.Fatal("Group MUST be soft deleted", group.GetSoftDeletedAt())
	}

	if !groupFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("Group MUST be soft deleted")
	}
}

func TestStoreGroupUpdate(t *testing.T) {
	store, err := initStore(":memory:")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	group := NewGroup().
		SetStatus(GROUP_STATUS_ACTIVE).
		SetHandle("staff").
		SetName("Staff")

	err = store.GroupCreate(context.Background(), group)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	group.SetName("All Staff").
		SetStatus(GROUP_STATUS_INACTIVE)

	err = store.GroupUpdate(context.Background(), group)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	groupFound, err := store.GroupFindByID(context.Background(), group.GetID())
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if groupFound == nil {
		t.Fatal("Group MUST NOT be nil")
	}

	if groupFound.GetName() != "All Staff" {
		t.Fatal("Name MUST be 'All Staff', found:", groupFound.GetName())
	}

	if groupFound.GetStatus() != GROUP_STATUS_INACTIVE {
		t.Fatal("Status MUST be inactive, found:", groupFound.GetStatus())
	}
}
