package userstore

import (
	"encoding/json"
	"time"

	"github.com/dracory/neat/database/orm"
	"github.com/dracory/neat/database/soft_delete"
	"github.com/dromara/carbon/v2"
)

// == TYPE ====================================================================

type roleImplementation struct {
	orm.ShortID

	StatusField    string `db:"status"`
	HandleField    string `db:"handle"`
	NameField      string `db:"name"`
	MemoField      string `db:"memo"`
	MetasField     string `db:"metas"`
	CreatedAtField orm.CreatedAt
	UpdatedAtField orm.UpdatedAt
	soft_delete.SoftDeletesMaxDate
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
		SetSoftDeletedAt(MAX_DATETIME)

	err := o.SetMetas(map[string]string{})

	if err != nil {
		return o
	}

	return o
}

func NewRoleFromExistingData(data map[string]string) RoleInterface {
	o := &roleImplementation{}
	if v, ok := data[COLUMN_ID]; ok {
		o.SetID(v)
	}
	if v, ok := data[COLUMN_STATUS]; ok {
		o.SetStatus(v)
	}
	if v, ok := data[COLUMN_HANDLE]; ok {
		o.SetHandle(v)
	}
	if v, ok := data[COLUMN_NAME]; ok {
		o.SetName(v)
	}
	if v, ok := data[COLUMN_MEMO]; ok {
		o.SetMemo(v)
	}
	if v, ok := data[COLUMN_METAS]; ok {
		o.SetMetas(map[string]string{})
		o.UpsertMetas(map[string]string{v: v})
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

func RoleNoImageUrl() string {
	return "/role/default.png"
}

func (o *roleImplementation) IsActive() bool {
	return o.GetStatus() == USER_STATUS_ACTIVE
}

func (o *roleImplementation) IsSoftDeleted() bool {
	return o.SoftDeletedAt.Before(time.Now().UTC())
}

func (o *roleImplementation) IsInactive() bool {
	return o.GetStatus() == USER_STATUS_INACTIVE
}

// == DATAOBJECT COMPATIBILITY ================================================

// Data returns all fields as a map.
func (o *roleImplementation) Data() map[string]string {
	return map[string]string{
		COLUMN_ID:              o.GetID(),
		COLUMN_STATUS:          o.GetStatus(),
		COLUMN_HANDLE:          o.GetHandle(),
		COLUMN_NAME:            o.GetName(),
		COLUMN_MEMO:            o.GetMemo(),
		COLUMN_METAS:           o.MetasField,
		COLUMN_CREATED_AT:      o.GetCreatedAt(),
		COLUMN_UPDATED_AT:      o.GetUpdatedAt(),
		COLUMN_SOFT_DELETED_AT: o.GetSoftDeletedAt(),
	}
}

// DataChanged returns all fields as a map (dirty tracking disabled).
func (o *roleImplementation) DataChanged() map[string]string {
	return o.Data()
}

// MarkAsNotDirty is a no-op (dirty tracking disabled).
func (o *roleImplementation) MarkAsNotDirty() {}

// ToMap returns a DB-ready map[string]any of the role.
func (o *roleImplementation) ToMap() map[string]any {
	return map[string]any{
		COLUMN_ID:              o.GetID(),
		COLUMN_STATUS:          o.GetStatus(),
		COLUMN_HANDLE:          o.GetHandle(),
		COLUMN_NAME:            o.GetName(),
		COLUMN_MEMO:            o.GetMemo(),
		COLUMN_METAS:           o.MetasField,
		COLUMN_CREATED_AT:      o.GetCreatedAtCarbon().StdTime(),
		COLUMN_UPDATED_AT:      o.GetUpdatedAtCarbon().StdTime(),
		COLUMN_SOFT_DELETED_AT: o.GetSoftDeletedAtCarbon().StdTime(),
	}
}

// Get returns the value of the specified column.
func (o *roleImplementation) Get(columnName string) string {
	switch columnName {
	case COLUMN_ID:
		return o.GetID()
	case COLUMN_STATUS:
		return o.GetStatus()
	case COLUMN_HANDLE:
		return o.GetHandle()
	case COLUMN_NAME:
		return o.GetName()
	case COLUMN_MEMO:
		return o.GetMemo()
	case COLUMN_METAS:
		return o.MetasField
	case COLUMN_CREATED_AT:
		return o.GetCreatedAt()
	case COLUMN_UPDATED_AT:
		return o.GetUpdatedAt()
	case COLUMN_SOFT_DELETED_AT:
		return o.GetSoftDeletedAt()
	}
	return ""
}

// Set sets the value of the specified column.
func (o *roleImplementation) Set(columnName string, value string) {
	switch columnName {
	case COLUMN_ID:
		o.SetID(value)
	case COLUMN_STATUS:
		o.SetStatus(value)
	case COLUMN_HANDLE:
		o.SetHandle(value)
	case COLUMN_NAME:
		o.SetName(value)
	case COLUMN_MEMO:
		o.SetMemo(value)
	case COLUMN_METAS:
		o.MetasField = value
	case COLUMN_CREATED_AT:
		o.SetCreatedAt(value)
	case COLUMN_UPDATED_AT:
		o.SetUpdatedAt(value)
	case COLUMN_SOFT_DELETED_AT:
		o.SetSoftDeletedAt(value)
	}
}

// == SETTERS AND GETTERS =====================================================

func (o *roleImplementation) GetCreatedAt() string {
	if o.CreatedAtField.CreatedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(o.CreatedAtField.CreatedAt).ToDateTimeString()
}

func (o *roleImplementation) GetCreatedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(o.CreatedAtField.CreatedAt)
}

func (o *roleImplementation) SetCreatedAt(createdAt string) RoleInterface {
	if createdAt == "" {
		return o
	}
	o.CreatedAtField.CreatedAt = carbon.Parse(createdAt, carbon.UTC).StdTime()
	return o
}

func (o *roleImplementation) GetHandle() string {
	return o.HandleField
}

func (o *roleImplementation) SetHandle(handle string) RoleInterface {
	o.HandleField = handle
	return o
}

func (o *roleImplementation) GetID() string {
	return o.ShortID.ID
}

func (o *roleImplementation) SetID(id string) RoleInterface {
	o.ShortID.ID = id
	return o
}

func (o *roleImplementation) GetMemo() string {
	return o.MemoField
}

func (o *roleImplementation) SetMemo(memo string) RoleInterface {
	o.MemoField = memo
	return o
}

func (o *roleImplementation) GetMetas() (map[string]string, error) {
	metasStr := o.MetasField

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

func (o *roleImplementation) GetMeta(name string) string {
	metas, err := o.GetMetas()

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
	o.MetasField = string(mapString)
	return nil
}

func (o *roleImplementation) UpsertMetas(metas map[string]string) error {
	currentMetas, err := o.GetMetas()

	if err != nil {
		return err
	}

	for k, v := range metas {
		currentMetas[k] = v
	}

	return o.SetMetas(currentMetas)
}

func (o *roleImplementation) GetName() string {
	return o.NameField
}

func (o *roleImplementation) SetName(name string) RoleInterface {
	o.NameField = name
	return o
}

func (o *roleImplementation) GetSoftDeletedAt() string {
	if o.SoftDeletedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(o.SoftDeletedAt).ToDateTimeString()
}

func (o *roleImplementation) GetSoftDeletedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(o.SoftDeletedAt)
}

func (o *roleImplementation) SetSoftDeletedAt(deletedAt string) RoleInterface {
	if deletedAt == "" {
		return o
	}
	o.SoftDeletedAt = carbon.Parse(deletedAt, carbon.UTC).StdTime()
	return o
}

func (o *roleImplementation) GetStatus() string {
	return o.StatusField
}

func (o *roleImplementation) SetStatus(status string) RoleInterface {
	o.StatusField = status
	return o
}

func (o *roleImplementation) GetUpdatedAt() string {
	if o.UpdatedAtField.UpdatedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(o.UpdatedAtField.UpdatedAt).ToDateTimeString()
}

func (o *roleImplementation) GetUpdatedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(o.UpdatedAtField.UpdatedAt)
}

func (o *roleImplementation) SetUpdatedAt(updatedAt string) RoleInterface {
	if updatedAt == "" {
		return o
	}
	o.UpdatedAtField.UpdatedAt = carbon.Parse(updatedAt, carbon.UTC).StdTime()
	return o
}
