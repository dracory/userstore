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

type groupRow struct {
	ID            string    `db:"id"`
	Status        string    `db:"status"`
	Handle        string    `db:"handle"`
	Name          string    `db:"name"`
	Memo          string    `db:"memo"`
	Metas         string    `db:"metas"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
	SoftDeletedAt time.Time `db:"soft_deleted_at"`
}

func (store *storeImplementation) GroupCount(ctx context.Context, options GroupQueryInterface) (int64, error) {
	q := store.buildGroupQuery(options)
	var count int64
	err := q.Table(store.groupTableName).Count(&count)
	return count, err
}

func (store *storeImplementation) GroupCreate(ctx context.Context, group GroupInterface) error {
	if group == nil {
		return errors.New("group is nil")
	}

	if group.GetCreatedAt() == "" {
		group.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	}
	if group.GetUpdatedAt() == "" {
		group.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	}
	if group.GetSoftDeletedAt() == "" {
		group.SetSoftDeletedAt(MAX_DATETIME)
	}

	return store.db.Query().Table(store.groupTableName).Create(group.ToMap())
}

func (store *storeImplementation) GroupDelete(ctx context.Context, group GroupInterface) error {
	if group == nil {
		return errors.New("group is nil")
	}

	return store.GroupDeleteByID(ctx, group.GetID())
}

func (store *storeImplementation) GroupDeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("group id is empty")
	}

	_, err := store.db.Query().
		Table(store.groupTableName).
		Where(COLUMN_ID+" = ?", id).
		Delete()

	return err
}

func (store *storeImplementation) GroupFindByHandle(ctx context.Context, handle string) (group GroupInterface, err error) {
	if handle == "" {
		return nil, errors.New("group handle is empty")
	}

	query := NewGroupQuery().SetHandle(handle).SetLimit(1)

	list, err := store.GroupList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

// GroupFindByHandleOrCreate - finds by handle or creates a group (with the given status)
func (store *storeImplementation) GroupFindByHandleOrCreate(ctx context.Context, handle, createStatus string) (GroupInterface, error) {
	existingGroup, errGroup := store.GroupFindByHandle(ctx, handle)

	if errGroup != nil {
		return nil, errGroup
	}

	if existingGroup != nil {
		return existingGroup, nil
	}

	newGroup := NewGroup().
		SetHandle(handle).
		SetStatus(createStatus)

	errCreate := store.GroupCreate(ctx, newGroup)

	if errCreate != nil {
		return nil, errCreate
	}

	return newGroup, nil
}

func (store *storeImplementation) GroupFindByID(ctx context.Context, id string) (group GroupInterface, err error) {
	if id == "" {
		return nil, errors.New("group id is empty")
	}

	query := NewGroupQuery().SetID(id).SetLimit(1)

	list, err := store.GroupList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *storeImplementation) GroupList(ctx context.Context, query GroupQueryInterface) ([]GroupInterface, error) {
	if query == nil {
		return []GroupInterface{}, errors.New("group list > group query is nil")
	}

	q := store.buildGroupQuery(query)

	var rows []groupRow
	if err := q.Table(store.groupTableName).Get(&rows); err != nil {
		return []GroupInterface{}, err
	}

	list := make([]GroupInterface, 0, len(rows))
	for _, r := range rows {
		group := &groupImplementation{}
		group.SetID(r.ID)
		group.SetStatus(r.Status)
		group.SetHandle(r.Handle)
		group.SetName(r.Name)
		group.SetMemo(r.Memo)
		group.MetasField = r.Metas
		group.CreatedAtField.CreatedAt = r.CreatedAt
		group.UpdatedAtField.UpdatedAt = r.UpdatedAt
		group.SoftDeletedAt = r.SoftDeletedAt
		list = append(list, group)
	}

	return list, nil
}

func (store *storeImplementation) GroupSoftDelete(ctx context.Context, group GroupInterface) error {
	if group == nil {
		return errors.New("group soft delete > group is nil")
	}

	group.SetSoftDeletedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	row := map[string]any{
		COLUMN_SOFT_DELETED_AT: group.GetSoftDeletedAtCarbon().StdTime(),
		COLUMN_UPDATED_AT:      carbon.Now(carbon.UTC).StdTime(),
	}

	_, err := store.db.Query().
		Table(store.groupTableName).
		Where(COLUMN_ID+" = ?", group.GetID()).
		Update(row)

	return err
}

func (store *storeImplementation) GroupSoftDeleteByID(ctx context.Context, id string) error {
	group, err := store.GroupFindByID(ctx, id)

	if err != nil {
		return err
	}

	return store.GroupSoftDelete(ctx, group)
}

func (store *storeImplementation) GroupUpdate(ctx context.Context, group GroupInterface) error {
	if group == nil {
		return errors.New("group update > group is nil")
	}

	group.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	row := group.ToMap()

	delete(row, COLUMN_ID)
	delete(row, COLUMN_CREATED_AT)

	_, err := store.db.Query().
		Table(store.groupTableName).
		Where(COLUMN_ID+" = ?", group.GetID()).
		Update(row)

	return err
}

// == QUERY BUILDER ==========================================================

func (store *storeImplementation) buildGroupQuery(options GroupQueryInterface) contractsorm.Query {
	// Use Model() to enable neat's automatic soft delete handling via SoftDeletesMaxDate
	q := store.db.Query().Model(&groupImplementation{})

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

	if options.HasHandle() && options.Handle() != "" {
		q = q.Where(COLUMN_HANDLE+" = ?", options.Handle())
	}

	if options.HasTitleLike() && options.TitleLike() != "" {
		q = q.Where(COLUMN_NAME+" LIKE ?", `%`+options.TitleLike()+`%`)
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
