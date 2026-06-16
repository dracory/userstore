package userstore

import (
	"encoding/json"
	"time"

	"github.com/dracory/neat/database/orm"
	"github.com/dracory/neat/database/soft_delete"
	"github.com/dracory/str"
	"github.com/dromara/carbon/v2"
)

// == TYPE ====================================================================

type userImplementation struct {
	orm.ShortID

	StatusField          string `db:"status"`
	FirstNameField       string `db:"first_name"`
	MiddleNamesField     string `db:"middle_names"`
	LastNameField        string `db:"last_name"`
	BusinessNameField    string `db:"business_name"`
	PhoneField           string `db:"phone"`
	EmailField           string `db:"email"`
	PasswordField        string `db:"password"`
	RoleField            string `db:"role"`
	CountryField         string `db:"country"`
	TimezoneField        string `db:"timezone"`
	ProfileImageUrlField string `db:"profile_image_url"`
	MetasField           string `db:"metas"`
	MemoField            string `db:"memo"`
	CreatedAtField       orm.CreatedAt
	UpdatedAtField       orm.UpdatedAt
	soft_delete.SoftDeletesMaxDate
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
		SetSoftDeletedAt(MAX_DATETIME)

	err := o.SetMetas(map[string]string{})

	if err != nil {
		return o
	}

	return o
}

func NewUserFromExistingData(data map[string]string) UserInterface {
	o := &userImplementation{}
	if v, ok := data[COLUMN_ID]; ok {
		o.SetID(v)
	}
	if v, ok := data[COLUMN_STATUS]; ok {
		o.SetStatus(v)
	}
	if v, ok := data[COLUMN_FIRST_NAME]; ok {
		o.SetFirstName(v)
	}
	if v, ok := data[COLUMN_MIDDLE_NAMES]; ok {
		o.SetMiddleNames(v)
	}
	if v, ok := data[COLUMN_LAST_NAME]; ok {
		o.SetLastName(v)
	}
	if v, ok := data[COLUMN_BUSINESS_NAME]; ok {
		o.SetBusinessName(v)
	}
	if v, ok := data[COLUMN_PHONE]; ok {
		o.SetPhone(v)
	}
	if v, ok := data[COLUMN_EMAIL]; ok {
		o.SetEmail(v)
	}
	if v, ok := data[COLUMN_PASSWORD]; ok {
		o.SetPassword(v)
	}
	if v, ok := data[COLUMN_ROLE]; ok {
		o.SetRole(v)
	}
	if v, ok := data[COLUMN_COUNTRY]; ok {
		o.SetCountry(v)
	}
	if v, ok := data[COLUMN_TIMEZONE]; ok {
		o.SetTimezone(v)
	}
	if v, ok := data[COLUMN_PROFILE_IMAGE_URL]; ok {
		o.SetProfileImageUrl(v)
	}
	if v, ok := data[COLUMN_METAS]; ok {
		o.SetMetas(map[string]string{})
		o.UpsertMetas(map[string]string{v: v})
	}
	if v, ok := data[COLUMN_MEMO]; ok {
		o.SetMemo(v)
	}
	if v, ok := data[COLUMN_CREATED_AT]; ok {
		o.SetCreatedAt(v)
	}
	if v, ok := data[COLUMN_UPDATED_AT]; ok {
		o.SetUpdatedAt(v)
	}
	if v, ok := data[COLUMN_SOFT_DELETED_AT]; ok {
		o.SetSoftDeletedAt(v)
	}
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
	return o.SoftDeletedAt.Before(time.Now().UTC())
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

// == DATAOBJECT COMPATIBILITY ================================================

// Get returns the value of the specified column.
func (o *userImplementation) Get(columnName string) string {
	switch columnName {
	case COLUMN_ID:
		return o.GetID()
	case COLUMN_STATUS:
		return o.GetStatus()
	case COLUMN_FIRST_NAME:
		return o.GetFirstName()
	case COLUMN_MIDDLE_NAMES:
		return o.GetMiddleNames()
	case COLUMN_LAST_NAME:
		return o.GetLastName()
	case COLUMN_BUSINESS_NAME:
		return o.GetBusinessName()
	case COLUMN_PHONE:
		return o.GetPhone()
	case COLUMN_EMAIL:
		return o.GetEmail()
	case COLUMN_PASSWORD:
		return o.GetPassword()
	case COLUMN_ROLE:
		return o.GetRole()
	case COLUMN_COUNTRY:
		return o.GetCountry()
	case COLUMN_TIMEZONE:
		return o.GetTimezone()
	case COLUMN_PROFILE_IMAGE_URL:
		return o.GetProfileImageUrl()
	case COLUMN_METAS:
		return o.MetasField
	case COLUMN_MEMO:
		return o.GetMemo()
	case COLUMN_CREATED_AT:
		return o.GetCreatedAt()
	case COLUMN_UPDATED_AT:
		return o.GetUpdatedAt()
	case COLUMN_SOFT_DELETED_AT:
		return o.GetSoftDeletedAt()
	}
	return ""
}

// Set sets the value of the specified column.
func (o *userImplementation) Set(columnName string, value string) {
	switch columnName {
	case COLUMN_ID:
		o.SetID(value)
	case COLUMN_STATUS:
		o.SetStatus(value)
	case COLUMN_FIRST_NAME:
		o.SetFirstName(value)
	case COLUMN_MIDDLE_NAMES:
		o.SetMiddleNames(value)
	case COLUMN_LAST_NAME:
		o.SetLastName(value)
	case COLUMN_BUSINESS_NAME:
		o.SetBusinessName(value)
	case COLUMN_PHONE:
		o.SetPhone(value)
	case COLUMN_EMAIL:
		o.SetEmail(value)
	case COLUMN_PASSWORD:
		o.SetPassword(value)
	case COLUMN_ROLE:
		o.SetRole(value)
	case COLUMN_COUNTRY:
		o.SetCountry(value)
	case COLUMN_TIMEZONE:
		o.SetTimezone(value)
	case COLUMN_PROFILE_IMAGE_URL:
		o.SetProfileImageUrl(value)
	case COLUMN_METAS:
		o.MetasField = value
	case COLUMN_MEMO:
		o.SetMemo(value)
	case COLUMN_CREATED_AT:
		o.SetCreatedAt(value)
	case COLUMN_UPDATED_AT:
		o.SetUpdatedAt(value)
	case COLUMN_SOFT_DELETED_AT:
		o.SetSoftDeletedAt(value)
	}
}

// Data returns all fields as a map.
func (o *userImplementation) Data() map[string]string {
	return map[string]string{
		COLUMN_ID:                o.GetID(),
		COLUMN_STATUS:            o.GetStatus(),
		COLUMN_FIRST_NAME:        o.GetFirstName(),
		COLUMN_MIDDLE_NAMES:      o.GetMiddleNames(),
		COLUMN_LAST_NAME:         o.GetLastName(),
		COLUMN_BUSINESS_NAME:     o.GetBusinessName(),
		COLUMN_PHONE:             o.GetPhone(),
		COLUMN_EMAIL:             o.GetEmail(),
		COLUMN_PASSWORD:          o.GetPassword(),
		COLUMN_ROLE:              o.GetRole(),
		COLUMN_COUNTRY:           o.GetCountry(),
		COLUMN_TIMEZONE:          o.GetTimezone(),
		COLUMN_PROFILE_IMAGE_URL: o.GetProfileImageUrl(),
		COLUMN_METAS:             o.MetasField,
		COLUMN_MEMO:              o.GetMemo(),
		COLUMN_CREATED_AT:        o.GetCreatedAt(),
		COLUMN_UPDATED_AT:        o.GetUpdatedAt(),
		COLUMN_SOFT_DELETED_AT:   o.GetSoftDeletedAt(),
	}
}

// DataChanged returns all fields as a map (dirty tracking disabled).
func (o *userImplementation) DataChanged() map[string]string {
	return o.Data()
}

// MarkAsNotDirty is a no-op (dirty tracking disabled).
func (o *userImplementation) MarkAsNotDirty() {}

// == SETTERS AND GETTERS =====================================================

func (o *userImplementation) GetBusinessName() string {
	return o.BusinessNameField
}

func (o *userImplementation) SetBusinessName(businessName string) UserInterface {
	o.BusinessNameField = businessName
	return o
}

func (o *userImplementation) GetCountry() string {
	return o.CountryField
}

func (o *userImplementation) SetCountry(country string) UserInterface {
	o.CountryField = country
	return o
}

func (o *userImplementation) GetCreatedAt() string {
	if o.CreatedAtField.CreatedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(o.CreatedAtField.CreatedAt).ToDateTimeString()
}

func (o *userImplementation) GetCreatedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(o.CreatedAtField.CreatedAt)
}

func (o *userImplementation) SetCreatedAt(createdAt string) UserInterface {
	if createdAt == "" {
		return o
	}
	o.CreatedAtField.CreatedAt = carbon.Parse(createdAt, carbon.UTC).StdTime()
	return o
}

func (o *userImplementation) GetSoftDeletedAt() string {
	if o.SoftDeletedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(o.SoftDeletedAt).ToDateTimeString()
}

func (o *userImplementation) GetSoftDeletedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(o.SoftDeletedAt)
}

func (o *userImplementation) SetSoftDeletedAt(deletedAt string) UserInterface {
	if deletedAt == "" {
		return o
	}
	o.SoftDeletedAt = carbon.Parse(deletedAt, carbon.UTC).StdTime()
	return o
}

func (o *userImplementation) GetEmail() string {
	return o.EmailField
}

func (o *userImplementation) SetEmail(email string) UserInterface {
	o.EmailField = email
	return o
}

func (o *userImplementation) GetFirstName() string {
	return o.FirstNameField
}

func (o *userImplementation) SetFirstName(firstName string) UserInterface {
	o.FirstNameField = firstName
	return o
}

func (o *userImplementation) GetID() string {
	return o.ShortID.ID
}

func (o *userImplementation) SetID(id string) UserInterface {
	o.ShortID.ID = id
	return o
}

func (o *userImplementation) GetLastName() string {
	return o.LastNameField
}

func (o *userImplementation) SetLastName(lastName string) UserInterface {
	o.LastNameField = lastName
	return o
}

func (o *userImplementation) GetMemo() string {
	return o.MemoField
}

func (o *userImplementation) SetMemo(memo string) UserInterface {
	o.MemoField = memo
	return o
}

func (o *userImplementation) GetMiddleNames() string {
	return o.MiddleNamesField
}

func (o *userImplementation) SetMiddleNames(middleNames string) UserInterface {
	o.MiddleNamesField = middleNames
	return o
}

func (o *userImplementation) GetMetas() (map[string]string, error) {
	metasStr := o.MetasField

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
	o.MetasField = string(mapString)
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
	return o.PasswordField
}

func (o *userImplementation) PasswordCompare(password string) bool {
	hash := o.GetPassword()
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
	o.PasswordField = password
	return o
}

func (o *userImplementation) GetPhone() string {
	return o.PhoneField
}

func (o *userImplementation) SetPhone(phone string) UserInterface {
	o.PhoneField = phone
	return o
}

func (o *userImplementation) GetProfileImageUrl() string {
	return o.ProfileImageUrlField
}

func (o *userImplementation) ProfileImageOrDefaultUrl() string {
	defaultURL := UserNoImageUrl()

	if o.GetProfileImageUrl() != "" {
		return o.GetProfileImageUrl()
	}

	return defaultURL
}

func (o *userImplementation) SetProfileImageUrl(profileImageUrl string) UserInterface {
	o.ProfileImageUrlField = profileImageUrl
	return o
}

func (o *userImplementation) GetRole() string {
	return o.RoleField
}

func (o *userImplementation) SetRole(role string) UserInterface {
	o.RoleField = role
	return o
}

func (o *userImplementation) GetStatus() string {
	return o.StatusField
}

func (o *userImplementation) SetStatus(status string) UserInterface {
	o.StatusField = status
	return o
}

func (o *userImplementation) GetTimezone() string {
	return o.TimezoneField
}

func (o *userImplementation) SetTimezone(timezone string) UserInterface {
	o.TimezoneField = timezone
	return o
}

func (o *userImplementation) GetUpdatedAt() string {
	if o.UpdatedAtField.UpdatedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(o.UpdatedAtField.UpdatedAt).ToDateTimeString()
}

func (o *userImplementation) GetUpdatedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(o.UpdatedAtField.UpdatedAt)
}

func (o *userImplementation) SetUpdatedAt(updatedAt string) UserInterface {
	if updatedAt == "" {
		return o
	}
	o.UpdatedAtField.UpdatedAt = carbon.Parse(updatedAt, carbon.UTC).StdTime()
	return o
}
