package front_routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsaid/hotel-bookings/configs"
	"github.com/lightsaid/hotel-bookings/pkg/fileupload"
	"github.com/lightsaid/hotel-bookings/routers/middleware"
)

func FrontendRouter() *gin.Engine {
	gin.SetMode(configs.Cfg.Server.AppMode)
	mux := gin.New()

	// 初始化上传文件工具，初始化后可以直接使用 fileupload.Local 对象
	fileupload.NewLocalUploader(configs.Cfg.Uploader.SaveDir, configs.Cfg.Uploader.MaxMB, configs.Cfg.Uploader.AllowExts...)

	// TODO: 自己定义
	mux.Use(gin.Logger())
	mux.Use(gin.Recovery())

	mux.Static("/static", "./static")

	v1 := mux.Group("/v1")
	{
		v1.GET("/healthz", healthzApi.HealthZ) // 健康检查
		v1.GET("/sms/send", smsApi.Send)
		v1.POST("/register", authApi.Register)     // 注册
		v1.POST("/login", authApi.Login)           // 登录
		v1.GET("/hotels", homeApi.GetHotels)       // 获取酒店
		v1.POST("/queryRooms", homeApi.QueryRooms) // 查询客房
	}

	auth := v1.Group("")
	auth.Use(middleware.RequireAuth(configs.TokenMaker)) // 需要登录认证
	{
		v1.GET("/profile", authApi.GetProfile)           // 获取个人信息
		v1.POST("/bookings", authApi.BookingRoom)        // 预订客房
		v1.GET("/bookings", authApi.GetBookings)         // 获取个人订房记录
		v1.GET("/bookings/:id", authApi.GetBookingsByID) // 获取个人单个订房记录
		v1.DELETE("/bookings", authApi.DeleteBooking)    // 删除记录
		v1.POST("/payment", authApi.UserPayment)         // 用户支付
	}
	return mux
}
