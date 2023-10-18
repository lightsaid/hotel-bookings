package errs

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)

// HandleSQLError 统一处理mysql错误
func HandleSQLError(err error) *ApiError {
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound.AsException(err)
		}

		var mysqlErr *mysql.MySQLError
		// 重复键错误, 具体哪个字段重复，由具体业务判断
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			if strings.Contains(err.Error(), "unq_hotelCode") {
				return ErrRecordExists.AsMessage(MsgHotelCodeExists).AsException(err)
			}

			if strings.Contains(err.Error(), "unq_roomNumber_hotelId") {
				return ErrRecordExists.AsMessage(MsgHotelRoomNumberExists).AsException(err)
			}

			// NOTE: 根据业务错误处理
			return ErrRecordExists.AsException(err)
		}

		return ErrServerError.AsException(err)
	}

	return nil
}
