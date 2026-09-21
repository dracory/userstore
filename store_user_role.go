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

type userRoleRow struct {
	ID            string    `db:"id"`
	UserID        string    `db:"user_id"`
	RoleID        string    `db:"role_id"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
	SoftDeletedAt time.Time `db:"soft_deleted_at"`
}

func (store *storeImplementation) UserRoleCount(ctx context.Context, options UserRoleQueryInterface) (int64, error) {
	q := store.buildUserRoleQuery(options)
	var count int64
	err := q.Table(store.userRoleTableName).Count(&count)
	return count, err
}

func (store *storeImplementation) UserRoleCreate(ctx context.Context, userRole UserRoleInterface) error {
	if userRole == nil {
		return errors.New("user role is nil")
	}

	if userRole.GetCreatedAt() == "" {
		userRole.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	}
	if userRole.GetUpdatedAt() == "" {
		userRole.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	}
	if userRole.GetSoftDeletedAt() == "" {
		userRole.SetSoftDeletedAt(MAX_DATETIME)
	}

	return store.db.Query().Table(store.userRoleTableName).Create(userRole.ToMap())
}

func (store *storeImplementation) UserRoleDelete(ctx context.Context, userRole UserRoleInterface) error {
	if userRole == nil {
		return errors.New("user role is nil")
	}

	return store.UserRoleDeleteByID(ctx, userRole.GetID())
}

func (store *storeImplementation) UserRoleDeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("user role id is empty")
	}

	_, err := store.db.Query().
		Table(store.userRoleTableName).
		Where(COLUMN_ID+" = ?", id).
		Delete()

	return err
}

func (store *storeImplementation) UserRoleFindByID(ctx context.Context, id string) (userRole UserRoleInterface, err error) {
	if id == "" {
		return nil, errors.New("user role id is empty")
	}

	query := NewUserRoleQuery().SetID(id).SetLimit(1)

	list, err := store.UserRoleList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *storeImplementation) UserRoleFindByUserIDAndRoleID(ctx context.Context, userID, roleID string) (userRole UserRoleInterface, err error) {
	if userID == "" {
		return nil, errors.New("user id is empty")
	}

	if roleID == "" {
		return nil, errors.New("role id is empty")
	}

	query := NewUserRoleQuery().SetUserID(userID).SetRoleID(roleID).SetLimit(1)

	list, err := store.UserRoleList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

// UserRoleFindByUserIDAndRoleIDOrCreate - finds by user ID and role ID or creates the assignment
func (store *storeImplementation) UserRoleFindByUserIDAndRoleIDOrCreate(ctx context.Context, userID, roleID string) (UserRoleInterface, error) {
	existingUserRole, errUserRole := store.UserRoleFindByUserIDAndRoleID(ctx, userID, roleID)

	if errUserRole != nil {
		return nil, errUserRole
	}

	if existingUserRole != nil {
		return existingUserRole, nil
	}

	newUserRole := NewUserRole().
		SetUserID(userID).
		SetRoleID(roleID)

	errCreate := store.UserRoleCreate(ctx, newUserRole)

	if errCreate != nil {
		return nil, errCreate
	}

	return newUserRole, nil
}

func (store *storeImplementation) UserRoleList(ctx context.Context, query UserRoleQueryInterface) ([]UserRoleInterface, error) {
	if query == nil {
		return []UserRoleInterface{}, errors.New("user role list > user role query is nil")
	}

	q := store.buildUserRoleQuery(query)

	var rows []userRoleRow
	if err := q.Table(store.userRoleTableName).Get(&rows); err != nil {
		return []UserRoleInterface{}, err
	}

	list := make([]UserRoleInterface, 0, len(rows))
	for _, r := range rows {
		userRole := &userRoleImplementation{}
		userRole.SetID(r.ID)
		userRole.SetUserID(r.UserID)
		userRole.SetRoleID(r.RoleID)
		userRole.CreatedAtField.CreatedAt = r.CreatedAt
		userRole.UpdatedAtField.UpdatedAt = r.UpdatedAt
		userRole.SoftDeletedAt = r.SoftDeletedAt
		list = append(list, userRole)
	}

	return list, nil
}

// UserRoles returns the roles assigned to a user
func (store *storeImplementation) UserRoles(ctx context.Context, userID string) ([]RoleInterface, error) {
	if userID == "" {
		return []RoleInterface{}, errors.New("user id is empty")
	}

	userRoles, err := store.UserRoleList(ctx, NewUserRoleQuery().SetUserID(userID))

	if err != nil {
		return []RoleInterface{}, err
	}

	roleIDs := lo.Map(userRoles, func(userRole UserRoleInterface, _ int) string {
		return userRole.GetRoleID()
	})

	if len(roleIDs) == 0 {
		return []RoleInterface{}, nil
	}

	return store.RoleList(ctx, NewRoleQuery().SetIDIn(roleIDs))
}

// UserHasRoles returns true if the user has all the given roles
func (store *storeImplementation) UserHasRoles(ctx context.Context, userID string, roleIDs []string) (bool, error) {
	if userID == "" {
		return false, errors.New("user id is empty")
	}

	roleIDs = lo.Uniq(roleIDs)

	if len(roleIDs) == 0 {
		return false, errors.New("role ids are empty")
	}

	count, err := store.UserRoleCount(ctx, NewUserRoleQuery().
		SetUserID(userID).
		SetRoleIDIn(roleIDs))

	if err != nil {
		return false, err
	}

	return count == int64(len(roleIDs)), nil
}

func (store *storeImplementation) UserRoleSoftDelete(ctx context.Context, userRole UserRoleInterface) error {
	if userRole == nil {
		return errors.New("user role soft delete > user role is nil")
	}

	userRole.SetSoftDeletedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	row := map[string]any{
		COLUMN_SOFT_DELETED_AT: userRole.GetSoftDeletedAtCarbon().StdTime(),
		COLUMN_UPDATED_AT:      carbon.Now(carbon.UTC).StdTime(),
	}

	_, err := store.db.Query().
		Table(store.userRoleTableName).
		Where(COLUMN_ID+" = ?", userRole.GetID()).
		Update(row)

	return err
}

func (store *storeImplementation) UserRoleSoftDeleteByID(ctx context.Context, id string) error {
	userRole, err := store.UserRoleFindByID(ctx, id)

	if err != nil {
		return err
	}

	return store.UserRoleSoftDelete(ctx, userRole)
}

func (store *storeImplementation) UserRoleUpdate(ctx context.Context, userRole UserRoleInterface) error {
	if userRole == nil {
		return errors.New("user role update > user role is nil")
	}

	userRole.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	row := userRole.ToMap()

	delete(row, COLUMN_ID)
	delete(row, COLUMN_CREATED_AT)

	_, err := store.db.Query().
		Table(store.userRoleTableName).
		Where(COLUMN_ID+" = ?", userRole.GetID()).
		Update(row)

	return err
}

// == QUERY BUILDER ==========================================================

func (store *storeImplementation) buildUserRoleQuery(options UserRoleQueryInterface) contractsorm.Query {
	// Use Model() to enable neat's automatic soft delete handling via SoftDeletesMaxDate
	q := store.db.Query().Model(&userRoleImplementation{})

	if options == nil {
		return q
	}

	if options.HasID() && options.GetID() != "" {
		q = q.Where(COLUMN_ID+" = ?", options.GetID())
	}

	if options.HasIDIn() && len(options.IDIn()) > 0 {
		q = q.Where(COLUMN_ID+" IN ?", options.IDIn())
	}

	if options.HasUserID() && options.UserID() != "" {
		q = q.Where(COLUMN_USER_ID+" = ?", options.UserID())
	}

	if options.HasUserIDIn() && len(options.UserIDIn()) > 0 {
		q = q.Where(COLUMN_USER_ID+" IN ?", options.UserIDIn())
	}

	if options.HasRoleID() && options.RoleID() != "" {
		q = q.Where(COLUMN_ROLE_ID+" = ?", options.RoleID())
	}

	if options.HasRoleIDIn() && len(options.RoleIDIn()) > 0 {
		q = q.Where(COLUMN_ROLE_ID+" IN ?", options.RoleIDIn())
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

	// Handle soft delete filtering via neat's automatic handling (SoftDeletesMaxDate)
	if options.HasSoftDeletedIncluded() && options.SoftDeletedIncluded() {
		q = q.WithSoftDeleted()
	}

	return q
}
