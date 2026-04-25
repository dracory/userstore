package userstore

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	user := NewUser()

	if user == nil {
		t.Fatal("NewUser should not return nil")
	}

	if user.GetID() == "" {
		t.Error("ID should not be empty")
	}

	if user.GetStatus() != USER_STATUS_UNVERIFIED {
		t.Errorf("Expected status %s, got %s", USER_STATUS_UNVERIFIED, user.GetStatus())
	}

	if user.GetRole() != USER_ROLE_USER {
		t.Errorf("Expected role %s, got %s", USER_ROLE_USER, user.GetRole())
	}

	if user.GetFirstName() != "" {
		t.Error("FirstName should be empty")
	}

	if user.GetLastName() != "" {
		t.Error("LastName should be empty")
	}

	if user.GetEmail() != "" {
		t.Error("Email should be empty")
	}
}

func TestNewUserFromExistingData(t *testing.T) {
	data := map[string]string{
		COLUMN_ID:            "test123",
		COLUMN_EMAIL:         "test@example.com",
		COLUMN_FIRST_NAME:    "John",
		COLUMN_LAST_NAME:     "Doe",
		COLUMN_STATUS:        USER_STATUS_ACTIVE,
		COLUMN_ROLE:          USER_ROLE_ADMINISTRATOR,
		COLUMN_BUSINESS_NAME: "Acme Inc",
	}

	user := NewUserFromExistingData(data)

	if user == nil {
		t.Fatal("NewUserFromExistingData should not return nil")
	}

	if user.GetID() != "test123" {
		t.Errorf("Expected ID test123, got %s", user.GetID())
	}

	if user.GetEmail() != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %s", user.GetEmail())
	}

	if user.GetFirstName() != "John" {
		t.Errorf("Expected FirstName John, got %s", user.GetFirstName())
	}

	if user.GetLastName() != "Doe" {
		t.Errorf("Expected LastName Doe, got %s", user.GetLastName())
	}

	if user.GetStatus() != USER_STATUS_ACTIVE {
		t.Errorf("Expected status %s, got %s", USER_STATUS_ACTIVE, user.GetStatus())
	}

	if user.GetRole() != USER_ROLE_ADMINISTRATOR {
		t.Errorf("Expected role %s, got %s", USER_ROLE_ADMINISTRATOR, user.GetRole())
	}

	if user.GetBusinessName() != "Acme Inc" {
		t.Errorf("Expected BusinessName Acme Inc, got %s", user.GetBusinessName())
	}
}

func TestUserIsActive(t *testing.T) {
	user := NewUser().SetStatus(USER_STATUS_ACTIVE)
	if !user.IsActive() {
		t.Error("User with ACTIVE status should be active")
	}

	user = NewUser().SetStatus(USER_STATUS_INACTIVE)
	if user.IsActive() {
		t.Error("User with INACTIVE status should not be active")
	}
}

func TestUserIsInactive(t *testing.T) {
	user := NewUser().SetStatus(USER_STATUS_INACTIVE)
	if !user.IsInactive() {
		t.Error("User with INACTIVE status should be inactive")
	}

	user = NewUser().SetStatus(USER_STATUS_ACTIVE)
	if user.IsInactive() {
		t.Error("User with ACTIVE status should not be inactive")
	}
}

func TestUserIsUnverified(t *testing.T) {
	user := NewUser().SetStatus(USER_STATUS_UNVERIFIED)
	if !user.IsUnverified() {
		t.Error("User with UNVERIFIED status should be unverified")
	}

	user = NewUser().SetStatus(USER_STATUS_ACTIVE)
	if user.IsUnverified() {
		t.Error("User with ACTIVE status should not be unverified")
	}
}

func TestUserIsAdministrator(t *testing.T) {
	user := NewUser().SetRole(USER_ROLE_ADMINISTRATOR)
	if !user.IsAdministrator() {
		t.Error("User with ADMINISTRATOR role should be administrator")
	}

	user = NewUser().SetRole(USER_ROLE_USER)
	if user.IsAdministrator() {
		t.Error("User with USER role should not be administrator")
	}
}

func TestUserIsManager(t *testing.T) {
	user := NewUser().SetRole(USER_ROLE_MANAGER)
	if !user.IsManager() {
		t.Error("User with MANAGER role should be manager")
	}

	user = NewUser().SetRole(USER_ROLE_USER)
	if user.IsManager() {
		t.Error("User with USER role should not be manager")
	}
}

func TestUserIsSuperuser(t *testing.T) {
	user := NewUser().SetRole(USER_ROLE_SUPERUSER)
	if !user.IsSuperuser() {
		t.Error("User with SUPERUSER role should be superuser")
	}

	user = NewUser().SetRole(USER_ROLE_USER)
	if user.IsSuperuser() {
		t.Error("User with USER role should not be superuser")
	}
}

func TestUserIsRegistrationCompleted(t *testing.T) {
	user := NewUser().SetFirstName("John").SetLastName("Doe")
	if !user.IsRegistrationCompleted() {
		t.Error("User with first and last name should have completed registration")
	}

	user = NewUser().SetFirstName("John")
	if user.IsRegistrationCompleted() {
		t.Error("User with only first name should not have completed registration")
	}

	user = NewUser().SetLastName("Doe")
	if user.IsRegistrationCompleted() {
		t.Error("User with only last name should not have completed registration")
	}

	user = NewUser()
	if user.IsRegistrationCompleted() {
		t.Error("User with no names should not have completed registration")
	}
}

func TestUserSettersAndGetters(t *testing.T) {
	user := NewUser()

	user.SetID("test123")
	if user.GetID() != "test123" {
		t.Errorf("Expected ID test123, got %s", user.GetID())
	}

	user.SetEmail("test@example.com")
	if user.GetEmail() != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %s", user.GetEmail())
	}

	user.SetFirstName("John")
	if user.GetFirstName() != "John" {
		t.Errorf("Expected FirstName John, got %s", user.GetFirstName())
	}

	user.SetLastName("Doe")
	if user.GetLastName() != "Doe" {
		t.Errorf("Expected LastName Doe, got %s", user.GetLastName())
	}

	user.SetMiddleNames("William")
	if user.GetMiddleNames() != "William" {
		t.Errorf("Expected MiddleNames William, got %s", user.GetMiddleNames())
	}

	user.SetPhone("+1234567890")
	if user.GetPhone() != "+1234567890" {
		t.Errorf("Expected Phone +1234567890, got %s", user.GetPhone())
	}

	user.SetCountry("US")
	if user.GetCountry() != "US" {
		t.Errorf("Expected Country US, got %s", user.GetCountry())
	}

	user.SetTimezone("UTC")
	if user.GetTimezone() != "UTC" {
		t.Errorf("Expected Timezone UTC, got %s", user.GetTimezone())
	}

	user.SetBusinessName("Acme Inc")
	if user.GetBusinessName() != "Acme Inc" {
		t.Errorf("Expected BusinessName Acme Inc, got %s", user.GetBusinessName())
	}

	user.SetMemo("Test memo")
	if user.GetMemo() != "Test memo" {
		t.Errorf("Expected Memo Test memo, got %s", user.GetMemo())
	}

	user.SetProfileImageUrl("https://example.com/image.jpg")
	if user.GetProfileImageUrl() != "https://example.com/image.jpg" {
		t.Errorf("Expected ProfileImageUrl https://example.com/image.jpg, got %s", user.GetProfileImageUrl())
	}
}

func TestUserMetas(t *testing.T) {
	user := NewUser()

	// Test setting metas
	err := user.SetMetas(map[string]string{"key1": "value1", "key2": "value2"})
	if err != nil {
		t.Errorf("SetMetas should not return error, got %v", err)
	}

	metas, err := user.GetMetas()
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
	if user.GetMeta("key1") != "value1" {
		t.Errorf("Expected Meta key1 to return value1, got %s", user.GetMeta("key1"))
	}

	// Test non-existent meta
	if user.GetMeta("nonexistent") != "" {
		t.Error("Meta should return empty string for non-existent key")
	}

	// Test upsert metas
	err = user.UpsertMetas(map[string]string{"key1": "newvalue1", "key3": "value3"})
	if err != nil {
		t.Errorf("UpsertMetas should not return error, got %v", err)
	}

	metas, err = user.GetMetas()
	if err != nil {
		t.Errorf("Metas should not return error, got %v", err)
	}

	if metas["key1"] != "newvalue1" {
		t.Errorf("Expected key1 to be newvalue1, got %s", metas["key1"])
	}

	if metas["key2"] != "value2" {
		t.Errorf("Expected key2 to still be value2, got %s", metas["key2"])
	}

	if metas["key3"] != "value3" {
		t.Errorf("Expected key3 to be value3, got %s", metas["key3"])
	}

	// Test SetMeta
	err = user.SetMeta("key4", "value4")
	if err != nil {
		t.Errorf("SetMeta should not return error, got %v", err)
	}

	if user.GetMeta("key4") != "value4" {
		t.Errorf("Expected Meta key4 to return value4, got %s", user.GetMeta("key4"))
	}
}

func TestUserPassword(t *testing.T) {
	user := NewUser()

	// Test SetPasswordAndHash
	password := "secretpassword"
	err := user.SetPasswordAndHash(password)
	if err != nil {
		t.Errorf("SetPasswordAndHash should not return error, got %v", err)
	}

	// Password should be hashed, not plain text
	if user.GetPassword() == password {
		t.Error("Password should be hashed, not plain text")
	}

	// Test PasswordCompare
	if !user.PasswordCompare(password) {
		t.Error("PasswordCompare should return true for correct password")
	}

	if user.PasswordCompare("wrongpassword") {
		t.Error("PasswordCompare should return false for incorrect password")
	}

	// Test SetPassword (without hashing)
	user.SetPassword("plainpassword")
	if user.GetPassword() != "plainpassword" {
		t.Error("SetPassword should set plain password without hashing")
	}
}

func TestUserNoImageUrl(t *testing.T) {
	url := UserNoImageUrl()
	if url != "/user/default.png" {
		t.Errorf("Expected /user/default.png, got %s", url)
	}
}

func TestUserChaining(t *testing.T) {
	user := NewUser().
		SetEmail("test@example.com").
		SetFirstName("John").
		SetLastName("Doe").
		SetStatus(USER_STATUS_ACTIVE).
		SetRole(USER_ROLE_ADMINISTRATOR)

	if user.GetEmail() != "test@example.com" {
		t.Error("Method chaining should work")
	}

	if user.GetFirstName() != "John" {
		t.Error("Method chaining should work")
	}

	if user.GetLastName() != "Doe" {
		t.Error("Method chaining should work")
	}

	if user.GetStatus() != USER_STATUS_ACTIVE {
		t.Error("Method chaining should work")
	}

	if user.GetRole() != USER_ROLE_ADMINISTRATOR {
		t.Error("Method chaining should work")
	}
}
