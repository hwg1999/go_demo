package grpc_test

import (
	"context"
	"errors"
	"fmt"
	"net"
	"testing"
	"time"

	pb "github.com/hwg1999/go_demo/rpc/grpc/pb/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func Test_grpc_server(t *testing.T) {
	// 监听127.0.0.1:8090
	ln, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println("建立监听出错")
		return
	}
	// 实例化gRPC服务端
	server := grpc.NewServer()
	// 注册UserService服务
	pb.RegisterUserServiceServer(server, &GrpcUserService{})
	// 向gRPC服务端注册反射服务
	reflection.Register(server)
	// 启动gRPC服务
	if err := server.Serve(ln); err != nil {
		fmt.Println("启动gRPC服务失败")
	}
}

func Test_grpc_client(t *testing.T) {
	// 远程连接凭证，insecure模式下禁用了传输安全认证
	credentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	// 连接gRPC服务器
	conn, err := grpc.NewClient(":8090", credentials)
	if err != nil {
		fmt.Printf("连接失败：%v\n", err)
	}
	// 延迟关闭连接
	defer conn.Close()
	// 初始化UserService客户端
	userClient := pb.NewUserServiceClient(conn)
	//定义超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 远程调用服务端方法
	response, err := userClient.GetUserInfo(ctx, &pb.UserRequest{Name: "golang"})
	if err != nil {
		fmt.Printf("调用GetUserInfo失败：%v\n", err)
		return
	}

	fmt.Printf("用户ID：%+v\n", response.Id)
	fmt.Printf("用户名：%+v\n", response.Username)
	fmt.Printf("用户昵称：%+v\n", response.Nickname)
}

type GrpcUserService struct {
	pb.UnimplementedUserServiceServer
}

var _ pb.UserServiceServer = &GrpcUserService{}

func (service *GrpcUserService) GetUserInfo(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	fmt.Printf("try to get user info by %s", request.GetName())

	if len(request.GetName()) == 0 {
		return nil, errors.New("the query name cannot be empty")
	}

	// 通常逻辑是从数据库或者缓存中查询
	response := pb.UserResponse{Id: 123, Username: request.GetName(), Nickname: "尊敬的" + request.GetName()}
	return &response, nil
}
