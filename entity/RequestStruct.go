package entity

// RequestBodyUserLogin
// Maintainers:贺胜 Times：2023-05-20
// Part 1:用户登陆Post请求方式结构体
type RequestBodyUserLogin struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

// RequestBodyUserRegister
// Maintainers:贺胜 Times: 2023-05-20
// Part 1: 用户注册Post请求方式结构体
type RequestBodyUserRegister struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Signature string `json:"signature"`
}
