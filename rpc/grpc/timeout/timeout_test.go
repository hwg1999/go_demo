package timeout_test

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

func Test_timeout_server(t *testing.T) {
	s := grpc.NewServer(
		grpc.ConnectionTimeout(3*time.Second),
		grpc.UnaryInterceptor(unaryServerInterceptor),
	)
	pb.RegisterOrderManagement6Server(s, &OrderManagement6Impl{})

	lit, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lit); err != nil {
		panic(err)
	}
}

func Test_timeout_client(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	conn, err := grpc.NewClient("127.0.0.1:8090",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(unaryClientInterceptor))
	if err != nil {
		if err == context.DeadlineExceeded {
			panic(err)
		}
		panic(err)
	}

	c := pb.NewOrderManagement6Client(conn)

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	order := pb.Order6{
		Id:          "101",
		Items:       []string{"iPhone XS", "Mac Book Pro"},
		Destination: wrapperspb.String("San Jose, CA"),
		Price:       2300.00,
	}
	res, err := c.AddOrder(ctx, &order)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.DeadlineExceeded {
			panic(err)
		}
		panic(err)
	}

	log.Println(res)
}

func unaryClientInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := invoker(ctx, method, req, reply, cc, opts...)

	return err
}

func unaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	m, err := handler(ctx, req)

	return m, err
}

var _ pb.OrderManagement6Server = &OrderManagement6Impl{}

var orders6 = make(map[string]pb.Order6, 0)

type OrderManagement6Impl struct {
	pb.UnimplementedOrderManagement6Server
}

// Simple RPC
func (s *OrderManagement6Impl) AddOrder(ctx context.Context, orderReq *pb.Order6) (*wrapperspb.StringValue, error) {
	log.Printf("Order Added. ID : %v", orderReq.Id)

	select {
	case <-ctx.Done():
		return nil, status.Errorf(codes.Canceled, "Client cancelled, abandoning.")
	default:
	}

	orders6[orderReq.Id] = *orderReq

	return &wrapperspb.StringValue{Value: "Order Added: " + orderReq.Id}, nil
}
