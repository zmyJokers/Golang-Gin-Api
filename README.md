# Project Middleground（中台）
>该项目为golang语言开发的纯API中台项目，请严格遵守[代码开发规范](#代码开发规范)

### 目录树
***
```
|- app
  |- controller  控制层
  |- dao  数据层
  |- middleware  中间件
  |- request 参数初始化
     |- logRequest 日志参数处理
     |- init.go 参数整合
  |- service 业务层
  |- init.go 项目初始化
|- common 公共组件
|- conf 配置文件 后缀名.go
|- database 模型层
    |- init.go 数据库连接初始化
|- routes 路由
|- main.go 入口文件
```
# 使用流程介绍
```
# git clone下载项目
`git clone git@git.transecs.com:middleground/middleground.git`

# 提示
首次安装项目，请先运行以下命令 安装必要第三方包
go get -u github.com/gin-gonic/gin v1.8.1
go get -u github.com/go-sql-driver/mysql v1.6.0
go get -u github.com/streadway/amqp v1.0.0

# 运行项目 
`go run middleground` 运行项目（最好使用包的形式运行）
* 运行失败请检查相关提示

# 地址 
127.0.0.1:8800

注意 : go mod init middleground （middleware是包名不能修改）
```
### 代码开发规范
> 注意：各个目录只负责自身业务不能混乱使用！！！这点很重要。
***
+ controller目录，文件名称必须结尾包含Controller (testController), func 方法以struct结构体引用结合interface进行调用
+ dao目录，数据层负责sql CURD, 文件名称结尾必须包含Dao (testDao)
+ request目录， API Func Params初始化
+ service目录，负责业务逻辑处理、结合dao目录进行开发

# 发布Linux流程
+ `go env -w GOOS=linux`
+ `go build -o ProjectName(同项目名称)`
+ `nohup ./main &`

