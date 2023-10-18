package back

import "github.com/lightsaid/hotel-bookings/service"

var svc service.BackService

// InitService 初始化 svc 提供给本包使用
func InitService(s service.BackService) {
	svc = s
}
