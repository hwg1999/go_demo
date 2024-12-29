package communicationmode_test

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	pb "github.com/hwg1999/go_demo/rpc/grpc/pb/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func Test_simple_rpc_server(t *testing.T) {
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &OrderManagementImpl{})
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func Test_simple_rpc_client(t *testing.T) {
	conn, err := grpc.NewClient("127.0.0.1:8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewOrderManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Get Order
	retrievedOrder, err := client.GetOrder(ctx, &wrapperspb.StringValue{Value: "101"})
	if err != nil {
		panic(err)
	}

	log.Print("GetOrder Response -> : ", retrievedOrder)
}

var _ pb.OrderManagementServer = &OrderManagementImpl{}

var orders = make(map[string]pb.Order)

type OrderManagementImpl struct {
	pb.UnimplementedOrderManagementServer
}

func (s *OrderManagementImpl) GetOrder(ctx context.Context, orderId *wrapperspb.StringValue) (*pb.Order, error) {
	ord, exists := orders[orderId.Value]
	if exists {
		return &ord, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Order does not exist. : %s", orderId)
}

func init() {
	orders["101"] = pb.Order{
		Id:          "101",
		Items:       []string{"apple", "banana"},
		Description: "fruits",
		Price:       20.0,
		Destination: "none",
	}
}
