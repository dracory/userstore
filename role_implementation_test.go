package userstore

import (
	"testing"
)

func TestNewRole(t *testing.T) {
	role := NewRole()

	if role == nil {
		t.Fatal("NewRole should not return nil")
	}

	if role.ID() == "" {
		t.Error("ID should not be empty")
	}

	if role.Status() != USER_STATUS_UNVERIFIED {
		t.Errorf("Expected status %s, got %s", USER_STATUS_UNVERIFIED, role.Status())
	}

	if role.Name() != "" {
		t.Error("Name should be empty")
	}

	if role.Handle() != "" {
		t.Error("Handle should be empty")
	}

	if role.Memo() != "" {
		t.Error("Memo should be empty")
	}
}

func TestNewRoleFromExistingData(t *testing.T) {
	data := map[string]string{
		COLUMN_ID:     "test123",
		COLUMN_NAME:   "Administrator",
		COLUMN_HANDLE: "admin",
		COLUMN_STATUS: USER_STATUS_ACTIVE,
		COLUMN_MEMO:   "Test memo",
	}

	role := NewRoleFromExistingData(data)

	if role == nil {
		t.Fatal("NewRoleFromExistingData should not return nil")
	}

	if role.ID() != "test123" {
		t.Errorf("Expected ID test123, got %s", role.ID())
	}

	if role.Name() != "Administrator" {
		t.Errorf("Expected Name Administrator, got %s", role.Name())
	}

	if role.Handle() != "admin" {
		t.Errorf("Expected Handle admin, got %s", role.Handle())
	}

	if role.Status() != USER_STATUS_ACTIVE {
		t.Errorf("Expected status %s, got %s", USER_STATUS_ACTIVE, role.Status())
	}

	if role.Memo() != "Test memo" {
		t.Errorf("Expected Memo Test memo, got %s", role.Memo())
	}
}

func TestRoleSettersAndGetters(t *testing.T) {
	role := NewRole()

	role.SetID("test123")
	if role.ID() != "test123" {
		t.Errorf("Expected ID test123, got %s", role.ID())
	}

	role.SetName("Administrator")
	if role.Name() != "Administrator" {
		t.Errorf("Expected Name Administrator, got %s", role.Name())
	}

	role.SetHandle("admin")
	if role.Handle() != "admin" {
		t.Errorf("Expected Handle admin, got %s", role.Handle())
	}

	role.SetMemo("Test memo")
	if role.Memo() != "Test memo" {
		t.Errorf("Expected Memo Test memo, got %s", role.Memo())
	}

	role.SetStatus(USER_STATUS_ACTIVE)
	if role.Status() != USER_STATUS_ACTIVE {
		t.Errorf("Expected status %s, got %s", USER_STATUS_ACTIVE, role.Status())
	}
}

func TestRoleMetas(t *testing.T) {
	role := NewRole()

	// Test setting metas
	err := role.SetMetas(map[string]string{"key1": "value1", "key2": "value2"})
	if err != nil {
		t.Errorf("SetMetas should not return error, got %v", err)
	}

	metas, err := role.Metas()
	if err != nil {
		t.Errorf("Metas should not return error, got %v", err)
	}

	if metas["key1"] != "value1" {
		t.Errorf("Expected key1 to be value1, got %s", metas["key1"])
	}

	if metas["key2"] != "value2" {
		t.Errorf("Expected key2 to be value2, got %s", metas["key2"])
	}

	// Test getting a single meta
	if role.Meta("key1") != "value1" {
		t.Errorf("Expected Meta key1 to return value1, got %s", role.Meta("key1"))
	}

	// Test non-existent meta
	if role.Meta("nonexistent") != "" {
		t.Error("Meta should return empty string for non-existent key")
	}

	// Test SetMeta
	err = role.SetMeta("key4", "value4")
	if err != nil {
		t.Errorf("SetMeta should not return error, got %v", err)
	}

	if role.Meta("key4") != "value4" {
		t.Errorf("Expected Meta key4 to return value4, got %s", role.Meta("key4"))
	}
}

func TestRoleNoImageUrl(t *testing.T) {
	url := RoleNoImageUrl()
	if url != "/role/default.png" {
		t.Errorf("Expected /role/default.png, got %s", url)
	}
}

func TestRoleChaining(t *testing.T) {
	role := NewRole().
		SetName("Administrator").
		SetHandle("admin").
		SetStatus(USER_STATUS_ACTIVE).
		SetMemo("Test memo")

	if role.Name() != "Administrator" {
		t.Error("Method chaining should work")
	}

	if role.Handle() != "admin" {
		t.Error("Method chaining should work")
	}

	if role.Status() != USER_STATUS_ACTIVE {
		t.Error("Method chaining should work")
	}

	if role.Memo() != "Test memo" {
		t.Error("Method chaining should work")
	}
}
