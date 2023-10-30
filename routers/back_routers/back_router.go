package back_routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lightsaid/hotel-bookings/configs"
	"github.com/lightsaid/hotel-bookings/pkg/fileupload"
	"github.com/lightsaid/hotel-bookings/routers/middleware"
)

// BackendRouter 后台路由
func BackendRouter() *gin.Engine {
	gin.SetMode(configs.Cfg.Server.AppMode)
	mux := gin.New()

	// 初始化上传文件工具，初始化后可以直接使用 fileupload.Local 对象
	fileupload.NewLocalUploader(configs.Cfg.Uploader.SaveDir, configs.Cfg.Uploader.MaxMB, configs.Cfg.Uploader.AllowExts...)

	// TODO: 自己定义
	// mux.Use(gin.Logger())
	mux.Use(middleware.GenRequestID())
	mux.Use(gin.Recovery())

	mux.Static("/static", "./static")
	mux.LoadHTMLGlob("views/*.html")

	if configs.Cfg.Server.AppMode == configs.DebugMode {
		mux.GET("/upload_test", func(c *gin.Context) {
			c.HTML(http.StatusOK, "upload_test.html", nil)
		})
	}

	r := mux.Group("/api")
	{
		r.GET("/healthz", healthzApi.HealthZ) // 健康检查
		r.POST("/login", authApi.Login)       // 登录
	}

	auth := r.Group("")
	auth.Use(middleware.RequireAuth(configs.TokenMaker)) // 需要登录认证
	{
		auth.POST("/fileUpload", uploadApi.UploadFile)

		auth.POST("/hotels", hotelApi.CreateHotel)       // 创建
		auth.PUT("/hotels", hotelApi.UpdateHotel)        // 更新
		auth.GET("/hotels/:id", hotelApi.GetHotel)       // 单个
		auth.GET("/hotels", hotelApi.ListHotels)         // 列表
		auth.DELETE("/hotels/:id", hotelApi.DeleteHotel) // 删除

		auth.POST("/rooms", roomApi.CreateRoom)          // 创建
		auth.PUT("/rooms", roomApi.UpdateRoom)           // 更新
		auth.GET("/rooms/:id", roomApi.GetRoom)          // 单个
		auth.GET("/rooms", roomApi.ListRooms)            // 列表
		auth.DELETE("/rooms/:id", roomApi.DeleteRoom)    // 删除
		auth.PUT("/update_type", roomApi.UpdateType)     // 更新客房类型
		auth.PUT("/update_status", roomApi.UpdateStatus) // 更新客房预订状态
		auth.GET("/query_rooms", roomApi.QueryRooms)

		auth.GET("/room_types", roomTypeApi.ListRoomTypes)              // 客房类型列表
		auth.GET("/booking_status", bookingStatusApi.ListBookingStatus) // 预定状态列表

		auth.GET("/users", userApi.ListUsers)

		auth.GET("/bookings")
	}

	return mux
}
