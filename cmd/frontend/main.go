package main

import (
	"log"

	_ "github.com/golang-migrate/migrate/v4"
	"github.com/lightsaid/hotel-bookings/configs"
	"github.com/lightsaid/hotel-bookings/pkg/env"
	"github.com/lightsaid/hotel-bookings/platform"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_, err := env.LoadingEnv("front.toml", &configs.Cfg, "./configs")
	if err != nil {
		log.Fatal("加载 front.toml ", err)
	}

	platform.NewApp(platform.Frontend).Start(&configs.Cfg)
}
