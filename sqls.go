package userstore

import (
	"github.com/dracory/sb"
)

// sqlRoleTableCreate returns a SQL string for creating the role table
// func (st *store) sqlRoleTableCreate() string {
// 	sql := sb.NewBuilder(sb.DatabaseDriverName(st.db)).
// 		Table(st.roleTableName).
// 		Column(sb.Column{
// 			Name:       COLUMN_ID,
// 			Type:       sb.COLUMN_TYPE_STRING,
// 			PrimaryKey: true,
// 			Length:     40,
// 		}).
// 		Column(sb.Column{
// 			Name:   COLUMN_HANDLE,
// 			Type:   sb.COLUMN_TYPE_STRING,
// 			Length: 50,
// 		}).
// 		Column(sb.Column{
// 			Name:   COLUMN_NAME,
// 			Type:   sb.COLUMN_TYPE_STRING,
// 			Length: 100,
// 		}).
// 		Column(sb.Column{
// 			Name:   COLUMN_CREATED_AT,
// 			Type:   sb.COLUMN_TYPE_DATETIME,
// 			Length: 0,
// 		}).
// 		Column(sb.Column{
// 			Name:   COLUMN_UPDATED_AT,
// 			Type:   sb.COLUMN_TYPE_DATETIME,
// 			Length: 0,
// 		}).
// 		Column(sb.Column{
// 			Name:   COLUMN_SOFT_DELETED_AT,
// 			Type:   sb.COLUMN_TYPE_DATETIME,
// 			Length: 0,
// 		}).
// 		CreateIfNotExists()

// 	return sql
// }

// sqlUserTableCreate returns a SQL string for creating the user table
func (st *store) sqlUserTableCreate() string {
	sql := sb.NewBuilder(sb.DatabaseDriverName(st.db)).
		Table(st.userTableName).
		Column(sb.Column{
			Name:       COLUMN_ID,
			Type:       sb.COLUMN_TYPE_STRING,
			PrimaryKey: true,
			Length:     40,
		}).
		Column(sb.Column{
			Name:   COLUMN_STATUS,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 40,
		}).
		Column(sb.Column{
			Name:   COLUMN_FIRST_NAME,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 50,
		}).
		Column(sb.Column{
			Name:   COLUMN_MIDDLE_NAMES,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 50,
		}).
		Column(sb.Column{
			Name:   COLUMN_LAST_NAME,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 50,
		}).
		Column(sb.Column{
			Name:   COLUMN_BUSINESS_NAME,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 100,
		}).
		Column(sb.Column{
			Name:   COLUMN_PHONE,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 20,
		}).
		Column(sb.Column{
			Name:   COLUMN_EMAIL,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 100,
		}).
		Column(sb.Column{
			Name:   COLUMN_PASSWORD,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 255,
		}).
		Column(sb.Column{
			Name:   COLUMN_ROLE,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 50,
		}).
		Column(sb.Column{
			Name:   COLUMN_COUNTRY,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 2,
		}).
		Column(sb.Column{
			Name:   COLUMN_TIMEZONE,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 40,
		}).
		Column(sb.Column{
			Name:   COLUMN_PROFILE_IMAGE_URL,
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 255,
		}).
		Column(sb.Column{
			Name: COLUMN_METAS,
			Type: sb.COLUMN_TYPE_TEXT,
		}).
		Column(sb.Column{
			Name: COLUMN_MEMO,
			Type: sb.COLUMN_TYPE_TEXT,
		}).
		Column(sb.Column{
			Name: COLUMN_CREATED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: COLUMN_UPDATED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: COLUMN_SOFT_DELETED_AT,
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		CreateIfNotExists()

	return sql
}
