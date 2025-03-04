package userstore

import "errors"

type UserQueryInterface interface {
	Validate() error

	Columns() []string
	SetColumns(columns []string) UserQueryInterface

	HasCountOnly() bool
	IsCountOnly() bool
	SetCountOnly(countOnly bool) UserQueryInterface

	HasCreatedAtGte() bool
	CreatedAtGte() string
	SetCreatedAtGte(createdAtGte string) UserQueryInterface

	HasCreatedAtLte() bool
	CreatedAtLte() string
	SetCreatedAtLte(createdAtLte string) UserQueryInterface

	HasEmail() bool
	Email() string
	SetEmail(email string) UserQueryInterface

	HasID() bool
	ID() string
	SetID(id string) UserQueryInterface

	HasIDIn() bool
	IDIn() []string
	SetIDIn(idIn []string) UserQueryInterface

	HasMetaLike() bool
	MetaLike() string
	SetMetaLike(metaLike string) UserQueryInterface

	HasLimit() bool
	Limit() int
	SetLimit(limit int) UserQueryInterface

	HasOffset() bool
	Offset() int
	SetOffset(offset int) UserQueryInterface

	HasOrderBy() bool
	OrderBy() string
	SetOrderBy(orderBy string) UserQueryInterface

	HasSortDirection() bool
	SortDirection() string
	SetSortDirection(sortDirection string) UserQueryInterface

	HasSoftDeletedIncluded() bool
	SoftDeletedIncluded() bool
	SetSoftDeletedIncluded(softDeletedIncluded bool) UserQueryInterface

	HasStatus() bool
	Status() string
	SetStatus(status string) UserQueryInterface

	HasStatusIn() bool
	StatusIn() []string
	SetStatusIn(statusIn []string) UserQueryInterface

	hasProperty(name string) bool
}

func NewUserQuery() UserQueryInterface {
	return &userQueryImplementation{
		properties: map[string]any{},
	}
}

type userQueryImplementation struct {
	properties map[string]any
}

func (c *userQueryImplementation) Validate() error {
	if c.HasCreatedAtGte() && c.CreatedAtGte() == "" {
		return errors.New("user query. created_at_gte cannot be empty")
	}

	if c.HasCreatedAtLte() && c.CreatedAtLte() == "" {
		return errors.New("user query. created_at_lte cannot be empty")
	}

	if c.HasEmail() && c.Email() == "" {
		return errors.New("user query. email cannot be empty")
	}

	if c.HasID() && c.ID() == "" {
		return errors.New("user query. id cannot be empty")
	}

	if c.HasIDIn() && len(c.IDIn()) == 0 {
		return errors.New("user query. id_in cannot be empty")
	}

	if c.HasMetaLike() && c.MetaLike() == "" {
		return errors.New("user query. meta_like cannot be empty")
	}

	if c.HasStatus() && c.Status() == "" {
		return errors.New("user query. status cannot be empty")
	}

	if c.HasStatusIn() && len(c.StatusIn()) == 0 {
		return errors.New("user query. status_in cannot be empty")
	}

	// if c.HasTitleLike() && c.TitleLike() == "" {
	// 	return errors.New("user query. title_like cannot be empty")
	// }

	if c.HasOrderBy() && c.OrderBy() == "" {
		return errors.New("user query. order_by cannot be empty")
	}

	if c.HasSortDirection() && c.SortDirection() == "" {
		return errors.New("user query. sort_direction cannot be empty")
	}

	if c.HasLimit() && c.Limit() <= 0 {
		return errors.New("user query. limit must be greater than 0")
	}

	if c.HasOffset() && c.Offset() < 0 {
		return errors.New("user query. offset must be greater than or equal to 0")
	}

	return nil
}

func (c *userQueryImplementation) Columns() []string {
	if !c.hasProperty("columns") {
		return []string{}
	}

	return c.properties["columns"].([]string)
}

func (c *userQueryImplementation) SetColumns(columns []string) UserQueryInterface {
	c.properties["columns"] = columns

	return c
}

func (c *userQueryImplementation) HasCountOnly() bool {
	return c.hasProperty("count_only")
}

func (c *userQueryImplementation) IsCountOnly() bool {
	if !c.HasCountOnly() {
		return false
	}

	return c.properties["count_only"].(bool)
}

func (c *userQueryImplementation) SetCountOnly(countOnly bool) UserQueryInterface {
	c.properties["count_only"] = countOnly

	return c
}

func (c *userQueryImplementation) HasCreatedAtGte() bool {
	return c.hasProperty("created_at_gte")
}

func (c *userQueryImplementation) CreatedAtGte() string {
	if !c.HasCreatedAtGte() {
		return ""
	}

	return c.properties["created_at_gte"].(string)
}

func (c *userQueryImplementation) SetCreatedAtGte(createdAtGte string) UserQueryInterface {
	c.properties["created_at_gte"] = createdAtGte

	return c
}

func (c *userQueryImplementation) HasCreatedAtLte() bool {
	return c.hasProperty("created_at_lte")
}

func (c *userQueryImplementation) CreatedAtLte() string {
	if !c.HasCreatedAtLte() {
		return ""
	}

	return c.properties["created_at_lte"].(string)
}

func (c *userQueryImplementation) SetCreatedAtLte(createdAtLte string) UserQueryInterface {
	c.properties["created_at_lte"] = createdAtLte

	return c
}

func (c *userQueryImplementation) HasEmail() bool {
	return c.hasProperty("email")
}

func (c *userQueryImplementation) Email() string {
	if !c.HasEmail() {
		return ""
	}

	return c.properties["email"].(string)
}

func (c *userQueryImplementation) SetEmail(email string) UserQueryInterface {
	c.properties["email"] = email

	return c
}

func (c *userQueryImplementation) HasID() bool {
	return c.hasProperty("id")
}

func (c *userQueryImplementation) ID() string {
	if !c.HasID() {
		return ""
	}

	return c.properties["id"].(string)
}

func (c *userQueryImplementation) SetID(id string) UserQueryInterface {
	c.properties["id"] = id

	return c
}

func (c *userQueryImplementation) HasIDIn() bool {
	return c.hasProperty("id_in")
}

func (c *userQueryImplementation) IDIn() []string {
	if !c.HasIDIn() {
		return []string{}
	}

	return c.properties["id_in"].([]string)
}

func (c *userQueryImplementation) SetIDIn(idIn []string) UserQueryInterface {
	c.properties["id_in"] = idIn

	return c
}

func (c *userQueryImplementation) HasLimit() bool {
	return c.hasProperty("limit")
}

func (c *userQueryImplementation) Limit() int {
	if !c.HasLimit() {
		return 0
	}

	return c.properties["limit"].(int)
}

func (c *userQueryImplementation) SetLimit(limit int) UserQueryInterface {
	c.properties["limit"] = limit

	return c
}

func (c *userQueryImplementation) HasMetaLike() bool {
	return c.hasProperty("meta_like")
}

func (c *userQueryImplementation) MetaLike() string {
	if !c.HasMetaLike() {
		return ""
	}

	return c.properties["meta_like"].(string)
}

func (c *userQueryImplementation) SetMetaLike(metaLike string) UserQueryInterface {
	c.properties["meta_like"] = metaLike

	return c
}

func (c *userQueryImplementation) HasOffset() bool {
	return c.hasProperty("offset")
}

func (c *userQueryImplementation) Offset() int {
	if !c.HasOffset() {
		return 0
	}

	return c.properties["offset"].(int)
}

func (c *userQueryImplementation) SetOffset(offset int) UserQueryInterface {
	c.properties["offset"] = offset

	return c
}

func (c *userQueryImplementation) HasOrderBy() bool {
	return c.hasProperty("order_by")
}

func (c *userQueryImplementation) OrderBy() string {
	if !c.HasOrderBy() {
		return ""
	}

	return c.properties["order_by"].(string)
}

func (c *userQueryImplementation) SetOrderBy(orderBy string) UserQueryInterface {
	c.properties["order_by"] = orderBy

	return c
}

func (c *userQueryImplementation) HasSortDirection() bool {
	return c.hasProperty("sort_direction")
}

func (c *userQueryImplementation) SortDirection() string {
	if !c.HasSortDirection() {
		return ""
	}

	return c.properties["sort_direction"].(string)
}

func (c *userQueryImplementation) SetSortDirection(sortDirection string) UserQueryInterface {
	c.properties["sort_direction"] = sortDirection

	return c
}

func (c *userQueryImplementation) HasSoftDeletedIncluded() bool {
	return c.hasProperty("with_soft_deleted")
}

func (c *userQueryImplementation) SoftDeletedIncluded() bool {
	if !c.HasSoftDeletedIncluded() {
		return false
	}

	return c.properties["with_soft_deleted"].(bool)
}

func (c *userQueryImplementation) SetSoftDeletedIncluded(softDeletedIncluded bool) UserQueryInterface {
	c.properties["with_soft_deleted"] = softDeletedIncluded

	return c
}

func (c *userQueryImplementation) HasStatus() bool {
	return c.hasProperty("status")
}

func (c *userQueryImplementation) Status() string {
	if !c.HasStatus() {
		return ""
	}

	return c.properties["status"].(string)
}

func (c *userQueryImplementation) SetStatus(status string) UserQueryInterface {
	c.properties["status"] = status

	return c
}

func (c *userQueryImplementation) HasStatusIn() bool {
	return c.hasProperty("status_in")
}

func (c *userQueryImplementation) StatusIn() []string {
	if !c.HasStatusIn() {
		return []string{}
	}

	return c.properties["status_in"].([]string)
}

func (c *userQueryImplementation) SetStatusIn(statusIn []string) UserQueryInterface {
	c.properties["status_in"] = statusIn

	return c
}

func (c *userQueryImplementation) hasProperty(name string) bool {
	_, ok := c.properties[name]
	return ok
}
