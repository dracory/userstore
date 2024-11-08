package userstore

type UserQueryOptions struct {
	ID           string
	IDIn         []string
	Status       string
	StatusIn     []string
	Email        string
	CreatedAtGte string
	CreatedAtLte string
	Offset       int
	Limit        int
	SortOrder    string
	OrderBy      string
	CountOnly    bool
	WithDeleted  bool
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
