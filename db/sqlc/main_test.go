package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/lightsaid/hotel-bookings/config"
	"github.com/lightsaid/hotel-bookings/pkg/env"
)

var testDB *sql.DB
var testQueries *Queries

func TestMain(m *testing.M) {
	var config config.Config
	_, err := env.LoadingEnv("config.toml", &config, "../../configs")
	if err != nil {
		log.Fatal("加载配置错误: ", err)
	}

	testDB, err = sql.Open("mysql", config.MySQL.DBSource())
	if err != nil {
		log.Fatal("连接MySQL失败: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
