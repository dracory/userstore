package userstore

import (
	"context"
	"database/sql"

	"github.com/dromara/carbon/v2"
)

// StoreInterface defines the interface for the user store
type StoreInterface interface {
	storeBaseInterface
	storeUserInterface
	storeRoleInterface
	storeUserRoleInterface
	storeGroupInterface
	storeUserGroupInterface
}

// storeBaseInterface defines the base store methods
type storeBaseInterface interface {
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

	// GetGroupTableName returns the group table name
	GetGroupTableName() string
	// SetGroupTableName sets the group table name
	SetGroupTableName(tableName string)

	// GetUserGroupTableName returns the user group table name
	GetUserGroupTableName() string
	// SetUserGroupTableName sets the user group table name
	SetUserGroupTableName(tableName string)

	// MigrateDown drops the user table
	MigrateDown(ctx context.Context, tx ...*sql.Tx) error
	// MigrateUp creates the user table
	MigrateUp(ctx context.Context, tx ...*sql.Tx) error

	// EnableDebug enables debug mode
	EnableDebug(debug bool)

	// GetDB returns the database connection
	GetDB() *sql.DB
}

// storeUserInterface defines the user store methods
type storeUserInterface interface {
	// UserCreate creates a new user
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

// storeRoleInterface defines the role store methods
type storeRoleInterface interface {
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
}

// storeUserRoleInterface defines the user role store methods
type storeUserRoleInterface interface {
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
	// UserRoles returns the roles assigned to a user
	UserRoles(ctx context.Context, userID string) ([]RoleInterface, error)
	// UserHasRoles returns true if the user has all the given roles
	UserHasRoles(ctx context.Context, userID string, roleIDs []string) (bool, error)
	// UserRoleSoftDelete soft deletes a user role
	UserRoleSoftDelete(ctx context.Context, userRole UserRoleInterface) error
	// UserRoleSoftDeleteByID soft deletes a user role by ID
	UserRoleSoftDeleteByID(ctx context.Context, id string) error
	// UserRoleUpdate updates a user role
	UserRoleUpdate(ctx context.Context, userRole UserRoleInterface) error
}

// storeGroupInterface defines the group store methods
type storeGroupInterface interface {
	// GroupCreate creates a new group
	GroupCreate(ctx context.Context, group GroupInterface) error
	// GroupCount returns the number of groups
	GroupCount(ctx context.Context, options GroupQueryInterface) (int64, error)
	// GroupDelete deletes a group
	GroupDelete(ctx context.Context, group GroupInterface) error
	// GroupDeleteByID deletes a group by ID
	GroupDeleteByID(ctx context.Context, id string) error
	// GroupFindByHandle finds a group by handle
	GroupFindByHandle(ctx context.Context, handle string) (GroupInterface, error)
	// GroupFindByHandleOrCreate finds a group by handle or creates it
	GroupFindByHandleOrCreate(ctx context.Context, handle, createStatus string) (GroupInterface, error)
	// GroupFindByID finds a group by ID
	GroupFindByID(ctx context.Context, id string) (GroupInterface, error)
	// GroupList returns a list of groups
	GroupList(ctx context.Context, query GroupQueryInterface) ([]GroupInterface, error)
	// GroupSoftDelete soft deletes a group
	GroupSoftDelete(ctx context.Context, group GroupInterface) error
	// GroupSoftDeleteByID soft deletes a group by ID
	GroupSoftDeleteByID(ctx context.Context, id string) error
	// GroupUpdate updates a group
	GroupUpdate(ctx context.Context, group GroupInterface) error
}

// storeUserGroupInterface defines the user group store methods
type storeUserGroupInterface interface {
	// UserGroupCreate creates a new user group
	UserGroupCreate(ctx context.Context, userGroup UserGroupInterface) error
	// UserGroupCount returns the number of user groups
	UserGroupCount(ctx context.Context, options UserGroupQueryInterface) (int64, error)
	// UserGroupDelete deletes a user group
	UserGroupDelete(ctx context.Context, userGroup UserGroupInterface) error
	// UserGroupDeleteByID deletes a user group by ID
	UserGroupDeleteByID(ctx context.Context, id string) error
	// UserGroupFindByID finds a user group by ID
	UserGroupFindByID(ctx context.Context, id string) (UserGroupInterface, error)
	// UserGroupFindByUserIDAndGroupID finds a user group by user ID and group ID
	UserGroupFindByUserIDAndGroupID(ctx context.Context, userID, groupID string) (UserGroupInterface, error)
	// UserGroupFindByUserIDAndGroupIDOrCreate finds a user group by user ID and group ID or creates it
	UserGroupFindByUserIDAndGroupIDOrCreate(ctx context.Context, userID, groupID string) (UserGroupInterface, error)
	// UserGroupList returns a list of user groups
	UserGroupList(ctx context.Context, query UserGroupQueryInterface) ([]UserGroupInterface, error)
	// UserGroups returns the groups a user belongs to
	UserGroups(ctx context.Context, userID string) ([]GroupInterface, error)
	// UserHasGroups returns true if the user belongs to all the given groups
	UserHasGroups(ctx context.Context, userID string, groupIDs []string) (bool, error)
	// UserGroupSoftDelete soft deletes a user group
	UserGroupSoftDelete(ctx context.Context, userGroup UserGroupInterface) error
	// UserGroupSoftDeleteByID soft deletes a user group by ID
	UserGroupSoftDeleteByID(ctx context.Context, id string) error
	// UserGroupUpdate updates a user group
	UserGroupUpdate(ctx context.Context, userGroup UserGroupInterface) error
}

// RoleInterface defines the interface for a role
type RoleInterface interface {
	// from dataobject

	// Data returns all fields as a map
	Data() map[string]string
	// DataChanged returns the changed fields as a map
	DataChanged() map[string]string
	// MarkAsNotDirty marks all fields as not changed
	MarkAsNotDirty()
	// ToMap returns a DB-ready map of the role
	ToMap() map[string]any

	// methods

	// IsActive returns true if the role is active
	IsActive() bool
	// IsInactive returns true if the role is inactive
	IsInactive() bool
	// IsSoftDeleted returns true if the role is soft deleted
	IsSoftDeleted() bool

	// setters and getters

	// GetCreatedAt returns the created at datetime
	GetCreatedAt() string
	// GetCreatedAtCarbon returns the created at datetime as carbon
	GetCreatedAtCarbon() *carbon.Carbon
	// SetCreatedAt sets the created at datetime
	SetCreatedAt(createdAt string) RoleInterface

	// GetHandle returns the handle
	GetHandle() string
	// SetHandle sets the handle
	SetHandle(handle string) RoleInterface

	// GetID returns the ID
	GetID() string
	// SetID sets the ID
	SetID(id string) RoleInterface

	// GetName returns the name
	GetName() string
	// SetName sets the name
	SetName(name string) RoleInterface

	// GetMemo returns the memo
	GetMemo() string
	// SetMemo sets the memo
	SetMemo(memo string) RoleInterface

	// GetMeta returns the value of the meta with the given name
	GetMeta(name string) string
	// SetMeta sets the value of the meta with the given name
	SetMeta(name string, value string) error
	// GetMetas returns all metas
	GetMetas() (map[string]string, error)
	// SetMetas sets all metas, overwriting existing ones
	SetMetas(metas map[string]string) error

	// GetStatus returns the status
	GetStatus() string
	// SetStatus sets the status
	SetStatus(status string) RoleInterface

	// GetSoftDeletedAt returns the soft deleted at datetime
	GetSoftDeletedAt() string
	// GetSoftDeletedAtCarbon returns the soft deleted at datetime as carbon
	GetSoftDeletedAtCarbon() *carbon.Carbon
	// SetSoftDeletedAt sets the soft deleted at datetime
	SetSoftDeletedAt(softDeletedAt string) RoleInterface

	// GetUpdatedAt returns the updated at datetime
	GetUpdatedAt() string
	// GetUpdatedAtCarbon returns the updated at datetime as carbon
	GetUpdatedAtCarbon() *carbon.Carbon
	// SetUpdatedAt sets the updated at datetime
	SetUpdatedAt(updatedAt string) RoleInterface
}

// UserRoleInterface defines the interface for a user role assignment
type UserRoleInterface interface {
	// from dataobject

	// Data returns all fields as a map
	Data() map[string]string
	// DataChanged returns the changed fields as a map
	DataChanged() map[string]string
	// MarkAsNotDirty marks all fields as not changed
	MarkAsNotDirty()
	// ToMap returns a DB-ready map of the user role
	ToMap() map[string]any

	// methods

	// IsSoftDeleted returns true if the user role is soft deleted
	IsSoftDeleted() bool

	// setters and getters

	// GetCreatedAt returns the created at datetime
	GetCreatedAt() string
	// GetCreatedAtCarbon returns the created at datetime as carbon
	GetCreatedAtCarbon() *carbon.Carbon
	// SetCreatedAt sets the created at datetime
	SetCreatedAt(createdAt string) UserRoleInterface

	// GetID returns the ID
	GetID() string
	// SetID sets the ID
	SetID(id string) UserRoleInterface

	// GetUserID returns the user ID
	GetUserID() string
	// SetUserID sets the user ID
	SetUserID(userID string) UserRoleInterface

	// GetRoleID returns the role ID
	GetRoleID() string
	// SetRoleID sets the role ID
	SetRoleID(roleID string) UserRoleInterface

	// GetSoftDeletedAt returns the soft deleted at datetime
	GetSoftDeletedAt() string
	// GetSoftDeletedAtCarbon returns the soft deleted at datetime as carbon
	GetSoftDeletedAtCarbon() *carbon.Carbon
	// SetSoftDeletedAt sets the soft deleted at datetime
	SetSoftDeletedAt(softDeletedAt string) UserRoleInterface

	// GetUpdatedAt returns the updated at datetime
	GetUpdatedAt() string
	// GetUpdatedAtCarbon returns the updated at datetime as carbon
	GetUpdatedAtCarbon() *carbon.Carbon
	// SetUpdatedAt sets the updated at datetime
	SetUpdatedAt(updatedAt string) UserRoleInterface
}

// UserInterface defines the interface for a user
type UserInterface interface {
	// from dataobject

	// Data returns all fields as a map
	Data() map[string]string
	// DataChanged returns the changed fields as a map
	DataChanged() map[string]string
	// MarkAsNotDirty marks all fields as not changed
	MarkAsNotDirty()
	// Get returns the value of the specified column
	Get(columnName string) string
	// Set sets the value of the specified column
	Set(columnName string, value string)
	// ToMap returns a DB-ready map of the user
	ToMap() map[string]any

	// methods

	// IsActive returns true if the user is active
	IsActive() bool
	// IsInactive returns true if the user is inactive
	IsInactive() bool
	// IsSoftDeleted returns true if the user is soft deleted
	IsSoftDeleted() bool
	// IsUnverified returns true if the user is unverified
	IsUnverified() bool

	// IsAdministrator returns true if the user has the administrator role
	IsAdministrator() bool
	// IsManager returns true if the user has the manager role
	IsManager() bool
	// IsSuperuser returns true if the user has the superuser role
	IsSuperuser() bool

	// IsRegistrationCompleted returns true if the user has completed registration
	IsRegistrationCompleted() bool

	// setters and getters

	// GetBusinessName returns the business name
	GetBusinessName() string
	// SetBusinessName sets the business name
	SetBusinessName(businessName string) UserInterface

	// GetCountry returns the country
	GetCountry() string
	// SetCountry sets the country
	SetCountry(country string) UserInterface

	// GetCreatedAt returns the created at datetime
	GetCreatedAt() string
	// GetCreatedAtCarbon returns the created at datetime as carbon
	GetCreatedAtCarbon() *carbon.Carbon
	// SetCreatedAt sets the created at datetime
	SetCreatedAt(createdAt string) UserInterface

	// GetEmail returns the email
	GetEmail() string
	// SetEmail sets the email
	SetEmail(email string) UserInterface

	// GetID returns the ID
	GetID() string
	// SetID sets the ID
	SetID(id string) UserInterface

	// GetFirstName returns the first name
	GetFirstName() string
	// SetFirstName sets the first name
	SetFirstName(firstName string) UserInterface

	// GetLastName returns the last name
	GetLastName() string
	// SetLastName sets the last name
	SetLastName(lastName string) UserInterface

	// GetMemo returns the memo
	GetMemo() string
	// SetMemo sets the memo
	SetMemo(memo string) UserInterface

	// GetMeta returns the value of the meta with the given name
	GetMeta(name string) string
	// SetMeta sets the value of the meta with the given name
	SetMeta(name string, value string) error
	// GetMetas returns all metas
	GetMetas() (map[string]string, error)
	// SetMetas sets all metas, overwriting existing ones
	SetMetas(metas map[string]string) error
	// UpsertMetas merges the given metas with the existing ones
	UpsertMetas(metas map[string]string) error

	// GetMiddleNames returns the middle names
	GetMiddleNames() string
	// SetMiddleNames sets the middle names
	SetMiddleNames(middleNames string) UserInterface

	// GetPassword returns the password hash
	GetPassword() string
	// PasswordCompare compares the given password with the stored hash
	PasswordCompare(password string) bool
	// SetPassword sets the password hash
	SetPassword(password string) UserInterface
	// SetPasswordAndHash hashes and sets the given plain text password
	SetPasswordAndHash(password string) error

	// GetPhone returns the phone
	GetPhone() string
	// SetPhone sets the phone
	SetPhone(phone string) UserInterface

	// GetProfileImageUrl returns the profile image URL
	GetProfileImageUrl() string
	// ProfileImageOrDefaultUrl returns the profile image URL or the default URL
	ProfileImageOrDefaultUrl() string
	// SetProfileImageUrl sets the profile image URL
	SetProfileImageUrl(profileImageUrl string) UserInterface

	// GetRole returns the role
	GetRole() string
	// SetRole sets the role
	SetRole(role string) UserInterface

	// GetSoftDeletedAt returns the soft deleted at datetime
	GetSoftDeletedAt() string
	// GetSoftDeletedAtCarbon returns the soft deleted at datetime as carbon
	GetSoftDeletedAtCarbon() *carbon.Carbon
	// SetSoftDeletedAt sets the soft deleted at datetime
	SetSoftDeletedAt(deletedAt string) UserInterface

	// GetTimezone returns the timezone
	GetTimezone() string
	// SetTimezone sets the timezone
	SetTimezone(timezone string) UserInterface

	// GetStatus returns the status
	GetStatus() string
	// SetStatus sets the status
	SetStatus(status string) UserInterface

	// GetUpdatedAt returns the updated at datetime
	GetUpdatedAt() string
	// GetUpdatedAtCarbon returns the updated at datetime as carbon
	GetUpdatedAtCarbon() *carbon.Carbon
	// SetUpdatedAt sets the updated at datetime
	SetUpdatedAt(updatedAt string) UserInterface
}

// GroupInterface defines the interface for a group
type GroupInterface interface {
	// from dataobject

	// Data returns all fields as a map
	Data() map[string]string
	// DataChanged returns the changed fields as a map
	DataChanged() map[string]string
	// MarkAsNotDirty marks all fields as not changed
	MarkAsNotDirty()
	// ToMap returns a DB-ready map of the group
	ToMap() map[string]any

	// methods

	// IsActive returns true if the group is active
	IsActive() bool
	// IsInactive returns true if the group is inactive
	IsInactive() bool
	// IsSoftDeleted returns true if the group is soft deleted
	IsSoftDeleted() bool

	// setters and getters

	// GetCreatedAt returns the created at datetime
	GetCreatedAt() string
	// GetCreatedAtCarbon returns the created at datetime as carbon
	GetCreatedAtCarbon() *carbon.Carbon
	// SetCreatedAt sets the created at datetime
	SetCreatedAt(createdAt string) GroupInterface

	// GetHandle returns the handle
	GetHandle() string
	// SetHandle sets the handle
	SetHandle(handle string) GroupInterface

	// GetID returns the ID
	GetID() string
	// SetID sets the ID
	SetID(id string) GroupInterface

	// GetName returns the name
	GetName() string
	// SetName sets the name
	SetName(name string) GroupInterface

	// GetMemo returns the memo
	GetMemo() string
	// SetMemo sets the memo
	SetMemo(memo string) GroupInterface

	// GetMeta returns the value of the meta with the given name
	GetMeta(name string) string
	// SetMeta sets the value of the meta with the given name
	SetMeta(name string, value string) error
	// GetMetas returns all metas
	GetMetas() (map[string]string, error)
	// SetMetas sets all metas, overwriting existing ones
	SetMetas(metas map[string]string) error
	// UpsertMetas merges the given metas with the existing ones
	UpsertMetas(metas map[string]string) error

	// GetStatus returns the status
	GetStatus() string
	// SetStatus sets the status
	SetStatus(status string) GroupInterface

	// GetSoftDeletedAt returns the soft deleted at datetime
	GetSoftDeletedAt() string
	// GetSoftDeletedAtCarbon returns the soft deleted at datetime as carbon
	GetSoftDeletedAtCarbon() *carbon.Carbon
	// SetSoftDeletedAt sets the soft deleted at datetime
	SetSoftDeletedAt(softDeletedAt string) GroupInterface

	// GetUpdatedAt returns the updated at datetime
	GetUpdatedAt() string
	// GetUpdatedAtCarbon returns the updated at datetime as carbon
	GetUpdatedAtCarbon() *carbon.Carbon
	// SetUpdatedAt sets the updated at datetime
	SetUpdatedAt(updatedAt string) GroupInterface
}

// UserGroupInterface defines the interface for a user group membership
type UserGroupInterface interface {
	// from dataobject

	// Data returns all fields as a map
	Data() map[string]string
	// DataChanged returns the changed fields as a map
	DataChanged() map[string]string
	// MarkAsNotDirty marks all fields as not changed
	MarkAsNotDirty()
	// ToMap returns a DB-ready map of the user group
	ToMap() map[string]any

	// methods

	// IsSoftDeleted returns true if the user group is soft deleted
	IsSoftDeleted() bool

	// setters and getters

	// GetCreatedAt returns the created at datetime
	GetCreatedAt() string
	// GetCreatedAtCarbon returns the created at datetime as carbon
	GetCreatedAtCarbon() *carbon.Carbon
	// SetCreatedAt sets the created at datetime
	SetCreatedAt(createdAt string) UserGroupInterface

	// GetID returns the ID
	GetID() string
	// SetID sets the ID
	SetID(id string) UserGroupInterface

	// GetUserID returns the user ID
	GetUserID() string
	// SetUserID sets the user ID
	SetUserID(userID string) UserGroupInterface

	// GetGroupID returns the group ID
	GetGroupID() string
	// SetGroupID sets the group ID
	SetGroupID(groupID string) UserGroupInterface

	// GetSoftDeletedAt returns the soft deleted at datetime
	GetSoftDeletedAt() string
	// GetSoftDeletedAtCarbon returns the soft deleted at datetime as carbon
	GetSoftDeletedAtCarbon() *carbon.Carbon
	// SetSoftDeletedAt sets the soft deleted at datetime
	SetSoftDeletedAt(softDeletedAt string) UserGroupInterface

	// GetUpdatedAt returns the updated at datetime
	GetUpdatedAt() string
	// GetUpdatedAtCarbon returns the updated at datetime as carbon
	GetUpdatedAtCarbon() *carbon.Carbon
	// SetUpdatedAt sets the updated at datetime
	SetUpdatedAt(updatedAt string) UserGroupInterface
}
