package back

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsaid/hotel-bookings/api/errs"
	"github.com/lightsaid/hotel-bookings/api/request"
	reps "github.com/lightsaid/hotel-bookings/api/response"
)

type HotelApi struct{}

func (*HotelApi) CreateHotel(c *gin.Context) {
	var req request.HotelRequest
	if ok := request.ShouldBind(c, &req); !ok {
		return
	}
	newID, err := svc.CreateHotel(c, req)
	if err != nil {
		reps.FAIL(c, errs.HandleSQLError(err))
		return
	}
	reps.OK(c, newID)
}

func (*HotelApi) UpdateHotel(c *gin.Context) {
	var req request.HotelRequest
	if ok := request.ShouldBind(c, &req); !ok {
		return
	}

	err := svc.UpdateHotel(c, req)
	if err != nil {
		reps.FAIL(c, errs.HandleSQLError(err))
		return
	}
	reps.OK(c, req.ID)
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

	// meta := reps.CalculateMetadata(total, req.PageNum, req.PageSize)
	// data := reps.ToListResponse(list, meta, errs.ErrOK)
	// reps.JSON(c, errs.ErrOK, data)

	reps.PAGE(c, list, total, req.PageNum, req.PageSize)
}

func (*HotelApi) GetHotel(c *gin.Context) {
	var req request.URIRequest
	if ok := request.ShouldBind(c, &req, request.BindUri); !ok {
		return
	}

	hotel, err := svc.GetHotelByID(c, req.ID)
	if err := errs.HandleSQLError(err); err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.OK(c, hotel)
}

func (*HotelApi) DeleteHotel(c *gin.Context) {
	var req request.URIRequest
	if ok := request.ShouldBind(c, &req, request.BindUri); !ok {
		return
	}

	err := svc.DeleteHotelByID(c, req.ID)
	if err := errs.HandleSQLError(err); err != nil {
		reps.FAIL(c, err)
		return
	}

	reps.OK(c, req.ID)
}
