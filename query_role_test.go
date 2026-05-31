package userstore

import (
	"testing"
)

func TestRoleQueryHandle(t *testing.T) {
	query := NewRoleQuery()

	// Test HasHandle returns false initially
	if query.HasHandle() {
		t.Fatal("HasHandle should return false initially")
	}

	// Test Handle returns empty string initially
	if query.Handle() != "" {
		t.Fatal("Handle should return empty string initially")
	}

	// Test SetHandle
	query = query.SetHandle("admin")
	if !query.HasHandle() {
		t.Fatal("HasHandle should return true after SetHandle")
	}
	if query.Handle() != "admin" {
		t.Fatalf("Handle should return 'admin', got '%s'", query.Handle())
	}
}

func TestRoleQueryID(t *testing.T) {
	query := NewRoleQuery()

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
	if err.Error() != "category query. id cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestRoleQueryStatus(t *testing.T) {
	query := NewRoleQuery()

	// Test HasStatus returns false initially
	if query.HasStatus() {
		t.Fatal("HasStatus should return false initially")
	}

	// Test Status returns empty string initially
	if query.Status() != "" {
		t.Fatal("Status should return empty string initially")
	}

	// Test SetStatus
	query = query.SetStatus("active")
	if !query.HasStatus() {
		t.Fatal("HasStatus should return true after SetStatus")
	}
	if query.Status() != "active" {
		t.Fatalf("Status should return 'active', got '%s'", query.Status())
	}

	// Test validation with valid value
	err := query.Validate()
	if err != nil {
		t.Fatal("unexpected validation error:", err)
	}

	// Test validation with empty value
	query = query.SetStatus("")
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for empty status")
	}
	if err.Error() != "category query. status cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestRoleQueryTitleLike(t *testing.T) {
	query := NewRoleQuery()

	// Test HasTitleLike returns false initially
	if query.HasTitleLike() {
		t.Fatal("HasTitleLike should return false initially")
	}

	// Test TitleLike returns empty string initially
	if query.TitleLike() != "" {
		t.Fatal("TitleLike should return empty string initially")
	}

	// Test SetTitleLike
	query = query.SetTitleLike("admin")
	if !query.HasTitleLike() {
		t.Fatal("HasTitleLike should return true after SetTitleLike")
	}
	if query.TitleLike() != "admin" {
		t.Fatalf("TitleLike should return 'admin', got '%s'", query.TitleLike())
	}

	// Test validation with valid value
	err := query.Validate()
	if err != nil {
		t.Fatal("unexpected validation error:", err)
	}

	// Test validation with empty value
	query = query.SetTitleLike("")
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for empty title_like")
	}
	if err.Error() != "category query. title_like cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestRoleQueryLimit(t *testing.T) {
	query := NewRoleQuery()

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
	if err.Error() != "category query. limit must be greater than 0" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}

	// Test validation with invalid value (negative)
	query = query.SetLimit(-5)
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for negative limit")
	}
}

func TestRoleQueryOffset(t *testing.T) {
	query := NewRoleQuery()

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
	if err.Error() != "category query. offset must be greater than or equal to 0" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestRoleQueryOrderBy(t *testing.T) {
	query := NewRoleQuery()

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
	if err.Error() != "category query. order_by cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestRoleQuerySortDirection(t *testing.T) {
	query := NewRoleQuery()

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
	if err.Error() != "category query. sort_direction cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestRoleQueryColumns(t *testing.T) {
	query := NewRoleQuery()

	// Test Columns returns empty slice initially
	if len(query.Columns()) != 0 {
		t.Fatal("Columns should return empty slice initially")
	}

	// Test SetColumns
	columns := []string{"id", "name", "handle"}
	query = query.SetColumns(columns)
	if len(query.Columns()) != 3 {
		t.Fatalf("Columns should return 3 items, got %d", len(query.Columns()))
	}
}

func TestRoleQueryCountOnly(t *testing.T) {
	query := NewRoleQuery()

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

func TestRoleQuerySoftDeletedIncluded(t *testing.T) {
	query := NewRoleQuery()

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

func TestRoleQueryChaining(t *testing.T) {
	query := NewRoleQuery()

	// Test method chaining
	query = query.
		SetID("123").
		SetHandle("admin").
		SetStatus("active").
		SetTitleLike("adm").
		SetLimit(10).
		SetOffset(5).
		SetOrderBy("created_at").
		SetSortDirection("DESC")

	if query.GetID() != "123" {
		t.Fatalf("GetID should return '123', got '%s'", query.GetID())
	}
	if query.Handle() != "admin" {
		t.Fatalf("Handle should return 'admin', got '%s'", query.Handle())
	}
	if query.Status() != "active" {
		t.Fatalf("Status should return 'active', got '%s'", query.Status())
	}
	if query.TitleLike() != "adm" {
		t.Fatalf("TitleLike should return 'adm', got '%s'", query.TitleLike())
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
