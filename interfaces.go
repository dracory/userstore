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

	// GetRoleTableName returns the role table name
	GetRoleTableName() string
	// SetRoleTableName sets the role table name
	SetRoleTableName(tableName string)

	// GetUserRoleTableName returns the user role table name
	GetUserRoleTableName() string
	// SetUserRoleTableName sets the user role table name
	SetUserRoleTableName(tableName string)

	// MigrateDown drops the user table
	MigrateDown(ctx context.Context, tx ...*sql.Tx) error
	// MigrateUp creates the user table
	MigrateUp(ctx context.Context, tx ...*sql.Tx) error

	// EnableDebug enables debug mode
	EnableDebug(debug bool)

	// GetDB returns the database connection
	GetDB() *sql.DB

	// RoleCreate creates a new role
	RoleCreate(ctx context.Context, role RoleInterface) error
	// RoleCount returns the number of roles
	RoleCount(ctx context.Context, options RoleQueryInterface) (int64, error)
	// RoleDelete deletes a role
	RoleDelete(ctx context.Context, role RoleInterface) error
	// RoleDeleteByID deletes a role by ID
	RoleDeleteByID(ctx context.Context, id string) error
	// RoleFindByHandle finds a role by handle
	RoleFindByHandle(ctx context.Context, handle string) (RoleInterface, error)
	// RoleFindByHandleOrCreate finds a role by handle or creates it
	RoleFindByHandleOrCreate(ctx context.Context, handle, createStatus string) (RoleInterface, error)
	// RoleFindByID finds a role by ID
	RoleFindByID(ctx context.Context, id string) (RoleInterface, error)
	// RoleList returns a list of roles
	RoleList(ctx context.Context, query RoleQueryInterface) ([]RoleInterface, error)
	// RoleSoftDelete soft deletes a role
	RoleSoftDelete(ctx context.Context, role RoleInterface) error
	// RoleSoftDeleteByID soft deletes a role by ID
	RoleSoftDeleteByID(ctx context.Context, id string) error
	// RoleUpdate updates a role
	RoleUpdate(ctx context.Context, role RoleInterface) error

	// UserRoleCreate creates a new user role
	UserRoleCreate(ctx context.Context, userRole UserRoleInterface) error
	// UserRoleCount returns the number of user roles
	UserRoleCount(ctx context.Context, options UserRoleQueryInterface) (int64, error)
	// UserRoleDelete deletes a user role
	UserRoleDelete(ctx context.Context, userRole UserRoleInterface) error
	// UserRoleDeleteByID deletes a user role by ID
	UserRoleDeleteByID(ctx context.Context, id string) error
	// UserRoleFindByID finds a user role by ID
	UserRoleFindByID(ctx context.Context, id string) (UserRoleInterface, error)
	// UserRoleFindByUserIDAndRoleID finds a user role by user ID and role ID
	UserRoleFindByUserIDAndRoleID(ctx context.Context, userID, roleID string) (UserRoleInterface, error)
	// UserRoleFindByUserIDAndRoleIDOrCreate finds a user role by user ID and role ID or creates it
	UserRoleFindByUserIDAndRoleIDOrCreate(ctx context.Context, userID, roleID string) (UserRoleInterface, error)
	// UserRoleList returns a list of user roles
	UserRoleList(ctx context.Context, query UserRoleQueryInterface) ([]UserRoleInterface, error)
	// UserRoleSoftDelete soft deletes a user role
	UserRoleSoftDelete(ctx context.Context, userRole UserRoleInterface) error
	// UserRoleSoftDeleteByID soft deletes a user role by ID
	UserRoleSoftDeleteByID(ctx context.Context, id string) error
	// UserRoleUpdate updates a user role
	UserRoleUpdate(ctx context.Context, userRole UserRoleInterface) error

	UserCreate(ctx context.Context, user UserInterface) error
	// UserCount returns the number of users
	UserCount(ctx context.Context, options UserQueryInterface) (int64, error)
	// UserDelete deletes a user
	UserDelete(ctx context.Context, user UserInterface) error
	// UserDeleteByID deletes a user by ID
	UserDeleteByID(ctx context.Context, id string) error
	// UserFindByEmail finds a user by email
	UserFindByEmail(ctx context.Context, email string) (UserInterface, error)
	// UserFindByEmailOrCreate finds a user by email or creates it
	UserFindByEmailOrCreate(ctx context.Context, email, createStatus string) (UserInterface, error)
	// UserFindByID finds a user by ID
	UserFindByID(ctx context.Context, userID string) (UserInterface, error)
	// UserList returns a list of users
	UserList(ctx context.Context, query UserQueryInterface) ([]UserInterface, error)
	// UserSoftDelete soft deletes a user
	UserSoftDelete(ctx context.Context, user UserInterface) error
	// UserSoftDeleteByID soft deletes a user by ID
	UserSoftDeleteByID(ctx context.Context, id string) error
	// UserUpdate updates a user
	UserUpdate(ctx context.Context, user UserInterface) error
}

type RoleInterface interface {
	// from dataobject

	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()
	ToMap() map[string]any

	// methods

	// IsActive returns true if the role is active
	IsActive() bool
	// IsInactive returns true if the role is inactive
	IsInactive() bool
	// IsSoftDeleted returns true if the role is soft deleted
	IsSoftDeleted() bool

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

type UserRoleInterface interface {
	// from dataobject

	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()
	ToMap() map[string]any

	// methods

	// IsSoftDeleted returns true if the user role is soft deleted
	IsSoftDeleted() bool

	// setters and getters

	GetCreatedAt() string
	GetCreatedAtCarbon() *carbon.Carbon
	SetCreatedAt(createdAt string) UserRoleInterface

	GetID() string
	SetID(id string) UserRoleInterface

	GetUserID() string
	SetUserID(userID string) UserRoleInterface

	GetRoleID() string
	SetRoleID(roleID string) UserRoleInterface

	GetSoftDeletedAt() string
	GetSoftDeletedAtCarbon() *carbon.Carbon
	SetSoftDeletedAt(softDeletedAt string) UserRoleInterface

	GetUpdatedAt() string
	GetUpdatedAtCarbon() *carbon.Carbon
	SetUpdatedAt(updatedAt string) UserRoleInterface
}

type UserInterface interface {
	// from dataobject

	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()
	Get(columnName string) string
	Set(columnName string, value string)
	ToMap() map[string]any

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
	ProfileImageOrDefaultUrl() string
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
