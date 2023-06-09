package entity

import "time"

// User
// Maintainers:贺胜  Times:2023-05-06
// Part 1:用户表
// Part 2:定义用户表，包括用户ID、在线状态、昵称、密码、个性签名、用户Token、日关注时间、周关注时间、月关注时间
// Bug修复:修复原WeakFocusTime拼写错误，改为WeekFocusTime
type User struct {
	UserID         int    `gorm:"column:user_id;type:int;not null;primary key auto_increment"`
	UserAvatar     string `gorm:"column:user_avatar;type:varchar(100)"`
	OnlineStatus   int    `gorm:"column:online_status;type:int;default:0"`
	Nickname       string `gorm:"column:nickname;type:nvarchar(12);not null"`
	Password       string `gorm:"column:password;type:varchar(40);not null"`
	Signature      string `gorm:"column:person_signature;type:nvarchar(50);"`
	UserToken      string `gorm:"column:user_token;type:varchar(400)"`
	DayFocusTime   int64  `gorm:"column:day_focus_time;type:bigint;"`
	WeekFocusTime  int64  `gorm:"column:week_focus_time;type:bigint;"`
	MonthFocusTime int64  `gorm:"column:month_focus_time;type:bigint;"`
}

// Group
// Maintainers:贺胜  Times:2023-04-13
// Part 1:群组表
// Part 2:定义群组表，包括群组ID、群组名称、创建者ID、创建时间
type Group struct {
	GroupID    int       `gorm:"column:group_id;type:int;not null;primary key auto_increment"`
	GroupName  string    `gorm:"column:group_name;type:nvarchar(10);not null"`
	CreateUser int       `gorm:"column:create_user_id;type:int;not null"`
	CreateTime time.Time `gorm:"column:create_Time;type:time;not null"`
}

// SendUserMessage
// Maintainers:贺胜  Times:2023-04-13
// Part 1:发送用户消息表
// Part 2:定义发送用户消息表，包括发送用户ID、接收用户ID、消息ID、消息内容、发送时间、阅读状态
type SendUserMessage struct {
	SendUser      int       `gorm:"column:send_user_id;type:int;not null;primary key"`
	ReceiveUser   int       `gorm:"column:receive_user_id;type:int;not null;primary key"`
	UserMessageID int       `gorm:"column:message_id;type:int;default:1;primary key auto_increment"`
	Message       string    `gorm:"column:send_message;type:text;not null"`
	SendTime      time.Time `gorm:"column:send_time;type:datetime;not null"`
	ReadStatus    int       `gorm:"column:read_status;type:int;default:0"`
}

// SendGroupMessage
// Maintainers:贺胜  Times:2023-04-13
// Part 1:发送群组消息表
// Part 2:定义发送群组消息表，包括发送用户ID、接收群组ID、消息ID、消息内容、发送时间、阅读状态、评价信息
type SendGroupMessage struct {
	SendUser              int       `gorm:"column:send_user_id;type:int;primary key"`
	ReceiveGroup          int       `gorm:"column:receive_group_id;type:int;primary key"`
	GroupMessageID        int       `gorm:"column:group_message_id;type:int;default:1;primary key auto_increment"`
	Message               string    `gorm:"column:send_message;type:text;not null"`
	SendTime              time.Time `gorm:"column:send_time;type:datetime;not null"`
	ReadStatus            int       `gorm:"column:read_status;type:int;default:1"`
	EvaluationInformation string    `gorm:"column:evaluation_information;type:text"`
}

// AchievementList
// Maintainers:贺胜  Times:2023-04-13
// Part 1:成就列表
// Part 2:定义成就列表，包括成就ID、成就名称、成就内容
type AchievementList struct {
	AchievementID      int    `gorm:"column:achievement_id;type:int;primary key auto_increment"`
	AchievementTitle   string `gorm:"column:achievement_title;type:nvarchar(15);not null"`
	AchievementContact string `gorm:"column:achievement_contact;type:nvarchar(30);not null"`
}

// GetAchievement
// Maintainers:贺胜  Times:2023-04-13
// Part 1:获取成就表
// Part 2:定义获取成就表，包括获取用户ID、获取成就ID、获取时间、获取状态
type GetAchievement struct {
	GetUserID        int       `gorm:"column:get_user_id;type:int;primary key"`
	GetAchievementID int       `gorm:"column:get_achievement_id;type:int;primary key"`
	GetTime          time.Time `gorm:"column:get_time;type:datetime;not null"`
	GetStatus        int       `gorm:"column:get_status;type:int;default:1"`
}

// GroupFile
// Maintainers:贺胜  Times:2023-04-13
// Part 1:群组文件表
// Part 2:定义群组文件表，包括文件ID、文件名称、文件类型、文件保存路径、文件上传用户ID、文件下载次数、群组ID、上传时间
type GroupFile struct {
	FileID            int       `gorm:"column:file_id;type:int;primary key auto_increment"`
	FileName          string    `gorm:"column:file_name;type:nvarchar(100);not null"`
	FileType          string    `gorm:"column:file_type;type:nvarchar(10);not null"`
	FilePath          string    `gorm:"column:file_save_path;type:nvarchar(30);not null"`
	FileUploadUserID  int       `gorm:"column:file_upload_users_id;type:int;not null"`
	FileDownloadTimes int       `gorm:"column:file_download_times;type:int;default:0"`
	GroupID           int       `gorm:"column:group_id;type:int;not null"`
	UploadTime        time.Time `gorm:"column:upload_time;type:datetime;not null"`
}

// GroupMembers
// Maintainers:贺胜  Times:2023-04-13
// Part 1:群组成员表
// Part 2:定义群组成员表，包括用户ID、群组ID、加入时间
type GroupMembers struct {
	UserID  int       `gorm:"column:user_id;type:int;not null;primary key"`
	GroupID int       `gorm:"column:group_id;type:int;not null;primary key"`
	AddTime time.Time `gorm:"column:add_time;type:datetime;not null"`
}

// AddFriends
// Maintainers:贺胜  Times:2023-04-13
// Part 1:添加好友表
// Part 2:定义添加好友表，包括用户ID、好友ID
type AddFriends struct {
	UserID   int `gorm:"column:user_id;not null;primary key"`
	FriendID int `gorm:"column:friend_id;not null;primary key"`
}

// FriendsList
// Maintainers:贺胜  Times:2023-04-13
// Part 1:好友列表
// Part 2:定义好友列表，包括好友列表ID、用户ID、好友ID、好友备注、好友状态
type FriendsList struct {
	FriendListID int    `gorm:"column:friend_list_id;type:int;primary key auto_increment"`
	UserID       int    `gorm:"column:user_id;type:int;not null"`
	FriendID     int    `gorm:"column:friend_id;type:int;not null"`
	FriendRemark string `gorm:"column:friend_remark;type:nvarchar(10)"`
	FriendStatus int    `gorm:"column:friend_status;type:int;default:0"`
}

// PlanList
// Maintainers:贺胜  Times:2023-04-13
// Part 1:计划列表
// Part 2:定义计划列表，包括计划列表ID、用户ID、计划开始时间、计划内容、计划标志、计划完成状态、添加计划时间
type PlanList struct {
	PlanListID     int       `gorm:"column:plan_list_id;type:int;primary key auto_increment"`
	UserID         int       `gorm:"column:user_id;type:int;not null"`
	PlanBeginTime  time.Time `gorm:"column:begin_time;type:datetime"`
	PlanContent    string    `gorm:"column:content;type:text;not null"`
	Signify        int       `gorm:"column:signify;type:tinyint;default:0"`
	FinishedStatus int       `gorm:"column:finished_status;type:tinyint;default:0"`
	AddPlanTime    time.Time `gorm:"column:add_plan_time;type:datetime;not null"`
}
