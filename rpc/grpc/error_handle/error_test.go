package errorhandle_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	pb "github.com/hwg1999/go_demo/rpc/grpc/pb/model"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func Test_error_handle_server(t *testing.T) {
	s := grpc.NewServer()
	pb.RegisterOrderManagement5Server(s, &OrderManagement5Impl{})
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func Test_error_handle_client(t *testing.T) {
	conn, err := grpc.NewClient("127.0.0.1:8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewOrderManagement5Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	order, err := client.GetOrder(ctx, &wrapperspb.StringValue{Value: ""})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			log.Println(err)
			return
		}

		switch st.Code() {
		case codes.InvalidArgument:
			for _, d := range st.Details() {
				switch info := d.(type) {
				case *epb.BadRequest_FieldViolation:
					log.Printf("Request Field Invalid: %s", info)
				default:
					log.Printf("Unexpected error type: %s", info)
				}
			}
		default:
			log.Printf("Unhandled error : %s ", st.String())
		}

		return
	}

	log.Print("GetOrder Response -> : ", order)
}

var _ pb.OrderManagement3Server = &OrderManagement3Impl{}

type OrderManagement3Impl struct {
	pb.UnimplementedOrderManagement3Server
}

var _ pb.OrderManagement5Server = &OrderManagement5Impl{}

var orders5 = map[string]pb.Order5{
	"101": {
		Id: "101",
		Items: []string{
			"Google",
			"Baidu",
		},
		Description: "example",
		Price:       0,
		Destination: "example",
	},
}

type OrderManagement5Impl struct {
	pb.UnimplementedOrderManagement5Server
}

// Simple RPC
func (s *OrderManagement5Impl) GetOrder(ctx context.Context, orderId *wrapperspb.StringValue) (*pb.Order5, error) {
	ord, exists := orders5[orderId.Value]
	if exists {
		return &ord, status.New(codes.OK, "ok").Err()
	}

	st := status.New(codes.InvalidArgument,
		"Order does not exist. order id: "+orderId.Value)

	details, err := st.WithDetails(
		&epb.BadRequest_FieldViolation{
			Field:       "ID",
			Description: fmt.Sprintf("Order ID received is not valid"),
		},
	)
	if err == nil {
		return nil, details.Err()
	}

	return nil, st.Err()
}
