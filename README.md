# HolaaPlanet

# 

# Project Struct

项目所需用到的数据库为MySQL数据库，需要提前安装好，默认的数据库配置为
```yaml
mysql:
  user: root
  password: 123456
  host: 127.0.0.1
  port: 3306
  database: Holla
```
项目的代码结构如下，在进行开发时需要严格将自己的代码进行分类不能将代码所有功能全部揉捏在一个源代码文件之中。

```Shell
.
├── LICENSE
├── README.md
├── conf.yaml //项目配置文件
├── configs //配置目录
├── controllers //HTTP路由处理
├── go.mod
├── go.sum
├── main.go
├── models //数据库处理
├── services //业务逻辑代码
├── static //存放静态文件
├── tests //测试文件
├── utils //工具函数
└── web //前端源码
```

