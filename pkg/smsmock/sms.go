package smsmock

import (
	"log/slog"
	"strconv"
	"sync"
	"time"

	"github.com/lightsaid/hotel-bookings/pkg/random"
)

// 模拟短信验证码，不会发送，仅仅以接口方式返回

type SMS struct {
	store    map[string]SMSCode // 存储短信验证码
	duration time.Duration      // 有效期
	mu       sync.RWMutex
}

func NewSMS(duration time.Duration) *SMS {
	sms := &SMS{
		store:    make(map[string]SMSCode),
		duration: duration,
		mu:       sync.RWMutex{},
	}

	// 定时检查清理
	go func() {
		tk := time.NewTimer(10 * time.Second)
		defer tk.Stop()
		for range tk.C {
			for key := range sms.store {
				val := sms.store[key]
				if val.Status == StatusExpired {
					delete(sms.store, key)
				}
			}
		}
	}()

	return sms
}

func (sms *SMS) Length() int {
	return len(sms.store)
}

// GenSMSCode 生成随机验证码, 如果是旧的验证码exists=true
func (sms *SMS) GenSMSCode(phoneNumber string) (smscode SMSCode, exists bool) {
	var err error
	smscode, err = sms.getSMSCode(phoneNumber)
	if err == nil {
		exists = true
		return
	}

	code := random.RandomInt(1000, 9999)
	smscode = SMSCode{
		PhoneNumber: phoneNumber,
		Code:        strconv.Itoa(code),
		ExpiresAt:   time.Now().Add(sms.duration),
	}

	sms.mu.Lock()
	sms.store[phoneNumber] = smscode
	sms.mu.Unlock()

	go sms.watch(phoneNumber)

	return
}

// GetSMSCode 获取验证码并修改状态为已使用
func (sms *SMS) GetSMSCode(phoneNumber string) (smscode SMSCode, err error) {
	smscode, err = sms.getSMSCode(phoneNumber)
	if err != nil {
		return
	}

	newCode := SMSCode{
		PhoneNumber: smscode.PhoneNumber,
		Code:        smscode.Code,
		ExpiresAt:   smscode.ExpiresAt,
		Status:      StatusInValid,
	}

	sms.mu.Lock()
	sms.store[phoneNumber] = newCode
	sms.mu.Unlock()
	return
}

// GetSMSCode 获取一个 SMSCode
func (sms *SMS) getSMSCode(phoneNumber string) (smscode SMSCode, err error) {
	sms.mu.Lock()
	defer sms.mu.Unlock()

	var ok bool
	smscode, ok = sms.store[phoneNumber]
	if !ok {
		return smscode, ErrInNotExists
	}

	if smscode.Status == StatusExpired {
		return smscode, ErrInExpired
	}

	if smscode.Status == StatusInValid {
		return smscode, ErrInValid
	}

	if smscode.Status != 0 {
		return smscode, ErrInExpired
	}

	return
}

// watch 定时检查smscode，更新状态
func (sms *SMS) watch(phoneNumber string) {
	// 创建一个定时器
	ticker := time.NewTicker(sms.duration)
	for range ticker.C {
		defer ticker.Stop()

		smscode, err := sms.getSMSCode(phoneNumber)
		if err != nil && err == ErrInNotExists {
			return
		}

		slog.Info(
			"update smscode status",
			slog.String("phone_number", phoneNumber),
			slog.Int("org_status", int(smscode.Status)),
			slog.Int("new_status", int(StatusExpired)),
		)

		sms.mu.Lock()
		smscode.Status = StatusExpired
		sms.store[phoneNumber] = smscode
		sms.mu.Unlock()
	}
}
