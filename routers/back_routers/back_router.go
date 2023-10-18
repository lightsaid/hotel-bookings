package back_routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsaid/hotel-bookings/config"
)

// BackendRouter 后台路由
func BackendRouter() *gin.Engine {
	gin.SetMode(config.Cfg.Server.AppMode)
	mux := gin.New()

	// TODO: 自己定义
	mux.Use(gin.Logger())
	mux.Use(gin.Recovery())

	r := mux.Group("/api")
	{
		r.GET("/healthz", healthzApi.HealthZ) // 健康检查
		r.POST("/login", authApi.Login)       // 登录
	}

	auth := r.Group("")
	// auth.Use() // TODO: 中间件 需要登录认证
	{
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

		auth.GET("/room_types", roomTypeApi.ListRoomTypes)              // 客房类型列表
		auth.GET("/booking_status", bookingStatusApi.ListBookingStatus) // 预定状态列表
	}

	return mux
}
