package userstore

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/golang-module/carbon/v2"
	"github.com/gouniverse/sb"
	"github.com/samber/lo"
)

// == TYPE ====================================================================

type store struct {
	userTableName      string
	db                 *sql.DB
	dbDriverName       string
	automigrateEnabled bool
	debugEnabled       bool
}

// == INTERFACE ===============================================================

var _ StoreInterface = (*store)(nil) // verify it extends the interface

// PUBLIC METHODS ============================================================

// AutoMigrate auto migrate
func (store *store) AutoMigrate() error {
	sql := store.sqlUserTableCreate()

	if sql == "" {
		return errors.New("user table create sql is empty")
	}

	if store.db == nil {
		return errors.New("userstore: database is nil")
	}

	_, err := store.db.Exec(sql)

	if err != nil {
		return err
	}

	return nil
}

// EnableDebug - enables the debug option
func (st *store) EnableDebug(debug bool) {
	st.debugEnabled = debug
}

func (store *store) UserCount(options UserQueryInterface) (int64, error) {
	options.SetCountOnly(true)

	q := store.userSelectQuery(options)

	sqlStr, params, errSql := q.Prepared(true).
		Limit(1).
		Select(goqu.COUNT(goqu.Star()).As("count")).
		ToSQL()

	if errSql != nil {
		return -1, nil
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	db := sb.NewDatabase(store.db, store.dbDriverName)
	mapped, err := db.SelectToMapString(sqlStr, params...)
	if err != nil {
		return -1, err
	}

	if len(mapped) < 1 {
		return -1, nil
	}

	countStr := mapped[0]["count"]

	i, err := strconv.ParseInt(countStr, 10, 64)

	if err != nil {
		return -1, err

	}

	return i, nil
}

func (store *store) UserCreate(user UserInterface) error {
	user.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	user.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	data := user.Data()

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Insert(store.userTableName).
		Prepared(true).
		Rows(data).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	if store.db == nil {
		return errors.New("userstore: database is nil")
	}

	_, err := store.db.Exec(sqlStr, params...)

	if err != nil {
		return err
	}

	user.MarkAsNotDirty()

	return nil
}

func (store *store) UserDelete(user UserInterface) error {
	if user == nil {
		return errors.New("user is nil")
	}

	return store.UserDeleteByID(user.ID())
}

func (store *store) UserDeleteByID(id string) error {
	if id == "" {
		return errors.New("user id is empty")
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Delete(store.userTableName).
		Prepared(true).
		Where(goqu.C("id").Eq(id)).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	_, err := store.db.Exec(sqlStr, params...)

	return err
}

func (store *store) UserFindByEmail(email string) (user UserInterface, err error) {
	if email == "" {
		return nil, errors.New("user email is empty")
	}

	query := NewUserQuery()
	query, err = query.SetEmail(email)

	if err != nil {
		return nil, err
	}

	query, err = query.SetLimit(1)

	if err != nil {
		return nil, err
	}

	list, err := store.UserList(query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

// UserFindByEmailOrCreate - finds by email or creates a user (with active status)
func (store *store) UserFindByEmailOrCreate(email string, createStatus string) (UserInterface, error) {
	existingUser, errUser := store.UserFindByEmail(email)

	if errUser != nil {
		return nil, errUser
	}

	if existingUser != nil {
		return existingUser, nil
	}

	newUser := NewUser().
		SetEmail(email).
		SetStatus(createStatus)

	errCreate := store.UserCreate(newUser)

	if errCreate != nil {
		return nil, errCreate
	}

	return newUser, nil
}

func (store *store) UserFindByID(id string) (user UserInterface, err error) {
	if id == "" {
		return nil, errors.New("user id is empty")
	}

	query := NewUserQuery()

	query, err = query.SetID(id)

	if err != nil {
		return nil, err
	}

	query, err = query.SetLimit(1)

	if err != nil {
		return nil, err
	}

	list, err := store.UserList(query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) UserList(query UserQueryInterface) ([]UserInterface, error) {
	q := store.userSelectQuery(query)

	sqlStr, _, errSql := q.Select().ToSQL()

	if errSql != nil {
		return []UserInterface{}, nil
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	if store.db == nil {
		return []UserInterface{}, errors.New("userstore: database is nil")
	}

	db := sb.NewDatabase(store.db, store.dbDriverName)

	if db == nil {
		return []UserInterface{}, errors.New("userstore: database is nil")
	}

	modelMaps, err := db.SelectToMapString(sqlStr)

	if err != nil {
		return []UserInterface{}, err
	}

	list := []UserInterface{}

	lo.ForEach(modelMaps, func(modelMap map[string]string, index int) {
		model := NewUserFromExistingData(modelMap)
		list = append(list, model)
	})

	return list, nil
}

func (store *store) UserSoftDelete(user UserInterface) error {
	if user == nil {
		return errors.New("user is nil")
	}

	user.SetDeletedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	return store.UserUpdate(user)
}

func (store *store) UserSoftDeleteByID(id string) error {
	user, err := store.UserFindByID(id)

	if err != nil {
		return err
	}

	return store.UserSoftDelete(user)
}

func (store *store) UserUpdate(user UserInterface) error {
	if user == nil {
		return errors.New("user is nil")
	}

	user.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	dataChanged := user.DataChanged()

	delete(dataChanged, COLUMN_ID) // ID is not updateable

	if len(dataChanged) < 1 {
		return nil
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Update(store.userTableName).
		Prepared(true).
		Set(dataChanged).
		Where(goqu.C(COLUMN_ID).Eq(user.ID())).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	if store.debugEnabled {
		log.Println(sqlStr)
	}

	if store.db == nil {
		return errors.New("userstore: database is nil")
	}

	_, err := store.db.Exec(sqlStr, params...)

	user.MarkAsNotDirty()

	return err
}

func (store *store) userSelectQuery(options UserQueryInterface) *goqu.SelectDataset {
	q := goqu.Dialect(store.dbDriverName).From(store.userTableName)

	if options.ID() != "" {
		q = q.Where(goqu.C(COLUMN_ID).Eq(options.ID()))
	}

	if len(options.IDIn()) > 0 {
		q = q.Where(goqu.C(COLUMN_ID).In(options.IDIn()))
	}

	if options.Status() != "" {
		q = q.Where(goqu.C(COLUMN_STATUS).Eq(options.Status()))
	}

	if len(options.StatusIn()) > 0 {
		q = q.Where(goqu.C(COLUMN_STATUS).In(options.StatusIn()))
	}

	if options.Email() != "" {
		q = q.Where(goqu.C(COLUMN_EMAIL).Eq(options.Email()))
	}

	if options.CreatedAtGte() != "" && options.CreatedAtLte() != "" {
		q = q.Where(
			goqu.C(COLUMN_CREATED_AT).Gte(options.CreatedAtGte()),
			goqu.C(COLUMN_CREATED_AT).Lte(options.CreatedAtLte()),
		)
	} else if options.CreatedAtGte() != "" {
		q = q.Where(goqu.C(COLUMN_CREATED_AT).Gte(options.CreatedAtGte()))
	} else if options.CreatedAtLte() != "" {
		q = q.Where(goqu.C(COLUMN_CREATED_AT).Lte(options.CreatedAtLte()))
	}

	if !options.CountOnly() {
		if options.Limit() > 0 {
			q = q.Limit(uint(options.Limit()))
		}

		if options.Offset() > 0 {
			q = q.Offset(uint(options.Offset()))
		}
	}

	sortOrder := sb.DESC
	if options.SortOrder() != "" {
		sortOrder = options.SortOrder()
	}

	if options.OrderBy() != "" {
		if strings.EqualFold(sortOrder, sb.ASC) {
			q = q.Order(goqu.I(options.OrderBy()).Asc())
		} else {
			q = q.Order(goqu.I(options.OrderBy()).Desc())
		}
	}

	if options.WithSoftDeleted() {
		return q // soft deleted users requested specifically
	}

	softDeleted := goqu.C(COLUMN_DELETED_AT).
		Gt(carbon.Now(carbon.UTC).ToDateTimeString())

	return q.Where(softDeleted)
}
