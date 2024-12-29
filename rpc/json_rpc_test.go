package rpc_test

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

func Test_json_rpc_server(t *testing.T) {
	// 注册服务端对象
	rpc.Register(&UserHandler{})
	// 同样利用TCP进行消息传输
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println("监听TCP端口出错")
		return
	}

	// 循环处理网络获取，并处理网络连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("建立TCP连接失败")
			continue
		}

		go func(conn net.Conn) {
			// 利用jsonrpc中的ServeConn(conn)来处理网络连接
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}

func Test_json_rpc_client(t *testing.T) {
	// 调用jsonrpc.Dial()函数，与服务端建立连接
	client, err := jsonrpc.Dial("tcp", ":8090")
	if err != nil {
		fmt.Println("与服务端建立连接出错")
		return
	}
	defer client.Close()

	user := &UserDetail{}
	// 调用远程RPC服务，并传递参数
	err = client.Call("UserHandler.GetUserInfo", "golang", user)
	if err != nil {
		fmt.Println("获得用户信息出错")
		return
	}

	fmt.Printf("用户ID：%+v\n", user.Id)
	fmt.Printf("用户名：%+v\n", user.Username)
	fmt.Printf("用户昵称：%+v\n", user.Nickname)
}
