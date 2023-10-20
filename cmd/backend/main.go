package main

import (
	"log"

	_ "github.com/golang-migrate/migrate/v4"
	"github.com/lightsaid/hotel-bookings/config"
	"github.com/lightsaid/hotel-bookings/pkg/env"
	"github.com/lightsaid/hotel-bookings/platform"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_, err := env.LoadingEnv("back.toml", &config.Cfg, "./configs")
	if err != nil {
		log.Fatal("加载 back.toml ", err)
	}

	platform.NewApp(platform.Backend).Start(&config.Cfg)
}
