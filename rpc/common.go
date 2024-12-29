package rpc_test

import (
	"errors"
	"fmt"
)

type UserDetail struct {
	Id       int
	Username string
	Nickname string
}

// 自定义结构体，用于绑定用户处理方法
type UserHandler struct {
}

// 获取用户信息的方法绑定到结构体UserHandler的指针类型上
func (handler *UserHandler) GetUserInfo(name string, reply *UserDetail) (err error) {
	// 打印一行注释信息，一般程序入口都会如此处理，方便出现错误时检查执行进度
	fmt.Printf("try to get user info by %s", name)
	// 校验入参的合法性，对于空参数，直接返回错误
	if len(name) == 0 {
		return errors.New("the query name cannot be empty")
	}

	// 通常逻辑是从数据库或者缓存中查询，此处是直接赋值。RPC响应通过reply来传递到客户端
	reply.Id = 123
	reply.Username = name
	reply.Nickname = "尊敬的" + name

	// 函数声明中，返回值是错误对象。当一切正常时，错误为空
	return nil
}
