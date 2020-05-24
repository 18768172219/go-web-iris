package service

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/hero"
	"main/modules/model"
	"main/modules/repository"
	"time"
)

/**
用户路由 路由
 */
func userRouter(party iris.Party)  {
	//user
	user := party.Party("/user")
	//userMenu
	userService := UserService{userRepository:repository.UserRepository{}}
	//addUser
	user.Post("/addUser", hero.Handler(userService.saveUser))
	//getAll
	user.Get("/getUsers", hero.Handler(userService.all))

}

/**
 用户 struct
 */
type UserService struct {
  userRepository repository.UserRepository
}

/**
保存用户信息
 */
func (us *UserService) saveUser(ctx iris.Context)  {
	var (
		err error
		user = new(model.User)
	)
	err = ctx.ReadJSON(&user)
	id := time.Now().UnixNano() / 1e6
	user.Id = uint64(id)
	if err != nil {
		goto ERR
	}
	if err = us.userRepository.SaveUser(user); err != nil {
		goto ERR
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(Code(0))
	return
ERR:
	golog.Errorf("保存用户失败, 用户信息：%s, 错误: %s",  user.Name, err)
}

/**
查询全部用户
 */
func (us *UserService) all(ctx iris.Context)  {
	var (
		users []*model.User
		err error
	)
	if users, err = us.userRepository.AllUsers(); err != nil {
		goto ERR
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(json(Code(0), users))
	return
ERR:
	golog.Errorf("查询所有用户失败，错误：%s", err)
}

type Code int

/**
  json包装
 */
func json(code Code, data interface{}) context.Map {
	return iris.Map{
		"code": fmt.Sprintf("%d", code),
		"msg":  code,
		"data": data,
	}
}