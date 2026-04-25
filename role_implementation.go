package userstore

import (
	"encoding/json"

	"github.com/dracory/dataobject"
	"github.com/dracory/sb"
	"github.com/dromara/carbon/v2"
)

// == CLASS ===================================================================

type roleImplementation struct {
	dataobject.DataObject
}

var _ RoleInterface = (*roleImplementation)(nil)

// == CONSTRUCTORS ============================================================

func NewRole() RoleInterface {
	o := (&roleImplementation{}).
		SetID(GenerateShortID()).
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
	o := &roleImplementation{}
	o.Hydrate(data)
	return o
}

// == METHODS =================================================================

func RoleNoImageUrl() string {
	return "/role/default.png"
}

func (o *roleImplementation) IsActive() bool {
	return o.Status() == USER_STATUS_ACTIVE
}

func (o *roleImplementation) IsSoftDeleted() bool {
	return o.SoftDeletedAtCarbon().Compare("<", carbon.Now(carbon.UTC))
}

func (o *roleImplementation) IsInactive() bool {
	return o.Status() == USER_STATUS_INACTIVE
}

// == SETTERS AND GETTERS =====================================================

func (o *roleImplementation) CreatedAt() string {
	return o.Get(COLUMN_CREATED_AT)
}

func (o *roleImplementation) CreatedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.CreatedAt(), carbon.UTC)
}

func (o *roleImplementation) SetCreatedAt(createdAt string) RoleInterface {
	o.Set(COLUMN_CREATED_AT, createdAt)
	return o
}

func (o *roleImplementation) Handle() string {
	return o.Get(COLUMN_HANDLE)
}

func (o *roleImplementation) SetHandle(handle string) RoleInterface {
	o.Set(COLUMN_HANDLE, handle)
	return o
}

func (o *roleImplementation) ID() string {
	return o.Get(COLUMN_ID)
}

func (o *roleImplementation) SetID(id string) RoleInterface {
	o.Set(COLUMN_ID, id)
	return o
}

func (o *roleImplementation) Memo() string {
	return o.Get(COLUMN_MEMO)
}

func (o *roleImplementation) SetMemo(memo string) RoleInterface {
	o.Set(COLUMN_MEMO, memo)
	return o
}

func (o *roleImplementation) Metas() (map[string]string, error) {
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

func (o *roleImplementation) Meta(name string) string {
	metas, err := o.Metas()

	if err != nil {
		return ""
	}

	if value, exists := metas[name]; exists {
		return value
	}

	return ""
}

func (o *roleImplementation) SetMeta(name, value string) error {
	return o.UpsertMetas(map[string]string{name: value})
}

// SetMetas stores metas as json string
// Warning: it overwrites any existing metas
func (o *roleImplementation) SetMetas(metas map[string]string) error {
	mapString, err := json.Marshal(metas)
	if err != nil {
		return err
	}
	o.Set(COLUMN_METAS, string(mapString))
	return nil
}

func (o *roleImplementation) UpsertMetas(metas map[string]string) error {
	currentMetas, err := o.Metas()

	if err != nil {
		return err
	}

	for k, v := range metas {
		currentMetas[k] = v
	}

	return o.SetMetas(currentMetas)
}

func (o *roleImplementation) Name() string {
	return o.Get(COLUMN_NAME)
}

func (o *roleImplementation) SetName(name string) RoleInterface {
	o.Set(COLUMN_NAME, name)
	return o
}

func (o *roleImplementation) SoftDeletedAt() string {
	return o.Get(COLUMN_SOFT_DELETED_AT)
}

func (o *roleImplementation) SoftDeletedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.SoftDeletedAt(), carbon.UTC)
}

func (o *roleImplementation) SetSoftDeletedAt(deletedAt string) RoleInterface {
	o.Set(COLUMN_SOFT_DELETED_AT, deletedAt)
	return o
}

func (o *roleImplementation) Status() string {
	return o.Get(COLUMN_STATUS)
}

func (o *roleImplementation) SetStatus(status string) RoleInterface {
	o.Set(COLUMN_STATUS, status)
	return o
}

func (o *roleImplementation) UpdatedAt() string {
	return o.Get(COLUMN_UPDATED_AT)
}

func (o *roleImplementation) UpdatedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.Get(COLUMN_UPDATED_AT), carbon.UTC)
}

func (o *roleImplementation) SetUpdatedAt(updatedAt string) RoleInterface {
	o.Set(COLUMN_UPDATED_AT, updatedAt)
	return o
}
