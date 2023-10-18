package service

import (
	"context"
	"database/sql"

	db "github.com/lightsaid/hotel-bookings/db/sqlc"
)

type insertParams interface {
	db.InsertHotelParams |
		db.InsertRoomParams
}

// HandleInsert 插入数据公共函数
func HandleInsert[T insertParams](
	c context.Context,
	arg T,
	fn func(c context.Context, arg T) (sql.Result, error)) (uint32, error) {
	result, err := fn(c, arg)
	if err != nil {
		return 0, err
	}

	newID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint32(newID), nil
}
