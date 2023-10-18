package request

type CreateHotelRequest struct {
	Title   string `json:"title"`
	Code    string `json:"code"`
	Address string `json:"address"`
}

type UpdateHotelRequest struct {
	Title   string `json:"title"`
	Code    string `json:"code"`
	Address string `json:"address"`
	ID      uint32 `json:"id"`
}
