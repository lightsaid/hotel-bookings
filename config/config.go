package config

import (
	"fmt"
	"time"
)

// 将配置解析到Cfg成全局变量
var Cfg Config

type Config struct {
	Server Server
	MySQL  MySQL
	Redis  Redis
}

type Server struct {
	AppMode   string
	BackPort  int
	FrontPort int
}

type MySQL struct {
	Host         string
	Port         int
	User         string
	Password     string
	DbName       string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  time.Duration
}

func (db *MySQL) DBSource() string {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.DbName,
	)
	return dsn
}

type Redis struct {
	Addr     string
	DB       int
	Password string
}
