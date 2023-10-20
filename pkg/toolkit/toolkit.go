package toolkit

import (
	"crypto/md5"
	"encoding/hex"
)

type toolkit struct{}

var TKs = new(toolkit)

// MD5 将 val 加密成hash, 密钥是 secret （可选）
func (*toolkit) MD5(val string, secret ...byte) (string, error) {
	h := md5.New()
	_, err := h.Write([]byte(val))
	if err != nil {
		return "", err
	}
	// 编码成十六进制 string
	result := hex.EncodeToString(h.Sum(secret))
	return result, nil
}
