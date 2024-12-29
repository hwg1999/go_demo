package rpc_test

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"testing"
)

func Test_http_rpc_server(t *testing.T) {
	// 将UserHandler的一个实例对象注册为RPC服务
	rpc.Register(&UserHandler{})
	// RPC通信，使用HTTP协议进行处理
	rpc.HandleHTTP()
	// 监听8090端口，并获得监听器对象
	listener, err := net.Listen("tcp", ":8090")
	// 如果发生错误，则打印错误信息，并直接退出程序
	if err != nil {
		fmt.Println("监听TCP端口出错")
		return
	}

	// 启动监听服务，准备接收来自客户端的RPC请求
	err = http.Serve(listener, nil)
	if err != nil {
		fmt.Println("获取TCP连接出错")
	}
}

func Test_http_rpc_client(t *testing.T) {
	// 以HTTP方式连接127.0.0.1:8090，并获得连接客户端对象client
	client, err := rpc.DialHTTP("tcp", ":8090")
	// 如果出现错误，则打印错误提示，并直接退出
	if err != nil {
		fmt.Println("RPC拨号连接出错")
		return
	}

	// 创建UserDetail对象
	user := &UserDetail{}
	// 利用客户端对象client调用远程方法，并传入参数
	err = client.Call("UserHandler.GetUserInfo", "golang", user)
	// 如果出现错误，则打印错误信息，并退出程序
	if err != nil {
		fmt.Println("获取用户信息出错")
		return
	}

	// 打印从服务端获得的响应信息
	fmt.Printf("用户ID：%+v\n", user.Id)
	fmt.Printf("用户名：%+v\n", user.Username)
	fmt.Printf("用户昵称：%+v\n", user.Nickname)
}
