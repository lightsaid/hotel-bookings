package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	ut "github.com/go-playground/universal-translator"
	"github.com/lightsaid/hotel-bookings/pkg/token"
)

// 成全局变量
var (
	Cfg        Config
	DB         *sql.DB
	Trans      ut.Translator
	TokenMaker token.TokenMaker
)

type Config struct {
	Server Server
	MySQL  MySQL
	Redis  Redis
	Token  Token
}

// JSON 以 JSON 格式打印配置
func (c *Config) JSON() {
	dataByte, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataByte))
}

type Server struct {
	AppMode string
	Host    string
	Port    int
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

type Token struct {
	TokenScretkey        string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}
