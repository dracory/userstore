package userstore

import "errors"

type UserQueryInterface interface {
	// Validate validates the query
	Validate() error

	// Columns returns the columns to return from the database
	Columns() []string

	// SetColumns sets the columns to return from the database
	SetColumns(columns []string) UserQueryInterface

	// HasCountOnly returns true if the query is for counting only
	HasCountOnly() bool

	// IsCountOnly returns true if the query is for counting only
	IsCountOnly() bool

	// SetCountOnly sets the query to count only
	SetCountOnly(countOnly bool) UserQueryInterface

	// HasCreatedAtGte returns true if the query has a created at greater than or equal to
	HasCreatedAtGte() bool

	// CreatedAtGte returns the created at greater than or equal to
	CreatedAtGte() string

	// SetCreatedAtGte sets the created at greater than or equal to
	SetCreatedAtGte(createdAtGte string) UserQueryInterface

	// HasCreatedAtLte returns true if the query has a created at less than or equal to
	HasCreatedAtLte() bool

	// CreatedAtLte returns the created at less than or equal to
	CreatedAtLte() string

	// SetCreatedAtLte sets the created at less than or equal to
	SetCreatedAtLte(createdAtLte string) UserQueryInterface

	// HasEmail returns true if the query has an email
	HasEmail() bool

	// Email returns the email
	Email() string

	// SetEmail sets the email
	SetEmail(email string) UserQueryInterface

	// HasEmailLike returns true if the query has an email like
	HasEmailLike() bool

	// EmailLike returns the email like
	EmailLike() string

	// SetEmailLike sets the email like
	SetEmailLike(emailLike string) UserQueryInterface

	// HasFirstName returns true if the query has a first name
	HasFirstName() bool

	// FirstName returns the first name
	FirstName() string

	// SetFirstName sets the first name
	SetFirstName(firstName string) UserQueryInterface

	// HasFirstNameLike returns true if the query has a first name like
	HasFirstNameLike() bool

	// FirstNameLike returns the first name like
	FirstNameLike() string

	// SetFirstNameLike sets the first name like
	SetFirstNameLike(firstNameLike string) UserQueryInterface

	// HasLastName returns true if the query has a last name
	HasLastName() bool

	// LastName returns the last name
	LastName() string

	// SetLastName sets the last name
	SetLastName(lastName string) UserQueryInterface

	// HasLastNameLike returns true if the query has a last name like
	HasLastNameLike() bool

	// LastNameLike returns the last name like
	LastNameLike() string

	// SetLastNameLike sets the last name like
	SetLastNameLike(lastNameLike string) UserQueryInterface

	// HasID returns true if the query has an ID
	HasID() bool

	// GetID returns the ID
	GetID() string

	// SetID sets the ID
	SetID(id string) UserQueryInterface

	// HasIDIn returns true if the query has an ID in
	HasIDIn() bool

	// IDIn returns the ID in
	IDIn() []string

	// SetIDIn sets the ID in
	SetIDIn(idIn []string) UserQueryInterface

	// HasMetaLike returns true if the query has a meta like
	HasMetaLike() bool

	// MetaLike returns the meta like
	MetaLike() string

	// SetMetaLike sets the meta like
	SetMetaLike(metaLike string) UserQueryInterface

	// HasLimit returns true if the query has a limit
	HasLimit() bool

	// Limit returns the limit
	Limit() int

	// SetLimit sets the limit
	SetLimit(limit int) UserQueryInterface

	// HasOffset returns true if the query has an offset
	HasOffset() bool

	// Offset returns the offset
	Offset() int

	// SetOffset sets the offset
	SetOffset(offset int) UserQueryInterface

	// HasOrderBy returns true if the query has an order by
	HasOrderBy() bool

	// OrderBy returns the order by
	OrderBy() string

	// SetOrderBy sets the order by
	SetOrderBy(orderBy string) UserQueryInterface

	// HasSortDirection returns true if the query has a sort direction
	HasSortDirection() bool

	// SortDirection returns the sort direction
	SortDirection() string

	// SetSortDirection sets the sort direction
	SetSortDirection(sortDirection string) UserQueryInterface

	// HasSoftDeletedIncluded returns true if the query has soft deleted included
	HasSoftDeletedIncluded() bool

	// SoftDeletedIncluded returns the soft deleted included
	SoftDeletedIncluded() bool

	// SetSoftDeletedIncluded sets the soft deleted included
	SetSoftDeletedIncluded(softDeletedIncluded bool) UserQueryInterface

	// HasStatus returns true if the query has a status
	HasStatus() bool

	// Status returns the status
	Status() string

	// SetStatus sets the status
	SetStatus(status string) UserQueryInterface

	// HasStatusIn returns true if the query has a status in
	HasStatusIn() bool

	// StatusIn returns the status in
	StatusIn() []string

	// SetStatusIn sets the status in
	SetStatusIn(statusIn []string) UserQueryInterface

	// hasProperty returns true if the query has a property
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

	if c.HasEmailLike() && c.EmailLike() == "" {
		return errors.New("user query. email_like cannot be empty")
	}

	if c.HasFirstName() && c.FirstName() == "" {
		return errors.New("user query. first_name cannot be empty")
	}

	if c.HasFirstNameLike() && c.FirstNameLike() == "" {
		return errors.New("user query. first_name_like cannot be empty")
	}

	if c.HasLastName() && c.LastName() == "" {
		return errors.New("user query. last_name cannot be empty")
	}

	if c.HasLastNameLike() && c.LastNameLike() == "" {
		return errors.New("user query. last_name_like cannot be empty")
	}

	if c.HasID() && c.GetID() == "" {
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

func (c *userQueryImplementation) HasEmailLike() bool {
	return c.hasProperty("email_like")
}

func (c *userQueryImplementation) EmailLike() string {
	if !c.HasEmailLike() {
		return ""
	}

	return c.properties["email_like"].(string)
}

func (c *userQueryImplementation) SetEmailLike(emailLike string) UserQueryInterface {
	c.properties["email_like"] = emailLike

	return c
}

func (c *userQueryImplementation) HasFirstName() bool {
	return c.hasProperty("first_name")
}

func (c *userQueryImplementation) FirstName() string {
	if !c.HasFirstName() {
		return ""
	}

	return c.properties["first_name"].(string)
}

func (c *userQueryImplementation) SetFirstName(firstName string) UserQueryInterface {
	c.properties["first_name"] = firstName

	return c
}

func (c *userQueryImplementation) HasFirstNameLike() bool {
	return c.hasProperty("first_name_like")
}

func (c *userQueryImplementation) FirstNameLike() string {
	if !c.HasFirstNameLike() {
		return ""
	}

	return c.properties["first_name_like"].(string)
}

func (c *userQueryImplementation) SetFirstNameLike(firstNameLike string) UserQueryInterface {
	c.properties["first_name_like"] = firstNameLike

	return c
}

func (c *userQueryImplementation) HasLastName() bool {
	return c.hasProperty("last_name")
}

func (c *userQueryImplementation) LastName() string {
	if !c.HasLastName() {
		return ""
	}

	return c.properties["last_name"].(string)
}

func (c *userQueryImplementation) SetLastName(lastName string) UserQueryInterface {
	c.properties["last_name"] = lastName

	return c
}

func (c *userQueryImplementation) HasLastNameLike() bool {
	return c.hasProperty("last_name_like")
}

func (c *userQueryImplementation) LastNameLike() string {
	if !c.HasLastNameLike() {
		return ""
	}

	return c.properties["last_name_like"].(string)
}

func (c *userQueryImplementation) SetLastNameLike(lastNameLike string) UserQueryInterface {
	c.properties["last_name_like"] = lastNameLike

	return c
}

func (c *userQueryImplementation) HasID() bool {
	return c.hasProperty("id")
}

func (c *userQueryImplementation) GetID() string {
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
