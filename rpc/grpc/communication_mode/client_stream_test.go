package communicationmode_test

import (
	"context"
	"io"
	"log"
	"net"
	"testing"

	pb "github.com/hwg1999/go_demo/rpc/grpc/pb/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func Test_client_stream_server(t *testing.T) {
	s := grpc.NewServer()

	pb.RegisterOrderManagement3Server(s, &OrderManagement3Impl{})

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func Test_client_stream_client(t *testing.T) {
	conn, err := grpc.NewClient("127.0.0.1:8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewOrderManagement3Client(conn)
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	stream, err := c.UpdateOrders(ctx)
	if err != nil {
		panic(err)
	}

	if err := stream.Send(&pb.Order3{
		Id:          "00",
		Items:       []string{"A", "B"},
		Description: "A with B",
		Price:       0.11,
		Destination: "ABC",
	}); err != nil {
		panic(err)
	}

	if err := stream.Send(&pb.Order3{
		Id:          "01",
		Items:       []string{"C", "D"},
		Description: "C with D",
		Price:       1.11,
		Destination: "ABCDEFG",
	}); err != nil {
		panic(err)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}

	log.Printf("Update Orders Res : %s", res)
}

var _ pb.OrderManagement3Server = &OrderManagement3Impl{}

type OrderManagement3Impl struct {
	pb.UnimplementedOrderManagement3Server
}

var orders3 = make(map[string]pb.Order3)

func (s *OrderManagement3Impl) UpdateOrders(stream pb.OrderManagement3_UpdateOrdersServer) error {
	ordersStr := "Updated Order IDs : "
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&wrapperspb.StringValue{Value: "orders processed" + ordersStr})
		}

		orders3[order.Id] = *order

		log.Println("Order ID ", order.Id, ": Updated")
		ordersStr += order.Id + ", "
	}
}

func init() {
	orders3["00"] = pb.Order3{
		Id:          "00",
		Items:       []string{"C", "D"},
		Description: "C with D",
		Price:       1.11,
		Destination: "ABCDEFG",
	}

	orders3["01"] = pb.Order3{
		Id:          "01",
		Items:       []string{"C", "D"},
		Description: "C with D",
		Price:       1.11,
		Destination: "ABCDEFG",
	}
}
