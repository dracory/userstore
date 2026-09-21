package userstore

import (
	"testing"
)

func TestUserRoleQueryID(t *testing.T) {
	query := NewUserRoleQuery()

	// Test HasID returns false initially
	if query.HasID() {
		t.Fatal("HasID should return false initially")
	}

	// Test GetID returns empty string initially
	if query.GetID() != "" {
		t.Fatal("GetID should return empty string initially")
	}

	// Test SetID
	query = query.SetID("123")
	if !query.HasID() {
		t.Fatal("HasID should return true after SetID")
	}
	if query.GetID() != "123" {
		t.Fatalf("GetID should return '123', got '%s'", query.GetID())
	}

	// Test validation with valid value
	err := query.Validate()
	if err != nil {
		t.Fatal("unexpected validation error:", err)
	}

	// Test validation with empty value
	query = query.SetID("")
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for empty id")
	}
	if err.Error() != "user role query. id cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestUserRoleQueryIDIn(t *testing.T) {
	query := NewUserRoleQuery()

	// Test HasIDIn returns false initially
	if query.HasIDIn() {
		t.Fatal("HasIDIn should return false initially")
	}

	// Test IDIn returns empty slice initially
	if len(query.IDIn()) != 0 {
		t.Fatal("IDIn should return empty slice initially")
	}

	// Test SetIDIn
	query = query.SetIDIn([]string{"1", "2"})
	if !query.HasIDIn() {
		t.Fatal("HasIDIn should return true after SetIDIn")
	}
	if len(query.IDIn()) != 2 {
		t.Fatalf("IDIn should return 2 items, got %d", len(query.IDIn()))
	}

	// Test validation with valid value
	err := query.Validate()
	if err != nil {
		t.Fatal("unexpected validation error:", err)
	}

	// Test validation with empty value
	query = query.SetIDIn([]string{})
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for empty id_in")
	}
	if err.Error() != "user role query. id_in cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestUserRoleQueryUserID(t *testing.T) {
	query := NewUserRoleQuery()

	// Test HasUserID returns false initially
	if query.HasUserID() {
		t.Fatal("HasUserID should return false initially")
	}

	// Test UserID returns empty string initially
	if query.UserID() != "" {
		t.Fatal("UserID should return empty string initially")
	}

	// Test SetUserID
	query = query.SetUserID("user_1")
	if !query.HasUserID() {
		t.Fatal("HasUserID should return true after SetUserID")
	}
	if query.UserID() != "user_1" {
		t.Fatalf("UserID should return 'user_1', got '%s'", query.UserID())
	}

	// Test validation with valid value
	err := query.Validate()
	if err != nil {
		t.Fatal("unexpected validation error:", err)
	}

	// Test validation with empty value
	query = query.SetUserID("")
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for empty user_id")
	}
	if err.Error() != "user role query. user_id cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestUserRoleQueryUserIDIn(t *testing.T) {
	query := NewUserRoleQuery()

	// Test HasUserIDIn returns false initially
	if query.HasUserIDIn() {
		t.Fatal("HasUserIDIn should return false initially")
	}

	// Test UserIDIn returns empty slice initially
	if len(query.UserIDIn()) != 0 {
		t.Fatal("UserIDIn should return empty slice initially")
	}

	// Test SetUserIDIn
	query = query.SetUserIDIn([]string{"user_1", "user_2"})
	if !query.HasUserIDIn() {
		t.Fatal("HasUserIDIn should return true after SetUserIDIn")
	}
	if len(query.UserIDIn()) != 2 {
		t.Fatalf("UserIDIn should return 2 items, got %d", len(query.UserIDIn()))
	}

	// Test validation with valid value
	err := query.Validate()
	if err != nil {
		t.Fatal("unexpected validation error:", err)
	}

	// Test validation with empty value
	query = query.SetUserIDIn([]string{})
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for empty user_id_in")
	}
	if err.Error() != "user role query. user_id_in cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestUserRoleQueryRoleID(t *testing.T) {
	query := NewUserRoleQuery()

	// Test HasRoleID returns false initially
	if query.HasRoleID() {
		t.Fatal("HasRoleID should return false initially")
	}

	// Test RoleID returns empty string initially
	if query.RoleID() != "" {
		t.Fatal("RoleID should return empty string initially")
	}

	// Test SetRoleID
	query = query.SetRoleID("role_1")
	if !query.HasRoleID() {
		t.Fatal("HasRoleID should return true after SetRoleID")
	}
	if query.RoleID() != "role_1" {
		t.Fatalf("RoleID should return 'role_1', got '%s'", query.RoleID())
	}

	// Test validation with valid value
	err := query.Validate()
	if err != nil {
		t.Fatal("unexpected validation error:", err)
	}

	// Test validation with empty value
	query = query.SetRoleID("")
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for empty role_id")
	}
	if err.Error() != "user role query. role_id cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestUserRoleQueryRoleIDIn(t *testing.T) {
	query := NewUserRoleQuery()

	// Test HasRoleIDIn returns false initially
	if query.HasRoleIDIn() {
		t.Fatal("HasRoleIDIn should return false initially")
	}

	// Test RoleIDIn returns empty slice initially
	if len(query.RoleIDIn()) != 0 {
		t.Fatal("RoleIDIn should return empty slice initially")
	}

	// Test SetRoleIDIn
	query = query.SetRoleIDIn([]string{"role_1", "role_2"})
	if !query.HasRoleIDIn() {
		t.Fatal("HasRoleIDIn should return true after SetRoleIDIn")
	}
	if len(query.RoleIDIn()) != 2 {
		t.Fatalf("RoleIDIn should return 2 items, got %d", len(query.RoleIDIn()))
	}

	// Test validation with valid value
	err := query.Validate()
	if err != nil {
		t.Fatal("unexpected validation error:", err)
	}

	// Test validation with empty value
	query = query.SetRoleIDIn([]string{})
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for empty role_id_in")
	}
	if err.Error() != "user role query. role_id_in cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestUserRoleQueryLimit(t *testing.T) {
	query := NewUserRoleQuery()

	// Test HasLimit returns false initially
	if query.HasLimit() {
		t.Fatal("HasLimit should return false initially")
	}

	// Test Limit returns 0 initially
	if query.Limit() != 0 {
		t.Fatal("Limit should return 0 initially")
	}

	// Test SetLimit
	query = query.SetLimit(10)
	if !query.HasLimit() {
		t.Fatal("HasLimit should return true after SetLimit")
	}
	if query.Limit() != 10 {
		t.Fatalf("Limit should return 10, got %d", query.Limit())
	}

	// Test validation with valid value
	err := query.Validate()
	if err != nil {
		t.Fatal("unexpected validation error:", err)
	}

	// Test validation with invalid value (0)
	query = query.SetLimit(0)
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for limit <= 0")
	}
	if err.Error() != "user role query. limit must be greater than 0" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}

	// Test validation with invalid value (negative)
	query = query.SetLimit(-5)
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for negative limit")
	}
}

func TestUserRoleQueryOffset(t *testing.T) {
	query := NewUserRoleQuery()

	// Test HasOffset returns false initially
	if query.HasOffset() {
		t.Fatal("HasOffset should return false initially")
	}

	// Test Offset returns 0 initially
	if query.Offset() != 0 {
		t.Fatal("Offset should return 0 initially")
	}

	// Test SetOffset
	query = query.SetOffset(5)
	if !query.HasOffset() {
		t.Fatal("HasOffset should return true after SetOffset")
	}
	if query.Offset() != 5 {
		t.Fatalf("Offset should return 5, got %d", query.Offset())
	}

	// Test validation with valid value
	err := query.Validate()
	if err != nil {
		t.Fatal("unexpected validation error:", err)
	}

	// Test validation with invalid value (negative)
	query = query.SetOffset(-1)
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for negative offset")
	}
	if err.Error() != "user role query. offset must be greater than or equal to 0" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestUserRoleQueryOrderBy(t *testing.T) {
	query := NewUserRoleQuery()

	// Test HasOrderBy returns false initially
	if query.HasOrderBy() {
		t.Fatal("HasOrderBy should return false initially")
	}

	// Test OrderBy returns empty string initially
	if query.OrderBy() != "" {
		t.Fatal("OrderBy should return empty string initially")
	}

	// Test SetOrderBy
	query = query.SetOrderBy("created_at")
	if !query.HasOrderBy() {
		t.Fatal("HasOrderBy should return true after SetOrderBy")
	}
	if query.OrderBy() != "created_at" {
		t.Fatalf("OrderBy should return 'created_at', got '%s'", query.OrderBy())
	}

	// Test validation with valid value
	err := query.Validate()
	if err != nil {
		t.Fatal("unexpected validation error:", err)
	}

	// Test validation with empty value
	query = query.SetOrderBy("")
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for empty order_by")
	}
	if err.Error() != "user role query. order_by cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestUserRoleQuerySortDirection(t *testing.T) {
	query := NewUserRoleQuery()

	// Test HasSortDirection returns false initially
	if query.HasSortDirection() {
		t.Fatal("HasSortDirection should return false initially")
	}

	// Test SortDirection returns empty string initially
	if query.SortDirection() != "" {
		t.Fatal("SortDirection should return empty string initially")
	}

	// Test SetSortDirection
	query = query.SetSortDirection("ASC")
	if !query.HasSortDirection() {
		t.Fatal("HasSortDirection should return true after SetSortDirection")
	}
	if query.SortDirection() != "ASC" {
		t.Fatalf("SortDirection should return 'ASC', got '%s'", query.SortDirection())
	}

	// Test validation with valid value
	err := query.Validate()
	if err != nil {
		t.Fatal("unexpected validation error:", err)
	}

	// Test validation with empty value
	query = query.SetSortDirection("")
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for empty sort_direction")
	}
	if err.Error() != "user role query. sort_direction cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestUserRoleQueryColumns(t *testing.T) {
	query := NewUserRoleQuery()

	// Test Columns returns empty slice initially
	if len(query.Columns()) != 0 {
		t.Fatal("Columns should return empty slice initially")
	}

	// Test SetColumns
	columns := []string{"id", "user_id", "role_id"}
	query = query.SetColumns(columns)
	if len(query.Columns()) != 3 {
		t.Fatalf("Columns should return 3 items, got %d", len(query.Columns()))
	}
}

func TestUserRoleQueryCountOnly(t *testing.T) {
	query := NewUserRoleQuery()

	// Test HasCountOnly returns false initially
	if query.HasCountOnly() {
		t.Fatal("HasCountOnly should return false initially")
	}

	// Test IsCountOnly returns false initially
	if query.IsCountOnly() {
		t.Fatal("IsCountOnly should return false initially")
	}

	// Test SetCountOnly
	query = query.SetCountOnly(true)
	if !query.HasCountOnly() {
		t.Fatal("HasCountOnly should return true after SetCountOnly")
	}
	if !query.IsCountOnly() {
		t.Fatal("IsCountOnly should return true after SetCountOnly(true)")
	}

	query = query.SetCountOnly(false)
	if query.IsCountOnly() {
		t.Fatal("IsCountOnly should return false after SetCountOnly(false)")
	}
}

func TestUserRoleQuerySoftDeletedIncluded(t *testing.T) {
	query := NewUserRoleQuery()

	// Test HasSoftDeletedIncluded returns false initially
	if query.HasSoftDeletedIncluded() {
		t.Fatal("HasSoftDeletedIncluded should return false initially")
	}

	// Test SoftDeletedIncluded returns false initially
	if query.SoftDeletedIncluded() {
		t.Fatal("SoftDeletedIncluded should return false initially")
	}

	// Test SetSoftDeletedIncluded
	query = query.SetSoftDeletedIncluded(true)
	if !query.HasSoftDeletedIncluded() {
		t.Fatal("HasSoftDeletedIncluded should return true after SetSoftDeletedIncluded")
	}
	if !query.SoftDeletedIncluded() {
		t.Fatal("SoftDeletedIncluded should return true after SetSoftDeletedIncluded(true)")
	}

	query = query.SetSoftDeletedIncluded(false)
	if query.SoftDeletedIncluded() {
		t.Fatal("SoftDeletedIncluded should return false after SetSoftDeletedIncluded(false)")
	}
}

func TestUserRoleQueryChaining(t *testing.T) {
	query := NewUserRoleQuery()

	// Test method chaining
	query = query.
		SetID("123").
		SetUserID("user_1").
		SetRoleID("role_1").
		SetLimit(10).
		SetOffset(5).
		SetOrderBy("created_at").
		SetSortDirection("DESC")

	if query.GetID() != "123" {
		t.Fatalf("GetID should return '123', got '%s'", query.GetID())
	}
	if query.UserID() != "user_1" {
		t.Fatalf("UserID should return 'user_1', got '%s'", query.UserID())
	}
	if query.RoleID() != "role_1" {
		t.Fatalf("RoleID should return 'role_1', got '%s'", query.RoleID())
	}
	if query.Limit() != 10 {
		t.Fatalf("Limit should return 10, got %d", query.Limit())
	}
	if query.Offset() != 5 {
		t.Fatalf("Offset should return 5, got %d", query.Offset())
	}
	if query.OrderBy() != "created_at" {
		t.Fatalf("OrderBy should return 'created_at', got '%s'", query.OrderBy())
	}
	if query.SortDirection() != "DESC" {
		t.Fatalf("SortDirection should return 'DESC', got '%s'", query.SortDirection())
	}
}
