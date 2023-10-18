package back

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsaid/hotel-bookings/api/errs"
	"github.com/lightsaid/hotel-bookings/api/request"
	reps "github.com/lightsaid/hotel-bookings/api/response"
)

type HotelApi struct{}

func (*HotelApi) CreateHotel(c *gin.Context) {

}

func (*HotelApi) UpdateHotel(c *gin.Context) {

}

func (*HotelApi) ListHotels(c *gin.Context) {
	var req request.ListRequest
	if ok := request.ShouldBind(c, &req, request.BindQuery); !ok {
		return
	}
	list, total, err := svc.GetListHotels(c, req)
	if err := errs.HandleSQLError(err); err != nil {
		reps.FAIL(c, err)
		return
	}

	meta := reps.CalculateMetadata(total, req.PageNum, req.PageSize)
	data := reps.ToListResponse(list, meta, errs.ErrOK)

	reps.JSON(c, errs.ErrOK, data)
}

func (*HotelApi) GetHotel(c *gin.Context) {

}

func (*HotelApi) DeleteHotel(c *gin.Context) {

}
