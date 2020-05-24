package main

import (
	"github.com/kataras/iris"
	"main/common"
	"main/db"
	"main/modules/service"
	)

func main() {
	//初始化配置文件
	common.InitConfig()
	//初始化数据库连接
	db.InitGorm()
	app := iris.Default()
	//fmt.Println("DDDDDDDDDDDD")
	//app.Handle("GET", "/", func(ctx iris.Context) {
	//	ctx.HTML("Hello World")
	//})
	//
	//app.Get("/ping", func(ctx iris.Context) {
	//	ctx.WriteString("pong")
	//})
	//
	//app.Get("/hello", func(ctx iris.Context) {
	//	ctx.JSON(iris.Map{"message": "Hello iris web framework"})
	//})
	//
	service.CommonHub(app)
	app.Run(iris.Addr(":8080"))
}