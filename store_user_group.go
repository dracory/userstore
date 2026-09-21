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

type userGroupRow struct {
	ID            string    `db:"id"`
	UserID        string    `db:"user_id"`
	GroupID       string    `db:"group_id"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
	SoftDeletedAt time.Time `db:"soft_deleted_at"`
}

func (store *storeImplementation) UserGroupCount(ctx context.Context, options UserGroupQueryInterface) (int64, error) {
	q := store.buildUserGroupQuery(options)
	var count int64
	err := q.Table(store.userGroupTableName).Count(&count)
	return count, err
}

func (store *storeImplementation) UserGroupCreate(ctx context.Context, userGroup UserGroupInterface) error {
	if userGroup == nil {
		return errors.New("user group is nil")
	}

	if userGroup.GetCreatedAt() == "" {
		userGroup.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	}
	if userGroup.GetUpdatedAt() == "" {
		userGroup.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	}
	if userGroup.GetSoftDeletedAt() == "" {
		userGroup.SetSoftDeletedAt(MAX_DATETIME)
	}

	return store.db.Query().Table(store.userGroupTableName).Create(userGroup.ToMap())
}

func (store *storeImplementation) UserGroupDelete(ctx context.Context, userGroup UserGroupInterface) error {
	if userGroup == nil {
		return errors.New("user group is nil")
	}

	return store.UserGroupDeleteByID(ctx, userGroup.GetID())
}

func (store *storeImplementation) UserGroupDeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("user group id is empty")
	}

	_, err := store.db.Query().
		Table(store.userGroupTableName).
		Where(COLUMN_ID+" = ?", id).
		Delete()

	return err
}

func (store *storeImplementation) UserGroupFindByID(ctx context.Context, id string) (userGroup UserGroupInterface, err error) {
	if id == "" {
		return nil, errors.New("user group id is empty")
	}

	query := NewUserGroupQuery().SetID(id).SetLimit(1)

	list, err := store.UserGroupList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *storeImplementation) UserGroupFindByUserIDAndGroupID(ctx context.Context, userID, groupID string) (userGroup UserGroupInterface, err error) {
	if userID == "" {
		return nil, errors.New("user id is empty")
	}

	if groupID == "" {
		return nil, errors.New("group id is empty")
	}

	query := NewUserGroupQuery().SetUserID(userID).SetGroupID(groupID).SetLimit(1)

	list, err := store.UserGroupList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

// UserGroupFindByUserIDAndGroupIDOrCreate - finds by user ID and group ID or creates the membership
func (store *storeImplementation) UserGroupFindByUserIDAndGroupIDOrCreate(ctx context.Context, userID, groupID string) (UserGroupInterface, error) {
	existingUserGroup, errUserGroup := store.UserGroupFindByUserIDAndGroupID(ctx, userID, groupID)

	if errUserGroup != nil {
		return nil, errUserGroup
	}

	if existingUserGroup != nil {
		return existingUserGroup, nil
	}

	newUserGroup := NewUserGroup().
		SetUserID(userID).
		SetGroupID(groupID)

	errCreate := store.UserGroupCreate(ctx, newUserGroup)

	if errCreate != nil {
		return nil, errCreate
	}

	return newUserGroup, nil
}

func (store *storeImplementation) UserGroupList(ctx context.Context, query UserGroupQueryInterface) ([]UserGroupInterface, error) {
	if query == nil {
		return []UserGroupInterface{}, errors.New("user group list > user group query is nil")
	}

	q := store.buildUserGroupQuery(query)

	var rows []userGroupRow
	if err := q.Table(store.userGroupTableName).Get(&rows); err != nil {
		return []UserGroupInterface{}, err
	}

	list := make([]UserGroupInterface, 0, len(rows))
	for _, r := range rows {
		userGroup := &userGroupImplementation{}
		userGroup.SetID(r.ID)
		userGroup.SetUserID(r.UserID)
		userGroup.SetGroupID(r.GroupID)
		userGroup.CreatedAtField.CreatedAt = r.CreatedAt
		userGroup.UpdatedAtField.UpdatedAt = r.UpdatedAt
		userGroup.SoftDeletedAt = r.SoftDeletedAt
		list = append(list, userGroup)
	}

	return list, nil
}

// UserGroups returns the groups a user belongs to
func (store *storeImplementation) UserGroups(ctx context.Context, userID string) ([]GroupInterface, error) {
	if userID == "" {
		return []GroupInterface{}, errors.New("user id is empty")
	}

	userGroups, err := store.UserGroupList(ctx, NewUserGroupQuery().SetUserID(userID))

	if err != nil {
		return []GroupInterface{}, err
	}

	groupIDs := lo.Map(userGroups, func(userGroup UserGroupInterface, _ int) string {
		return userGroup.GetGroupID()
	})

	if len(groupIDs) == 0 {
		return []GroupInterface{}, nil
	}

	return store.GroupList(ctx, NewGroupQuery().SetIDIn(groupIDs))
}

// UserHasGroups returns true if the user belongs to all the given groups
func (store *storeImplementation) UserHasGroups(ctx context.Context, userID string, groupIDs []string) (bool, error) {
	if userID == "" {
		return false, errors.New("user id is empty")
	}

	groupIDs = lo.Uniq(groupIDs)

	if len(groupIDs) == 0 {
		return false, errors.New("group ids are empty")
	}

	count, err := store.UserGroupCount(ctx, NewUserGroupQuery().
		SetUserID(userID).
		SetGroupIDIn(groupIDs))

	if err != nil {
		return false, err
	}

	return count == int64(len(groupIDs)), nil
}

func (store *storeImplementation) UserGroupSoftDelete(ctx context.Context, userGroup UserGroupInterface) error {
	if userGroup == nil {
		return errors.New("user group soft delete > user group is nil")
	}

	userGroup.SetSoftDeletedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	row := map[string]any{
		COLUMN_SOFT_DELETED_AT: userGroup.GetSoftDeletedAtCarbon().StdTime(),
		COLUMN_UPDATED_AT:      carbon.Now(carbon.UTC).StdTime(),
	}

	_, err := store.db.Query().
		Table(store.userGroupTableName).
		Where(COLUMN_ID+" = ?", userGroup.GetID()).
		Update(row)

	return err
}

func (store *storeImplementation) UserGroupSoftDeleteByID(ctx context.Context, id string) error {
	userGroup, err := store.UserGroupFindByID(ctx, id)

	if err != nil {
		return err
	}

	return store.UserGroupSoftDelete(ctx, userGroup)
}

func (store *storeImplementation) UserGroupUpdate(ctx context.Context, userGroup UserGroupInterface) error {
	if userGroup == nil {
		return errors.New("user group update > user group is nil")
	}

	userGroup.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	row := userGroup.ToMap()

	delete(row, COLUMN_ID)
	delete(row, COLUMN_CREATED_AT)

	_, err := store.db.Query().
		Table(store.userGroupTableName).
		Where(COLUMN_ID+" = ?", userGroup.GetID()).
		Update(row)

	return err
}

// == QUERY BUILDER ==========================================================

func (store *storeImplementation) buildUserGroupQuery(options UserGroupQueryInterface) contractsorm.Query {
	// Use Model() to enable neat's automatic soft delete handling via SoftDeletesMaxDate
	q := store.db.Query().Model(&userGroupImplementation{})

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

	if options.HasGroupID() && options.GroupID() != "" {
		q = q.Where(COLUMN_GROUP_ID+" = ?", options.GroupID())
	}

	if options.HasGroupIDIn() && len(options.GroupIDIn()) > 0 {
		q = q.Where(COLUMN_GROUP_ID+" IN ?", options.GroupIDIn())
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
