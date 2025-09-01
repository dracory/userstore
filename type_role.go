package userstore

import (
	"encoding/json"

	"github.com/dracory/dataobject"
	"github.com/dracory/uid"
	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/sb"
)

// == CLASS ===================================================================

type role struct {
	dataobject.DataObject
}

var _ RoleInterface = (*role)(nil)

// == CONSTRUCTORS ============================================================

func NewRole() RoleInterface {
	o := (&role{}).
		SetID(uid.HumanUid()).
		SetStatus(USER_STATUS_UNVERIFIED).
		SetMemo("").
		SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetSoftDeletedAt(sb.MAX_DATETIME)

	err := o.SetMetas(map[string]string{})

	if err != nil {
		return o
	}

	return o
}

func NewRoleFromExistingData(data map[string]string) RoleInterface {
	o := &role{}
	o.Hydrate(data)
	return o
}

// == METHODS =================================================================

func RoleNoImageUrl() string {
	return "/role/default.png"
}

func (o *role) IsActive() bool {
	return o.Status() == USER_STATUS_ACTIVE
}

func (o *role) IsSoftDeleted() bool {
	return o.SoftDeletedAtCarbon().Compare("<", carbon.Now(carbon.UTC))
}

func (o *role) IsInactive() bool {
	return o.Status() == USER_STATUS_INACTIVE
}

// == SETTERS AND GETTERS =====================================================

func (o *role) CreatedAt() string {
	return o.Get(COLUMN_CREATED_AT)
}

func (o *role) CreatedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.CreatedAt(), carbon.UTC)
}

func (o *role) SetCreatedAt(createdAt string) RoleInterface {
	o.Set(COLUMN_CREATED_AT, createdAt)
	return o
}

func (o *role) Handle() string {
	return o.Get(COLUMN_HANDLE)
}

func (o *role) SetHandle(handle string) RoleInterface {
	o.Set(COLUMN_HANDLE, handle)
	return o
}

func (o *role) ID() string {
	return o.Get(COLUMN_ID)
}

func (o *role) SetID(id string) RoleInterface {
	o.Set(COLUMN_ID, id)
	return o
}

func (o *role) Memo() string {
	return o.Get(COLUMN_MEMO)
}

func (o *role) SetMemo(memo string) RoleInterface {
	o.Set(COLUMN_MEMO, memo)
	return o
}

func (o *role) Metas() (map[string]string, error) {
	metasStr := o.Get(COLUMN_METAS)

	if metasStr == "" {
		metasStr = "{}"
	}

	metasJson := map[string]string{}
	errJson := json.Unmarshal([]byte(metasStr), &metasJson)
	if errJson != nil {
		return map[string]string{}, errJson
	}

	return metasJson, nil
}

func (o *role) Meta(name string) string {
	metas, err := o.Metas()

	if err != nil {
		return ""
	}

	if value, exists := metas[name]; exists {
		return value
	}

	return ""
}

func (o *role) SetMeta(name, value string) error {
	return o.UpsertMetas(map[string]string{name: value})
}

// SetMetas stores metas as json string
// Warning: it overwrites any existing metas
func (o *role) SetMetas(metas map[string]string) error {
	mapString, err := json.Marshal(metas)
	if err != nil {
		return err
	}
	o.Set(COLUMN_METAS, string(mapString))
	return nil
}

func (o *role) UpsertMetas(metas map[string]string) error {
	currentMetas, err := o.Metas()

	if err != nil {
		return err
	}

	for k, v := range metas {
		currentMetas[k] = v
	}

	return o.SetMetas(currentMetas)
}

func (o *role) Name() string {
	return o.Get(COLUMN_NAME)
}

func (o *role) SetName(name string) RoleInterface {
	o.Set(COLUMN_NAME, name)
	return o
}

func (o *role) SoftDeletedAt() string {
	return o.Get(COLUMN_SOFT_DELETED_AT)
}

func (o *role) SoftDeletedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.SoftDeletedAt(), carbon.UTC)
}

func (o *role) SetSoftDeletedAt(deletedAt string) RoleInterface {
	o.Set(COLUMN_SOFT_DELETED_AT, deletedAt)
	return o
}

func (o *role) Status() string {
	return o.Get(COLUMN_STATUS)
}

func (o *role) SetStatus(status string) RoleInterface {
	o.Set(COLUMN_STATUS, status)
	return o
}

func (o *role) UpdatedAt() string {
	return o.Get(COLUMN_UPDATED_AT)
}

func (o *role) UpdatedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.Get(COLUMN_UPDATED_AT), carbon.UTC)
}

func (o *role) SetUpdatedAt(updatedAt string) RoleInterface {
	o.Set(COLUMN_UPDATED_AT, updatedAt)
	return o
}
