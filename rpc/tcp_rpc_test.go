package rpc_test

import (
	"fmt"
	"net"
	"net/rpc"
	"testing"
)

func Test_rcp_rpc_server(t *testing.T) {
	rpc.Register(&UserHandler{})

	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println("监听TCP端口出错")
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("建立TCP连接失败")
			continue
		}

		go func(conn net.Conn) {
			rpc.ServeConn(conn)
		}(conn)
	}
}

func Test_tcp_rpc_client(t *testing.T) {
	// 利用TCP连接远程服务端，并获得客户端对象
	client, err := rpc.Dial("tcp", ":8090")
	if err != nil {
		fmt.Println("与服务端建立连接出错")
		return
	}
	defer client.Close()

	user := &UserDetail{}
	err = client.Call("UserHandler.GetUserInfo", "golang", user)
	if err != nil {
		fmt.Println("获取用户信息出错")
		return
	}

	fmt.Printf("用户ID：%+v\n", user.Id)
	fmt.Printf("用户名：%+v\n", user.Username)
	fmt.Printf("用户昵称：%+v\n", user.Nickname)
}
