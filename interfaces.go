package userstore

import "github.com/golang-module/carbon/v2"

type StoreInterface interface {
	AutoMigrate() error
	EnableDebug(debug bool)
	UserCreate(user UserInterface) error
	UserCount(options UserQueryInterface) (int64, error)
	UserDelete(user UserInterface) error
	UserDeleteByID(id string) error
	UserFindByEmail(email string) (UserInterface, error)
	UserFindByID(userID string) (UserInterface, error)
	UserList(query UserQueryInterface) ([]UserInterface, error)
	UserSoftDelete(user UserInterface) error
	UserSoftDeleteByID(id string) error
	UserUpdate(user UserInterface) error
}

type UserInterface interface {
	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()

	IsActive() bool
	IsInactive() bool
	IsSoftDeleted() bool
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
	UpsertMetas(metas map[string]string) error

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

	SoftDeletedAt() string
	SoftDeletedAtCarbon() carbon.Carbon
	SetSoftDeletedAt(deletedAt string) UserInterface

	Timezone() string
	SetTimezone(timezone string) UserInterface

	Status() string
	SetStatus(status string) UserInterface

	PasswordCompare(password string) bool

	UpdatedAt() string
	UpdatedAtCarbon() carbon.Carbon
	SetUpdatedAt(updatedAt string) UserInterface
}

type UserQueryInterface interface {
	ID() string
	SetID(id string) (UserQueryInterface, error)

	IDIn() []string
	SetIDIn(idIn []string) (UserQueryInterface, error)

	Status() string
	SetStatus(status string) (UserQueryInterface, error)

	StatusIn() []string
	SetStatusIn(statusIn []string) (UserQueryInterface, error)

	Email() string
	SetEmail(email string) (UserQueryInterface, error)

	CreatedAtGte() string
	SetCreatedAtGte(createdAtGte string) (UserQueryInterface, error)

	CreatedAtLte() string
	SetCreatedAtLte(createdAtLte string) (UserQueryInterface, error)

	Offset() int
	SetOffset(offset int) (UserQueryInterface, error)

	Limit() int
	SetLimit(limit int) (UserQueryInterface, error)

	SortOrder() string
	SetSortOrder(sortOrder string) (UserQueryInterface, error)

	OrderBy() string
	SetOrderBy(orderBy string) (UserQueryInterface, error)

	CountOnly() bool
	SetCountOnly(countOnly bool) UserQueryInterface

	WithSoftDeleted() bool
	SetWithSoftDeleted(withDeleted bool) UserQueryInterface
}
