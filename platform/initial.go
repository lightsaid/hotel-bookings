package platform

import (
	"database/sql"
	"io"
	"log"
	"os"

	"github.com/lightsaid/hotel-bookings/configs"
	"github.com/lightsaid/hotel-bookings/pkg/logger"
	"github.com/lightsaid/hotel-bookings/pkg/token"
)

func initMySQL() {
	var err error
	configs.DB, err = sql.Open("mysql", configs.Cfg.MySQL.DBSource())
	if err != nil {
		log.Fatal("initMySQL() ", err)
	}
	err = configs.DB.Ping()
	if err != nil {
		log.Fatal("Ping MySQL ", err)
	}
}

func initTokenMaker() {
	var err error
	configs.TokenMaker, err = token.NewJWTMaker(configs.Cfg.Token.TokenScretkey, configs.Cfg.Server.Host)
	if err != nil {
		log.Fatal("initTokenMaker() ", err)
	}
}

func initLogger() {
	var output io.Writer = os.Stdout
	if configs.Cfg.Server.AppMode == configs.ReleaseMode {
		output = logger.DefaultOutput(configs.Cfg.Logger.LogFile)
	}
	logger.NewLogger(configs.Cfg.Logger.Level, output)
}
