package entity

import (
	"errors"
	"fmt"
)

// ErrorInfo
// Maintainers:贺胜 Times:2021-04-15
// Part 1:错误信息结构体
type ErrorInfo struct {
	UserNotFound      error
	UserPasswordError error
}

// ErrorToString
// Maintainers:贺胜 Times:2021-04-15
// Part 1:错误信息转换为字符串
func (t *ErrorInfo) ErrorToString(err error) string {
	errMessage := fmt.Sprintf("%s", err)
	return errMessage
}

var ErrorUser ErrorInfo

// init
// Maintainers:贺胜 Times:2021-04-15
// Part 1:初始化错误信息
func init() {
	ErrorUser = ErrorInfo{
		UserNotFound:      errors.New("user not found"),
		UserPasswordError: errors.New("user password error"),
	}
}
