package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// GetSha256
// Maintainers:贺胜 Times:2023-04-13
// Part 1:获取字符串的SHA256值
// Part 2:获取字符串的SHA256值，返回32位字符串
func GetSha256(str string) string {
	srcByte := []byte(str)
	sha256New := sha256.New()
	sha256Bytes := sha256New.Sum(srcByte)
	sha256String := hex.EncodeToString(sha256Bytes)
	sha256String = sha256String[0:32]
	return sha256String
}
