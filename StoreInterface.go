package userstore

type StoreInterface interface {
	AutoMigrate() error
	EnableDebug(debug bool)
	UserCreate(user UserInterface) error
	UserDelete(user UserInterface) error
	UserDeleteByID(id string) error
	UserFindByID(userID string) (UserInterface, error)
	UserList(options UserQueryOptions) ([]UserInterface, error)
	UserSoftDelete(user UserInterface) error
	UserSoftDeleteByID(id string) error
	UserUpdate(user UserInterface) error
}
