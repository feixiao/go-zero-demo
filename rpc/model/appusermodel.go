package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	appUserFieldNames          = builderx.RawFieldNames(&AppUser{})
	appUserRows                = strings.Join(appUserFieldNames, ",")
	appUserRowsExpectAutoSet   = strings.Join(stringx.Remove(appUserFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	appUserRowsWithPlaceHolder = strings.Join(stringx.Remove(appUserFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAppUserIdPrefix = "cache:appUser:id:"
)

type (
	AppUserModel interface {
		Insert(data AppUser) (sql.Result, error)
		FindOne(id int64) (*AppUser, error)
		Update(data AppUser) error
		Delete(id int64) error
	}

	defaultAppUserModel struct {
		sqlc.CachedConn
		table string
	}

	AppUser struct {
		Id         int64     `db:"id"`
		Username   string    `db:"username"`
		Password   string    `db:"password"`
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"`
		DeletedAt  time.Time `db:"deleted_at"`
	}
)

func NewAppUserModel(conn sqlx.SqlConn, c cache.CacheConf) AppUserModel {
	return &defaultAppUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`app_user`",
	}
}

func (m *defaultAppUserModel) Insert(data AppUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, appUserRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Username, data.Password, data.DeletedAt)

	return ret, err
}

func (m *defaultAppUserModel) FindOne(id int64) (*AppUser, error) {
	appUserIdKey := fmt.Sprintf("%s%v", cacheAppUserIdPrefix, id)
	var resp AppUser
	err := m.QueryRow(&resp, appUserIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appUserRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAppUserModel) Update(data AppUser) error {
	appUserIdKey := fmt.Sprintf("%s%v", cacheAppUserIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, appUserRowsWithPlaceHolder)
		return conn.Exec(query, data.Username, data.Password, data.DeletedAt, data.Id)
	}, appUserIdKey)
	return err
}

func (m *defaultAppUserModel) Delete(id int64) error {

	appUserIdKey := fmt.Sprintf("%s%v", cacheAppUserIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, appUserIdKey)
	return err
}

func (m *defaultAppUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAppUserIdPrefix, primary)
}

func (m *defaultAppUserModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appUserRows, m.table)
	return conn.QueryRow(v, query, primary)
}
