package userstore

import "github.com/golang-module/carbon/v2"

type UserInterface interface {
	// From dataobject

	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()

	IsActive() bool
	IsDeleted() bool
	IsInactive() bool
	IsUnverified() bool
	IsAdministrator() bool
	IsManager() bool
	IsSuperuser() bool
	IsRegistrationCompleted() bool
	BusinessName() string
	SetBusinessName(businessName string) UserInterface
	Country() string
	SetCountry(country string) UserInterface
	CreatedAt() string
	CreatedAtCarbon() carbon.Carbon
	SetCreatedAt(createdAt string) UserInterface
	DeletedAt() string
	SetDeletedAt(deletedAt string) UserInterface
	Email() string
	SetEmail(email string) UserInterface
	ID() string
	SetID(id string) UserInterface
	FirstName() string
	SetFirstName(firstName string) UserInterface
	LastName() string
	SetLastName(lastName string) UserInterface
	Memo() string
	SetMemo(memo string) UserInterface
	Meta(name string) string
	SetMeta(name string, value string) error
	Metas() (map[string]string, error)
	SetMetas(metas map[string]string) error
	MiddleNames() string
	SetMiddleNames(middleNames string) UserInterface
	Password() string
	SetPassword(password string) UserInterface
	Phone() string
	SetPhone(phone string) UserInterface
	ProfileImageUrl() string
	SetProfileImageUrl(profileImageUrl string) UserInterface
	Role() string
	SetRole(role string) UserInterface
	Timezone() string
	SetTimezone(timezone string) UserInterface
	Status() string
	SetStatus(status string) UserInterface
	PasswordCompare(password string) bool
	UpdatedAt() string
	UpdatedAtCarbon() carbon.Carbon
	SetUpdatedAt(updatedAt string) UserInterface
}
