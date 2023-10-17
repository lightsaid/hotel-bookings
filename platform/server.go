package platform

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/lightsaid/hotel-bookings/config"
	"github.com/lightsaid/hotel-bookings/routers/back_routers"
	"github.com/lightsaid/hotel-bookings/routers/front_routers"
)

// 做前、端端 API 服务启动时的公共事情

type AppType string

const (
	Frontend = "F"
	Backend  = "B"
)

type App struct {
	appType AppType
}

func NewApp(appType AppType) *App {
	return &App{
		appType: appType,
	}
}

func (app *App) serve() {
	var mux http.Handler
	if app.appType == Backend {
		mux = back_routers.BackendRouter()
	} else {
		mux = front_routers.FrontendRouter()
	}

	addr := fmt.Sprintf("0.0.0.0:%d", config.Cfg.Server.Port)
	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  time.Minute,
	}

	fmt.Println("HTTP server on ", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Println("HTTP server ", err)
	}
}

func (app *App) Start(config *config.Config) {
	app.serve()
}
