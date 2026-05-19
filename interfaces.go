package userstore

import (
	"context"
	"database/sql"

	"github.com/dromara/carbon/v2"
)

type StoreInterface interface {
	// GetUserTableName returns the user table name
	GetUserTableName() string
	// SetUserTableName sets the user table name
	SetUserTableName(tableName string)

	// MigrateDown drops the user table
	MigrateDown(tx ...*sql.Tx) error
	// MigrateUp creates the user table
	MigrateUp(tx ...*sql.Tx) error

	EnableDebug(debug bool)
	GetDB() *sql.DB

	// RoleCreate(ctx context.Context, role RoleInterface) error
	// RoleDelete(ctx context.Context, role RoleInterface) error
	// RoleDeleteByID(ctx context.Context, id string) error
	// RoleFindByID(ctx context.Context, id string) (RoleInterface, error)
	// RoleList(ctx context.Context, query RoleQueryInterface) ([]RoleInterface, error)
	// RoleSoftDelete(ctx context.Context, role RoleInterface) error
	// RoleSoftDeleteByID(ctx context.Context, id string) error
	// RoleUpdate(ctx context.Context, role RoleInterface) error

	UserCreate(ctx context.Context, user UserInterface) error
	UserCount(ctx context.Context, options UserQueryInterface) (int64, error)
	UserDelete(ctx context.Context, user UserInterface) error
	UserDeleteByID(ctx context.Context, id string) error
	UserFindByEmail(ctx context.Context, email string) (UserInterface, error)
	UserFindByID(ctx context.Context, userID string) (UserInterface, error)
	UserList(ctx context.Context, query UserQueryInterface) ([]UserInterface, error)
	UserSoftDelete(ctx context.Context, user UserInterface) error
	UserSoftDeleteByID(ctx context.Context, id string) error
	UserUpdate(ctx context.Context, user UserInterface) error
}

type RoleInterface interface {
	// from dataobject

	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()

	// setters and getters

	GetCreatedAt() string
	GetCreatedAtCarbon() *carbon.Carbon
	SetCreatedAt(createdAt string) RoleInterface

	GetHandle() string
	SetHandle(handle string) RoleInterface

	GetID() string
	SetID(id string) RoleInterface

	GetName() string
	SetName(name string) RoleInterface

	GetMemo() string
	SetMemo(memo string) RoleInterface

	GetMeta(name string) string
	SetMeta(name string, value string) error
	GetMetas() (map[string]string, error)
	SetMetas(metas map[string]string) error

	GetStatus() string
	SetStatus(status string) RoleInterface

	GetSoftDeletedAt() string
	GetSoftDeletedAtCarbon() *carbon.Carbon
	SetSoftDeletedAt(softDeletedAt string) RoleInterface

	GetUpdatedAt() string
	GetUpdatedAtCarbon() *carbon.Carbon
	SetUpdatedAt(updatedAt string) RoleInterface
}

type UserInterface interface {
	// from dataobject

	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()
	Get(columnName string) string
	Set(columnName string, value string)

	// methods

	IsActive() bool
	IsInactive() bool
	IsSoftDeleted() bool
	IsUnverified() bool

	IsAdministrator() bool
	IsManager() bool
	IsSuperuser() bool

	IsRegistrationCompleted() bool

	// setters and getters

	GetBusinessName() string
	SetBusinessName(businessName string) UserInterface

	GetCountry() string
	SetCountry(country string) UserInterface

	GetCreatedAt() string
	GetCreatedAtCarbon() *carbon.Carbon
	SetCreatedAt(createdAt string) UserInterface

	GetEmail() string
	SetEmail(email string) UserInterface

	GetID() string
	SetID(id string) UserInterface

	GetFirstName() string
	SetFirstName(firstName string) UserInterface

	GetLastName() string
	SetLastName(lastName string) UserInterface

	GetMemo() string
	SetMemo(memo string) UserInterface

	GetMeta(name string) string
	SetMeta(name string, value string) error
	GetMetas() (map[string]string, error)
	SetMetas(metas map[string]string) error
	UpsertMetas(metas map[string]string) error

	GetMiddleNames() string
	SetMiddleNames(middleNames string) UserInterface

	GetPassword() string
	PasswordCompare(password string) bool
	SetPassword(password string) UserInterface
	SetPasswordAndHash(password string) error

	GetPhone() string
	SetPhone(phone string) UserInterface

	GetProfileImageUrl() string
	SetProfileImageUrl(profileImageUrl string) UserInterface

	GetRole() string
	SetRole(role string) UserInterface

	GetSoftDeletedAt() string
	GetSoftDeletedAtCarbon() *carbon.Carbon
	SetSoftDeletedAt(deletedAt string) UserInterface

	GetTimezone() string
	SetTimezone(timezone string) UserInterface

	GetStatus() string
	SetStatus(status string) UserInterface

	GetUpdatedAt() string
	GetUpdatedAtCarbon() *carbon.Carbon
	SetUpdatedAt(updatedAt string) UserInterface
}
