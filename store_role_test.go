package userstore

import (
	"context"
	"strings"
	"testing"
)

func TestStoreRoleCount(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	count, err := store.RoleCount(context.Background(), NewRoleQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 0 {
		t.Fatal("unexpected count:", count)
	}

	err = store.RoleCreate(context.Background(), NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("administrator").
		SetName("Administrator"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	count, err = store.RoleCount(context.Background(), NewRoleQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 1 {
		t.Fatal("unexpected count:", count)
	}

	err = store.RoleCreate(context.Background(), NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("manager").
		SetName("Manager"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	count, err = store.RoleCount(context.Background(), NewRoleQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 2 {
		t.Fatal("unexpected count:", count)
	}
}

func TestStoreRoleCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	role := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("administrator").
		SetName("Administrator")

	err = store.RoleCreate(context.Background(), role)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func TestStoreRoleDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	role := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("administrator").
		SetName("Administrator")

	err = store.RoleCreate(context.Background(), role)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.RoleDelete(context.Background(), role)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	roleFound, err := store.RoleFindByID(context.Background(), role.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if roleFound != nil {
		t.Fatal("Role MUST be nil")
	}

	roleFindWithDeleted, err := store.RoleList(context.Background(), NewRoleQuery().
		SetID(role.GetID()).
		SetSoftDeletedIncluded(true))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(roleFindWithDeleted) != 0 {
		t.Fatal("Role MUST be nil")
	}
}

func TestStoreRoleDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	role := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("administrator").
		SetName("Administrator")

	err = store.RoleCreate(context.Background(), role)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.RoleDeleteByID(context.Background(), role.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	roleFound, err := store.RoleFindByID(context.Background(), role.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if roleFound != nil {
		t.Fatal("Role MUST be nil")
	}

	roleFindWithDeleted, err := store.RoleList(context.Background(), NewRoleQuery().
		SetID(role.GetID()).
		SetSoftDeletedIncluded(true))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(roleFindWithDeleted) != 0 {
		t.Fatal("Role MUST NOT be found")
	}
}

func TestStoreRoleFindByHandle(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	role := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("administrator").
		SetName("Administrator")

	err = role.SetMetas(map[string]string{
		"permission_1": "Permission 1",
		"permission_2": "Permission 2",
		"permission_3": "Permission 3",
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.RoleCreate(context.Background(), role)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	roleFound, errFind := store.RoleFindByHandle(context.Background(), role.GetHandle())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if roleFound == nil {
		t.Fatal("Role MUST NOT be nil")
	}

	if roleFound.GetID() != role.GetID() {
		t.Fatal("IDs do not match")
	}

	if roleFound.GetHandle() != role.GetHandle() {
		t.Fatal("Handles do not match")
	}

	if roleFound.GetName() != role.GetName() {
		t.Fatal("Names do not match")
	}

	if roleFound.GetStatus() != role.GetStatus() {
		t.Fatal("Statuses do not match")
	}

	if roleFound.GetMeta("permission_1") != role.GetMeta("permission_1") {
		t.Fatal("Metas do not match")
	}

	if roleFound.GetMeta("permission_2") != role.GetMeta("permission_2") {
		t.Fatal("Metas do not match")
	}

	if roleFound.GetMeta("permission_3") != role.GetMeta("permission_3") {
		t.Fatal("Metas do not match")
	}
}

func TestStoreRoleFindByHandleOrCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	role, err := store.RoleFindByHandleOrCreate(context.Background(), "administrator", ROLE_STATUS_ACTIVE)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if role == nil {
		t.Fatal("Role MUST NOT be nil")
	}

	if role.GetHandle() != "administrator" {
		t.Fatal("Handles do not match")
	}

	if role.GetStatus() != ROLE_STATUS_ACTIVE {
		t.Fatal("Statuses do not match")
	}

	roleFound, err := store.RoleFindByHandleOrCreate(context.Background(), "administrator", ROLE_STATUS_ACTIVE)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if roleFound.GetID() != role.GetID() {
		t.Fatal("Role MUST be the existing one")
	}
}

func TestStoreRoleFindByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	role := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("administrator").
		SetName("Administrator")

	err = role.SetMetas(map[string]string{
		"permission_1": "Permission 1",
		"permission_2": "Permission 2",
		"permission_3": "Permission 3",
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	ctx := context.Background()
	err = store.RoleCreate(ctx, role)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	roleFound, errFind := store.RoleFindByID(ctx, role.GetID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if roleFound == nil {
		t.Fatal("Role MUST NOT be nil")
	}

	if roleFound.GetID() != role.GetID() {
		t.Fatal("IDs do not match")
	}

	if roleFound.GetHandle() != role.GetHandle() {
		t.Fatal("Handles do not match")
	}

	if roleFound.GetName() != role.GetName() {
		t.Fatal("Names do not match")
	}

	if roleFound.GetStatus() != role.GetStatus() {
		t.Fatal("Statuses do not match")
	}

	if roleFound.GetMeta("permission_1") != role.GetMeta("permission_1") {
		t.Fatal("Metas do not match")
	}

	if roleFound.GetMeta("permission_2") != role.GetMeta("permission_2") {
		t.Fatal("Metas do not match")
	}

	if roleFound.GetMeta("permission_3") != role.GetMeta("permission_3") {
		t.Fatal("Metas do not match")
	}
}

func TestStoreRoleList(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	roles := []RoleInterface{
		NewRole().
			SetStatus(ROLE_STATUS_ACTIVE).
			SetHandle("administrator").
			SetName("Administrator"),
		NewRole().
			SetStatus(ROLE_STATUS_INACTIVE).
			SetHandle("manager").
			SetName("Manager"),
	}

	for _, role := range roles {
		err = store.RoleCreate(context.Background(), role)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	}

	listActive, err := store.RoleList(context.Background(), NewRoleQuery().SetStatus(ROLE_STATUS_ACTIVE))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(listActive) != 1 {
		t.Fatal("unexpected list length:", len(listActive))
	}

	listHandle, err := store.RoleList(context.Background(), NewRoleQuery().SetHandle("manager"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(listHandle) != 1 {
		t.Fatal("unexpected list length:", len(listHandle))
	}
}

func TestStoreRoleSoftDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	role := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("administrator").
		SetName("Administrator")

	err = store.RoleCreate(context.Background(), role)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.RoleSoftDelete(context.Background(), role)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if role.GetSoftDeletedAt() == MAX_DATETIME {
		t.Fatal("Role MUST be soft deleted")
	}

	roleFound, errFind := store.RoleFindByID(context.Background(), role.GetID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if roleFound != nil {
		t.Fatal("Role MUST be soft deleted, so MUST be nil")
	}

	roleFindWithDeleted, err := store.RoleList(context.Background(), NewRoleQuery().
		SetSoftDeletedIncluded(true).
		SetID(role.GetID()).
		SetLimit(1))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(roleFindWithDeleted) == 0 {
		t.Fatal("Role MUST be soft deleted")
	}

	if strings.Contains(roleFindWithDeleted[0].GetSoftDeletedAt(), MAX_DATETIME) {
		t.Fatal("Role MUST be soft deleted", role.GetSoftDeletedAt())
	}

	if !roleFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("Role MUST be soft deleted")
	}
}

func TestStoreRoleSoftDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	role := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("administrator").
		SetName("Administrator")

	err = store.RoleCreate(context.Background(), role)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.RoleSoftDeleteByID(context.Background(), role.GetID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if role.GetSoftDeletedAt() != MAX_DATETIME {
		t.Fatal("Role MUST NOT be soft deleted, as it was soft deleted by ID")
	}

	roleFound, errFind := store.RoleFindByID(context.Background(), role.GetID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if roleFound != nil {
		t.Fatal("Role MUST be nil")
	}

	query := NewRoleQuery().
		SetSoftDeletedIncluded(true).
		SetID(role.GetID()).
		SetLimit(1)

	roleFindWithDeleted, err := store.RoleList(context.Background(), query)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(roleFindWithDeleted) == 0 {
		t.Fatal("Role MUST be soft deleted")
	}

	if strings.Contains(roleFindWithDeleted[0].GetSoftDeletedAt(), MAX_DATETIME) {
		t.Fatal("Role MUST be soft deleted", role.GetSoftDeletedAt())
	}

	if !roleFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("Role MUST be soft deleted")
	}
}

func TestStoreRoleQueryTitleLike(t *testing.T) {
	store, err := initStore(":memory:")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	roles := []RoleInterface{
		NewRole().
			SetStatus(ROLE_STATUS_ACTIVE).
			SetHandle("administrator").
			SetName("Administrator"),
		NewRole().
			SetStatus(ROLE_STATUS_ACTIVE).
			SetHandle("admin_assistant").
			SetName("Admin Assistant"),
		NewRole().
			SetStatus(ROLE_STATUS_ACTIVE).
			SetHandle("manager").
			SetName("Manager"),
	}

	for _, role := range roles {
		err = store.RoleCreate(context.Background(), role)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
	}

	// Test partial match on name (search "admin" should find "Administrator" and "Admin Assistant")
	list, err := store.RoleList(context.Background(), NewRoleQuery().SetTitleLike("admin"))
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(list) != 2 {
		t.Fatalf("Expected 2 roles with name containing 'admin', got %d", len(list))
	}

	for _, role := range list {
		if !strings.Contains(strings.ToLower(role.GetName()), "admin") {
			t.Fatalf("Expected name containing 'admin', got '%s'", role.GetName())
		}
	}
}

func TestStoreRoleUpdate(t *testing.T) {
	store, err := initStore(":memory:")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	defer func() {
		if err := store.GetDB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	role := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("administrator").
		SetName("Administrator")

	err = store.RoleCreate(context.Background(), role)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	role.SetName("Super Administrator").
		SetStatus(ROLE_STATUS_INACTIVE)

	err = store.RoleUpdate(context.Background(), role)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	roleFound, err := store.RoleFindByID(context.Background(), role.GetID())
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if roleFound == nil {
		t.Fatal("Role MUST NOT be nil")
	}

	if roleFound.GetName() != "Super Administrator" {
		t.Fatal("Name MUST be 'Super Administrator', found:", roleFound.GetName())
	}

	if roleFound.GetStatus() != ROLE_STATUS_INACTIVE {
		t.Fatal("Status MUST be inactive, found:", roleFound.GetStatus())
	}
}
