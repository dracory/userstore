package userstore

import "errors"

// UserGroupQueryInterface defines the interface for user group queries
type UserGroupQueryInterface interface {
	// Validate validates the query
	Validate() error

	// Columns returns the columns to return from the database
	Columns() []string
	// SetColumns sets the columns to return from the database
	SetColumns(columns []string) UserGroupQueryInterface

	// HasCountOnly returns true if the query is for counting only
	HasCountOnly() bool
	// IsCountOnly returns true if the query is for counting only
	IsCountOnly() bool
	// SetCountOnly sets the query to count only
	SetCountOnly(countOnly bool) UserGroupQueryInterface

	// HasID returns true if the query has an ID
	HasID() bool
	// GetID returns the ID
	GetID() string
	// SetID sets the ID
	SetID(id string) UserGroupQueryInterface

	// HasIDIn returns true if the query has an ID in
	HasIDIn() bool
	// IDIn returns the ID in
	IDIn() []string
	// SetIDIn sets the ID in
	SetIDIn(idIn []string) UserGroupQueryInterface

	// HasUserID returns true if the query has a user ID
	HasUserID() bool
	// UserID returns the user ID
	UserID() string
	// SetUserID sets the user ID
	SetUserID(userID string) UserGroupQueryInterface

	// HasUserIDIn returns true if the query has a user ID in
	HasUserIDIn() bool
	// UserIDIn returns the user ID in
	UserIDIn() []string
	// SetUserIDIn sets the user ID in
	SetUserIDIn(userIDIn []string) UserGroupQueryInterface

	// HasGroupID returns true if the query has a group ID
	HasGroupID() bool
	// GroupID returns the group ID
	GroupID() string
	// SetGroupID sets the group ID
	SetGroupID(groupID string) UserGroupQueryInterface

	// HasGroupIDIn returns true if the query has a group ID in
	HasGroupIDIn() bool
	// GroupIDIn returns the group ID in
	GroupIDIn() []string
	// SetGroupIDIn sets the group ID in
	SetGroupIDIn(groupIDIn []string) UserGroupQueryInterface

	// HasLimit returns true if the query has a limit
	HasLimit() bool
	// Limit returns the limit
	Limit() int
	// SetLimit sets the limit
	SetLimit(limit int) UserGroupQueryInterface

	// HasOffset returns true if the query has an offset
	HasOffset() bool
	// Offset returns the offset
	Offset() int
	// SetOffset sets the offset
	SetOffset(offset int) UserGroupQueryInterface

	// HasOrderBy returns true if the query has an order by
	HasOrderBy() bool
	// OrderBy returns the order by
	OrderBy() string
	// SetOrderBy sets the order by
	SetOrderBy(orderBy string) UserGroupQueryInterface

	// HasSortDirection returns true if the query has a sort direction
	HasSortDirection() bool
	// SortDirection returns the sort direction
	SortDirection() string
	// SetSortDirection sets the sort direction
	SetSortDirection(sortDirection string) UserGroupQueryInterface

	// HasSoftDeletedIncluded returns true if the query has soft deleted included
	HasSoftDeletedIncluded() bool
	// SoftDeletedIncluded returns the soft deleted included
	SoftDeletedIncluded() bool
	// SetSoftDeletedIncluded sets the soft deleted included
	SetSoftDeletedIncluded(softDeletedIncluded bool) UserGroupQueryInterface

	// hasProperty returns true if the query has a property
	hasProperty(name string) bool
}

func NewUserGroupQuery() UserGroupQueryInterface {
	return &userGroupQueryImplementation{
		properties: make(map[string]any),
	}
}

type userGroupQueryImplementation struct {
	properties map[string]any
}

func (c *userGroupQueryImplementation) Validate() error {
	if c.HasID() && c.GetID() == "" {
		return errors.New("user group query. id cannot be empty")
	}

	if c.HasIDIn() && len(c.IDIn()) == 0 {
		return errors.New("user group query. id_in cannot be empty")
	}

	if c.HasUserID() && c.UserID() == "" {
		return errors.New("user group query. user_id cannot be empty")
	}

	if c.HasUserIDIn() && len(c.UserIDIn()) == 0 {
		return errors.New("user group query. user_id_in cannot be empty")
	}

	if c.HasGroupID() && c.GroupID() == "" {
		return errors.New("user group query. group_id cannot be empty")
	}

	if c.HasGroupIDIn() && len(c.GroupIDIn()) == 0 {
		return errors.New("user group query. group_id_in cannot be empty")
	}

	if c.HasOrderBy() && c.OrderBy() == "" {
		return errors.New("user group query. order_by cannot be empty")
	}

	if c.HasSortDirection() && c.SortDirection() == "" {
		return errors.New("user group query. sort_direction cannot be empty")
	}

	if c.HasLimit() && c.Limit() <= 0 {
		return errors.New("user group query. limit must be greater than 0")
	}

	if c.HasOffset() && c.Offset() < 0 {
		return errors.New("user group query. offset must be greater than or equal to 0")
	}

	return nil
}

func (c *userGroupQueryImplementation) Columns() []string {
	if !c.hasProperty("columns") {
		return []string{}
	}

	return c.properties["columns"].([]string)
}

func (c *userGroupQueryImplementation) SetColumns(columns []string) UserGroupQueryInterface {
	c.properties["columns"] = columns

	return c
}

func (c *userGroupQueryImplementation) HasCountOnly() bool {
	return c.hasProperty("count_only")
}

func (c *userGroupQueryImplementation) IsCountOnly() bool {
	if !c.HasCountOnly() {
		return false
	}

	return c.properties["count_only"].(bool)
}

func (c *userGroupQueryImplementation) SetCountOnly(countOnly bool) UserGroupQueryInterface {
	c.properties["count_only"] = countOnly

	return c
}

func (c *userGroupQueryImplementation) HasID() bool {
	return c.hasProperty("id")
}

func (c *userGroupQueryImplementation) GetID() string {
	if !c.HasID() {
		return ""
	}

	return c.properties["id"].(string)
}

func (c *userGroupQueryImplementation) SetID(id string) UserGroupQueryInterface {
	c.properties["id"] = id

	return c
}

func (c *userGroupQueryImplementation) HasIDIn() bool {
	return c.hasProperty("id_in")
}

func (c *userGroupQueryImplementation) IDIn() []string {
	if !c.HasIDIn() {
		return []string{}
	}

	return c.properties["id_in"].([]string)
}

func (c *userGroupQueryImplementation) SetIDIn(idIn []string) UserGroupQueryInterface {
	c.properties["id_in"] = idIn

	return c
}

func (c *userGroupQueryImplementation) HasUserID() bool {
	return c.hasProperty("user_id")
}

func (c *userGroupQueryImplementation) UserID() string {
	if !c.HasUserID() {
		return ""
	}

	return c.properties["user_id"].(string)
}

func (c *userGroupQueryImplementation) SetUserID(userID string) UserGroupQueryInterface {
	c.properties["user_id"] = userID

	return c
}

func (c *userGroupQueryImplementation) HasUserIDIn() bool {
	return c.hasProperty("user_id_in")
}

func (c *userGroupQueryImplementation) UserIDIn() []string {
	if !c.HasUserIDIn() {
		return []string{}
	}

	return c.properties["user_id_in"].([]string)
}

func (c *userGroupQueryImplementation) SetUserIDIn(userIDIn []string) UserGroupQueryInterface {
	c.properties["user_id_in"] = userIDIn

	return c
}

func (c *userGroupQueryImplementation) HasGroupID() bool {
	return c.hasProperty("group_id")
}

func (c *userGroupQueryImplementation) GroupID() string {
	if !c.HasGroupID() {
		return ""
	}

	return c.properties["group_id"].(string)
}

func (c *userGroupQueryImplementation) SetGroupID(groupID string) UserGroupQueryInterface {
	c.properties["group_id"] = groupID

	return c
}

func (c *userGroupQueryImplementation) HasGroupIDIn() bool {
	return c.hasProperty("group_id_in")
}

func (c *userGroupQueryImplementation) GroupIDIn() []string {
	if !c.HasGroupIDIn() {
		return []string{}
	}

	return c.properties["group_id_in"].([]string)
}

func (c *userGroupQueryImplementation) SetGroupIDIn(groupIDIn []string) UserGroupQueryInterface {
	c.properties["group_id_in"] = groupIDIn

	return c
}

func (c *userGroupQueryImplementation) HasLimit() bool {
	return c.hasProperty("limit")
}

func (c *userGroupQueryImplementation) Limit() int {
	if !c.HasLimit() {
		return 0
	}

	return c.properties["limit"].(int)
}

func (c *userGroupQueryImplementation) SetLimit(limit int) UserGroupQueryInterface {
	c.properties["limit"] = limit

	return c
}

func (c *userGroupQueryImplementation) HasOffset() bool {
	return c.hasProperty("offset")
}

func (c *userGroupQueryImplementation) Offset() int {
	if !c.HasOffset() {
		return 0
	}

	return c.properties["offset"].(int)
}

func (c *userGroupQueryImplementation) SetOffset(offset int) UserGroupQueryInterface {
	c.properties["offset"] = offset

	return c
}

func (c *userGroupQueryImplementation) HasOrderBy() bool {
	return c.hasProperty("order_by")
}

func (c *userGroupQueryImplementation) OrderBy() string {
	if !c.HasOrderBy() {
		return ""
	}

	return c.properties["order_by"].(string)
}

func (c *userGroupQueryImplementation) SetOrderBy(orderBy string) UserGroupQueryInterface {
	c.properties["order_by"] = orderBy

	return c
}

func (c *userGroupQueryImplementation) HasSortDirection() bool {
	return c.hasProperty("sort_direction")
}

func (c *userGroupQueryImplementation) SortDirection() string {
	if !c.HasSortDirection() {
		return ""
	}

	return c.properties["sort_direction"].(string)
}

func (c *userGroupQueryImplementation) SetSortDirection(sortDirection string) UserGroupQueryInterface {
	c.properties["sort_direction"] = sortDirection

	return c
}

func (c *userGroupQueryImplementation) HasSoftDeletedIncluded() bool {
	return c.hasProperty("soft_deleted_included")
}

func (c *userGroupQueryImplementation) SoftDeletedIncluded() bool {
	if !c.HasSoftDeletedIncluded() {
		return false
	}

	return c.properties["soft_deleted_included"].(bool)
}

func (c *userGroupQueryImplementation) SetSoftDeletedIncluded(softDeletedIncluded bool) UserGroupQueryInterface {
	c.properties["soft_deleted_included"] = softDeletedIncluded

	return c
}

func (c *userGroupQueryImplementation) hasProperty(name string) bool {
	_, ok := c.properties[name]
	return ok
}
