package main

import (
	"fmt"
	"go-copyright-p1/eths"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"go-copyright-p1/configs"
	"go-copyright-p1/routes"
)

var EchoObj *echo.Echo //echo框架对象全局定义

//静态文件处理
func staticFile() {
	//设置根目录
	EchoObj.Static("/", "static/pc/home")
	//全路径处理
	EchoObj.Static("/static", "static")
	//其余网页文件夹加载地址处理
	EchoObj.Static("/upload", "static/pc/upload")
	EchoObj.Static("/css", "static/pc/css")
	EchoObj.Static("/assets", "static/pc/assets")
	EchoObj.Static("/user", "static/pc/user")
}

func main() {

	fmt.Printf("get config %v ,%v\n", configs.Config.Common.Port, configs.Config.Db.ConnStr)
	EchoObj = echo.New()             //创建echo对象
	EchoObj.Use(middleware.Logger()) //安装日志中间件
	EchoObj.Use(middleware.Recover())
	EchoObj.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	EchoObj.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	go eths.EventSubscrib("ws://localhost:8546", configs.Config.Eth.QaAddr)
	staticFile()

	//注册
	EchoObj.POST("/account", routes.Register)
	//session处理
	EchoObj.GET("/session", routes.GetSession)
	//登录
	EchoObj.POST("/login", routes.Login)
	//上传资产
	EchoObj.POST("/content", routes.Upload)
	//查看当前用户资产
	EchoObj.GET("/content", routes.GetContents)
	//查看图片二进制数据
	//EchoObj.GET("/content/:title",routes.GetContent)
	//发起交易
	EchoObj.POST("/auction", routes.Auction)
	//查看资产交易记录
	EchoObj.POST("/tran", routes.Tranrecode)
	//查看版权交易记录
	EchoObj.POST("/copytran", routes.CopyTranrecode)
	//获取交易
	EchoObj.GET("/auctions", routes.GetAuction)
	//购买资产
	EchoObj.GET("/buy", routes.Buy)
	//提交申请
	EchoObj.GET("/apply", routes.Apply)
	//取消交易
	EchoObj.GET("/revoke", routes.Revoke)
	//查询待确认交易
	EchoObj.GET("/oklist", routes.OkList)
	//拒绝交易版权
	EchoObj.GET("/no", routes.Notrans)
	//确认版权转让
	EchoObj.GET("/ok", routes.ConfirmTran)
	//写文件到服务器本地
	EchoObj.GET("/down", routes.Download)
	//写文件到服务器本地
	EchoObj.GET("/downfile", routes.DownloadFile)

	//以下与秘钥相关
	//设置秘钥
	EchoObj.GET("/setpass", routes.SetPass)
	//查看秘钥
	EchoObj.GET("/lookpass", routes.GetPass)

	EchoObj.GET("/ping", routes.PingHandler)                        //路由测试函数
	EchoObj.Logger.Fatal(EchoObj.Start(configs.Config.Common.Port)) //启动服务
}
