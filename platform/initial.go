package platform

import (
	"database/sql"
	"log"

	"github.com/lightsaid/hotel-bookings/config"
)

func initMySQL() {
	var err error
	config.DB, err = sql.Open("mysql", config.Cfg.MySQL.DBSource())
	if err != nil {
		log.Fatal("open mysql ", err)
	}
}
