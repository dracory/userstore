# UserStore <a href="https://gitpod.io/#https://github.com/dracory/userstore" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>


[![Tests Status](https://github.com/dracory/userstore/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/dracory/userstore/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dracory/userstore)](https://goreportcard.com/report/github.com/dracory/userstore)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/dracory/userstore)](https://pkg.go.dev/github.com/dracory/userstore)

UserStore is a robust user management package.

Supports multiple database storages (SQLite, MySQL, or PostgreSQL)

## Features

- User and Role management
- Soft delete support
- Meta data storage for custom fields
- Password hashing and verification
- Method chaining for fluent API
- Transaction support
- Query builder for complex searches

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

err := userStore.UserCreate(user)

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
    SetStatus(userstore.USER_STATUS_ACTIVE)

err := userStore.RoleCreate(role)

if err != nil {
	return errors.New("role failed to create")
}
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
