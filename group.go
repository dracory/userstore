package userstore

import (
	"encoding/json"
	"time"

	"github.com/dracory/neat/database/orm"
	"github.com/dracory/neat/database/soft_delete"
	"github.com/dracory/neat/support/uid"
	"github.com/dromara/carbon/v2"
)

// == TYPE ====================================================================

type groupImplementation struct {
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

var _ GroupInterface = (*groupImplementation)(nil)

// == CONSTRUCTORS ============================================================

func NewGroup() GroupInterface {
	o := (&groupImplementation{}).
		SetID(uid.GenerateShortID()).
		SetStatus(GROUP_STATUS_ACTIVE).
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

func NewGroupFromExistingData(data map[string]string) GroupInterface {
	o := &groupImplementation{}
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

func (o *groupImplementation) IsActive() bool {
	return o.GetStatus() == GROUP_STATUS_ACTIVE
}

func (o *groupImplementation) IsSoftDeleted() bool {
	return o.SoftDeletedAt.Before(time.Now().UTC())
}

func (o *groupImplementation) IsInactive() bool {
	return o.GetStatus() == GROUP_STATUS_INACTIVE
}

// == DATAOBJECT COMPATIBILITY ================================================

// Data returns all fields as a map.
func (o *groupImplementation) Data() map[string]string {
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
func (o *groupImplementation) DataChanged() map[string]string {
	return o.Data()
}

// MarkAsNotDirty is a no-op (dirty tracking disabled).
func (o *groupImplementation) MarkAsNotDirty() {}

// ToMap returns a DB-ready map[string]any of the group.
func (o *groupImplementation) ToMap() map[string]any {
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

// == SETTERS AND GETTERS =====================================================

func (o *groupImplementation) GetCreatedAt() string {
	if o.CreatedAtField.CreatedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(o.CreatedAtField.CreatedAt).ToDateTimeString()
}

func (o *groupImplementation) GetCreatedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(o.CreatedAtField.CreatedAt)
}

func (o *groupImplementation) SetCreatedAt(createdAt string) GroupInterface {
	if createdAt == "" {
		return o
	}
	o.CreatedAtField.CreatedAt = carbon.Parse(createdAt, carbon.UTC).StdTime()
	return o
}

func (o *groupImplementation) GetHandle() string {
	return o.HandleField
}

func (o *groupImplementation) SetHandle(handle string) GroupInterface {
	o.HandleField = handle
	return o
}

func (o *groupImplementation) GetID() string {
	return o.ShortID.ID
}

func (o *groupImplementation) SetID(id string) GroupInterface {
	o.ShortID.ID = id
	return o
}

func (o *groupImplementation) GetMemo() string {
	return o.MemoField
}

func (o *groupImplementation) SetMemo(memo string) GroupInterface {
	o.MemoField = memo
	return o
}

func (o *groupImplementation) GetMetas() (map[string]string, error) {
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

func (o *groupImplementation) GetMeta(name string) string {
	metas, err := o.GetMetas()

	if err != nil {
		return ""
	}

	if value, exists := metas[name]; exists {
		return value
	}

	return ""
}

func (o *groupImplementation) SetMeta(name, value string) error {
	return o.UpsertMetas(map[string]string{name: value})
}

// SetMetas stores metas as json string
// Warning: it overwrites any existing metas
func (o *groupImplementation) SetMetas(metas map[string]string) error {
	mapString, err := json.Marshal(metas)
	if err != nil {
		return err
	}
	o.MetasField = string(mapString)
	return nil
}

func (o *groupImplementation) UpsertMetas(metas map[string]string) error {
	currentMetas, err := o.GetMetas()

	if err != nil {
		return err
	}

	for k, v := range metas {
		currentMetas[k] = v
	}

	return o.SetMetas(currentMetas)
}

func (o *groupImplementation) GetName() string {
	return o.NameField
}

func (o *groupImplementation) SetName(name string) GroupInterface {
	o.NameField = name
	return o
}

func (o *groupImplementation) GetSoftDeletedAt() string {
	if o.SoftDeletedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(o.SoftDeletedAt).ToDateTimeString()
}

func (o *groupImplementation) GetSoftDeletedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(o.SoftDeletedAt)
}

func (o *groupImplementation) SetSoftDeletedAt(deletedAt string) GroupInterface {
	if deletedAt == "" {
		return o
	}
	o.SoftDeletedAt = carbon.Parse(deletedAt, carbon.UTC).StdTime()
	return o
}

func (o *groupImplementation) GetStatus() string {
	return o.StatusField
}

func (o *groupImplementation) SetStatus(status string) GroupInterface {
	o.StatusField = status
	return o
}

func (o *groupImplementation) GetUpdatedAt() string {
	if o.UpdatedAtField.UpdatedAt.IsZero() {
		return ""
	}
	return carbon.CreateFromStdTime(o.UpdatedAtField.UpdatedAt).ToDateTimeString()
}

func (o *groupImplementation) GetUpdatedAtCarbon() *carbon.Carbon {
	return carbon.CreateFromStdTime(o.UpdatedAtField.UpdatedAt)
}

func (o *groupImplementation) SetUpdatedAt(updatedAt string) GroupInterface {
	if updatedAt == "" {
		return o
	}
	o.UpdatedAtField.UpdatedAt = carbon.Parse(updatedAt, carbon.UTC).StdTime()
	return o
}
