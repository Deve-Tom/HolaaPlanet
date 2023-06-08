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

// RequestPerTime
// Maintainers:邵洁 Times: 2023-06-08
// Part 1: 专注时长Put请求方式结构体
type RequestPerTime struct {
	UserID    int    `json:"user_id"`
	DayTime   string `json:"ay_focus_time"`
	WeekTime  string `json:"week_focus_time"`
	MonthTime string `json:"month_focus_time"`
}
