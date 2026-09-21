package userstore

import "errors"

// GroupQueryInterface defines the interface for group queries
type GroupQueryInterface interface {
	// Validate validates the query
	Validate() error

	// Columns returns the columns to return from the database
	Columns() []string
	// SetColumns sets the columns to return from the database
	SetColumns(columns []string) GroupQueryInterface

	// HasCountOnly returns true if the query is for counting only
	HasCountOnly() bool
	// IsCountOnly returns true if the query is for counting only
	IsCountOnly() bool
	// SetCountOnly sets the query to count only
	SetCountOnly(countOnly bool) GroupQueryInterface

	// HasHandle returns true if the query has a handle
	HasHandle() bool
	// Handle returns the handle
	Handle() string
	// SetHandle sets the handle
	SetHandle(handle string) GroupQueryInterface

	// HasID returns true if the query has an ID
	HasID() bool
	// GetID returns the ID
	GetID() string
	// SetID sets the ID
	SetID(id string) GroupQueryInterface

	// HasIDIn returns true if the query has an ID in
	HasIDIn() bool
	// IDIn returns the ID in
	IDIn() []string
	// SetIDIn sets the ID in
	SetIDIn(idIn []string) GroupQueryInterface

	// HasLimit returns true if the query has a limit
	HasLimit() bool
	// Limit returns the limit
	Limit() int
	// SetLimit sets the limit
	SetLimit(limit int) GroupQueryInterface

	// HasOffset returns true if the query has an offset
	HasOffset() bool
	// Offset returns the offset
	Offset() int
	// SetOffset sets the offset
	SetOffset(offset int) GroupQueryInterface

	// HasOrderBy returns true if the query has an order by
	HasOrderBy() bool
	// OrderBy returns the order by
	OrderBy() string
	// SetOrderBy sets the order by
	SetOrderBy(orderBy string) GroupQueryInterface

	// HasSortDirection returns true if the query has a sort direction
	HasSortDirection() bool
	// SortDirection returns the sort direction
	SortDirection() string
	// SetSortDirection sets the sort direction
	SetSortDirection(sortDirection string) GroupQueryInterface

	// HasSoftDeletedIncluded returns true if the query has soft deleted included
	HasSoftDeletedIncluded() bool
	// SoftDeletedIncluded returns the soft deleted included
	SoftDeletedIncluded() bool
	// SetSoftDeletedIncluded sets the soft deleted included
	SetSoftDeletedIncluded(softDeletedIncluded bool) GroupQueryInterface

	// HasStatus returns true if the query has a status
	HasStatus() bool
	// Status returns the status
	Status() string
	// SetStatus sets the status
	SetStatus(status string) GroupQueryInterface

	// HasTitleLike returns true if the query has a title like
	HasTitleLike() bool
	// TitleLike returns the title like
	TitleLike() string
	// SetTitleLike sets the title like
	SetTitleLike(titleLike string) GroupQueryInterface

	// hasProperty returns true if the query has a property
	hasProperty(name string) bool
}

func NewGroupQuery() GroupQueryInterface {
	return &groupQueryImplementation{
		properties: make(map[string]any),
	}
}

type groupQueryImplementation struct {
	properties map[string]any
}

func (c *groupQueryImplementation) Validate() error {
	if c.HasID() && c.GetID() == "" {
		return errors.New("group query. id cannot be empty")
	}

	if c.HasIDIn() && len(c.IDIn()) == 0 {
		return errors.New("group query. id_in cannot be empty")
	}

	if c.HasStatus() && c.Status() == "" {
		return errors.New("group query. status cannot be empty")
	}

	if c.HasTitleLike() && c.TitleLike() == "" {
		return errors.New("group query. title_like cannot be empty")
	}

	if c.HasOrderBy() && c.OrderBy() == "" {
		return errors.New("group query. order_by cannot be empty")
	}

	if c.HasSortDirection() && c.SortDirection() == "" {
		return errors.New("group query. sort_direction cannot be empty")
	}

	if c.HasLimit() && c.Limit() <= 0 {
		return errors.New("group query. limit must be greater than 0")
	}

	if c.HasOffset() && c.Offset() < 0 {
		return errors.New("group query. offset must be greater than or equal to 0")
	}

	return nil
}

func (c *groupQueryImplementation) Columns() []string {
	if !c.hasProperty("columns") {
		return []string{}
	}

	return c.properties["columns"].([]string)
}

func (c *groupQueryImplementation) SetColumns(columns []string) GroupQueryInterface {
	c.properties["columns"] = columns

	return c
}

func (c *groupQueryImplementation) HasCountOnly() bool {
	return c.hasProperty("count_only")
}

func (c *groupQueryImplementation) IsCountOnly() bool {
	if !c.HasCountOnly() {
		return false
	}

	return c.properties["count_only"].(bool)
}

func (c *groupQueryImplementation) SetCountOnly(countOnly bool) GroupQueryInterface {
	c.properties["count_only"] = countOnly

	return c
}

func (c *groupQueryImplementation) HasID() bool {
	return c.hasProperty("id")
}

func (c *groupQueryImplementation) HasHandle() bool {
	return c.hasProperty("handle")
}

func (c *groupQueryImplementation) Handle() string {
	if !c.HasHandle() {
		return ""
	}

	return c.properties["handle"].(string)
}

func (c *groupQueryImplementation) SetHandle(handle string) GroupQueryInterface {
	c.properties["handle"] = handle

	return c
}

func (c *groupQueryImplementation) GetID() string {
	if !c.HasID() {
		return ""
	}

	return c.properties["id"].(string)
}

func (c *groupQueryImplementation) SetID(id string) GroupQueryInterface {
	c.properties["id"] = id

	return c
}

func (c *groupQueryImplementation) HasIDIn() bool {
	return c.hasProperty("id_in")
}

func (c *groupQueryImplementation) IDIn() []string {
	if !c.HasIDIn() {
		return []string{}
	}

	return c.properties["id_in"].([]string)
}

func (c *groupQueryImplementation) SetIDIn(idIn []string) GroupQueryInterface {
	c.properties["id_in"] = idIn

	return c
}

func (c *groupQueryImplementation) HasLimit() bool {
	return c.hasProperty("limit")
}

func (c *groupQueryImplementation) Limit() int {
	if !c.HasLimit() {
		return 0
	}

	return c.properties["limit"].(int)
}

func (c *groupQueryImplementation) SetLimit(limit int) GroupQueryInterface {
	c.properties["limit"] = limit

	return c
}

func (c *groupQueryImplementation) HasOffset() bool {
	return c.hasProperty("offset")
}

func (c *groupQueryImplementation) Offset() int {
	if !c.HasOffset() {
		return 0
	}

	return c.properties["offset"].(int)
}

func (c *groupQueryImplementation) SetOffset(offset int) GroupQueryInterface {
	c.properties["offset"] = offset

	return c
}

func (c *groupQueryImplementation) HasOrderBy() bool {
	return c.hasProperty("order_by")
}

func (c *groupQueryImplementation) OrderBy() string {
	if !c.HasOrderBy() {
		return ""
	}

	return c.properties["order_by"].(string)
}

func (c *groupQueryImplementation) SetOrderBy(orderBy string) GroupQueryInterface {
	c.properties["order_by"] = orderBy

	return c
}

func (c *groupQueryImplementation) HasSortDirection() bool {
	return c.hasProperty("sort_direction")
}

func (c *groupQueryImplementation) SortDirection() string {
	if !c.HasSortDirection() {
		return ""
	}

	return c.properties["sort_direction"].(string)
}

func (c *groupQueryImplementation) SetSortDirection(sortDirection string) GroupQueryInterface {
	c.properties["sort_direction"] = sortDirection

	return c
}

func (c *groupQueryImplementation) HasSoftDeletedIncluded() bool {
	return c.hasProperty("soft_deleted_included")
}

func (c *groupQueryImplementation) SoftDeletedIncluded() bool {
	if !c.HasSoftDeletedIncluded() {
		return false
	}

	return c.properties["soft_deleted_included"].(bool)
}

func (c *groupQueryImplementation) SetSoftDeletedIncluded(softDeletedIncluded bool) GroupQueryInterface {
	c.properties["soft_deleted_included"] = softDeletedIncluded

	return c
}

func (c *groupQueryImplementation) HasStatus() bool {
	return c.hasProperty("status")
}

func (c *groupQueryImplementation) Status() string {
	if !c.HasStatus() {
		return ""
	}

	return c.properties["status"].(string)
}

func (c *groupQueryImplementation) SetStatus(status string) GroupQueryInterface {
	c.properties["status"] = status

	return c
}

func (c *groupQueryImplementation) HasTitleLike() bool {
	return c.hasProperty("title_like")
}

func (c *groupQueryImplementation) TitleLike() string {
	if !c.HasTitleLike() {
		return ""
	}

	return c.properties["title_like"].(string)
}

func (c *groupQueryImplementation) SetTitleLike(titleLike string) GroupQueryInterface {
	c.properties["title_like"] = titleLike

	return c
}

func (c *groupQueryImplementation) hasProperty(name string) bool {
	_, ok := c.properties[name]
	return ok
}
