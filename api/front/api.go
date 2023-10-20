package front

import "github.com/lightsaid/hotel-bookings/service"

var svc service.FrontService

// InitService 初始化 svc 提供给本包使用
func InitService(s service.FrontService) {
	svc = s
}
