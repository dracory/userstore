package userstore

import (
	"github.com/golang-module/carbon/v2"
	"github.com/gouniverse/dataobject"
	"github.com/gouniverse/maputils"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/uid"
	"github.com/gouniverse/utils"
)

// == CLASS ===================================================================

type user struct {
	dataobject.DataObject
}

var _ UserInterface = (*user)(nil)

// == CONSTRUCTORS ============================================================

func NewUser() UserInterface {
	o := &user{}

	o.SetID(uid.HumanUid()).
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
	o := &user{}
	o.Hydrate(data)
	return o
}

// == METHODS =================================================================

func UserNoImageUrl() string {
	return "/user/default.png"
	//return config.MediaUrl + "/user/default.png"
}

func (o *user) IsActive() bool {
	return o.Status() == USER_STATUS_ACTIVE
}

func (o *user) IsDeleted() bool {
	return o.Status() == USER_STATUS_DELETED
}

func (o *user) IsInactive() bool {
	return o.Status() == USER_STATUS_INACTIVE
}

func (o *user) IsUnverified() bool {
	return o.Status() == USER_STATUS_UNVERIFIED
}

func (o *user) IsAdministrator() bool {
	return o.Role() == USER_ROLE_ADMINISTRATOR
}

func (o *user) IsManager() bool {
	return o.Role() == USER_ROLE_MANAGER
}

func (o *user) IsSuperuser() bool {
	return o.Role() == USER_ROLE_SUPERUSER
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
func (o *user) IsRegistrationCompleted() bool {
	return o.FirstName() != "" && o.LastName() != ""
}

// == SETTERS AND GETTERS =====================================================

func (o *user) BusinessName() string {
	return o.Get("business_name")
}

func (o *user) SetBusinessName(businessName string) UserInterface {
	o.Set("business_name", businessName)
	return o
}

func (o *user) Country() string {
	return o.Get("country")
}

func (o *user) SetCountry(country string) UserInterface {
	o.Set("country", country)
	return o
}

func (o *user) CreatedAt() string {
	return o.Get("created_at")
}

func (o *user) CreatedAtCarbon() carbon.Carbon {
	return carbon.Parse(o.CreatedAt(), carbon.UTC)
}

func (o *user) SetCreatedAt(createdAt string) UserInterface {
	o.Set("created_at", createdAt)
	return o
}

func (o *user) SoftDeletedAt() string {
	return o.Get(COLUMN_SOFT_DELETED_AT)
}

func (o *user) SoftDeletedAtCarbon() carbon.Carbon {
	return carbon.NewCarbon().Parse(o.SoftDeletedAt(), carbon.UTC)
}

func (o *user) SetSoftDeletedAt(deletedAt string) UserInterface {
	o.Set(COLUMN_SOFT_DELETED_AT, deletedAt)
	return o
}

func (o *user) Email() string {
	return o.Get(COLUMN_EMAIL)
}

func (o *user) SetEmail(email string) UserInterface {
	o.Set(COLUMN_EMAIL, email)
	return o
}

func (o *user) FirstName() string {
	return o.Get(COLUMN_FIRST_NAME)
}

func (o *user) SetFirstName(firstName string) UserInterface {
	o.Set(COLUMN_FIRST_NAME, firstName)
	return o
}

func (o *user) ID() string {
	return o.Get(COLUMN_ID)
}

func (o *user) SetID(id string) UserInterface {
	o.Set(COLUMN_ID, id)
	return o
}

func (o *user) LastName() string {
	return o.Get(COLUMN_LAST_NAME)
}

func (o *user) SetLastName(lastName string) UserInterface {
	o.Set(COLUMN_LAST_NAME, lastName)
	return o
}

func (o *user) Memo() string {
	return o.Get(COLUMN_MEMO)
}

func (o *user) SetMemo(memo string) UserInterface {
	o.Set(COLUMN_MEMO, memo)
	return o
}

func (o *user) MiddleNames() string {
	return o.Get(COLUMN_MIDDLE_NAMES)
}

func (o *user) SetMiddleNames(middleNames string) UserInterface {
	o.Set(COLUMN_MIDDLE_NAMES, middleNames)
	return o
}

func (o *user) Metas() (map[string]string, error) {
	metasStr := o.Get(COLUMN_METAS)

	if metasStr == "" {
		metasStr = "{}"
	}

	metasJson, errJson := utils.FromJSON(metasStr, map[string]string{})
	if errJson != nil {
		return map[string]string{}, errJson
	}

	return maputils.MapStringAnyToMapStringString(metasJson.(map[string]any)), nil
}

func (o *user) Meta(name string) string {
	metas, err := o.Metas()

	if err != nil {
		return ""
	}

	if value, exists := metas[name]; exists {
		return value
	}

	return ""
}

func (o *user) SetMeta(name string, value string) error {
	return o.UpsertMetas(map[string]string{name: value})
}

// SetMetas stores metas as json string
// Warning: it overwrites any existing metas
func (o *user) SetMetas(metas map[string]string) error {
	mapString, err := utils.ToJSON(metas)
	if err != nil {
		return err
	}
	o.Set(COLUMN_METAS, mapString)
	return nil
}

func (o *user) UpsertMetas(metas map[string]string) error {
	currentMetas, err := o.Metas()

	if err != nil {
		return err
	}

	for k, v := range metas {
		currentMetas[k] = v
	}

	return o.SetMetas(currentMetas)
}

func (o *user) Password() string {
	return o.Get(COLUMN_PASSWORD)
}

func (o *user) PasswordCompare(password string) bool {
	hash := o.Get(COLUMN_PASSWORD)
	return utils.StrToBcryptHashCompare(password, hash)
}

// SetPasswordAndHash hashes the password before saving
func (o *user) SetPasswordAndHash(password string) error {
	hash, err := utils.StrToBcryptHash(password)

	if err != nil {
		return err
	}

	o.SetPassword(hash)

	return nil
}

// SetPassword sets the password as provided, if you want it hashed use SetPasswordAndHash() method
func (o *user) SetPassword(password string) UserInterface {
	o.Set(COLUMN_PASSWORD, password)
	return o
}

func (o *user) Phone() string {
	return o.Get(COLUMN_PHONE)
}

func (o *user) SetPhone(phone string) UserInterface {
	o.Set(COLUMN_PHONE, phone)
	return o
}

func (o *user) ProfileImageUrl() string {
	return o.Get(COLUMN_PROFILE_IMAGE_URL)
}

func (o *user) ProfileImageOrDefaultUrl() string {
	defaultURL := UserNoImageUrl()

	if o.ProfileImageUrl() != "" {
		return o.ProfileImageUrl()
	}

	return defaultURL
}

func (o *user) SetProfileImageUrl(imageUrl string) UserInterface {
	o.Set(COLUMN_PROFILE_IMAGE_URL, imageUrl)
	return o
}

func (o *user) Role() string {
	return o.Get(COLUMN_ROLE)
}

func (o *user) SetRole(role string) UserInterface {
	o.Set(COLUMN_ROLE, role)
	return o
}

func (o *user) Status() string {
	return o.Get(COLUMN_STATUS)
}

func (o *user) SetStatus(status string) UserInterface {
	o.Set(COLUMN_STATUS, status)
	return o
}

func (o *user) Timezone() string {
	return o.Get(COLUMN_TIMEZONE)
}

func (o *user) SetTimezone(timezone string) UserInterface {
	o.Set(COLUMN_TIMEZONE, timezone)
	return o
}

func (o *user) UpdatedAt() string {
	return o.Get(COLUMN_UPDATED_AT)
}

func (o *user) UpdatedAtCarbon() carbon.Carbon {
	return carbon.NewCarbon().Parse(o.Get(COLUMN_UPDATED_AT), carbon.UTC)
}

func (o *user) SetUpdatedAt(updatedAt string) UserInterface {
	o.Set("updated_at", updatedAt)
	return o
}
