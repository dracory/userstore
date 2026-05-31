package userstore

import (
	"testing"
)

func TestUserQueryFirstName(t *testing.T) {
	query := NewUserQuery()

	// Test HasFirstName returns false initially
	if query.HasFirstName() {
		t.Fatal("HasFirstName should return false initially")
	}

	// Test FirstName returns empty string initially
	if query.FirstName() != "" {
		t.Fatal("FirstName should return empty string initially")
	}

	// Test SetFirstName
	query = query.SetFirstName("John")
	if !query.HasFirstName() {
		t.Fatal("HasFirstName should return true after SetFirstName")
	}
	if query.FirstName() != "John" {
		t.Fatalf("FirstName should return 'John', got '%s'", query.FirstName())
	}

	// Test validation with valid value
	err := query.Validate()
	if err != nil {
		t.Fatal("unexpected validation error:", err)
	}

	// Test validation with empty value
	query = query.SetFirstName("")
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for empty first_name")
	}
	if err.Error() != "user query. first_name cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestUserQueryLastName(t *testing.T) {
	query := NewUserQuery()

	// Test HasLastName returns false initially
	if query.HasLastName() {
		t.Fatal("HasLastName should return false initially")
	}

	// Test LastName returns empty string initially
	if query.LastName() != "" {
		t.Fatal("LastName should return empty string initially")
	}

	// Test SetLastName
	query = query.SetLastName("Doe")
	if !query.HasLastName() {
		t.Fatal("HasLastName should return true after SetLastName")
	}
	if query.LastName() != "Doe" {
		t.Fatalf("LastName should return 'Doe', got '%s'", query.LastName())
	}

	// Test validation with valid value
	err := query.Validate()
	if err != nil {
		t.Fatal("unexpected validation error:", err)
	}

	// Test validation with empty value
	query = query.SetLastName("")
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for empty last_name")
	}
	if err.Error() != "user query. last_name cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestUserQueryFirstNameLike(t *testing.T) {
	query := NewUserQuery()

	// Test HasFirstNameLike returns false initially
	if query.HasFirstNameLike() {
		t.Fatal("HasFirstNameLike should return false initially")
	}

	// Test FirstNameLike returns empty string initially
	if query.FirstNameLike() != "" {
		t.Fatal("FirstNameLike should return empty string initially")
	}

	// Test SetFirstNameLike
	query = query.SetFirstNameLike("mil")
	if !query.HasFirstNameLike() {
		t.Fatal("HasFirstNameLike should return true after SetFirstNameLike")
	}
	if query.FirstNameLike() != "mil" {
		t.Fatalf("FirstNameLike should return 'mil', got '%s'", query.FirstNameLike())
	}

	// Test validation with valid value
	err := query.Validate()
	if err != nil {
		t.Fatal("unexpected validation error:", err)
	}

	// Test validation with empty value
	query = query.SetFirstNameLike("")
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for empty first_name_like")
	}
	if err.Error() != "user query. first_name_like cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestUserQueryLastNameLike(t *testing.T) {
	query := NewUserQuery()

	// Test HasLastNameLike returns false initially
	if query.HasLastNameLike() {
		t.Fatal("HasLastNameLike should return false initially")
	}

	// Test LastNameLike returns empty string initially
	if query.LastNameLike() != "" {
		t.Fatal("LastNameLike should return empty string initially")
	}

	// Test SetLastNameLike
	query = query.SetLastNameLike("sm")
	if !query.HasLastNameLike() {
		t.Fatal("HasLastNameLike should return true after SetLastNameLike")
	}
	if query.LastNameLike() != "sm" {
		t.Fatalf("LastNameLike should return 'sm', got '%s'", query.LastNameLike())
	}

	// Test validation with valid value
	err := query.Validate()
	if err != nil {
		t.Fatal("unexpected validation error:", err)
	}

	// Test validation with empty value
	query = query.SetLastNameLike("")
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for empty last_name_like")
	}
	if err.Error() != "user query. last_name_like cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestUserQueryEmailLike(t *testing.T) {
	query := NewUserQuery()

	// Test HasEmailLike returns false initially
	if query.HasEmailLike() {
		t.Fatal("HasEmailLike should return false initially")
	}

	// Test EmailLike returns empty string initially
	if query.EmailLike() != "" {
		t.Fatal("EmailLike should return empty string initially")
	}

	// Test SetEmailLike
	query = query.SetEmailLike("test")
	if !query.HasEmailLike() {
		t.Fatal("HasEmailLike should return true after SetEmailLike")
	}
	if query.EmailLike() != "test" {
		t.Fatalf("EmailLike should return 'test', got '%s'", query.EmailLike())
	}

	// Test validation with valid value
	err := query.Validate()
	if err != nil {
		t.Fatal("unexpected validation error:", err)
	}

	// Test validation with empty value
	query = query.SetEmailLike("")
	err = query.Validate()
	if err == nil {
		t.Fatal("expected validation error for empty email_like")
	}
	if err.Error() != "user query. email_like cannot be empty" {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestUserQueryChaining(t *testing.T) {
	query := NewUserQuery()

	// Test method chaining
	query = query.
		SetFirstName("John").
		SetLastName("Doe").
		SetFirstNameLike("Joh").
		SetLastNameLike("Do").
		SetEmailLike("test")

	if query.FirstName() != "John" {
		t.Fatalf("FirstName should return 'John', got '%s'", query.FirstName())
	}
	if query.LastName() != "Doe" {
		t.Fatalf("LastName should return 'Doe', got '%s'", query.LastName())
	}
	if query.FirstNameLike() != "Joh" {
		t.Fatalf("FirstNameLike should return 'Joh', got '%s'", query.FirstNameLike())
	}
	if query.LastNameLike() != "Do" {
		t.Fatalf("LastNameLike should return 'Do', got '%s'", query.LastNameLike())
	}
	if query.EmailLike() != "test" {
		t.Fatalf("EmailLike should return 'test', got '%s'", query.EmailLike())
	}
}
