package userstore

import (
	"github.com/dracory/neat/database/orm"
	"github.com/dracory/neat/database/soft_delete"
	"github.com/dracory/neat/support/uid"
	"github.com/dromara/carbon/v2"
	"time"
)

// == TYPE ====================================================================

type userRoleImplementation struct {
	orm.ShortID

	UserIDField    string `db:"user_id"`
	RoleIDField    string `db:"role_id"`
	CreatedAtField orm.CreatedAt
	UpdatedAtField orm.UpdatedAt
	soft_delete.SoftDeletesMaxDate
}

var _ UserRoleInterface = (*userRoleImplementation)(nil)

// == CONSTRUCTORS ============================================================

func NewUserRole() UserRoleInterface {
	o := (&userRoleImplementation{}).
		SetID(uid.GenerateShortID()).
		SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetSoftDeletedAt(MAX_DATETIME)

	return o
}

func NewUserRoleFromExistingData(data map[string]string) UserRoleInterface {
	o := &userRoleImplementation{}
	if v, ok := data[COLUMN_ID]; ok {
		o.SetID(v)
	}
	if v, ok := data[COLUMN_USER_ID]; ok {
		o.SetUserID(v)
	}
	if v, ok := data[COLUMN_ROLE_ID]; ok {
		o.SetRoleID(v)
	}
	if v, ok := data[COLUMN_CREATED_AT]; ok {
		o.SetCreatedAt(v)
	}
	if v, ok := data[COLUMN_UPDATED_AT]; ok {
		o.SetUpdatedAt(v)
	}
	if v, ok := data[COLUMN_SOFT_DELETED_AT]; ok {
		o.SetSoftDeletedAt(v)
	}
	return o
}

// == METHODS =================================================================

func (o *userRoleImplementation) IsSoftDeleted() bool {
	return o.SoftDeletedAt.Before(time.Now().UTC())
}

// == DATAOBJECT COMPATIBILITY ================================================

// Data returns all fields as a map.
func (o *userRoleImplementation) Data() map[string]string {
	return map[string]string{
		COLUMN_ID:              o.GetID(),
		COLUMN_USER_ID:         o.GetUserID(),
		COLUMN_ROLE_ID:         o.GetRoleID(),
		COLUMN_CREATED_AT:      o.GetCreatedAt(),
		COLUMN_UPDATED_AT:      o.GetUpdatedAt(),
		COLUMN_SOFT_DELETED_AT: o.GetSoftDeletedAt(),
	}
}

// DataChanged returns all fields as a map (dirty tracking disabled).
func (o *userRoleImplementation) DataChanged() map[string]string {
	return o.Data()
}

// MarkAsNotDirty is a no-op (dirty tracking disabled).
func (o *userRoleImplementation) MarkAsNotDirty() {}

// ToMap returns a DB-ready map[string]any of the user role.
func (o *userRoleImplementation) ToMap() map[string]any {
	return map[string]any{
		COLUMN_ID:              o.GetID(),
		COLUMN_USER_ID:         o.GetUserID(),
		COLUMN_ROLE_ID:         o.GetRoleID(),
		COLUMN_CREATED_AT:      o.GetCreatedAtCarbon().StdTime(),
		COLUMN_UPDATED_AT:      o.GetUpdatedAtCarbon().StdTime(),
		COLUMN_SOFT_DELETED_AT: o.GetSoftDeletedAtCarbon().StdTime(),
	}
}

// == SETTERS AND GETTERS =====================================================

func (o *userRoleImplementation) GetCreatedAt() string {
	if o.CreatedAtField.CreatedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(o.CreatedAtField.CreatedAt).ToDateTimeString()
}

func (o *userRoleImplementation) GetCreatedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(o.CreatedAtField.CreatedAt)
}

func (o *userRoleImplementation) SetCreatedAt(createdAt string) UserRoleInterface {
	if createdAt == "" {
		return o
	}
	o.CreatedAtField.CreatedAt = carbon.Parse(createdAt, carbon.UTC).StdTime()
	return o
}

func (o *userRoleImplementation) GetID() string {
	return o.ShortID.ID
}

func (o *userRoleImplementation) SetID(id string) UserRoleInterface {
	o.ShortID.ID = id
	return o
}

func (o *userRoleImplementation) GetUserID() string {
	return o.UserIDField
}

func (o *userRoleImplementation) SetUserID(userID string) UserRoleInterface {
	o.UserIDField = userID
	return o
}

func (o *userRoleImplementation) GetRoleID() string {
	return o.RoleIDField
}

func (o *userRoleImplementation) SetRoleID(roleID string) UserRoleInterface {
	o.RoleIDField = roleID
	return o
}

func (o *userRoleImplementation) GetSoftDeletedAt() string {
	if o.SoftDeletedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(o.SoftDeletedAt).ToDateTimeString()
}

func (o *userRoleImplementation) GetSoftDeletedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(o.SoftDeletedAt)
}

func (o *userRoleImplementation) SetSoftDeletedAt(deletedAt string) UserRoleInterface {
	if deletedAt == "" {
		return o
	}
	o.SoftDeletedAt = carbon.Parse(deletedAt, carbon.UTC).StdTime()
	return o
}

func (o *userRoleImplementation) GetUpdatedAt() string {
	if o.UpdatedAtField.UpdatedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(o.UpdatedAtField.UpdatedAt).ToDateTimeString()
}

func (o *userRoleImplementation) GetUpdatedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(o.UpdatedAtField.UpdatedAt)
}

func (o *userRoleImplementation) SetUpdatedAt(updatedAt string) UserRoleInterface {
	if updatedAt == "" {
		return o
	}
	o.UpdatedAtField.UpdatedAt = carbon.Parse(updatedAt, carbon.UTC).StdTime()
	return o
}
