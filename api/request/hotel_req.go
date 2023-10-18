package request

type HotelRequest struct {
	Title   string `json:"title" binding:"required"`
	Code    string `json:"code" binding:"required,min=4,max=8"`
	Address string `json:"address" binding:"required,max=255"`
	// 如果id存在，则校验，用指针解决 0 不校验的问题。
	ID *uint32 `json:"id,omitempty" binding:"omitempty,min=1"`
}
