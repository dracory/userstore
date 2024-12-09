package userstore

import "errors"

type RoleQueryInterface interface {
	Validate() error

	Columns() []string
	SetColumns(columns []string) RoleQueryInterface

	HasCountOnly() bool
	IsCountOnly() bool
	SetCountOnly(countOnly bool) RoleQueryInterface

	HasHandle() bool
	Handle() string
	SetHandle(handle string) RoleQueryInterface

	HasID() bool
	ID() string
	SetID(id string) RoleQueryInterface

	HasLimit() bool
	Limit() int
	SetLimit(limit int) RoleQueryInterface

	HasOffset() bool
	Offset() int
	SetOffset(offset int) RoleQueryInterface

	HasOrderBy() bool
	OrderBy() string
	SetOrderBy(orderBy string) RoleQueryInterface

	HasSortDirection() bool
	SortDirection() string
	SetSortDirection(sortDirection string) RoleQueryInterface

	HasSoftDeletedIncluded() bool
	SoftDeletedIncluded() bool
	SetSoftDeletedIncluded(softDeletedIncluded bool) RoleQueryInterface

	HasStatus() bool
	Status() string
	SetStatus(status string) RoleQueryInterface

	HasTitleLike() bool
	TitleLike() string
	SetTitleLike(titleLike string) RoleQueryInterface

	hasProperty(name string) bool
}

func NewRoleQuery() RoleQueryInterface {
	return &roleQueryImplementation{
		properties: make(map[string]any),
	}
}

type roleQueryImplementation struct {
	properties map[string]any
}

func (c *roleQueryImplementation) Validate() error {
	if c.HasID() && c.ID() == "" {
		return errors.New("category query. id cannot be empty")
	}

	if c.HasIDIn() && len(c.IDIn()) == 0 {
		return errors.New("category query. id_in cannot be empty")
	}

	if c.HasParentID() && c.ParentID() == "" {
		return errors.New("category query. parent_id cannot be empty")
	}

	if c.HasStatus() && c.Status() == "" {
		return errors.New("category query. status cannot be empty")
	}

	if c.HasTitleLike() && c.TitleLike() == "" {
		return errors.New("category query. title_like cannot be empty")
	}

	if c.HasOrderBy() && c.OrderBy() == "" {
		return errors.New("category query. order_by cannot be empty")
	}

	if c.HasSortDirection() && c.SortDirection() == "" {
		return errors.New("category query. sort_direction cannot be empty")
	}

	if c.HasLimit() && c.Limit() <= 0 {
		return errors.New("category query. limit must be greater than 0")
	}

	if c.HasOffset() && c.Offset() < 0 {
		return errors.New("category query. offset must be greater than or equal to 0")
	}

	return nil
}

func (c *roleQueryImplementation) Columns() []string {
	if !c.hasProperty("columns") {
		return []string{}
	}

	return c.properties["columns"].([]string)
}

func (c *roleQueryImplementation) SetColumns(columns []string) RoleQueryInterface {
	c.properties["columns"] = columns

	return c
}

func (c *roleQueryImplementation) HasCountOnly() bool {
	return c.hasProperty("count_only")
}

func (c *roleQueryImplementation) IsCountOnly() bool {
	if !c.HasCountOnly() {
		return false
	}

	return c.properties["count_only"].(bool)
}

func (c *roleQueryImplementation) SetCountOnly(countOnly bool) RoleQueryInterface {
	c.properties["count_only"] = countOnly

	return c
}

func (c *roleQueryImplementation) HasID() bool {
	return c.hasProperty("id")
}

func (c *roleQueryImplementation) HasHandle() bool {
	return c.hasProperty("handle")
}

func (c *roleQueryImplementation) Handle() string {
	if !c.HasHandle() {
		return ""
	}

	return c.properties["handle"].(string)
}

func (c *roleQueryImplementation) SetHandle(handle string) RoleQueryInterface {
	c.properties["handle"] = handle

	return c
}

func (c *roleQueryImplementation) ID() string {
	if !c.HasID() {
		return ""
	}

	return c.properties["id"].(string)
}

func (c *roleQueryImplementation) SetID(id string) RoleQueryInterface {
	c.properties["id"] = id

	return c
}

func (c *roleQueryImplementation) HasIDIn() bool {
	return c.hasProperty("id_in")
}

func (c *roleQueryImplementation) IDIn() []string {
	if !c.HasIDIn() {
		return []string{}
	}

	return c.properties["id_in"].([]string)
}

func (c *roleQueryImplementation) SetIDIn(idIn []string) RoleQueryInterface {
	c.properties["id_in"] = idIn

	return c
}

func (c *roleQueryImplementation) HasLimit() bool {
	return c.hasProperty("limit")
}

func (c *roleQueryImplementation) Limit() int {
	if !c.HasLimit() {
		return 0
	}

	return c.properties["limit"].(int)
}

func (c *roleQueryImplementation) SetLimit(limit int) RoleQueryInterface {
	c.properties["limit"] = limit

	return c
}

func (c *roleQueryImplementation) HasOffset() bool {
	return c.hasProperty("offset")
}

func (c *roleQueryImplementation) Offset() int {
	if !c.HasOffset() {
		return 0
	}

	return c.properties["offset"].(int)
}

func (c *roleQueryImplementation) SetOffset(offset int) RoleQueryInterface {
	c.properties["offset"] = offset

	return c
}

func (c *roleQueryImplementation) HasOrderBy() bool {
	return c.hasProperty("order_by")
}

func (c *roleQueryImplementation) OrderBy() string {
	if !c.HasOrderBy() {
		return ""
	}

	return c.properties["order_by"].(string)
}

func (c *roleQueryImplementation) SetOrderBy(orderBy string) RoleQueryInterface {
	c.properties["order_by"] = orderBy

	return c
}

func (c *roleQueryImplementation) HasSortDirection() bool {
	return c.hasProperty("sort_direction")
}

func (c *roleQueryImplementation) SortDirection() string {
	if !c.HasSortDirection() {
		return ""
	}

	return c.properties["sort_direction"].(string)
}

func (c *roleQueryImplementation) SetSortDirection(sortDirection string) RoleQueryInterface {
	c.properties["sort_direction"] = sortDirection

	return c
}

func (c *roleQueryImplementation) HasParentID() bool {
	return c.hasProperty("parent_id")
}

func (c *roleQueryImplementation) ParentID() string {
	if !c.HasParentID() {
		return ""
	}

	return c.properties["parent_id"].(string)
}

func (c *roleQueryImplementation) SetParentID(parentID string) RoleQueryInterface {
	c.properties["parent_id"] = parentID

	return c
}

func (c *roleQueryImplementation) HasSoftDeletedIncluded() bool {
	return c.hasProperty("soft_deleted_included")
}

func (c *roleQueryImplementation) SoftDeletedIncluded() bool {
	if !c.HasSoftDeletedIncluded() {
		return false
	}

	return c.properties["soft_deleted_included"].(bool)
}

func (c *roleQueryImplementation) SetSoftDeletedIncluded(softDeletedIncluded bool) RoleQueryInterface {
	c.properties["soft_deleted_included"] = softDeletedIncluded

	return c
}

func (c *roleQueryImplementation) HasStatus() bool {
	return c.hasProperty("status")
}

func (c *roleQueryImplementation) Status() string {
	if !c.HasStatus() {
		return ""
	}

	return c.properties["status"].(string)
}

func (c *roleQueryImplementation) SetStatus(status string) RoleQueryInterface {
	c.properties["status"] = status

	return c
}

func (c *roleQueryImplementation) HasTitleLike() bool {
	return c.hasProperty("title_like")
}

func (c *roleQueryImplementation) TitleLike() string {
	if !c.HasTitleLike() {
		return ""
	}

	return c.properties["title_like"].(string)
}

func (c *roleQueryImplementation) SetTitleLike(titleLike string) RoleQueryInterface {
	c.properties["title_like"] = titleLike

	return c
}

func (c *roleQueryImplementation) hasProperty(name string) bool {
	_, ok := c.properties[name]
	return ok
}
