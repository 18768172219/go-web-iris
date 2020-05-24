package service

import (
	"github.com/kataras/iris"
	)

/**
 基础路由的配置信息
 */
func CommonHub(app *iris.Application)  {
	party := partySetting(app)
	//party
	userRouter(party)
}

/**
配置 文件 信息
 */
func partySetting(app *iris.Application) (party iris.Party)  {
	app.Logger().SetLevel("debug")
    //设置 party信息
	party = app.Party("/api/admin").AllowMethods(iris.MethodOptions)
	return party
}
