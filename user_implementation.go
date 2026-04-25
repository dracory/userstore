package userstore

import (
	"encoding/json"

	"github.com/dracory/dataobject"
	"github.com/dracory/sb"
	"github.com/dracory/str"
	"github.com/dromara/carbon/v2"
)

// == CLASS ===================================================================

type userImplementation struct {
	dataobject.DataObject
}

var _ UserInterface = (*userImplementation)(nil)

// == CONSTRUCTORS ============================================================

func NewUser() UserInterface {
	o := &userImplementation{}

	o.SetID(GenerateShortID()).
		SetStatus(USER_STATUS_UNVERIFIED).
		SetFirstName("").
		SetMiddleNames("").
		SetLastName("").
		SetEmail("").
		SetProfileImageUrl("").
		SetRole(USER_ROLE_USER).
		SetBusinessName("").
		SetPhone("").
		SetPassword("").
		SetTimezone("").
		SetCountry("").
		SetMemo("").
		SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetSoftDeletedAt(sb.MAX_DATETIME)

	err := o.SetMetas(map[string]string{})

	if err != nil {
		return o
	}

	return o
}

func NewUserFromExistingData(data map[string]string) UserInterface {
	o := &userImplementation{}
	o.Hydrate(data)
	return o
}

// == METHODS =================================================================

func UserNoImageUrl() string {
	return "/user/default.png"
}

func (o *userImplementation) IsActive() bool {
	return o.GetStatus() == USER_STATUS_ACTIVE
}

func (o *userImplementation) IsSoftDeleted() bool {
	return o.GetSoftDeletedAtCarbon().Compare("<", carbon.Now(carbon.UTC))
}

func (o *userImplementation) IsInactive() bool {
	return o.GetStatus() == USER_STATUS_INACTIVE
}

func (o *userImplementation) IsUnverified() bool {
	return o.GetStatus() == USER_STATUS_UNVERIFIED
}

func (o *userImplementation) IsAdministrator() bool {
	return o.GetRole() == USER_ROLE_ADMINISTRATOR
}

func (o *userImplementation) IsManager() bool {
	return o.GetRole() == USER_ROLE_MANAGER
}

func (o *userImplementation) IsSuperuser() bool {
	return o.GetRole() == USER_ROLE_SUPERUSER
}

// IsRegistrationCompleted checks if the user registration is incomplete.
//
// Registration is considered incomplete if the user's first name
// or last name is empty.
//
// Parameters:
// - authUser: a pointer to a userstore.User object representing the authenticated user.
//
// Returns:
// - bool: true if the user registration is incomplete, false otherwise.
func (o *userImplementation) IsRegistrationCompleted() bool {
	return o.GetFirstName() != "" && o.GetLastName() != ""
}

// == SETTERS AND GETTERS =====================================================

// Get returns the value of the specified column.
// Always prefers to use the existing Get* methods.
// func (o *userImplementation) GetValue(columnName string) string {
// 	return o.Get(columnName)
// }

// Set sets the value of the specified column.
// Always prefers to use the existing Set* methods.
// func (o *userImplementation) Set(columnName string, value string) UserInterface {
// 	o.Set(columnName, value)
// 	return o
// }

func (o *userImplementation) GetBusinessName() string {
	return o.Get(COLUMN_BUSINESS_NAME)
}

func (o *userImplementation) SetBusinessName(businessName string) UserInterface {
	o.Set(COLUMN_BUSINESS_NAME, businessName)
	return o
}

func (o *userImplementation) GetCountry() string {
	return o.Get(COLUMN_COUNTRY)
}

func (o *userImplementation) SetCountry(country string) UserInterface {
	o.Set(COLUMN_COUNTRY, country)
	return o
}

func (o *userImplementation) GetCreatedAt() string {
	return o.Get(COLUMN_CREATED_AT)
}

func (o *userImplementation) GetCreatedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.GetCreatedAt(), carbon.UTC)
}

func (o *userImplementation) SetCreatedAt(createdAt string) UserInterface {
	o.Set(COLUMN_CREATED_AT, createdAt)
	return o
}

func (o *userImplementation) GetSoftDeletedAt() string {
	return o.Get(COLUMN_SOFT_DELETED_AT)
}

func (o *userImplementation) GetSoftDeletedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.GetSoftDeletedAt(), carbon.UTC)
}

func (o *userImplementation) SetSoftDeletedAt(deletedAt string) UserInterface {
	o.Set(COLUMN_SOFT_DELETED_AT, deletedAt)
	return o
}

func (o *userImplementation) GetEmail() string {
	return o.Get(COLUMN_EMAIL)
}

func (o *userImplementation) SetEmail(email string) UserInterface {
	o.Set(COLUMN_EMAIL, email)
	return o
}

func (o *userImplementation) GetFirstName() string {
	return o.Get(COLUMN_FIRST_NAME)
}

func (o *userImplementation) SetFirstName(firstName string) UserInterface {
	o.Set(COLUMN_FIRST_NAME, firstName)
	return o
}

func (o *userImplementation) GetID() string {
	return o.Get(COLUMN_ID)
}

func (o *userImplementation) SetID(id string) UserInterface {
	o.Set(COLUMN_ID, id)
	return o
}

func (o *userImplementation) GetLastName() string {
	return o.Get(COLUMN_LAST_NAME)
}

func (o *userImplementation) SetLastName(lastName string) UserInterface {
	o.Set(COLUMN_LAST_NAME, lastName)
	return o
}

func (o *userImplementation) GetMemo() string {
	return o.Get(COLUMN_MEMO)
}

func (o *userImplementation) SetMemo(memo string) UserInterface {
	o.Set(COLUMN_MEMO, memo)
	return o
}

func (o *userImplementation) GetMiddleNames() string {
	return o.Get(COLUMN_MIDDLE_NAMES)
}

func (o *userImplementation) SetMiddleNames(middleNames string) UserInterface {
	o.Set(COLUMN_MIDDLE_NAMES, middleNames)
	return o
}

func (o *userImplementation) GetMetas() (map[string]string, error) {
	metasStr := o.Get(COLUMN_METAS)

	if metasStr == "" {
		metasStr = "{}"
	}

	var metas map[string]string
	if err := json.Unmarshal([]byte(metasStr), &metas); err != nil {
		return map[string]string{}, err
	}

	return metas, nil
}

func (o *userImplementation) GetMeta(name string) string {
	metas, err := o.GetMetas()

	if err != nil {
		return ""
	}

	if value, exists := metas[name]; exists {
		return value
	}

	return ""
}

func (o *userImplementation) SetMeta(name, value string) error {
	return o.UpsertMetas(map[string]string{name: value})
}

// SetMetas stores metas as json string
// Warning: it overwrites any existing metas
func (o *userImplementation) SetMetas(metas map[string]string) error {
	mapString, err := json.Marshal(metas)
	if err != nil {
		return err
	}
	o.Set(COLUMN_METAS, string(mapString))
	return nil
}

func (o *userImplementation) UpsertMetas(metas map[string]string) error {
	currentMetas, err := o.GetMetas()

	if err != nil {
		return err
	}

	for k, v := range metas {
		currentMetas[k] = v
	}

	return o.SetMetas(currentMetas)
}

func (o *userImplementation) GetPassword() string {
	return o.Get(COLUMN_PASSWORD)
}

func (o *userImplementation) PasswordCompare(password string) bool {
	hash := o.Get(COLUMN_PASSWORD)
	return str.BcryptHashCompare(password, hash)
}

// SetPasswordAndHash hashes the password before saving
func (o *userImplementation) SetPasswordAndHash(password string) error {
	hash, err := str.ToBcryptHash(password)

	if err != nil {
		return err
	}

	o.SetPassword(hash)

	return nil
}

// SetPassword sets the password as provided, if you want it hashed use SetPasswordAndHash() method
func (o *userImplementation) SetPassword(password string) UserInterface {
	o.Set(COLUMN_PASSWORD, password)
	return o
}

func (o *userImplementation) GetPhone() string {
	return o.Get(COLUMN_PHONE)
}

func (o *userImplementation) SetPhone(phone string) UserInterface {
	o.Set(COLUMN_PHONE, phone)
	return o
}

func (o *userImplementation) GetProfileImageUrl() string {
	return o.Get(COLUMN_PROFILE_IMAGE_URL)
}

func (o *userImplementation) ProfileImageOrDefaultUrl() string {
	defaultURL := UserNoImageUrl()

	if o.GetProfileImageUrl() != "" {
		return o.GetProfileImageUrl()
	}

	return defaultURL
}

func (o *userImplementation) SetProfileImageUrl(imageUrl string) UserInterface {
	o.Set(COLUMN_PROFILE_IMAGE_URL, imageUrl)
	return o
}

func (o *userImplementation) GetRole() string {
	return o.Get(COLUMN_ROLE)
}

func (o *userImplementation) SetRole(role string) UserInterface {
	o.Set(COLUMN_ROLE, role)
	return o
}

func (o *userImplementation) GetStatus() string {
	return o.Get(COLUMN_STATUS)
}

func (o *userImplementation) SetStatus(status string) UserInterface {
	o.Set(COLUMN_STATUS, status)
	return o
}

func (o *userImplementation) GetTimezone() string {
	return o.Get(COLUMN_TIMEZONE)
}

func (o *userImplementation) SetTimezone(timezone string) UserInterface {
	o.Set(COLUMN_TIMEZONE, timezone)
	return o
}

func (o *userImplementation) GetUpdatedAt() string {
	return o.Get(COLUMN_UPDATED_AT)
}

func (o *userImplementation) GetUpdatedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.Get(COLUMN_UPDATED_AT), carbon.UTC)
}

func (o *userImplementation) SetUpdatedAt(updatedAt string) UserInterface {
	o.Set(COLUMN_UPDATED_AT, updatedAt)
	return o
}
