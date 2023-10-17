package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/lightsaid/hotel-bookings/pkg/random"
	"github.com/stretchr/testify/require"

	_ "github.com/go-sql-driver/mysql"
)

func createRandomHotel(t *testing.T) *Hotel {
	arg := InsertHotelParams{
		Title:   random.RandomString(10),
		Code:    random.RandomString(6),
		Address: random.RandomString(16),
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 创建
	result, err := testQueries.InsertHotel(ctx, arg)
	require.NoError(t, err)

	newID, err := result.LastInsertId()
	require.NoError(t, err)
	require.True(t, newID > 0)

	// 获取
	hotel, err := testQueries.GetHotelByID(ctx, uint32(newID))
	require.NoError(t, err)

	// 对比前后参数是否匹配
	require.Equal(t, arg.Title, hotel.Title)
	require.Equal(t, arg.Code, hotel.Code)
	require.Equal(t, arg.Address, hotel.Address)
	require.Equal(t, uint32(newID), hotel.ID)
	require.Equal(t, hotel.IsDeleted, false)
	require.WithinDuration(t, hotel.CreatedAt, time.Now(), time.Second)
	require.WithinDuration(t, hotel.UpdatedAt, time.Now(), time.Second)

	return hotel
}

func TestInsertHotel(t *testing.T) {
	createRandomHotel(t)
}

func TestUpdateHotel(t *testing.T) {
	hotel := createRandomHotel(t)
	arg := UpdateHotelParams{
		Title:   sql.NullString{String: random.RandomString(10), Valid: true},
		Address: sql.NullString{String: random.RandomString(16), Valid: true},
		ID:      hotel.ID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := testQueries.UpdateHotel(ctx, arg)
	require.NoError(t, err)

	hotel2, err := testQueries.GetHotelByID(ctx, hotel.ID)
	require.NoError(t, err)

	require.Equal(t, arg.ID, hotel2.ID)
	require.Equal(t, arg.Title.String, hotel2.Title)
	require.Equal(t, arg.Address.String, hotel2.Address)
	require.WithinDuration(t, hotel2.UpdatedAt, time.Now(), time.Second)
}

func TestListHotels(t *testing.T) {
	// 创建10条数据，然后获取
	var n = 10
	for i := 0; i < n; i++ {
		createRandomHotel(t)
	}

	arg := GetHotelsParams{
		Limit:  5,
		Offset: 0,
	}
	ctx := context.Background()
	list, err := testQueries.GetHotels(ctx, arg)
	require.NoError(t, err)
	require.True(t, len(list) == 5)

	arg.Offset = 6
	list, err = testQueries.GetHotels(ctx, arg)
	require.NoError(t, err)
	require.True(t, len(list) == 5)
}

func TestGetHotelByTitle(t *testing.T) {
	hotel := createRandomHotel(t)

	runeTitle := []rune(hotel.Title)
	queryTitle := string(runeTitle[2:4])

	t.Log("queryTitle: ", queryTitle)

	arg := GetHotelsByTitleParams{
		Title:  fmt.Sprintf("%s%s%s", "%", queryTitle, "%"),
		Limit:  10,
		Offset: 0,
	}

	t.Log("arg.Title: ", arg.Title)

	list, err := testQueries.GetHotelsByTitle(context.TODO(), arg)
	require.NoError(t, err)

	t.Log("len: ", len(list))

	require.True(t, len(list) >= 1)
}

func TestDeleteHotel(t *testing.T) {
	hotel := createRandomHotel(t)
	err := testQueries.DeleteHotelByID(context.TODO(), hotel.ID)
	require.NoError(t, err)

	_, err = testQueries.GetHotelByID(context.TODO(), hotel.ID)
	require.Error(t, err)
	require.ErrorIs(t, err, sql.ErrNoRows)
}
