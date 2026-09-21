package userstore

import (
	"time"

	"github.com/dracory/neat/database/orm"
	"github.com/dracory/neat/database/soft_delete"
	"github.com/dracory/neat/support/uid"
	"github.com/dromara/carbon/v2"
)

// == TYPE ====================================================================

type userGroupImplementation struct {
	orm.ShortID

	UserIDField    string `db:"user_id"`
	GroupIDField   string `db:"group_id"`
	CreatedAtField orm.CreatedAt
	UpdatedAtField orm.UpdatedAt
	soft_delete.SoftDeletesMaxDate
}

var _ UserGroupInterface = (*userGroupImplementation)(nil)

// == CONSTRUCTORS ============================================================

func NewUserGroup() UserGroupInterface {
	o := (&userGroupImplementation{}).
		SetID(uid.GenerateShortID()).
		SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetSoftDeletedAt(MAX_DATETIME)

	return o
}

func NewUserGroupFromExistingData(data map[string]string) UserGroupInterface {
	o := &userGroupImplementation{}
	if v, ok := data[COLUMN_ID]; ok {
		o.SetID(v)
	}
	if v, ok := data[COLUMN_USER_ID]; ok {
		o.SetUserID(v)
	}
	if v, ok := data[COLUMN_GROUP_ID]; ok {
		o.SetGroupID(v)
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

func (o *userGroupImplementation) IsSoftDeleted() bool {
	return o.SoftDeletedAt.Before(time.Now().UTC())
}

// == DATAOBJECT COMPATIBILITY ================================================

// Data returns all fields as a map.
func (o *userGroupImplementation) Data() map[string]string {
	return map[string]string{
		COLUMN_ID:              o.GetID(),
		COLUMN_USER_ID:         o.GetUserID(),
		COLUMN_GROUP_ID:        o.GetGroupID(),
		COLUMN_CREATED_AT:      o.GetCreatedAt(),
		COLUMN_UPDATED_AT:      o.GetUpdatedAt(),
		COLUMN_SOFT_DELETED_AT: o.GetSoftDeletedAt(),
	}
}

// DataChanged returns all fields as a map (dirty tracking disabled).
func (o *userGroupImplementation) DataChanged() map[string]string {
	return o.Data()
}

// MarkAsNotDirty is a no-op (dirty tracking disabled).
func (o *userGroupImplementation) MarkAsNotDirty() {}

// ToMap returns a DB-ready map[string]any of the user group.
func (o *userGroupImplementation) ToMap() map[string]any {
	return map[string]any{
		COLUMN_ID:              o.GetID(),
		COLUMN_USER_ID:         o.GetUserID(),
		COLUMN_GROUP_ID:        o.GetGroupID(),
		COLUMN_CREATED_AT:      o.GetCreatedAtCarbon().StdTime(),
		COLUMN_UPDATED_AT:      o.GetUpdatedAtCarbon().StdTime(),
		COLUMN_SOFT_DELETED_AT: o.GetSoftDeletedAtCarbon().StdTime(),
	}
}

// == SETTERS AND GETTERS =====================================================

func (o *userGroupImplementation) GetCreatedAt() string {
	if o.CreatedAtField.CreatedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(o.CreatedAtField.CreatedAt).ToDateTimeString()
}

func (o *userGroupImplementation) GetCreatedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(o.CreatedAtField.CreatedAt)
}

func (o *userGroupImplementation) SetCreatedAt(createdAt string) UserGroupInterface {
	if createdAt == "" {
		return o
	}
	o.CreatedAtField.CreatedAt = carbon.Parse(createdAt, carbon.UTC).StdTime()
	return o
}

func (o *userGroupImplementation) GetID() string {
	return o.ShortID.ID
}

func (o *userGroupImplementation) SetID(id string) UserGroupInterface {
	o.ShortID.ID = id
	return o
}

func (o *userGroupImplementation) GetUserID() string {
	return o.UserIDField
}

func (o *userGroupImplementation) SetUserID(userID string) UserGroupInterface {
	o.UserIDField = userID
	return o
}

func (o *userGroupImplementation) GetGroupID() string {
	return o.GroupIDField
}

func (o *userGroupImplementation) SetGroupID(groupID string) UserGroupInterface {
	o.GroupIDField = groupID
	return o
}

func (o *userGroupImplementation) GetSoftDeletedAt() string {
	if o.SoftDeletedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(o.SoftDeletedAt).ToDateTimeString()
}

func (o *userGroupImplementation) GetSoftDeletedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(o.SoftDeletedAt)
}

func (o *userGroupImplementation) SetSoftDeletedAt(deletedAt string) UserGroupInterface {
	if deletedAt == "" {
		return o
	}
	o.SoftDeletedAt = carbon.Parse(deletedAt, carbon.UTC).StdTime()
	return o
}

func (o *userGroupImplementation) GetUpdatedAt() string {
	if o.UpdatedAtField.UpdatedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(o.UpdatedAtField.UpdatedAt).ToDateTimeString()
}

func (o *userGroupImplementation) GetUpdatedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(o.UpdatedAtField.UpdatedAt)
}

func (o *userGroupImplementation) SetUpdatedAt(updatedAt string) UserGroupInterface {
	if updatedAt == "" {
		return o
	}
	o.UpdatedAtField.UpdatedAt = carbon.Parse(updatedAt, carbon.UTC).StdTime()
	return o
}
