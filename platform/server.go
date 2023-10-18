package platform

import (
	"fmt"
	"log"
	"net/http"
	"time"

	api "github.com/lightsaid/hotel-bookings/api/back"
	"github.com/lightsaid/hotel-bookings/config"
	db "github.com/lightsaid/hotel-bookings/db/sqlc"
	"github.com/lightsaid/hotel-bookings/routers/back_routers"
	"github.com/lightsaid/hotel-bookings/routers/front_routers"
	"github.com/lightsaid/hotel-bookings/service/back"
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
		// 设置 back 服务需要的对象
		app.setupBack()
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

func (app *App) setupBack() {
	// 初始化数据库
	initMySQL()
	store := db.NewSQLStore(config.DB)
	// 初始化 back 服务
	api.InitService(back.NewService(store))
}

func (app *App) Start(config *config.Config) {
	app.serve()
}
