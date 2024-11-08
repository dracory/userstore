package userstore

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
