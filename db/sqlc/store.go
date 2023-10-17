package db

import "database/sql"

// Store 定义执行数据库查询和事务的所有方法
type Store interface {
	Querier
}

// SQLStore 提供执行SQL查询和事务的所有功能, 实现 Store 接口
type SQLStore struct {
	db *sql.DB
	*Queries
}

// NewSQLStore 创建一个 SQLStore 实现 Store 接口，操作DB
func NewSQLStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
