package common

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "bluebell"

// EncryptPassword 利用MD5值机进行密码加密
func EncryptPassword(password string) (string, error) {
	b := md5.New()
	if _, err := b.Write([]byte(secret)); err != nil {
		return "", err
	}

	return hex.EncodeToString(b.Sum([]byte(password))), nil
}
