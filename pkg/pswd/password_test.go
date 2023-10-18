package pswd_test

import (
	"fmt"
	"testing"

	"github.com/lightsaid/hotel-bookings/pkg/pswd"
)

func TestGenHash(t *testing.T) {
	str, _ := pswd.GenHashPassword("123456")
	fmt.Println(str)
}
