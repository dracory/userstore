package userstore

import "errors"

// UserRoleQueryInterface defines the interface for user role queries
type UserRoleQueryInterface interface {
	// Validate validates the query
	Validate() error

	// Columns returns the columns to return from the database
	Columns() []string
	// SetColumns sets the columns to return from the database
	SetColumns(columns []string) UserRoleQueryInterface

	// HasCountOnly returns true if the query is for counting only
	HasCountOnly() bool
	// IsCountOnly returns true if the query is for counting only
	IsCountOnly() bool
	// SetCountOnly sets the query to count only
	SetCountOnly(countOnly bool) UserRoleQueryInterface

	// HasID returns true if the query has an ID
	HasID() bool
	// GetID returns the ID
	GetID() string
	// SetID sets the ID
	SetID(id string) UserRoleQueryInterface

	// HasIDIn returns true if the query has an ID in
	HasIDIn() bool
	// IDIn returns the ID in
	IDIn() []string
	// SetIDIn sets the ID in
	SetIDIn(idIn []string) UserRoleQueryInterface

	// HasUserID returns true if the query has a user ID
	HasUserID() bool
	// UserID returns the user ID
	UserID() string
	// SetUserID sets the user ID
	SetUserID(userID string) UserRoleQueryInterface

	// HasUserIDIn returns true if the query has a user ID in
	HasUserIDIn() bool
	// UserIDIn returns the user ID in
	UserIDIn() []string
	// SetUserIDIn sets the user ID in
	SetUserIDIn(userIDIn []string) UserRoleQueryInterface

	// HasRoleID returns true if the query has a role ID
	HasRoleID() bool
	// RoleID returns the role ID
	RoleID() string
	// SetRoleID sets the role ID
	SetRoleID(roleID string) UserRoleQueryInterface

	// HasRoleIDIn returns true if the query has a role ID in
	HasRoleIDIn() bool
	// RoleIDIn returns the role ID in
	RoleIDIn() []string
	// SetRoleIDIn sets the role ID in
	SetRoleIDIn(roleIDIn []string) UserRoleQueryInterface

	// HasLimit returns true if the query has a limit
	HasLimit() bool
	// Limit returns the limit
	Limit() int
	// SetLimit sets the limit
	SetLimit(limit int) UserRoleQueryInterface

	// HasOffset returns true if the query has an offset
	HasOffset() bool
	// Offset returns the offset
	Offset() int
	// SetOffset sets the offset
	SetOffset(offset int) UserRoleQueryInterface

	// HasOrderBy returns true if the query has an order by
	HasOrderBy() bool
	// OrderBy returns the order by
	OrderBy() string
	// SetOrderBy sets the order by
	SetOrderBy(orderBy string) UserRoleQueryInterface

	// HasSortDirection returns true if the query has a sort direction
	HasSortDirection() bool
	// SortDirection returns the sort direction
	SortDirection() string
	// SetSortDirection sets the sort direction
	SetSortDirection(sortDirection string) UserRoleQueryInterface

	// HasSoftDeletedIncluded returns true if the query has soft deleted included
	HasSoftDeletedIncluded() bool
	// SoftDeletedIncluded returns the soft deleted included
	SoftDeletedIncluded() bool
	// SetSoftDeletedIncluded sets the soft deleted included
	SetSoftDeletedIncluded(softDeletedIncluded bool) UserRoleQueryInterface

	// hasProperty returns true if the query has a property
	hasProperty(name string) bool
}

func NewUserRoleQuery() UserRoleQueryInterface {
	return &userRoleQueryImplementation{
		properties: make(map[string]any),
	}
}

type userRoleQueryImplementation struct {
	properties map[string]any
}

func (c *userRoleQueryImplementation) Validate() error {
	if c.HasID() && c.GetID() == "" {
		return errors.New("user role query. id cannot be empty")
	}

	if c.HasIDIn() && len(c.IDIn()) == 0 {
		return errors.New("user role query. id_in cannot be empty")
	}

	if c.HasUserID() && c.UserID() == "" {
		return errors.New("user role query. user_id cannot be empty")
	}

	if c.HasUserIDIn() && len(c.UserIDIn()) == 0 {
		return errors.New("user role query. user_id_in cannot be empty")
	}

	if c.HasRoleID() && c.RoleID() == "" {
		return errors.New("user role query. role_id cannot be empty")
	}

	if c.HasRoleIDIn() && len(c.RoleIDIn()) == 0 {
		return errors.New("user role query. role_id_in cannot be empty")
	}

	if c.HasOrderBy() && c.OrderBy() == "" {
		return errors.New("user role query. order_by cannot be empty")
	}

	if c.HasSortDirection() && c.SortDirection() == "" {
		return errors.New("user role query. sort_direction cannot be empty")
	}

	if c.HasLimit() && c.Limit() <= 0 {
		return errors.New("user role query. limit must be greater than 0")
	}

	if c.HasOffset() && c.Offset() < 0 {
		return errors.New("user role query. offset must be greater than or equal to 0")
	}

	return nil
}

func (c *userRoleQueryImplementation) Columns() []string {
	if !c.hasProperty("columns") {
		return []string{}
	}

	return c.properties["columns"].([]string)
}

func (c *userRoleQueryImplementation) SetColumns(columns []string) UserRoleQueryInterface {
	c.properties["columns"] = columns

	return c
}

func (c *userRoleQueryImplementation) HasCountOnly() bool {
	return c.hasProperty("count_only")
}

func (c *userRoleQueryImplementation) IsCountOnly() bool {
	if !c.HasCountOnly() {
		return false
	}

	return c.properties["count_only"].(bool)
}

func (c *userRoleQueryImplementation) SetCountOnly(countOnly bool) UserRoleQueryInterface {
	c.properties["count_only"] = countOnly

	return c
}

func (c *userRoleQueryImplementation) HasID() bool {
	return c.hasProperty("id")
}

func (c *userRoleQueryImplementation) GetID() string {
	if !c.HasID() {
		return ""
	}

	return c.properties["id"].(string)
}

func (c *userRoleQueryImplementation) SetID(id string) UserRoleQueryInterface {
	c.properties["id"] = id

	return c
}

func (c *userRoleQueryImplementation) HasIDIn() bool {
	return c.hasProperty("id_in")
}

func (c *userRoleQueryImplementation) IDIn() []string {
	if !c.HasIDIn() {
		return []string{}
	}

	return c.properties["id_in"].([]string)
}

func (c *userRoleQueryImplementation) SetIDIn(idIn []string) UserRoleQueryInterface {
	c.properties["id_in"] = idIn

	return c
}

func (c *userRoleQueryImplementation) HasUserID() bool {
	return c.hasProperty("user_id")
}

func (c *userRoleQueryImplementation) UserID() string {
	if !c.HasUserID() {
		return ""
	}

	return c.properties["user_id"].(string)
}

func (c *userRoleQueryImplementation) SetUserID(userID string) UserRoleQueryInterface {
	c.properties["user_id"] = userID

	return c
}

func (c *userRoleQueryImplementation) HasUserIDIn() bool {
	return c.hasProperty("user_id_in")
}

func (c *userRoleQueryImplementation) UserIDIn() []string {
	if !c.HasUserIDIn() {
		return []string{}
	}

	return c.properties["user_id_in"].([]string)
}

func (c *userRoleQueryImplementation) SetUserIDIn(userIDIn []string) UserRoleQueryInterface {
	c.properties["user_id_in"] = userIDIn

	return c
}

func (c *userRoleQueryImplementation) HasRoleID() bool {
	return c.hasProperty("role_id")
}

func (c *userRoleQueryImplementation) RoleID() string {
	if !c.HasRoleID() {
		return ""
	}

	return c.properties["role_id"].(string)
}

func (c *userRoleQueryImplementation) SetRoleID(roleID string) UserRoleQueryInterface {
	c.properties["role_id"] = roleID

	return c
}

func (c *userRoleQueryImplementation) HasRoleIDIn() bool {
	return c.hasProperty("role_id_in")
}

func (c *userRoleQueryImplementation) RoleIDIn() []string {
	if !c.HasRoleIDIn() {
		return []string{}
	}

	return c.properties["role_id_in"].([]string)
}

func (c *userRoleQueryImplementation) SetRoleIDIn(roleIDIn []string) UserRoleQueryInterface {
	c.properties["role_id_in"] = roleIDIn

	return c
}

func (c *userRoleQueryImplementation) HasLimit() bool {
	return c.hasProperty("limit")
}

func (c *userRoleQueryImplementation) Limit() int {
	if !c.HasLimit() {
		return 0
	}

	return c.properties["limit"].(int)
}

func (c *userRoleQueryImplementation) SetLimit(limit int) UserRoleQueryInterface {
	c.properties["limit"] = limit

	return c
}

func (c *userRoleQueryImplementation) HasOffset() bool {
	return c.hasProperty("offset")
}

func (c *userRoleQueryImplementation) Offset() int {
	if !c.HasOffset() {
		return 0
	}

	return c.properties["offset"].(int)
}

func (c *userRoleQueryImplementation) SetOffset(offset int) UserRoleQueryInterface {
	c.properties["offset"] = offset

	return c
}

func (c *userRoleQueryImplementation) HasOrderBy() bool {
	return c.hasProperty("order_by")
}

func (c *userRoleQueryImplementation) OrderBy() string {
	if !c.HasOrderBy() {
		return ""
	}

	return c.properties["order_by"].(string)
}

func (c *userRoleQueryImplementation) SetOrderBy(orderBy string) UserRoleQueryInterface {
	c.properties["order_by"] = orderBy

	return c
}

func (c *userRoleQueryImplementation) HasSortDirection() bool {
	return c.hasProperty("sort_direction")
}

func (c *userRoleQueryImplementation) SortDirection() string {
	if !c.HasSortDirection() {
		return ""
	}

	return c.properties["sort_direction"].(string)
}

func (c *userRoleQueryImplementation) SetSortDirection(sortDirection string) UserRoleQueryInterface {
	c.properties["sort_direction"] = sortDirection

	return c
}

func (c *userRoleQueryImplementation) HasSoftDeletedIncluded() bool {
	return c.hasProperty("soft_deleted_included")
}

func (c *userRoleQueryImplementation) SoftDeletedIncluded() bool {
	if !c.HasSoftDeletedIncluded() {
		return false
	}

	return c.properties["soft_deleted_included"].(bool)
}

func (c *userRoleQueryImplementation) SetSoftDeletedIncluded(softDeletedIncluded bool) UserRoleQueryInterface {
	c.properties["soft_deleted_included"] = softDeletedIncluded

	return c
}

func (c *userRoleQueryImplementation) hasProperty(name string) bool {
	_, ok := c.properties[name]
	return ok
}
