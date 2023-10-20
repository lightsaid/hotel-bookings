package back

import (
	"github.com/gin-gonic/gin"
	reps "github.com/lightsaid/hotel-bookings/api/response"
	"github.com/lightsaid/hotel-bookings/pkg/errs"
	"github.com/lightsaid/hotel-bookings/pkg/fileupload"
)

type UplaodApi struct{}

func (*UplaodApi) UploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		reps.FAIL(c, errs.ErrServerError.AsException(err))
		return
	}
	result, err := fileupload.Local.SaveFile(fileHeader)
	if err != nil {
		if fileupload.IsUploaderError(err) {
			reps.FAIL(c, errs.ErrServerError.AsMessage(err.Error()))
			return
		}
		reps.FAIL(c, errs.ErrServerError.AsException(err))
		return
	}
	reps.OK(c, result)
}
