package userstore

import (
	"context"
	"errors"
	"strings"
	"time"

	contractsorm "github.com/dracory/neat/contracts/database/orm"
	"github.com/dromara/carbon/v2"
	"github.com/samber/lo"
)

type userRow struct {
	ID              string    `db:"id"`
	Status          string    `db:"status"`
	FirstName       string    `db:"first_name"`
	MiddleNames     string    `db:"middle_names"`
	LastName        string    `db:"last_name"`
	BusinessName    string    `db:"business_name"`
	Phone           string    `db:"phone"`
	Email           string    `db:"email"`
	Password        string    `db:"password"`
	Role            string    `db:"role"`
	Country         string    `db:"country"`
	Timezone        string    `db:"timezone"`
	ProfileImageUrl string    `db:"profile_image_url"`
	Metas           string    `db:"metas"`
	Memo            string    `db:"memo"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	SoftDeletedAt   time.Time `db:"soft_deleted_at"`
}

func (store *storeImplementation) UserCount(ctx context.Context, options UserQueryInterface) (int64, error) {
	q := store.buildQuery(options)
	var count int64
	err := q.Table(store.userTableName).Count(&count)
	return count, err
}

func (store *storeImplementation) UserCreate(ctx context.Context, user UserInterface) error {
	if user == nil {
		return errors.New("user is nil")
	}

	if user.GetCreatedAt() == "" {
		user.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	}
	if user.GetUpdatedAt() == "" {
		user.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	}
	if user.GetSoftDeletedAt() == "" {
		user.SetSoftDeletedAt(MAX_DATETIME)
	}

	row := map[string]any{
		COLUMN_ID:                user.GetID(),
		COLUMN_STATUS:            user.GetStatus(),
		COLUMN_FIRST_NAME:        user.GetFirstName(),
		COLUMN_MIDDLE_NAMES:      user.GetMiddleNames(),
		COLUMN_LAST_NAME:         user.GetLastName(),
		COLUMN_BUSINESS_NAME:     user.GetBusinessName(),
		COLUMN_PHONE:             user.GetPhone(),
		COLUMN_EMAIL:             user.GetEmail(),
		COLUMN_PASSWORD:          user.GetPassword(),
		COLUMN_ROLE:              user.GetRole(),
		COLUMN_COUNTRY:           user.GetCountry(),
		COLUMN_TIMEZONE:          user.GetTimezone(),
		COLUMN_PROFILE_IMAGE_URL: user.GetProfileImageUrl(),
		COLUMN_METAS:             user.Get(COLUMN_METAS),
		COLUMN_MEMO:              user.GetMemo(),
		COLUMN_CREATED_AT:        user.GetCreatedAtCarbon().StdTime(),
		COLUMN_UPDATED_AT:        user.GetUpdatedAtCarbon().StdTime(),
		COLUMN_SOFT_DELETED_AT:   user.GetSoftDeletedAtCarbon().StdTime(),
	}

	return store.db.Query().Table(store.userTableName).Create(row)
}

func (store *storeImplementation) UserDelete(ctx context.Context, user UserInterface) error {
	if user == nil {
		return errors.New("user is nil")
	}

	return store.UserDeleteByID(ctx, user.GetID())
}

func (store *storeImplementation) UserDeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("user id is empty")
	}

	_, err := store.db.Query().
		Table(store.userTableName).
		Where(COLUMN_ID+" = ?", id).
		Delete()

	return err
}

func (store *storeImplementation) UserFindByEmail(ctx context.Context, email string) (user UserInterface, err error) {
	if email == "" {
		return nil, errors.New("user email is empty")
	}

	query := NewUserQuery().SetEmail(email).SetLimit(1)

	list, err := store.UserList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

// UserFindByEmailOrCreate - finds by email or creates a user (with active status)
func (store *storeImplementation) UserFindByEmailOrCreate(ctx context.Context, email, createStatus string) (UserInterface, error) {
	existingUser, errUser := store.UserFindByEmail(ctx, email)

	if errUser != nil {
		return nil, errUser
	}

	if existingUser != nil {
		return existingUser, nil
	}

	newUser := NewUser().
		SetEmail(email).
		SetStatus(createStatus)

	errCreate := store.UserCreate(ctx, newUser)

	if errCreate != nil {
		return nil, errCreate
	}

	return newUser, nil
}

func (store *storeImplementation) UserFindByID(ctx context.Context, id string) (user UserInterface, err error) {
	if id == "" {
		return nil, errors.New("user id is empty")
	}

	query := NewUserQuery().SetID(id).SetLimit(1)

	list, err := store.UserList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *storeImplementation) UserList(ctx context.Context, query UserQueryInterface) ([]UserInterface, error) {
	if query == nil {
		return []UserInterface{}, errors.New("user list > user query is nil")
	}

	q := store.buildQuery(query)

	var rows []userRow
	if err := q.Table(store.userTableName).Get(&rows); err != nil {
		return []UserInterface{}, err
	}

	list := make([]UserInterface, 0, len(rows))
	for _, r := range rows {
		user := &userImplementation{}
		user.SetID(r.ID)
		user.SetStatus(r.Status)
		user.SetFirstName(r.FirstName)
		user.SetMiddleNames(r.MiddleNames)
		user.SetLastName(r.LastName)
		user.SetBusinessName(r.BusinessName)
		user.SetPhone(r.Phone)
		user.SetEmail(r.Email)
		user.SetPassword(r.Password)
		user.SetRole(r.Role)
		user.SetCountry(r.Country)
		user.SetTimezone(r.Timezone)
		user.SetProfileImageUrl(r.ProfileImageUrl)
		user.MetasField = r.Metas
		user.SetMemo(r.Memo)
		user.CreatedAtField.CreatedAt = r.CreatedAt
		user.UpdatedAtField.UpdatedAt = r.UpdatedAt
		user.SoftDeletedAt = r.SoftDeletedAt
		list = append(list, user)
	}

	return list, nil
}

func (store *storeImplementation) UserSoftDelete(ctx context.Context, user UserInterface) error {
	if user == nil {
		return errors.New("user soft delete > user is nil")
	}

	user.SetSoftDeletedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	row := map[string]any{
		COLUMN_SOFT_DELETED_AT: user.GetSoftDeletedAtCarbon().StdTime(),
		COLUMN_UPDATED_AT:      carbon.Now(carbon.UTC).StdTime(),
	}

	_, err := store.db.Query().
		Table(store.userTableName).
		Where(COLUMN_ID+" = ?", user.GetID()).
		Update(row)

	return err
}

func (store *storeImplementation) UserSoftDeleteByID(ctx context.Context, id string) error {
	user, err := store.UserFindByID(ctx, id)

	if err != nil {
		return err
	}

	return store.UserSoftDelete(ctx, user)
}

func (store *storeImplementation) UserUpdate(ctx context.Context, user UserInterface) error {
	if user == nil {
		return errors.New("user update > user is nil")
	}

	user.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	row := map[string]any{
		COLUMN_STATUS:            user.GetStatus(),
		COLUMN_FIRST_NAME:        user.GetFirstName(),
		COLUMN_MIDDLE_NAMES:      user.GetMiddleNames(),
		COLUMN_LAST_NAME:         user.GetLastName(),
		COLUMN_BUSINESS_NAME:     user.GetBusinessName(),
		COLUMN_PHONE:             user.GetPhone(),
		COLUMN_EMAIL:             user.GetEmail(),
		COLUMN_PASSWORD:          user.GetPassword(),
		COLUMN_ROLE:              user.GetRole(),
		COLUMN_COUNTRY:           user.GetCountry(),
		COLUMN_TIMEZONE:          user.GetTimezone(),
		COLUMN_PROFILE_IMAGE_URL: user.GetProfileImageUrl(),
		COLUMN_METAS:             user.Get(COLUMN_METAS),
		COLUMN_MEMO:              user.GetMemo(),
		COLUMN_UPDATED_AT:        user.GetUpdatedAtCarbon().StdTime(),
	}

	_, err := store.db.Query().
		Table(store.userTableName).
		Where(COLUMN_ID+" = ?", user.GetID()).
		Update(row)

	return err
}

// == QUERY BUILDER ==========================================================

func (store *storeImplementation) buildQuery(options UserQueryInterface) contractsorm.Query {
	q := store.db.Query()

	if options == nil {
		return q
	}

	if options.HasID() && options.GetID() != "" {
		q = q.Where(COLUMN_ID+" = ?", options.GetID())
	}

	if options.HasIDIn() && len(options.IDIn()) > 0 {
		q = q.Where(COLUMN_ID+" IN ?", options.IDIn())
	}

	if options.HasStatus() && options.Status() != "" {
		q = q.Where(COLUMN_STATUS+" = ?", options.Status())
	}

	if options.HasStatusIn() && len(options.StatusIn()) > 0 {
		q = q.Where(COLUMN_STATUS+" IN ?", options.StatusIn())
	}

	if options.HasEmail() && options.Email() != "" {
		q = q.Where(COLUMN_EMAIL+" = ?", options.Email())
	}

	if options.HasEmailLike() && options.EmailLike() != "" {
		q = q.Where(COLUMN_EMAIL+" LIKE ?", `%`+options.EmailLike()+`%`)
	}

	if options.HasFirstName() && options.FirstName() != "" {
		q = q.Where(COLUMN_FIRST_NAME+" = ?", options.FirstName())
	}

	if options.HasFirstNameLike() && options.FirstNameLike() != "" {
		q = q.Where(COLUMN_FIRST_NAME+" LIKE ?", `%`+options.FirstNameLike()+`%`)
	}

	if options.HasLastName() && options.LastName() != "" {
		q = q.Where(COLUMN_LAST_NAME+" = ?", options.LastName())
	}

	if options.HasLastNameLike() && options.LastNameLike() != "" {
		q = q.Where(COLUMN_LAST_NAME+" LIKE ?", `%`+options.LastNameLike()+`%`)
	}

	if options.HasMetaLike() && options.MetaLike() != "" {
		q = q.Where(COLUMN_METAS+" LIKE ?", `%`+options.MetaLike()+`%`)
	}

	if options.HasCreatedAtGte() && options.CreatedAtGte() != "" {
		q = q.Where(COLUMN_CREATED_AT+" >= ?", options.CreatedAtGte())
	}

	if options.HasCreatedAtLte() && options.CreatedAtLte() != "" {
		q = q.Where(COLUMN_CREATED_AT+" <= ?", options.CreatedAtLte())
	}

	if options.HasLimit() && options.Limit() > 0 {
		q = q.Limit(options.Limit())
	}

	if options.HasOffset() && options.Offset() > 0 {
		q = q.Offset(options.Offset())
	}

	if options.HasOrderBy() && options.OrderBy() != "" {
		sort := lo.Ternary(options.HasSortDirection(), options.SortDirection(), "DESC")
		if strings.EqualFold(sort, "ASC") {
			q = q.OrderBy(options.OrderBy(), "asc")
		} else {
			q = q.OrderBy(options.OrderBy(), "desc")
		}
	}

	if options.HasSoftDeletedIncluded() && options.SoftDeletedIncluded() {
		q = q.WithSoftDeleted()
	} else {
		q = q.Where(COLUMN_SOFT_DELETED_AT+" = ?", carbon.Parse(MAX_DATETIME, carbon.UTC).StdTime())
	}

	return q
}
