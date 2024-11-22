package userstore

import "errors"

type userQuery struct {
	id              string
	idIn            []string
	status          string
	statusIn        []string
	email           string
	createdAtGte    string
	createdAtLte    string
	countOnly       bool
	offset          int64
	limit           int
	sortOrder       string
	orderBy         string
	withSoftDeleted bool
}

func NewUserQuery() UserQueryInterface {
	return &userQuery{}
}

var _ UserQueryInterface = (*userQuery)(nil)

func (q *userQuery) ID() string {
	return q.id
}

func (q *userQuery) SetID(id string) (UserQueryInterface, error) {
	if id == "" {
		return q, errors.New(ERROR_EMPTY_STRING)
	}

	q.id = id

	return q, nil
}

func (q *userQuery) IDIn() []string {
	return q.idIn
}

func (q *userQuery) SetIDIn(idIn []string) (UserQueryInterface, error) {
	if len(idIn) < 1 {
		return q, errors.New(ERROR_EMPTY_ARRAY)
	}

	q.idIn = idIn

	return q, nil
}

func (q *userQuery) Status() string {
	return q.status
}

func (q *userQuery) SetStatus(status string) (UserQueryInterface, error) {
	if status == "" {
		return q, errors.New(ERROR_EMPTY_STRING)
	}
	q.status = status
	return q, nil
}

func (q *userQuery) StatusIn() []string {
	return q.statusIn
}

func (q *userQuery) SetStatusIn(statusIn []string) (UserQueryInterface, error) {
	if len(statusIn) < 1 {
		return q, errors.New(ERROR_EMPTY_ARRAY)
	}
	q.statusIn = statusIn
	return q, nil
}

func (q *userQuery) Email() string {
	return q.email
}

func (q *userQuery) SetEmail(email string) (UserQueryInterface, error) {
	if email == "" {
		return q, errors.New(ERROR_EMPTY_STRING)
	}
	q.email = email
	return q, nil
}

func (q *userQuery) CreatedAtGte() string {
	return q.createdAtGte
}

func (q *userQuery) SetCreatedAtGte(createdAtGte string) (UserQueryInterface, error) {
	if createdAtGte == "" {
		return q, errors.New(ERROR_EMPTY_STRING)
	}
	q.createdAtGte = createdAtGte
	return q, nil
}

func (q *userQuery) CreatedAtLte() string {
	return q.createdAtLte
}

func (q *userQuery) SetCreatedAtLte(createdAtLte string) (UserQueryInterface, error) {
	if createdAtLte == "" {
		return q, errors.New(ERROR_EMPTY_STRING)
	}
	q.createdAtLte = createdAtLte
	return q, nil
}

func (q *userQuery) Offset() int {
	return int(q.offset)
}

func (q *userQuery) SetOffset(offset int) (UserQueryInterface, error) {
	if offset < 0 {
		return q, errors.New(ERROR_NEGATIVE_NUMBER)
	}
	q.offset = int64(offset)
	return q, nil
}

func (q *userQuery) Limit() int {
	return q.limit
}

func (q *userQuery) SetLimit(limit int) (UserQueryInterface, error) {
	if limit < 1 {
		return q, errors.New(ERROR_NEGATIVE_NUMBER)
	}
	q.limit = limit
	return q, nil
}

func (q *userQuery) SortOrder() string {
	return q.sortOrder
}

func (q *userQuery) SetSortOrder(sortOrder string) (UserQueryInterface, error) {
	if sortOrder == "" {
		return q, errors.New(ERROR_EMPTY_STRING)
	}
	q.sortOrder = sortOrder
	return q, nil
}

func (q *userQuery) OrderBy() string {
	return q.orderBy
}

func (q *userQuery) SetOrderBy(orderBy string) (UserQueryInterface, error) {
	if orderBy == "" {
		return q, errors.New(ERROR_EMPTY_STRING)
	}
	q.orderBy = orderBy
	return q, nil
}

func (q *userQuery) CountOnly() bool {
	return q.countOnly
}

func (q *userQuery) SetCountOnly(countOnly bool) UserQueryInterface {
	q.countOnly = countOnly
	return q
}

func (q *userQuery) WithSoftDeleted() bool {
	return q.withSoftDeleted
}

func (q *userQuery) SetWithSoftDeleted(withSoftDeleted bool) UserQueryInterface {
	q.withSoftDeleted = withSoftDeleted
	return q
}
