package back

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lightsaid/hotel-bookings/api/request"
	reps "github.com/lightsaid/hotel-bookings/api/response"
)

type HotelApi struct{}

func (*HotelApi) CreateHotel(c *gin.Context) {

}

func (*HotelApi) UpdateHotel(c *gin.Context) {

}

func (*HotelApi) ListHotels(c *gin.Context) {
	req := request.ListRequest{
		PageNum:  1,
		PageSize: 10,
	}
	list, total, err := svc.GetListHotels(c, req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
		return
	}

	data := reps.ToListResponse(list, total, req.PageNum, req.PageSize)

	c.JSON(http.StatusOK, gin.H{"result": data})
}

func (*HotelApi) GetHotel(c *gin.Context) {

}

func (*HotelApi) DeleteHotel(c *gin.Context) {

}
