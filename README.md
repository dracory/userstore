# UserStore


[![Tests Status](https://github.com/dracory/userstore/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/dracory/userstore/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dracory/userstore)](https://goreportcard.com/report/github.com/dracory/userstore)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/dracory/userstore)](https://pkg.go.dev/github.com/dracory/userstore)

UserStore is a robust user management package.

Supports multiple database storages (SQLite, MySQL, or PostgreSQL)

## Features

- User, Role, and Group management
- Soft delete support
- Meta data storage for custom fields
- Password hashing and verification
- Method chaining for fluent API
- Transaction support
- Query builder for complex searches

## Documentation

View the complete documentation at: https://htmlpreview.github.io/?https://raw.githubusercontent.com/dracory/userstore/main/docs/livewiki-html/index.html

## License

This project is licensed under the GNU General Public License version 3 (GPL-3.0). You can find a copy of the license at https://www.gnu.org/licenses/gpl-3.0.en.html

For commercial use, please use my [contact page](https://lesichkov.co.uk/contact) to obtain a commercial license.

## Installation

```
go get github.com/dracory/userstore
```

## Setup

```golang
userStore, err = userstore.NewStore(userstore.NewStoreOptions{
	DB:                 databaseInstance,
    UserTableName:      "user",
	RoleTableName:      "role",       // required when RolesEnabled is true
	UserRoleTableName:  "user_role",  // required when RolesEnabled is true
	RolesEnabled:       true,
	GroupTableName:     "groups",     // required when GroupsEnabled is true ("group" is a reserved SQL keyword)
	UserGroupTableName: "user_group", // required when GroupsEnabled is true
	GroupsEnabled:      true,
	AutomigrateEnabled: true,
	DebugEnabled:       false,
})

if err != nil {
	return errors.Join(errors.New("userstore.NewStore"), err)
}
```

## API Convention

All getter methods use the `Get` prefix and all setter methods use the `Set` prefix:

- **Getters**: `GetID()`, `GetName()`, `GetEmail()`, `GetStatus()`, etc.
- **Setters**: `SetID()`, `SetName()`, `SetEmail()`, `SetStatus()`, etc.

## Examples

### Creating a User

```golang
user := userstore.NewUser().
    SetStatus(userstore.USER_STATUS_ACTIVE).
    SetFirstName("John").
    SetLastName("Doe").
    SetEmail("test@test.com")

err := userStore.UserCreate(context.Background(), user)

if err != nil {
	return errors.New("user failed to create")
}
```

### Reading User Properties

```golang
id := user.GetID()
email := user.GetEmail()
firstName := user.GetFirstName()
lastName := user.GetLastName()
status := user.GetStatus()
```

### Creating a Role

```golang
role := userstore.NewRole().
    SetName("Administrator").
    SetHandle("admin").
    SetStatus(userstore.ROLE_STATUS_ACTIVE)

err := userStore.RoleCreate(context.Background(), role)

if err != nil {
	return errors.New("role failed to create")
}
```

### Assigning a Role to a User

```golang
userRole := userstore.NewUserRole().
    SetUserID(user.GetID()).
    SetRoleID(role.GetID())

err := userStore.UserRoleCreate(context.Background(), userRole)

if err != nil {
	return errors.New("user role failed to create")
}

// Or find-or-create the assignment
userRole, err := userStore.UserRoleFindByUserIDAndRoleIDOrCreate(
    context.Background(), user.GetID(), role.GetID())

// List a user's roles
roles, err := userStore.UserRoles(context.Background(), user.GetID())

// Check if a user has one or more roles (all must match)
hasRoles, err := userStore.UserHasRoles(
    context.Background(), user.GetID(), []string{role.GetID()})
```

### Creating a Group

```golang
group := userstore.NewGroup().
    SetName("Staff").
    SetHandle("staff").
    SetStatus(userstore.GROUP_STATUS_ACTIVE)

err := userStore.GroupCreate(context.Background(), group)

if err != nil {
	return errors.New("group failed to create")
}
```

### Adding a User to a Group

```golang
userGroup := userstore.NewUserGroup().
    SetUserID(user.GetID()).
    SetGroupID(group.GetID())

err := userStore.UserGroupCreate(context.Background(), userGroup)

if err != nil {
	return errors.New("user group failed to create")
}

// Or find-or-create the membership
userGroup, err := userStore.UserGroupFindByUserIDAndGroupIDOrCreate(
    context.Background(), user.GetID(), group.GetID())

// List a user's groups
groups, err := userStore.UserGroups(context.Background(), user.GetID())

// Check if a user belongs to one or more groups (all must match)
isMember, err := userStore.UserHasGroups(
    context.Background(), user.GetID(), []string{group.GetID()})
```

### Finding Users

```golang
// Find by ID
user, err := userStore.UserFindByID(context.Background(), userID)

// List users with query
query := userstore.NewUserQuery().
    SetStatus(userstore.USER_STATUS_ACTIVE).
    SetLimit(10)

users, err := userStore.UserList(context.Background(), query)
```

### Updating Users

```golang
user.SetFirstName("Jane")
user.SetEmail("jane@example.com")

err := userStore.UserUpdate(context.Background(), user)
```

### Soft Deleting Users

```golang
err := userStore.UserSoftDelete(context.Background(), user)
```

## Runnable Examples

See the `examples/` directory for complete runnable programs:

- `examples/basic` — creating, finding, updating, and soft deleting users
- `examples/roles` — creating roles and assigning them to users
- `examples/groups` — creating groups and managing memberships

Run any of them with:

```
go run ./examples/basic
```
