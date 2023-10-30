package configs

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

const (
	// debug 开发环境
	DebugMode = "debug"

	// release 生产环境
	ReleaseMode = "release"

	// test        测试
	TestMode = "test"
)

type Config struct {
	Server   Server
	MySQL    MySQL
	Redis    Redis
	Token    Token
	Uploader Uploader
	Logger   Logger
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

type Uploader struct {
	SaveDir   string
	AllowExts []string
	MaxMB     int // 最大限制，单位: MB
}

type Logger struct {
	Level   string
	LogFile string
}
