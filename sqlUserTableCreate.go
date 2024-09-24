package userstore

import (
	"github.com/gouniverse/sb"
)

// SQLCreateTable returns a SQL string for creating the country table
func (st *Store) sqlUserTableCreate() string {
	sql := sb.NewBuilder(sb.DatabaseDriverName(st.db)).
		Table(st.userTableName).
		Column(sb.Column{
			Name:       "id",
			Type:       sb.COLUMN_TYPE_STRING,
			PrimaryKey: true,
			Length:     40,
		}).
		Column(sb.Column{
			Name:   "status",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 40,
		}).
		Column(sb.Column{
			Name:   "first_name",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 50,
		}).
		Column(sb.Column{
			Name:   "middle_names",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 50,
		}).
		Column(sb.Column{
			Name:   "last_name",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 50,
		}).
		Column(sb.Column{
			Name:   "business_name",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 100,
		}).
		Column(sb.Column{
			Name:   "phone",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 20,
		}).
		Column(sb.Column{
			Name:   "email",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 100,
		}).
		Column(sb.Column{
			Name:   "password",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 255,
		}).
		Column(sb.Column{
			Name:   "role",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 50,
		}).
		Column(sb.Column{
			Name:   "country",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 2,
		}).
		Column(sb.Column{
			Name:   "timezone",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 40,
		}).
		Column(sb.Column{
			Name:   "profile_image_url",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 255,
		}).
		Column(sb.Column{
			Name: "metas",
			Type: sb.COLUMN_TYPE_TEXT,
		}).
		Column(sb.Column{
			Name: "memo",
			Type: sb.COLUMN_TYPE_TEXT,
		}).
		Column(sb.Column{
			Name: "created_at",
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: "updated_at",
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: "deleted_at",
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		CreateIfNotExists()

	return sql
}
