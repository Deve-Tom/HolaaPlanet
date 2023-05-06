package configs

import (
	"HolaaPlanet/entity"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB // 数据库
	HostIP     string   // 本机IP
	ServerIP   string   // 服务器IP
	HostPort   string   // 本机端口
	ServerPort string   // 服务器端口
)

// InitRootConfig
// Maintainers:贺胜  Times:2023-04-05
// Part 1:配置文件加载
// Part 2:加载"./"目录下的"conf.yaml"配置文件，加载成功则将配置文件中的配置信息赋值给全局变量，否则报错"Load config file false"
// Part 3:加载配置文件并为后续操作建立基础
func InitRootConfig() {
	// 加载"./"目录下的"conf.yaml"配置文件
	viper.AddConfigPath("./")
	viper.SetConfigFile("conf.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Load config file false")
		return
	}
}

// LoadDB
// Maintainers:贺胜 Times:2023-04-09
// Part 1:数据库加载
// Part 2:使用Gorm框架加载数据库，加载成功则返回可操作的DB指针否则报错"Load DB false "并返回nil
// Part 3:加载数据库并为后续操作建立基础/*
func LoadDB() *gorm.DB {
	InitRootConfig()
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"))
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Load DB false")
	}

	return db
}

// InitDB
// Maintainers:贺胜 Times:2023-05-06
// Part 1:数据库初始化
// Part 2:使用Gorm框架初始化数据库，初始化成功则返回可操作的DB指针否则报错"Init DB false "并返回nil
// Part 3:初始化数据库，包括建表、建立主键和建立外键
// Bug修复：修改数据库中users表的时间数据类型从datetime改为time(hh:mm:ss) 以及删除组消息与用户消息/*
func InitDB() *gorm.DB {
	db := LoadDB()
	if db == nil {
		return nil
	}
	// 自动创建数据表
	createTable(db)
	// 建立主键
	createPrimaryKey(db)
	// 建立自增长
	changeTableAutoIncrement(db)
	// 建立外键
	createForeignKey(db)
	// 改变时间数据库类型
	changeTableUserFocusTimeType(db)

	return db
}

// createTable
// Maintainers:贺胜 Times:2023-04-13
// Part 1:创建数据表
// Part 2:使用Gorm框架创建数据表，创建成功则返回可操作的DB指针否则报错"Create table false "以及具体错误信息
// Part 3:创建数据表，不会重复创建，如果存在则跳过/*
func createTable(db *gorm.DB) {
	// 自动创建数据表
	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		fmt.Printf("Create user table false:%v", err)
	}

	err = db.AutoMigrate(&entity.Group{})
	if err != nil {
		fmt.Printf("Create Group table false:%v", err)
	}

	err = db.AutoMigrate(&entity.SendGroupMessage{})
	if err != nil {
		fmt.Printf("Create send group message table false:%v", err)
	}

	err = db.AutoMigrate(&entity.SendUserMessage{})
	if err != nil {
		fmt.Printf("Create send user message table false:%v", err)
	}

	err = db.AutoMigrate(&entity.AchievementList{})
	if err != nil {
		fmt.Printf("Create achievement list table false:%v", err)
	}

	err = db.AutoMigrate(&entity.GetAchievement{})
	if err != nil {
		fmt.Printf("Create get achievement table false:%v", err)
	}

	err = db.AutoMigrate(&entity.GroupFile{})
	if err != nil {
		fmt.Printf("Create group file table false:%v", err)
	}

	err = db.AutoMigrate(&entity.GroupMembers{})
	if err != nil {
		fmt.Printf("Create group members table false:%v", err)
	}

	err = db.AutoMigrate(&entity.AddFriends{})
	if err != nil {
		fmt.Printf("Create add friends table false:%v", err)
	}

	err = db.AutoMigrate(&entity.FriendsList{})
	if err != nil {
		fmt.Printf("Create firend list table false:%v", err)
	}

	err = db.AutoMigrate(&entity.PlanList{})
	if err != nil {
		fmt.Printf("Create plan list table false:%v", err)
	}
}

// createPrimaryKey
// Maintainers:贺胜 Times:2023-04-13
// Part 1:创建主键
// Part 2:使用Gorm框架创建主键
// Part 3:创建主键，不会重复创建，如果存在则跳过/*
func createPrimaryKey(db *gorm.DB) {
	// 建立主键
	db.Exec("alter table `users` add primary key (user_id)")
	db.Exec("alter table `groups` add primary key (group_id)")
	// db.Exec("alter table `send_group_messages` add primary key (send_user_id,receive_group_id,group_message_id)")
	// db.Exec("alter table `send_user_messages` add primary key (send_user_id,receive_user_id,message_id)")
	db.Exec("alter table `send_group_messages` add primary key (group_message_id)")
	db.Exec("alter table `send_user_messages` add primary key (message_id)")
	db.Exec("alter table `get_achievements` add primary key (get_user_id,get_achievement_id)")
	db.Exec("alter table `group_files` add primary key (group_id,file_id)")
	db.Exec("alter table `group_members` add primary key (group_id,user_id)")
	db.Exec("alter table `add_friends` add primary key (user_id,friend_id)")
	db.Exec("alter table `friends_lists` add primary key (friend_list_id)")
	db.Exec("alter table `plan_lists` add primary key (plan_list_id)")
	db.Exec("alter table `achievement_lists` add primary key (achievement_id)")
}

// changeTableAutoIncrement
// Maintainers:贺胜 Times:2023-04-13
// Part 1:修改表的自增长
// Part 2:使用Gorm框架修改表的自增长
// Part 3:修改表的自增长，不会重复修改，如果存在则跳过/*
func changeTableAutoIncrement(db *gorm.DB) {
	db.Exec("alter table `users` change user_id user_id bigint not null auto_increment;")
	db.Exec("alter table `groups` change group_id group_id bigint not null auto_increment;")

	db.Exec("alter table `send_group_messages` change group_message_id group_message_id bigint not null auto_increment;")
	db.Exec("alter table `send_user_messages` change message_id message_id bigint not null auto_increment;")

	db.Exec("alter table `achievement_lists` change achievement_id achievement_id bigint not null auto_increment;")
	// db.Exec("alter table `group_files` change file_id file_id bigint not null auto_increment;")
	db.Exec("alter table `friends_lists` change friend_list_id friend_list_id bigint not null auto_increment;")
	db.Exec("alter table `plan_lists` change plan_list_id plan_list_id bigint not null auto_increment;")
	db.Exec("alter table `achievement_lists` change achievement_id achievement_id bigint not null auto_increment;")
}

// changeTableUserFocusTimeType
// Maintainers:贺胜 Times:2023-05-06
// Part 1:修改表中用户关注时长的数据库数据格式为time(00:00:00)
// Part 2:使用Gorm框架修改表中用户关注时长的数据库数据格式为time(00:00:00)
// Part 3:修改表中用户关注时长的数据库数据格式为time(00:00:00)，不会重复修改，如果存在则跳过/*
func changeTableUserFocusTimeType(db *gorm.DB) {
	db.Exec("alter table `users` change day_focus_time day_focus_time time null;")
	db.Exec("alter table `users` change week_focus_time week_focus_time time null;")
	db.Exec("alter table `users` change month_focus_time month_focus_time time null;")
}

// createForeignKey
// Maintainers:贺胜 Times:2023-04-13
// Part 1:创建外键
// Part 2:使用Gorm框架创建外键
// Part 3:创建外键，不会重复创建，如果存在则跳过/*
func createForeignKey(db *gorm.DB) {
	// 建立外键
	db.Exec("alter table `groups` add constraint fk_groups_users1 foreign key (create_user_id)  references users(user_id) on delete cascade on update cascade;")
	db.Exec("alter table `send_group_messages` add constraint fk_send_group_messages_send_user_id1 foreign key (send_user_id) references users(user_id) on delete cascade on update cascade;")
	db.Exec("alter table `send_group_messages` add constraint fk_send_group_messages_send_receive_group_id1 foreign key (receive_group_id) references `groups`(group_id) on delete cascade on update cascade;")
	db.Exec("alter table `send_user_messages` add constraint fk_send_user_messages_send_user_id1 foreign key (send_user_id) references users(user_id) on delete cascade on update cascade;")
	db.Exec("alter table `send_user_messages` add constraint fk_send_user_messages_receive_user_id1 foreign key (receive_user_id) references users(user_id) on delete cascade on update cascade;")
	db.Exec("alter table `get_achievements` add constraint fk_send_user_get_achievements_get_user_id1 foreign key (get_user_id) references users(user_id) on delete cascade on update cascade;")
	db.Exec("alter table `get_achievements` add constraint fk_get_achievements_get_achievement_id1 foreign key (get_achievement_id) references achievement_lists(achievement_id) on delete cascade on update cascade;")
	db.Exec("alter table `group_members` add constraint fk_group_members_user_id1 foreign key (user_id) references users(user_id) on delete cascade on update cascade;")
	db.Exec("alter table `group_members` add constraint fk_group_members_group_id1 foreign key (group_id) references `groups`(group_id) on delete cascade on update cascade;")
	db.Exec("alter table `add_friends` add constraint fk_add_friends_user_id1 foreign key (user_id) references users(user_id) on delete cascade on update cascade;")
	db.Exec("alter table `add_friends` add constraint fk_add_friends_friend_id1 foreign key (friend_id) references users(user_id) on delete cascade on update cascade;")
	db.Exec("alter table `friends_lists` add constraint fk_friends_lists_friend_id1 foreign key (friend_id) references users(user_id) on delete cascade on update cascade;")
	db.Exec("alter table `friends_lists` add constraint fk_friends_lists_user_id1 foreign key (user_id) references users(`user_id`) on delete cascade on update cascade;")
	db.Exec("alter table `plan_lists` add constraint fk_plan_lists_user_id1 foreign key (user_id) references users(user_id) on delete cascade on update cascade;")
	db.Exec("alter table `group_files` add constraint fk_group_files_group_id1 foreign key (group_id) references `groups`(group_id) on delete cascade on update cascade;")
}

// init
// Maintainers:贺胜 Times:2023-04-13
// Part 1:初始化变量/*
func init() {
	InitRootConfig()
	ServerIP = viper.GetString("server.host")
	ServerPort = viper.GetString("server.port")
	HostIP = viper.GetString("mysql.host")
	HostPort = viper.GetString("mysql.port")
	DB = LoadDB()
}
