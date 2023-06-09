package entity

import "time"

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
	UserID int `json:"user_id"`
	Time   int `json:"time"`
}

// RequestUpdateInformation
// Maintainers:邵洁 Times: 2023-06-09
// Part 1: 用户信息修改Post请求方式结构体
type RequestUpdateInformation struct {
	UserID    int    `json:"user_id"`
	NickName  string `json:"nickname"`
	Signature string `json:"person_signature"`
	PassWord  string `json:"password"`
}

// RequestBodyGets
// Maintainers:宋昭城 Times: 2023-06-09
// Part 1: 用户获取名言警句、获取成就结构体
type RequestBodyGets struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}

// RequestBodyPlan
// Maintainers:宋昭城 Times: 2023-06-09
// Part 1: 用户创建计划列表结构体
type RequestBodyPlan struct {
	UserID    string    `json:"user_id"`
	BeginTime time.Time `json:"begin_time"`
	Content   string    `json:"content"`
	Token     string    `json:"token"`
}

// RequestMultiTime
// Maintainers:邵洁 Times: 2023-06-09
// Part 1: 多人专注时间Put请求方式结构体
type RequestMultiTime struct {
	User1ID int `json:"user1_id"`
	User2ID int `json:"user2_id"`
	Time    int `json:"time"`
}
