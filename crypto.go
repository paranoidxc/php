package php

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Md5(str string) string {
	// Md5函数接受一个字符串作为输入，并返回其MD5哈希值的十六进制字符串表示
	hash := md5.New()
	hash.Write([]byte(str))
	checksum := hash.Sum(nil)
	return hex.EncodeToString(checksum)
}

func Hash256(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
