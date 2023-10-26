package smsmock_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/lightsaid/hotel-bookings/pkg/smsmock"
	"github.com/stretchr/testify/require"
)

var sms *smsmock.SMS

func initSMS() {
	sms = smsmock.NewSMS(time.Duration(3) * time.Second)
}

func TestGenSMSCode(t *testing.T) {
	initSMS()

	var phoneNumber = "132"
	smscode, exists := sms.GenSMSCode(phoneNumber)
	require.True(t, !exists)
	require.NotEmpty(t, smscode)
	require.WithinDuration(t, time.Now().Add(3*time.Second), smscode.ExpiresAt, time.Second)
	require.True(t, smscode.Status == smsmock.StatusNormal)

	smscode2, err := sms.GetSMSCode(phoneNumber)
	require.NoError(t, err)
	require.Equal(t, smscode, smscode2)

	smscode3, err := sms.GetSMSCode(phoneNumber)
	require.ErrorIs(t, err, smsmock.ErrInValid)
	require.NotEqual(t, smscode3.Status, smscode2.Status)

	// ======= test 是否过期
	phoneNumber = "138"
	smscode, _ = sms.GenSMSCode(phoneNumber)

	time.Sleep(5 * time.Second)

	smscode4, err := sms.GetSMSCode(phoneNumber)
	require.ErrorIs(t, err, smsmock.ErrInExpired)
	require.Equal(t, int(smsmock.StatusExpired), int(smscode4.Status))

	fmt.Println("len1: ", sms.Length())
	time.Sleep(20 * time.Second)
	fmt.Println("len2: ", sms.Length())

	_, err = sms.GetSMSCode(phoneNumber)
	require.ErrorIs(t, smsmock.ErrInNotExists, err)

}
