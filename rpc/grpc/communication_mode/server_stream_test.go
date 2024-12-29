package communicationmode_test

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"testing"

	pb "github.com/hwg1999/go_demo/rpc/grpc/pb/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func Test_server_stream_server(t *testing.T) {
	s := grpc.NewServer()

	pb.RegisterOrderManagement2Server(s, &OrderManagement2Impl{})

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func Test_server_stream_client(t *testing.T) {
	conn, err := grpc.NewClient("127.0.0.1:8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewOrderManagement2Client(conn)
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	stream, err := c.SearchOrders(ctx, &wrapperspb.StringValue{Value: "Google"})
	if err != nil {
		panic(err)
	}

	for {
		order, err := stream.Recv()
		if err == io.EOF {
			break
		}
		log.Println("Search Result:", order)
	}
}

var _ pb.OrderManagement2Server = &OrderManagement2Impl{}

type OrderManagement2Impl struct {
	pb.UnimplementedOrderManagement2Server
}

var orders2 = make(map[string]pb.Order2)

func (s *OrderManagement2Impl) SearchOrders(query *wrapperspb.StringValue, stream pb.OrderManagement2_SearchOrdersServer) error {
	for _, order := range orders2 {
		for _, str := range order.Items {
			if strings.Contains(str, query.Value) {
				err := stream.Send(&order)
				if err != nil {
					return fmt.Errorf("error send: %v", err)
				}
			}
		}
	}

	return nil
}

func init() {
	orders2["Google"] = pb.Order2{
		Id:          "101",
		Items:       []string{"apple", "banana", "Google"},
		Description: "fruits",
		Price:       20.0,
		Destination: "none",
	}
}
