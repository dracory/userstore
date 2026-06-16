package userstore

import (
	"testing"
)

func TestNewRole(t *testing.T) {
	role := NewRole()

	if role == nil {
		t.Fatal("NewRole should not return nil")
	}

	if role.GetID() == "" {
		t.Error("ID should not be empty")
	}

	if role.GetStatus() != USER_STATUS_UNVERIFIED {
		t.Errorf("Expected status %s, got %s", USER_STATUS_UNVERIFIED, role.GetStatus())
	}

	if role.GetName() != "" {
		t.Error("Name should be empty")
	}

	if role.GetHandle() != "" {
		t.Error("Handle should be empty")
	}

	if role.GetMemo() != "" {
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

	if role.GetID() != "test123" {
		t.Errorf("Expected ID test123, got %s", role.GetID())
	}

	if role.GetName() != "Administrator" {
		t.Errorf("Expected Name Administrator, got %s", role.GetName())
	}

	if role.GetHandle() != "admin" {
		t.Errorf("Expected Handle admin, got %s", role.GetHandle())
	}

	if role.GetStatus() != USER_STATUS_ACTIVE {
		t.Errorf("Expected status %s, got %s", USER_STATUS_ACTIVE, role.GetStatus())
	}

	if role.GetMemo() != "Test memo" {
		t.Errorf("Expected Memo Test memo, got %s", role.GetMemo())
	}
}

func TestRoleSettersAndGetters(t *testing.T) {
	role := NewRole()

	role.SetID("test123")
	if role.GetID() != "test123" {
		t.Errorf("Expected ID test123, got %s", role.GetID())
	}

	role.SetName("Administrator")
	if role.GetName() != "Administrator" {
		t.Errorf("Expected Name Administrator, got %s", role.GetName())
	}

	role.SetHandle("admin")
	if role.GetHandle() != "admin" {
		t.Errorf("Expected Handle admin, got %s", role.GetHandle())
	}

	role.SetMemo("Test memo")
	if role.GetMemo() != "Test memo" {
		t.Errorf("Expected Memo Test memo, got %s", role.GetMemo())
	}

	role.SetStatus(USER_STATUS_ACTIVE)
	if role.GetStatus() != USER_STATUS_ACTIVE {
		t.Errorf("Expected status %s, got %s", USER_STATUS_ACTIVE, role.GetStatus())
	}
}

func TestRoleMetas(t *testing.T) {
	role := NewRole()

	// Test setting metas
	err := role.SetMetas(map[string]string{"key1": "value1", "key2": "value2"})
	if err != nil {
		t.Errorf("SetMetas should not return error, got %v", err)
	}

	metas, err := role.GetMetas()
	if err != nil {
		t.Errorf("GetMetas should not return error, got %v", err)
	}

	if metas["key1"] != "value1" {
		t.Errorf("Expected key1 to be value1, got %s", metas["key1"])
	}

	if metas["key2"] != "value2" {
		t.Errorf("Expected key2 to be value2, got %s", metas["key2"])
	}

	// Test getting a single meta
	if role.GetMeta("key1") != "value1" {
		t.Errorf("Expected GetMeta key1 to return value1, got %s", role.GetMeta("key1"))
	}

	// Test non-existent meta
	if role.GetMeta("nonexistent") != "" {
		t.Error("GetMeta should return empty string for non-existent key")
	}

	// Test SetMeta
	err = role.SetMeta("key4", "value4")
	if err != nil {
		t.Errorf("SetMeta should not return error, got %v", err)
	}

	if role.GetMeta("key4") != "value4" {
		t.Errorf("Expected GetMeta key4 to return value4, got %s", role.GetMeta("key4"))
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

	if role.GetName() != "Administrator" {
		t.Error("Method chaining should work")
	}

	if role.GetHandle() != "admin" {
		t.Error("Method chaining should work")
	}

	if role.GetStatus() != USER_STATUS_ACTIVE {
		t.Error("Method chaining should work")
	}

	if role.GetMemo() != "Test memo" {
		t.Error("Method chaining should work")
	}
}
