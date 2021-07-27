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
	tbUsersFieldNames          = builderx.RawFieldNames(&TbUsers{})
	tbUsersRows                = strings.Join(tbUsersFieldNames, ",")
	tbUsersRowsExpectAutoSet   = strings.Join(stringx.Remove(tbUsersFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	tbUsersRowsWithPlaceHolder = strings.Join(stringx.Remove(tbUsersFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheTbUsersIdPrefix       = "cache:tbUsers:id:"
	cacheTbUsersUsernamePrefix = "cache:tbUsers:username:"
)

type (
	TbUsersModel interface {
		Insert(data TbUsers) (sql.Result, error)
		FindOne(id int64) (*TbUsers, error)
		FindOneByUsername(username string) (*TbUsers, error)
		Update(data TbUsers) error
		Delete(id int64) error
	}

	defaultTbUsersModel struct {
		sqlc.CachedConn
		table string
	}

	TbUsers struct {
		Id        int64     `db:"id"`
		Username  string    `db:"username"`
		Password  string    `db:"password"`
		CreatedAt time.Time `db:"createdAt"`
		UpdatedAt time.Time `db:"updatedAt"`
		DeletedAt time.Time `db:"deletedAt"`
	}
)

func NewTbUsersModel(conn sqlx.SqlConn, c cache.CacheConf) TbUsersModel {
	return &defaultTbUsersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`tb_users`",
	}
}

func (m *defaultTbUsersModel) Insert(data TbUsers) (sql.Result, error) {
	tbUsersUsernameKey := fmt.Sprintf("%s%v", cacheTbUsersUsernamePrefix, data.Username)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, tbUsersRowsExpectAutoSet)
		return conn.Exec(query, data.Username, data.Password, data.CreatedAt, data.UpdatedAt, data.DeletedAt)
	}, tbUsersUsernameKey)
	return ret, err
}

func (m *defaultTbUsersModel) FindOne(id int64) (*TbUsers, error) {
	tbUsersIdKey := fmt.Sprintf("%s%v", cacheTbUsersIdPrefix, id)
	var resp TbUsers
	err := m.QueryRow(&resp, tbUsersIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tbUsersRows, m.table)
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

func (m *defaultTbUsersModel) FindOneByUsername(username string) (*TbUsers, error) {
	tbUsersUsernameKey := fmt.Sprintf("%s%v", cacheTbUsersUsernamePrefix, username)
	var resp TbUsers
	err := m.QueryRowIndex(&resp, tbUsersUsernameKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", tbUsersRows, m.table)
		if err := conn.QueryRow(&resp, query, username); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTbUsersModel) Update(data TbUsers) error {
	tbUsersIdKey := fmt.Sprintf("%s%v", cacheTbUsersIdPrefix, data.Id)
	tbUsersUsernameKey := fmt.Sprintf("%s%v", cacheTbUsersUsernamePrefix, data.Username)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tbUsersRowsWithPlaceHolder)
		return conn.Exec(query, data.Username, data.Password, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Id)
	}, tbUsersIdKey, tbUsersUsernameKey)
	return err
}

func (m *defaultTbUsersModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	tbUsersIdKey := fmt.Sprintf("%s%v", cacheTbUsersIdPrefix, id)
	tbUsersUsernameKey := fmt.Sprintf("%s%v", cacheTbUsersUsernamePrefix, data.Username)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, tbUsersIdKey, tbUsersUsernameKey)
	return err
}

func (m *defaultTbUsersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTbUsersIdPrefix, primary)
}

func (m *defaultTbUsersModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tbUsersRows, m.table)
	return conn.QueryRow(v, query, primary)
}
