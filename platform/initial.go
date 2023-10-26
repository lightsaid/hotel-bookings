package platform

import (
	"database/sql"
	"io"
	"log"
	"os"

	"github.com/lightsaid/hotel-bookings/config"
	"github.com/lightsaid/hotel-bookings/pkg/logger"
	"github.com/lightsaid/hotel-bookings/pkg/token"
)

func initMySQL() {
	var err error
	config.DB, err = sql.Open("mysql", config.Cfg.MySQL.DBSource())
	if err != nil {
		log.Fatal("initMySQL() ", err)
	}
}

func initTokenMaker() {
	var err error
	config.TokenMaker, err = token.NewJWTMaker(config.Cfg.Token.TokenScretkey, config.Cfg.Server.Host)
	if err != nil {
		log.Fatal("initTokenMaker() ", err)
	}
}

func initLogger() {
	var output io.Writer = os.Stdout
	if config.Cfg.Server.AppMode == config.ReleaseMode {
		output = logger.DefaultOutput(config.Cfg.Logger.LogFile)
	}
	logger.NewLogger(config.Cfg.Logger.Level, output)
}
