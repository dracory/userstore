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

type roleRow struct {
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

func (store *storeImplementation) RoleCount(ctx context.Context, options RoleQueryInterface) (int64, error) {
	q := store.buildRoleQuery(options)
	var count int64
	err := q.Table(store.roleTableName).Count(&count)
	return count, err
}

func (store *storeImplementation) RoleCreate(ctx context.Context, role RoleInterface) error {
	if role == nil {
		return errors.New("role is nil")
	}

	if role.GetCreatedAt() == "" {
		role.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	}
	if role.GetUpdatedAt() == "" {
		role.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	}
	if role.GetSoftDeletedAt() == "" {
		role.SetSoftDeletedAt(MAX_DATETIME)
	}

	return store.db.Query().Table(store.roleTableName).Create(role.ToMap())
}

func (store *storeImplementation) RoleDelete(ctx context.Context, role RoleInterface) error {
	if role == nil {
		return errors.New("role is nil")
	}

	return store.RoleDeleteByID(ctx, role.GetID())
}

func (store *storeImplementation) RoleDeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("role id is empty")
	}

	_, err := store.db.Query().
		Table(store.roleTableName).
		Where(COLUMN_ID+" = ?", id).
		Delete()

	return err
}

func (store *storeImplementation) RoleFindByHandle(ctx context.Context, handle string) (role RoleInterface, err error) {
	if handle == "" {
		return nil, errors.New("role handle is empty")
	}

	query := NewRoleQuery().SetHandle(handle).SetLimit(1)

	list, err := store.RoleList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

// RoleFindByHandleOrCreate - finds by handle or creates a role (with the given status)
func (store *storeImplementation) RoleFindByHandleOrCreate(ctx context.Context, handle, createStatus string) (RoleInterface, error) {
	existingRole, errRole := store.RoleFindByHandle(ctx, handle)

	if errRole != nil {
		return nil, errRole
	}

	if existingRole != nil {
		return existingRole, nil
	}

	newRole := NewRole().
		SetHandle(handle).
		SetStatus(createStatus)

	errCreate := store.RoleCreate(ctx, newRole)

	if errCreate != nil {
		return nil, errCreate
	}

	return newRole, nil
}

func (store *storeImplementation) RoleFindByID(ctx context.Context, id string) (role RoleInterface, err error) {
	if id == "" {
		return nil, errors.New("role id is empty")
	}

	query := NewRoleQuery().SetID(id).SetLimit(1)

	list, err := store.RoleList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *storeImplementation) RoleList(ctx context.Context, query RoleQueryInterface) ([]RoleInterface, error) {
	if query == nil {
		return []RoleInterface{}, errors.New("role list > role query is nil")
	}

	q := store.buildRoleQuery(query)

	var rows []roleRow
	if err := q.Table(store.roleTableName).Get(&rows); err != nil {
		return []RoleInterface{}, err
	}

	list := make([]RoleInterface, 0, len(rows))
	for _, r := range rows {
		role := &roleImplementation{}
		role.SetID(r.ID)
		role.SetStatus(r.Status)
		role.SetHandle(r.Handle)
		role.SetName(r.Name)
		role.SetMemo(r.Memo)
		role.MetasField = r.Metas
		role.CreatedAtField.CreatedAt = r.CreatedAt
		role.UpdatedAtField.UpdatedAt = r.UpdatedAt
		role.SoftDeletedAt = r.SoftDeletedAt
		list = append(list, role)
	}

	return list, nil
}

func (store *storeImplementation) RoleSoftDelete(ctx context.Context, role RoleInterface) error {
	if role == nil {
		return errors.New("role soft delete > role is nil")
	}

	role.SetSoftDeletedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	row := map[string]any{
		COLUMN_SOFT_DELETED_AT: role.GetSoftDeletedAtCarbon().StdTime(),
		COLUMN_UPDATED_AT:      carbon.Now(carbon.UTC).StdTime(),
	}

	_, err := store.db.Query().
		Table(store.roleTableName).
		Where(COLUMN_ID+" = ?", role.GetID()).
		Update(row)

	return err
}

func (store *storeImplementation) RoleSoftDeleteByID(ctx context.Context, id string) error {
	role, err := store.RoleFindByID(ctx, id)

	if err != nil {
		return err
	}

	return store.RoleSoftDelete(ctx, role)
}

func (store *storeImplementation) RoleUpdate(ctx context.Context, role RoleInterface) error {
	if role == nil {
		return errors.New("role update > role is nil")
	}

	role.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	row := role.ToMap()

	delete(row, COLUMN_ID)
	delete(row, COLUMN_CREATED_AT)

	_, err := store.db.Query().
		Table(store.roleTableName).
		Where(COLUMN_ID+" = ?", role.GetID()).
		Update(row)

	return err
}

// == QUERY BUILDER ==========================================================

func (store *storeImplementation) buildRoleQuery(options RoleQueryInterface) contractsorm.Query {
	// Use Model() to enable neat's automatic soft delete handling via SoftDeletesMaxDate
	q := store.db.Query().Model(&roleImplementation{})

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
