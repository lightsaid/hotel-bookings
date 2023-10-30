package platform

import (
	"fmt"
	"log"
	"net/http"
	"time"

	backApi "github.com/lightsaid/hotel-bookings/api/back"
	frontApi "github.com/lightsaid/hotel-bookings/api/front"
	"github.com/lightsaid/hotel-bookings/api/validate"
	"github.com/lightsaid/hotel-bookings/configs"
	db "github.com/lightsaid/hotel-bookings/db/sqlc"
	"github.com/lightsaid/hotel-bookings/routers/back_routers"
	"github.com/lightsaid/hotel-bookings/routers/front_routers"
	"github.com/lightsaid/hotel-bookings/service/back"
	"github.com/lightsaid/hotel-bookings/service/front"
)

// 做 前端、后台 API服务启动时的公共事情

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
	initLogger()

	// 初始化数据库
	initMySQL()
	defer configs.DB.Close()

	initTokenMaker()

	configs.Trans = validate.NewValidation("zh")

	var mux http.Handler
	if app.appType == Backend {
		// 设置 back 服务需要的对象
		app.setupBack()
		mux = back_routers.BackendRouter()
	} else {
		app.setupFront()
		mux = front_routers.FrontendRouter()
	}

	addr := fmt.Sprintf("0.0.0.0:%d", configs.Cfg.Server.Port)
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
	store := db.NewSQLStore(configs.DB)
	// 初始化 back 服务
	backApi.InitService(back.NewService(store))
}

func (app *App) setupFront() {
	store := db.NewSQLStore(configs.DB)
	// 初始化 front 服务
	frontApi.InitService(front.NewService(store))
}

func (app *App) Start(config *configs.Config) {
	app.serve()
}
