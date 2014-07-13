package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func LoginPassword(initPassword string) (password string) {
	md5_hash := md5.New()
	md5_hash.Write([]byte(initPassword))
	password = hex.EncodeToString(md5_hash.Sum(nil))
	return
}
