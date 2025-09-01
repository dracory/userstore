# UserStore <a href="https://gitpod.io/#https://github.com/dracory/userstore" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>


[![Tests Status](https://github.com/dracory/userstore/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/dracory/userstore/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dracory/userstore)](https://goreportcard.com/report/github.com/dracory/userstore)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/dracory/userstore)](https://pkg.go.dev/github.com/dracory/userstore)

UserStore is a robust user management package.

Supports multiple database storages (SQLite, MySQL, or PostgreSQL)

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

## Examples

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
